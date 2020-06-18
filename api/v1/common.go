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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

type DisplayIdenSpec struct {
	// Name
	// +optional
	Name string `json:"name,omitempty"`
	// NameCheck indicates whether to strictly check Name
	// +optional
	NameCheck *bool `json:"nameCheck,omitempty"`
	// +optional
	Desciption string `json:"description,omitempty"`
}

// ExternalInfoBase describe the corresponding resource's info in external system
type ExternalInfoBase struct {
	// +optional
	Id string `json:"id,omitempty"`
	// +optional
	Status string `json:"status,omitempty"`
	// Action indicate the latest action for external vm.
	// +optional
	Action string `json:"action,omitempty"`
}

// ResourcePhase is a label for the condition of a resource at the current time
type ResourcePhase string

const (
	// ReourcePending means the external resource in an unstable intermediate state.
	ResourcePending ResourcePhase = "Pending"
	// ResourceReady means the external resource in an normal state.
	ResourceReady ResourcePhase = "Ready"
	// ResourceRunning means the external resource has been enter a working state.
	ResourceRunning ResourcePhase = "Running"
	// ResourceFailed means the external resource has been unnormal and should be delete.
	ResourceFailed ResourcePhase = "Failed"
	// ResourceUnkown means the external system went wrong and controller should resync after some time.
	ResourceUnkown ResourcePhase = "Unkown"
	// ResourceInvalid means the resource is invalid that means user should edit the spec or recreate one.
	ResourceInvalid ResourcePhase = "Invalid"
	// ResourceWaiting means this resource is waiting for others
	ResourceWaiting ResourcePhase = "Waiting"
	// ResourceSucceeded means the resource completed its mission and finished.
	ResourceFinished ResourcePhase = "Finished"
)

type ObjectReference struct {
	// +optional
	Kind string `json:"kind,omitempty"`
	// +optional
	Namespace string `json:"namespace,omitempty"`
	// +optional
	Name string `json:"name,omitempty"`
	// +optional
	UID types.UID `json:"uid,omitempty"`
	// +optional
	APIVersion string `json:"apiVersion,omitempty"`
	// +optional
	ResourceVersion string `json:"resourceVersion,omitempty"`
}

// LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.
type LocalObjectReference struct {
	Name string `json:"name,omitempty"`
}

// +kubebuilder:object:generate=false
type IResource interface {
	runtime.Object
	metav1.Object
	GetResourceStatus() IResourceStatus
	GetResourceSpec() IResourceSpec
	SetResourceStatus(is IResourceStatus)
}

// +kubebuilder:object:generate=false
type IResourceStatus interface {
	DeepCopy2() IResourceStatus
	GetPhase() ResourcePhase
	SetPhase(phase ResourcePhase, reason string)
	GetTryTimes() int32
	SetTryTimes(i int32)
	GetBaseExternalInfo() ExternalInfoBase
	SetBaseExternalInfo(info ExternalInfoBase)
}

// +kubebuilder:object:generate=false
type IResourceSpec interface {
	GetMaxRetryTimes() int32
}

type ResourceSpecBase struct {
	// Nil or Non-positive number means unlimited.
	// +optional
	MaxRetryTimes *int32 `json:"maxRetryTimes,omitempty"`
}

func (rs *ResourceSpecBase) GetMaxRetryTimes() int32 {
	if rs.MaxRetryTimes == nil {
		return 5
	}
	return *rs.MaxRetryTimes
}

type ResourceStatusBase struct {
	// +optional
	Phase ResourcePhase `json:"phase,omitempty"`
	// A human readable message indicating details about why resource is in this phase.
	// +optional
	Reason string `json:"reason,omitempty"`
	// TryTimes record the continuous try times.
	TryTimes int32 `json:"tryTimes"`
}

func (rb *ResourceStatusBase) GetPhase() ResourcePhase {
	return rb.Phase
}

func (rb *ResourceStatusBase) SetPhase(phase ResourcePhase, reason string) {
	rb.Phase = phase
	rb.Reason = reason
}

func (rb *ResourceStatusBase) GetTryTimes() int32 {
	return rb.TryTimes
}

func (rb *ResourceStatusBase) SetTryTimes(i int32) {
	rb.TryTimes = i
}

func (vmStatus *VirtualMachineStatus) GetBaseExternalInfo() ExternalInfoBase {
	return vmStatus.ExternalInfo.ExternalInfoBase
}

func (vmStatus *VirtualMachineStatus) SetBaseExternalInfo(info ExternalInfoBase) {
	vmStatus.ExternalInfo.ExternalInfoBase = info
}

func (vmStatus *VirtualMachineStatus) DeepCopy2() IResourceStatus {
	return vmStatus.DeepCopy()
}

func (apStatus *AnsiblePlaybookStatus) GetBaseExternalInfo() ExternalInfoBase {
	return apStatus.ExternalInfo.ExternalInfoBase
}

func (apStatus *AnsiblePlaybookStatus) DeepCopy2() IResourceStatus {
	return apStatus.DeepCopy()
}

func (apStatus *AnsiblePlaybookStatus) SetBaseExternalInfo(info ExternalInfoBase) {
	apStatus.ExternalInfo.ExternalInfoBase = info
}

func (epStatus *EndpointStatus) GetBaseExternalInfo() ExternalInfoBase {
	return epStatus.ExternalInfo
}

func (epStatus *EndpointStatus) DeepCopy2() IResourceStatus {
	return epStatus.DeepCopy()
}

func (epStatus *EndpointStatus) SetBaseExternalInfo(info ExternalInfoBase) {
	epStatus.ExternalInfo = info
}

func (vm *VirtualMachine) GetResourceStatus() IResourceStatus {
	return &vm.Status
}

func (vm *VirtualMachine) GetResourceSpec() IResourceSpec {
	return &vm.Spec
}

func (vm *VirtualMachine) SetResourceStatus(is IResourceStatus) {
	vm.Status = *is.(*VirtualMachineStatus)
}

func (ap *AnsiblePlaybook) GetResourceStatus() IResourceStatus {
	return &ap.Status
}

func (ap *AnsiblePlaybook) SetResourceStatus(is IResourceStatus) {
	ap.Status = *is.(*AnsiblePlaybookStatus)
}

func (ap *AnsiblePlaybook) GetResourceSpec() IResourceSpec {
	return &ap.Spec
}

func (ep *Endpoint) GetResourceStatus() IResourceStatus {
	return &ep.Status
}

func (ep *Endpoint) SetResourceStatus(is IResourceStatus) {
	ep.Status = *is.(*EndpointStatus)
}

func (ep *Endpoint) GetResourceSpec() IResourceSpec {
	return &ep.Spec
}
