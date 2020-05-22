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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AnsiblePlaybookTemplateSpec defines the desired state of AnsiblePlaybookTemplate.
type AnsiblePlaybookTemplateSpec struct {
	// Playbook describe the main content of absible playbook which should be in yaml format.
	Playbook string `json:"playbook"`
	// Requirements describe the source of roles dependent on Playbook
	// +optional
	Requirements string `json:"requirements,omitempty"`
	// Files describe the associated file tree and file content which should be in json format.
	// +optional
	Files string `json:"files,omitempty"`
	// Vars describe the vars to apply this ansible playbook.
	// +optional
	Vars []AnsiblePlaybookTemplateVar `json:"vars,omitempty"`
}

// AnsiblePlaybookTemplateStatus defines the observed state of AnsiblePlaybookTemplate
type AnsiblePlaybookTemplateStatus struct {
}

type AnsiblePlaybookTemplateVar struct {
	Name string `json:"name"`
	// Required indicates whether this variable is required.
	// +optional
	Required *bool `json:"required,omitempty"`
	// Default describe the default value of this variable.
	// If it is empty, Required should be true.
	// +optional
	Default *IntOrString `json:"default,omitempty"`
}

// +kubebuilder:object:root=true

// AnsiblePlaybookTemplate is the Schema for the ansibleplaybooktemplates API
type AnsiblePlaybookTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AnsiblePlaybookTemplateSpec   `json:"spec,omitempty"`
	Status AnsiblePlaybookTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AnsiblePlaybookTemplateList contains a list of AnsiblePlaybookTemplate
type AnsiblePlaybookTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AnsiblePlaybookTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AnsiblePlaybookTemplate{}, &AnsiblePlaybookTemplateList{})
}
