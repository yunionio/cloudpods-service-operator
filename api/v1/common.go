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
	SetResourcePhase(phase ResourcePhase, reason string)
	SetExternalId(id string)
}

func (vm *VirtualMachine) SetResourcePhase(phase ResourcePhase, reason string) {
	vm.Status.Phase = phase
	vm.Status.Reason = reason
}

func (vm *VirtualMachine) SetExternalId(id string) {
	vm.Status.ExternalInfo.Id = id
}

func (ap *AnsiblePlaybook) SetResourcePhase(phase ResourcePhase, reason string) {
	ap.Status.Phase = phase
	ap.Status.Reason = reason
}

func (ap *AnsiblePlaybook) SetExternalId(id string) {
	ap.Status.ExternalInfo.Id = id
}
