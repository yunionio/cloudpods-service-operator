// Copyright 2020 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controllers

import (
	"context"
	"fmt"
	"strings"
	"time"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/source"

	onecloudv1 "yunion.io/x/onecloud-service-operator/api/v1"
	"yunion.io/x/onecloud-service-operator/pkg/options"
	"yunion.io/x/onecloud-service-operator/pkg/resources"
)

// AnsiblePlaybookReconciler reconciles a AnsiblePlaybook object
type AnsiblePlaybookReconciler struct {
	ReconcilerBase
	// Enable intensive information collection during the reconcile process
	Dense bool
}

// +kubebuilder:rbac:groups=onecloud.yunion.io,resources=ansibleplaybooks,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=onecloud.yunion.io,resources=ansibleplaybooks/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=onecloud.yunion.io,resources=virtualmachines,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=onecloud.yunion.io,resources=ansibleplaybooktemplates,verbs=get;list;watch;create;update;patch;delete

func (r *AnsiblePlaybookReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()

	var ansiblePlaybook onecloudv1.AnsiblePlaybook
	if err := r.Get(ctx, req.NamespacedName, &ansiblePlaybook); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	log := r.GetLog(&ansiblePlaybook)
	remoteAp := resources.NewAnisblePlaybook(&ansiblePlaybook)

	dealErr := func(err error) (ctrl.Result, error) {
		return r.dealErr(ctx, remoteAp, err)
	}

	var (
		apPendingAfter = time.Duration(options.Options.AnsiblePlaybookConfig.IntervalPending) * time.Second
		apWaitingAfter = time.Duration(options.Options.AnsiblePlaybookConfig.IntervalWaiting) * time.Second
		dense          = options.Options.AnsiblePlaybookConfig.Dense
	)

	has, ret, err := r.UseFinallizer(ctx, remoteAp)
	if !has {
		return ret, err
	}

	if ansiblePlaybook.Status.Phase == onecloudv1.ResourceInvalid {
		return ctrl.Result{}, nil
	}

	if ansiblePlaybook.Status.Phase == onecloudv1.ResourceFinished {
		if ansiblePlaybook.Status.ExternalInfo.Id == "" {
			return ctrl.Result{}, nil
		}
		// sync info first
		apStatus, err := remoteAp.Reconcile(ctx)
		if err != nil {
			return dealErr(err)
		}
		if r.requireUpdate(&ansiblePlaybook, apStatus) {
			ansiblePlaybook.Status = *apStatus
			return ctrl.Result{}, r.Status().Update(ctx, &ansiblePlaybook)
		}
		// clear this
		if err := r.clear(ctx, remoteAp); err != nil {
			return dealErr(err)
		}
		return ctrl.Result{RequeueAfter: time.Second, Requeue: true}, nil
	}

	var playbookTemplate onecloudv1.AnsiblePlaybookTemplate

	if ansiblePlaybook.Status.ExternalInfo.Id == "" {
		// wait for PlaybookTemplateRef
		if ansiblePlaybook.Spec.PlaybookTemplateRef != nil {
			nameSpacedName := types.NamespacedName{
				Namespace: req.Namespace,
				Name:      ansiblePlaybook.Spec.PlaybookTemplateRef.Name,
			}
			if err := r.Get(ctx, nameSpacedName, &playbookTemplate); err != nil {
				if !apierrors.IsNotFound(err) {
					log.Error(err, "unable to fetch ansibleplaybooktemplate")
					return ctrl.Result{}, err
				}
				return r.MarkWaiting(ctx, &ansiblePlaybook, fmt.Sprintf("wait for AnsiblePlaybookTemplate '%s': %s", nameSpacedName, "no such resource"), apWaitingAfter)
			}

		} else if ansiblePlaybook.Spec.PlaybookTemplate != nil {
			playbookTemplate = onecloudv1.AnsiblePlaybookTemplate{
				Spec: *ansiblePlaybook.Spec.PlaybookTemplate,
			}
		} else {
			log.V(0).Info("webhook will stop this but not")
			return ctrl.Result{}, nil
		}

		hosts := make([]resources.AnsiblePlaybookHost, 0, len(ansiblePlaybook.Spec.Inventory))
		// wair for all VitualMachines running
		for _, host := range ansiblePlaybook.Spec.Inventory {
			nameSpacedName := types.NamespacedName{
				Namespace: host.VirtualMachine.Namespace,
				Name:      host.VirtualMachine.Name,
			}
			if len(nameSpacedName.Namespace) == 0 {
				nameSpacedName.Namespace = ansiblePlaybook.Namespace
			}
			var vm onecloudv1.VirtualMachine
			if err := r.Get(ctx, nameSpacedName, &vm); err != nil {
				if !apierrors.IsNotFound(err) {
					log.Error(err, "unable to fetch virtualmachines")
					return ctrl.Result{}, err
				}
				return r.MarkWaiting(ctx, &ansiblePlaybook, fmt.Sprintf("wait for VirtualMachine '%s': %s", nameSpacedName, "no such resource"), apWaitingAfter)
			}
			if vm.Status.Phase != onecloudv1.ResourceRunning {
				return r.MarkWaiting(ctx, &ansiblePlaybook, fmt.Sprintf("wait for VirtualMachine '%s': %s", nameSpacedName, "need phase 'Running'"), apWaitingAfter)
			}
			if len(vm.Status.ExternalInfo.Eip) == 0 && len(vm.Status.ExternalInfo.Ips) == 0 {
				return r.MarkWaiting(ctx, &ansiblePlaybook, fmt.Sprintf("wait for VirtualMachine '%s': %s", nameSpacedName, "need eip or ips"), apWaitingAfter)
			}
			// build vars
			noVars := make([]string, 0)
			vars := make(map[string]interface{}, len(host.Vars))
			for _, temVar := range playbookTemplate.Spec.Vars {
				if value, ok := host.Vars[temVar.Name]; ok {
					v, err := value.GetValue(ctx)
					if err != nil {
						// invalid
						log.Error(err, "StringStore.GetValue")
						ansiblePlaybook.GetResourceStatus().SetPhase(onecloudv1.ResourceInvalid,
							fmt.Sprintf("The value of var '%s' is valid: %s", temVar.Name, err.Error()),
						)
						return ctrl.Result{}, r.Status().Update(ctx, &ansiblePlaybook)
					}
					if v == nil || v.IsZero() {
						// need to wait
						return r.MarkWaiting(ctx, &ansiblePlaybook, fmt.Sprintf("wait for var '%s'", temVar.Name), apWaitingAfter)
					}
					vars[temVar.Name] = v.Interface()
					continue
				}
				if _, ok := ansiblePlaybook.Spec.Vars[temVar.Name]; ok {
					// This variable exists in the public variable
					continue
				}
				if temVar.Required != nil && *temVar.Required {
					noVars = append(noVars, temVar.Name)
					continue
				}
				if temVar.Default != nil {
					vars[temVar.Name] = temVar.Default.Interface()
					continue
				}
			}
			// set phase invalid
			if len(noVars) > 0 {
				ansiblePlaybook.GetResourceStatus().SetPhase(onecloudv1.ResourceInvalid, fmt.Sprintf(
					"Required these missed variables: %s for virtualMachine '%s'",
					strings.Join(noVars, ", "), host.VirtualMachine.Name),
				)
				return ctrl.Result{}, r.Status().Update(ctx, &ansiblePlaybook)
			}
			hosts = append(hosts, resources.AnsiblePlaybookHost{
				VM:   &vm,
				Vars: vars,
			})
		}

		// build common vars
		commonVars := make(map[string]interface{}, len(ansiblePlaybook.Spec.Vars))
		for varName, sv := range ansiblePlaybook.Spec.Vars {
			vv, err := sv.GetValue(ctx)
			if err != nil {
				// invalid
				log.Error(err, "StringStore.GetValue")
				ansiblePlaybook.GetResourceStatus().SetPhase(onecloudv1.ResourceInvalid,
					fmt.Sprintf("The value of var '%s' is valid: %s", varName, err.Error()),
				)
				return ctrl.Result{}, r.Status().Update(ctx, &ansiblePlaybook)
			}
			if vv == nil || vv.IsZero() {
				// need to wait
				return r.MarkWaiting(ctx, &ansiblePlaybook, fmt.Sprintf("wait for var '%s'", varName), apWaitingAfter)
			}
			commonVars[varName] = vv.Interface()
		}

		// all other resources ready, create ansible playbook
		return r.Create(ctx, remoteAp, resources.APCreateParams{hosts, &playbookTemplate, commonVars}, true)
	}

	var recon func(ctx context.Context) (*onecloudv1.AnsiblePlaybookStatus, error)

	recon = func(ctx context.Context) (*onecloudv1.AnsiblePlaybookStatus, error) {
		rs, err := remoteAp.GetStatus(ctx)
		if err != nil {
			return nil, err
		}
		return rs.(*onecloudv1.AnsiblePlaybookStatus), err
	}
	if dense {
		recon = remoteAp.Reconcile
	}
	apStatus, err := recon(ctx)
	if err != nil {
		return dealErr(err)
	}
	if r.requireUpdate(&ansiblePlaybook, apStatus) {
		ansiblePlaybook.Status = *apStatus
		return ctrl.Result{}, r.Status().Update(ctx, &ansiblePlaybook)
	}

	if ansiblePlaybook.Status.Phase == onecloudv1.ResourceFailed {
		return r.Delete(ctx, remoteAp)
	}

	// Pending
	if ansiblePlaybook.Status.Phase == onecloudv1.ResourcePending {
		return ctrl.Result{Requeue: true, RequeueAfter: apPendingAfter}, nil
	}

	// Unkown
	if ansiblePlaybook.Status.Phase == onecloudv1.ResourceUnkown {
		return ctrl.Result{Requeue: true, RequeueAfter: 2 * time.Second}, nil
	}

	return ctrl.Result{}, nil
}

func (r *AnsiblePlaybookReconciler) clear(ctx context.Context, remoteAp resources.AnsiblePlaybook) error {
	_, err := remoteAp.Delete(ctx)
	return err
}

func (r *AnsiblePlaybookReconciler) requireUpdate(ap *onecloudv1.AnsiblePlaybook, newStatus *onecloudv1.AnsiblePlaybookStatus) bool {
	if newStatus == nil {
		return false
	}
	if ap.Status.Phase != newStatus.Phase || ap.Status.Reason != newStatus.Reason {
		return true
	}
	if ap.Status.ExternalInfo.Output != newStatus.ExternalInfo.Output {
		return true
	}
	return false
}

func (r *AnsiblePlaybookReconciler) SetupWithManager(mgr ctrl.Manager) error {
	ap := &onecloudv1.AnsiblePlaybook{}
	return ctrl.NewControllerManagedBy(mgr).
		For(ap).
		Watches(
			&source.Kind{Type: &onecloudv1.VirtualMachine{}},
			&handler.EnqueueRequestForOwner{
				OwnerType:    ap,
				IsController: false,
			},
		).
		Watches(
			&source.Kind{Type: &onecloudv1.AnsiblePlaybookTemplate{}},
			&handler.EnqueueRequestForOwner{
				OwnerType:    ap,
				IsController: false,
			}).
		Complete(r)
}
