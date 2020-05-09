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

package v1

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var virtualmachinelog = logf.Log.WithName("virtualmachine-resource")

func (r *VirtualMachine) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-onecloud-yunion-io-v1-virtualmachine,mutating=true,failurePolicy=fail,groups=onecloud.yunion.io,resources=virtualmachines,verbs=create;update,versions=v1,name=mvirtualmachine.kb.io

var _ webhook.Defaulter = &VirtualMachine{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *VirtualMachine) Default() {
	virtualmachinelog.Info("default", "name", r.Name)
	var (
		cOne int64 = 1
		c30G int64 = 30
	)
	if len(r.Spec.Name) == 0 {
		r.Spec.Name = r.ObjectMeta.Name
	}
	if len(r.Spec.VmConfig.InstanceType) == 0 {
		if r.Spec.VmConfig.VcpuCount == nil {
			r.Spec.VmConfig.VcpuCount = &cOne
		}
		if r.Spec.VmConfig.VmemSizeGB == nil {
			r.Spec.VmConfig.VmemSizeGB = &c30G
		}
	}
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// +kubebuilder:webhook:verbs=create;update,path=/validate-onecloud-yunion-io-v1-virtualmachine,mutating=false,failurePolicy=fail,groups=onecloud.yunion.io,resources=virtualmachines,versions=v1,name=vvirtualmachine.kb.io

var _ webhook.Validator = &VirtualMachine{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *VirtualMachine) ValidateCreate() error {
	virtualmachinelog.Info("validate create", "name", r.Name)
	if len(r.Spec.VmConfig.RootDisk.Image) == 0 {
		return fmt.Errorf("The RootDisk's 'Image' should not be empty")
	}
	return r.validte()
}

func (r *VirtualMachine) validte() error {
	if (r.Spec.VmConfig.VcpuCount != nil || r.Spec.VmConfig.VmemSizeGB != nil) && len(r.Spec.VmConfig.InstanceType) > 0 {
		return fmt.Errorf("'VmConfig.VcpuCount' or 'VmConfig.VmemSizeGB' conflict with 'VmConfig.InstanceType'")
	}
	if len(r.Spec.Eip) > 0 && r.Spec.NewEip != nil {
		return fmt.Errorf("'NewEip' and 'Eip' conflicts with each other")
	}
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *VirtualMachine) ValidateUpdate(old runtime.Object) error {
	virtualmachinelog.Info("validate update", "name", r.Name)
	if err := r.validte(); err != nil {
		return err
	}

	oldR := old.(*VirtualMachine)
	// ProjectSpec Check
	if r.Spec.Project != oldR.Spec.Project {
		return fmt.Errorf("Can not change 'project' in Spec")
	}

	// VmConfig Check
	vmConfig, oldVmConfig := r.Spec.VmConfig, oldR.Spec.VmConfig
	if vmConfig.VMPreferSpec != oldVmConfig.VMPreferSpec {
		return fmt.Errorf("Can not change Prefer config in Spec")
	}
	if vmConfig.Hypervisor != oldVmConfig.Hypervisor {
		return fmt.Errorf("Can not change 'vmConfig.hypervisor' in Spec")
	}
	if err := checkVMDiskSpec(vmConfig.RootDisk, oldVmConfig.RootDisk); err != nil {
		return err
	}
	if len(vmConfig.DataDisks) < len(oldVmConfig.DataDisks) {
		return fmt.Errorf("Can not reduce the length of dataDisks in Spec")
	}
	for i := 0; i < len(oldVmConfig.DataDisks); i++ {
		if err := checkVMDiskSpec(vmConfig.DataDisks[i], oldVmConfig.DataDisks[i]); err != nil {
			return err
		}
	}
	if !networkEqual(vmConfig.Networks, oldVmConfig.Networks) {
		return fmt.Errorf("Can not change 'vmConfig.networks' in Spec")
	}

	// PasswordSpec Check
	if r.Spec.VMPasswordSpec != oldR.Spec.VMPasswordSpec {
		return fmt.Errorf("Can not change Password config in Spec")
	}

	if len(oldR.Spec.Eip) == 0 && len(r.Spec.Eip) > 0 {
		return fmt.Errorf("Can not change Eip when choosing 'Creating new EIP to bind' at the start")
	}
	if oldR.Spec.NewEip == nil && r.Spec.NewEip != nil {
		return fmt.Errorf("Can not change NewEip when choosing 'Binding existing EIP' at the start")
	}
	if oldR.Spec.NewEip != nil && r.Spec.NewEip.ChargeType != r.Spec.NewEip.ChargeType {
		return fmt.Errorf("Can not change 'newEip.eipChargeType")
	}

	if r.Spec.BillDuration != oldR.Spec.BillDuration {
		return fmt.Errorf("Can not change 'billDuration'")
	}
	if r.Spec.AutoRenew != oldR.Spec.AutoRenew {
		return fmt.Errorf("Can not change 'autoRenew'")
	}

	return nil
}

func checkVMDiskSpec(ds1, ds2 VMDiskSpec) error {
	if ds1.Image != ds2.Image {
		return fmt.Errorf("Can not change 'image' of Disk in Spec")
	}
	if ds1.Driver != ds2.Driver {
		return fmt.Errorf("Can not change 'driver' of Disk in Spec")
	}
	if ds1.Storage != ds2.Storage {
		return fmt.Errorf("Can not change 'storage' of Disk in Spec")
	}
	if ds1.SizeGB < ds2.SizeGB {
		return fmt.Errorf("Can not reduce 'sizeMb' of Disk in Spec")
	}
	return nil
}

func networkEqual(ns1, ns2 []VMNetworkSpec) bool {
	if len(ns1) != len(ns2) {
		return false
	}
	for i := 0; i < len(ns1); i++ {
		if ns1[i] != ns2[i] {
			return false
		}
	}
	return true
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *VirtualMachine) ValidateDelete() error {
	virtualmachinelog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
