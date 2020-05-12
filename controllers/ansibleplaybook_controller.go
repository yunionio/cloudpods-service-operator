/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/go-logr/logr"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/source"

	"yunion.io/x/pkg/utils"

	onecloudv1 "yunion.io/x/onecloud-service-operator/api/v1"
	"yunion.io/x/onecloud-service-operator/provider"
)

const (
	PendingAfter = 15 * time.Second
	WaitingAfter = 15 * time.Second
)

// AnsiblePlaybookReconciler reconciles a AnsiblePlaybook object
type AnsiblePlaybookReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=onecloud.yunion.io,resources=ansibleplaybooks,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=onecloud.yunion.io,resources=ansibleplaybooks/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=onecloud.yunion.io,resources=virtualmachines,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=onecloud.yunion.io,resources=ansibleplaybooktemplates,verbs=get;list;watch;create;update;patch;delete

func (r *AnsiblePlaybookReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("ansibleplaybook", req.NamespacedName)

	var ansiblePlaybook onecloudv1.AnsiblePlaybook
	if err := r.Get(ctx, req.NamespacedName, &ansiblePlaybook); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	dealErr := func(err error) (ctrl.Result, error) {
		return dealErr(ctx, log, r, &ansiblePlaybook, provider.ResourceAP, err)
	}

	myFinalizerName := "virtualmachine.finalizers.onecloud.yunion.io"
	// add finalizer
	if ansiblePlaybook.ObjectMeta.DeletionTimestamp.IsZero() {
		if !utils.IsInStringArray(myFinalizerName, ansiblePlaybook.ObjectMeta.Finalizers) {
			ansiblePlaybook.ObjectMeta.Finalizers = append(ansiblePlaybook.ObjectMeta.Finalizers, myFinalizerName)
			if err := r.Update(ctx, &ansiblePlaybook); err != nil {
				return ctrl.Result{}, err
			}
			return ctrl.Result{}, nil
		}
	} else {
		if utils.IsInStringArray(myFinalizerName, ansiblePlaybook.ObjectMeta.Finalizers) {
			if len(ansiblePlaybook.Status.ExternalInfo.Id) == 0 {
				ansiblePlaybook.ObjectMeta.Finalizers = removeString(ansiblePlaybook.ObjectMeta.Finalizers,
					myFinalizerName)
				if err := r.Update(ctx, &ansiblePlaybook); err != nil {
					return ctrl.Result{}, err
				}
				return ctrl.Result{}, nil
			}
			ret, err := r.Delete(ctx, &ansiblePlaybook)
			if err != nil {
				return dealErr(err)
			}
			return ret, nil
		}
		return ctrl.Result{}, nil
	}

	if ansiblePlaybook.Status.Phase == onecloudv1.ResourceInvalid {
		return ctrl.Result{}, nil
	}

	if ansiblePlaybook.Status.Phase == onecloudv1.ResourceFinished {
		if ansiblePlaybook.Status.ExternalInfo.Id == "" {
			return ctrl.Result{}, nil
		}
		// sync info first
		apStatus, err := provider.Provider.APReconcile(ctx, &ansiblePlaybook)
		if err != nil {
			return dealErr(err)
		}
		if r.requireUpdate(&ansiblePlaybook, apStatus) {
			ansiblePlaybook.Status = *apStatus
			return ctrl.Result{}, r.Status().Update(ctx, &ansiblePlaybook)
		}
		// clear this
		if err := r.apClear(ctx, &ansiblePlaybook); err != nil {
			return dealErr(err)
		}
		return ctrl.Result{}, nil
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
				return r.markWaiting(ctx, log, &ansiblePlaybook,
					fmt.Sprintf("wait for AnsiblePlaybookTemplate '%s': %s", nameSpacedName, "no such resource"))
			}

		} else if ansiblePlaybook.Spec.PlaybookTemplate != nil {
			playbookTemplate = onecloudv1.AnsiblePlaybookTemplate{
				Spec: *ansiblePlaybook.Spec.PlaybookTemplate,
			}
		} else {
			log.V(0).Info("webhook will stop this but not")
			return ctrl.Result{}, nil
		}

		hosts := make([]provider.AnsiblePlaybookHost, 0, len(ansiblePlaybook.Spec.Inventory))
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
				return r.markWaiting(ctx, log, &ansiblePlaybook,
					fmt.Sprintf("wait for VirtualMachine '%s': %s", nameSpacedName, "no such resource"))
			}
			if vm.Status.Phase != onecloudv1.ResourceRunning {
				return r.markWaiting(ctx, log, &ansiblePlaybook,
					fmt.Sprintf("wait for VirtualMachine '%s': %s", nameSpacedName, "need phase 'Running'"))
			}
			if len(vm.Status.ExternalInfo.Eip) == 0 && len(vm.Status.ExternalInfo.Ips) == 0 {
				return r.markWaiting(ctx, log, &ansiblePlaybook,
					fmt.Sprintf("wait for VirtualMachine '%s': %s", nameSpacedName, "need eip or ips"))
			}
			// build vars
			noVars := make([]string, 0)
			vars := make(map[string]interface{}, len(host.Vars))
			for _, temVar := range playbookTemplate.Spec.Vars {
				if value, ok := host.Vars[temVar.Name]; ok {
					v, err := value.Value(ctx)
					if err != nil {
						// invalid
						log.Error(err, "StringValue.Value")
						ansiblePlaybook.SetResourcePhase(onecloudv1.ResourceInvalid,
							fmt.Sprintf("The value of var '%s' is valid: %s", temVar.Name, err.Error()),
						)
						return ctrl.Result{}, r.Status().Update(ctx, &ansiblePlaybook)
					}
					if v == nil {
						// need to wait
						return r.markWaiting(ctx, log, &ansiblePlaybook,
							fmt.Sprintf("wait for var '%s'", temVar.Name),
						)
					}
					vars[temVar.Name] = intOrString2interface(v)
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
					vars[temVar.Name] = intOrString2interface(temVar.Default)
					continue
				}
			}
			// set phase invalid
			if len(noVars) > 0 {
				ansiblePlaybook.SetResourcePhase(onecloudv1.ResourceInvalid, fmt.Sprintf(
					"Required these missed variables: %s for virtualMachine '%s'",
					strings.Join(noVars, ", "), host.VirtualMachine.Name),
				)
				return ctrl.Result{}, r.Status().Update(ctx, &ansiblePlaybook)
			}
			hosts = append(hosts, provider.AnsiblePlaybookHost{
				VM:   &vm,
				Vars: vars,
			})
		}

		// build common vars
		commonVars := make(map[string]interface{}, len(ansiblePlaybook.Spec.Vars))
		for varName, sv := range ansiblePlaybook.Spec.Vars {
			vv, err := sv.Value(ctx)
			if err != nil {
				// invalid
				log.Error(err, "StringValue.Value")
				ansiblePlaybook.SetResourcePhase(onecloudv1.ResourceInvalid,
					fmt.Sprintf("The value of var '%s' is valid: %s", varName, err.Error()),
				)
				return ctrl.Result{}, r.Status().Update(ctx, &ansiblePlaybook)
			}
			if vv == nil {
				// need to wait
				return r.markWaiting(ctx, log, &ansiblePlaybook,
					fmt.Sprintf("wait for var '%s'", varName),
				)
			}
			commonVars[varName] = intOrString2interface(vv)
		}

		// all other resources ready, create ansible playbook
		return ctrl.Result{}, r.apCreate(ctx, &ansiblePlaybook, hosts, &playbookTemplate, commonVars)
	}

	apStatus, err := provider.Provider.APGetStatus(ctx, &ansiblePlaybook)
	if err != nil {
		return dealErr(err)
	}
	if r.requireUpdate(&ansiblePlaybook, apStatus) {
		ansiblePlaybook.Status = *apStatus
		return ctrl.Result{}, r.Status().Update(ctx, &ansiblePlaybook)
	}

	if ansiblePlaybook.Status.Phase == onecloudv1.ResourceFailed {
		if err := r.apDelete(ctx, &ansiblePlaybook); err != nil {
			return dealErr(err)
		}
		return ctrl.Result{}, nil
	}

	// Pending
	if ansiblePlaybook.Status.Phase == onecloudv1.ResourcePending {
		return ctrl.Result{Requeue: true, RequeueAfter: PendingAfter}, nil
	}

	// Unkown
	if ansiblePlaybook.Status.Phase == onecloudv1.ResourceUnkown {
		return ctrl.Result{Requeue: true, RequeueAfter: 2 * time.Second}, nil
	}

	return ctrl.Result{}, nil
}

func intOrString2interface(is *intstr.IntOrString) interface{} {
	switch is.Type {
	case intstr.Int:
		return is.IntVal
	case intstr.String:
		return is.StrVal
	default:
		// not arrive
		return is
	}
}

func (r *AnsiblePlaybookReconciler) Delete(ctx context.Context, ap *onecloudv1.AnsiblePlaybook) (ctrl.Result, error) {
	apStatus, err := provider.Provider.APGetStatus(ctx, ap)
	if err != nil {
		return ctrl.Result{}, err
	}
	if r.requireUpdate(ap, apStatus) {
		ap.Status = *apStatus
		return ctrl.Result{}, r.Status().Update(ctx, ap)
	}
	// Pending
	if ap.Status.Phase == onecloudv1.ResourcePending {
		return ctrl.Result{Requeue: true, RequeueAfter: 5 * time.Second}, nil
	}

	// Unkown
	if ap.Status.Phase == onecloudv1.ResourceUnkown {
		return ctrl.Result{Requeue: true, RequeueAfter: 2 * time.Second}, nil
	}
	// delete this
	if err := r.apDelete(ctx, ap); err != nil {
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil
}

func (r *AnsiblePlaybookReconciler) apCreate(ctx context.Context, ap *onecloudv1.AnsiblePlaybook,
	hosts []provider.AnsiblePlaybookHost, apt *onecloudv1.AnsiblePlaybookTemplate, commonVars map[string]interface{}) error {
	// check if recreate items has reached the max limit
	maxRetryTimes := r.maxRetryTimes(ap)
	if ap.Status.TryTimes > maxRetryTimes {
		ap.Status.Phase = onecloudv1.ResourceInvalid
		ap.Status.Reason = fmt.Sprintf("The number of consecutive retry failures exceeds the maximum %d", maxRetryTimes)
		return r.Status().Update(ctx, ap)
	}
	extInfo, err := provider.Provider.APCreate(ctx, ap, hosts, apt, commonVars)
	if err != nil {
		return err
	}
	ap.Status.ExternalInfo.ExternalInfoBase = extInfo
	ap.Status.Phase = onecloudv1.ResourcePending
	ap.Status.TryTimes += 1
	return r.Status().Update(ctx, ap)
}

func (r *AnsiblePlaybookReconciler) apDelete(ctx context.Context, ap *onecloudv1.AnsiblePlaybook) error {
	isDelete, extInfo, err := provider.Provider.APDelete(ctx, ap)
	if err != nil {
		return err
	}
	if isDelete {
		ap.Status.Phase = onecloudv1.ResourcePending
	}
	ap.Status.ExternalInfo.ExternalInfoBase = extInfo
	return r.Status().Update(ctx, ap)
}

func (r *AnsiblePlaybookReconciler) apClear(ctx context.Context, ap *onecloudv1.AnsiblePlaybook) error {
	_, _, err := provider.Provider.APDelete(ctx, ap)
	return err
}

func (r *AnsiblePlaybookReconciler) maxRetryTimes(ap *onecloudv1.AnsiblePlaybook) int32 {
	if ap.Spec.MaxRetryTime == nil || *ap.Spec.MaxRetryTime <= 0 {
		return math.MaxInt32
	}
	return *ap.Spec.MaxRetryTime
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

func (r *AnsiblePlaybookReconciler) markWaiting(ctx context.Context, log logr.Logger, ap *onecloudv1.AnsiblePlaybook, msg string) (ctrl.Result, error) {
	newStatus := ap.Status.DeepCopy()
	newStatus.Phase = onecloudv1.ResourceWaiting
	newStatus.Reason = msg
	if !r.requireUpdate(ap, newStatus) {
		return ctrl.Result{Requeue: true, RequeueAfter: WaitingAfter}, nil
	}
	ap.Status = *newStatus
	if err := r.Status().Update(ctx, ap); err != nil {
		log.Error(err, "unablt to update ansibleplaybook")
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil
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
