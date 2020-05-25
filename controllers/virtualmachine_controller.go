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
	"time"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"yunion.io/x/pkg/utils"

	onecloudv1 "yunion.io/x/onecloud-resource-operator/api/v1"
	"yunion.io/x/onecloud-resource-operator/pkg/options"
	"yunion.io/x/onecloud-resource-operator/pkg/provider"
	"yunion.io/x/onecloud-resource-operator/pkg/util"
)

var (
	vmPendingAfter = time.Duration(options.Options.VirtualMachineConfig.IntervalPending) * time.Minute
)

// VirtualMachineReconciler reconciles a VirtualMachine object
type VirtualMachineReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=onecloud.yunion.io,resources=virtualmachines,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=onecloud.yunion.io,resources=virtualmachines/status,verbs=get;update;patch

func (r *VirtualMachineReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("virtualmachine", req.NamespacedName)

	var virtualMachine onecloudv1.VirtualMachine
	if err := r.Get(ctx, req.NamespacedName, &virtualMachine); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	dealErr := func(err error) (ctrl.Result, error) {
		return dealErr(ctx, log, r, &virtualMachine, provider.ResourceVM, err)
	}

	myFinalizerName := "virtualmachine.finalizers.onecloud.yunion.io"
	// add finalizer
	if virtualMachine.ObjectMeta.DeletionTimestamp.IsZero() {
		if !utils.IsInStringArray(myFinalizerName, virtualMachine.ObjectMeta.Finalizers) {
			virtualMachine.ObjectMeta.Finalizers = append(virtualMachine.ObjectMeta.Finalizers, myFinalizerName)
			if err := r.Update(ctx, &virtualMachine); err != nil {
				return ctrl.Result{}, err
			}
			return ctrl.Result{}, nil
		}
	} else {
		if utils.IsInStringArray(myFinalizerName, virtualMachine.ObjectMeta.Finalizers) {
			if len(virtualMachine.Status.ExternalInfo.Id) == 0 {
				virtualMachine.ObjectMeta.Finalizers = removeString(virtualMachine.ObjectMeta.Finalizers,
					myFinalizerName)
				if err := r.Update(ctx, &virtualMachine); err != nil {
					return ctrl.Result{}, err
				}
				return ctrl.Result{}, nil
			}
			ret, err := r.Delete(ctx, &virtualMachine)
			if err != nil {
				return dealErr(err)
			}
			return ret, nil
		}
		return ctrl.Result{}, nil
	}

	// Invalid resource
	if virtualMachine.Status.Phase == onecloudv1.ResourceInvalid {
		return ctrl.Result{}, nil
	}

	// That virtualMachine.RemoteStatus.VmId field is empty, indicating that there is no corresponding VM,
	// and we need to create a new one
	if len(virtualMachine.Status.ExternalInfo.Id) == 0 {
		err := r.vmCreate(ctx, &virtualMachine)
		if err != nil {
			return dealErr(err)
		}
		return ctrl.Result{}, nil
	}

	// VirutalMachine.RemoteStatus.VmId is not empty, sync status
	vmStatus, err := provider.Provider.VMGetStatus(ctx, &virtualMachine)
	if err != nil {
		return dealErr(err)
	}
	if r.requireUpdate(&virtualMachine, vmStatus) {
		virtualMachine.Status = *vmStatus
		return ctrl.Result{}, r.Status().Update(ctx, &virtualMachine)
	}

	// Running
	if virtualMachine.Status.Phase == onecloudv1.ResourceRunning || virtualMachine.Status.Phase == onecloudv1.ResourceReady {
		vmStatus, specPhase, err := r.vmReconcile(ctx, log, &virtualMachine)
		if err != nil {
			return dealErr(err)
		}
		if specPhase != virtualMachine.Status.Phase {
			var extInfoBase onecloudv1.ExternalInfoBase
			switch specPhase {
			case onecloudv1.ResourceRunning:
				//start
				extInfoBase, err = provider.Provider.VMStart(ctx, &virtualMachine)
				if err != nil {
					return dealErr(err)
				}
			case onecloudv1.ResourceReady:
				//stop
				extInfoBase, err = provider.Provider.VMStop(ctx, &virtualMachine)
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
		if err := r.vmDelete(ctx, &virtualMachine); err != nil {
			return dealErr(err)
		}
		return ctrl.Result{}, nil
	}

	// Pending
	if virtualMachine.Status.Phase == onecloudv1.ResourcePending {
		return ctrl.Result{Requeue: true, RequeueAfter: 5 * time.Second}, nil
	}

	// Unkown
	if virtualMachine.Status.Phase == onecloudv1.ResourceUnkown {
		return ctrl.Result{Requeue: true, RequeueAfter: 2 * time.Second}, nil
	}

	return ctrl.Result{}, nil
}

func (r *VirtualMachineReconciler) Delete(ctx context.Context, vm *onecloudv1.VirtualMachine) (ctrl.Result, error) {
	// sync status first
	vmStatus, err := provider.Provider.VMGetStatus(ctx, vm)
	if err != nil {
		return ctrl.Result{}, err
	}
	if r.requireUpdate(vm, vmStatus) {
		vm.Status = *vmStatus
		return ctrl.Result{}, r.Status().Update(ctx, vm)
	}
	// Pending
	if vm.Status.Phase == onecloudv1.ResourcePending {
		return ctrl.Result{Requeue: true, RequeueAfter: 5 * time.Second}, nil
	}
	// Unkown
	if vm.Status.Phase == onecloudv1.ResourceUnkown {
		return ctrl.Result{Requeue: true, RequeueAfter: 2 * time.Second}, nil
	}
	// Delete VM
	if err := r.vmDelete(ctx, vm); err != nil {
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil
}

func (r *VirtualMachineReconciler) vmReconcile(ctx context.Context, log logr.Logger,
	vm *onecloudv1.VirtualMachine) (vmStatus *onecloudv1.VirtualMachineStatus, specPhase onecloudv1.ResourcePhase, err error) {
	specPhase = onecloudv1.ResourceRunning
	oper, vmInfo, err := provider.Provider.VMReconcile(ctx, log, vm)
	if err != nil {
		return
	}
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
	if vm.Status.Phase != newStatus.Phase || vm.Status.CreateTimes != newStatus.CreateTimes {
		return true
	}
	if vm.Status.ExternalInfo.Eip != newStatus.ExternalInfo.Eip || !util.EqualStringSlices(vm.Status.ExternalInfo.Ips,
		newStatus.ExternalInfo.Ips) {
		return true
	}
	return false
}

func (r *VirtualMachineReconciler) vmCreate(ctx context.Context, vm *onecloudv1.VirtualMachine) error {
	// check if recreate times has reached the max limit
	recreateMaxTimes := r.maxRecreateTimes(vm)
	if vm.Status.CreateTimes-1 == recreateMaxTimes {
		vm.Status.Phase = onecloudv1.ResourceInvalid
		vm.Status.Reason = fmt.Sprintf("The number of consecutive retry creation failures exceeds the maximum %d",
			recreateMaxTimes)
		return r.Status().Update(ctx, vm)
	}
	extInfo, err := provider.Provider.VMCreate(ctx, vm)
	if err != nil {
		return err
	}
	vm.Status.ExternalInfo.ExternalInfoBase = extInfo
	vm.Status.Phase = onecloudv1.ResourcePending
	vm.Status.CreateTimes += 1
	return r.Status().Update(ctx, vm)
}

func (r *VirtualMachineReconciler) vmDelete(ctx context.Context, vm *onecloudv1.VirtualMachine) error {
	extInfo, err := provider.Provider.VMDelete(ctx, vm)
	if err != nil {
		return err
	}
	vm.Status.Phase = onecloudv1.ResourcePending
	vm.Status.ExternalInfo.ExternalInfoBase = extInfo

	return r.Status().Update(ctx, vm)
}

func (r *VirtualMachineReconciler) maxRecreateTimes(vm *onecloudv1.VirtualMachine) int32 {
	if vm.Spec.RecreatePolicy != nil {
		return vm.Spec.RecreatePolicy.MaxTimes
	}
	return provider.Provider.VMDefaultRecreatePolicy().MaxTimes
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
