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
	"time"

	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	onecloudv1 "yunion.io/x/onecloud-service-operator/api/v1"
	"yunion.io/x/onecloud-service-operator/pkg/options"
	"yunion.io/x/onecloud-service-operator/pkg/resources"
	"yunion.io/x/onecloud-service-operator/pkg/util"
)

// VirtualMachineReconciler reconciles a VirtualMachine object
type VirtualMachineReconciler struct {
	ReconcilerBase
}

// +kubebuilder:rbac:groups=onecloud.yunion.io,resources=virtualmachines,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=onecloud.yunion.io,resources=virtualmachines/status,verbs=get;update;patch

func (r *VirtualMachineReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()

	var virtualMachine onecloudv1.VirtualMachine
	if err := r.Get(ctx, req.NamespacedName, &virtualMachine); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	log := r.GetLog(&virtualMachine)
	remoteVm := resources.NewVirtualMachine(&virtualMachine)

	dealErr := func(err error) (ctrl.Result, error) {
		return r.dealErr(ctx, remoteVm, err)
	}

	var (
		vmPendingAfter = time.Duration(options.Options.VirtualMachineConfig.IntervalPending) * time.Minute
	)

	has, ret, err := r.UseFinallizer(ctx, remoteVm)
	if !has {
		return ret, err
	}

	// Invalid resource
	if virtualMachine.Status.Phase == onecloudv1.ResourceInvalid {
		return ctrl.Result{}, nil
	}

	// That virtualMachine.RemoteStatus.VmId field is empty, indicating that there is no corresponding VM,
	// and we need to create a new one
	if len(virtualMachine.Status.ExternalInfo.Id) == 0 {
		var in interface{}
		return r.Create(ctx, remoteVm, in, true)
	}

	// VirutalMachine.RemoteStatus.VmId is not empty, sync status
	update, ret, err := r.GetStatus(ctx, remoteVm)
	if update {
		return ret, err
	}

	// Running
	if virtualMachine.Status.Phase == onecloudv1.ResourceRunning || virtualMachine.Status.Phase == onecloudv1.ResourceReady {
		vmStatus, specPhase, err := r.reconcile(ctx, log, remoteVm)
		if err != nil {
			return dealErr(err)
		}
		if specPhase != virtualMachine.Status.Phase {
			var extInfoBase onecloudv1.ExternalInfoBase
			switch specPhase {
			case onecloudv1.ResourceRunning:
				//start
				extInfoBase, err = remoteVm.Start(ctx)
				if err != nil {
					return dealErr(err)
				}
			case onecloudv1.ResourceReady:
				//stop
				extInfoBase, err = remoteVm.Stop(ctx)
				if err != nil {
					return dealErr(err)
				}
			}
			vmStatus.Phase = onecloudv1.ResourcePending
			vmStatus.ExternalInfo.ExternalInfoBase = extInfoBase
			vmStatus.Reason = "Try reach the corresponding phase before performing the operation"
		}
		if r.requireUpdate(&virtualMachine, vmStatus) {
			virtualMachine.Status = *vmStatus
			return ctrl.Result{}, r.Status().Update(ctx, &virtualMachine)
		}
		return ctrl.Result{}, nil
	}

	// Failed
	if virtualMachine.Status.Phase == onecloudv1.ResourceFailed {
		// before delete, log the status of vm
		log.V(-1).Info(fmt.Sprintf("vm's externalInfoBase: %#v", virtualMachine.Status.ExternalInfo.ExternalInfoBase))
		return r.Delete(ctx, remoteVm)
	}

	// Pending
	if virtualMachine.Status.Phase == onecloudv1.ResourcePending {
		return ctrl.Result{Requeue: true, RequeueAfter: vmPendingAfter * time.Second}, nil
	}

	// Unkown
	if virtualMachine.Status.Phase == onecloudv1.ResourceUnkown {
		return ctrl.Result{Requeue: true, RequeueAfter: 2 * time.Second}, nil
	}

	return ctrl.Result{}, nil
}

func (r *VirtualMachineReconciler) reconcile(ctx context.Context, log logr.Logger, remoteVm resources.VirtualMachine) (vmStatus *onecloudv1.VirtualMachineStatus, specPhase onecloudv1.ResourcePhase, err error) {
	specPhase = onecloudv1.ResourceRunning
	oper, vmInfo, err := remoteVm.Reconcile(ctx, log)
	if err != nil {
		return
	}
	vm := remoteVm.VirtualMachine
	vmStatus = vm.Status.DeepCopy()
	vmStatus.ExternalInfo = *vmInfo
	// nothing to do
	if oper == nil {
		return
	}

	// execute operation
	if len(oper.PrePhase) > 0 {
		specPhase = oper.PrePhase[0]
		if vm.Status.Phase != specPhase {
			return
		}
	}
	extInfo, err := oper.Operator(ctx)
	if err != nil {
		return
	}
	vmStatus.ExternalInfo.ExternalInfoBase = extInfo
	vmStatus.Phase = onecloudv1.ResourcePending
	vmStatus.Reason = oper.OperDesc.String()
	return
}

func (r *VirtualMachineReconciler) requireUpdate(vm *onecloudv1.VirtualMachine,
	newStatus *onecloudv1.VirtualMachineStatus) bool {
	if newStatus == nil {
		return false
	}
	if vm.Status.Phase != newStatus.Phase || vm.Status.TryTimes != newStatus.TryTimes {
		return true
	}
	if vm.Status.ExternalInfo.Eip != newStatus.ExternalInfo.Eip || !util.EqualStringSlices(vm.Status.ExternalInfo.Ips,
		newStatus.ExternalInfo.Ips) {
		return true
	}
	return false
}

func (r *VirtualMachineReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&onecloudv1.VirtualMachine{}).
		Complete(r)
}

func removeString(slice []string, s string) (result []string) {
	for _, item := range slice {
		if item == s {
			continue
		}
		result = append(result, item)
	}
	return
}
