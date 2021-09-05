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
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AnsiblePlaybookSpec defines the desired state of AnsiblePlaybook
type AnsiblePlaybookSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// PlaybookTemplateRef specifies the AnsiblePlaybookTemplate.
	// +optional
	PlaybookTemplateRef *LocalObjectReference `json:"playbookTemplateRef,omitempty"`

	// PlaybookTemplate describe the ansible playbook
	// +optional
	PlaybookTemplate *AnsiblePlaybookTemplateSpec `json:"playbookTemplate,omitempty"`

	// VirtualMachines specifies the inventory of ansible playbook.
	Inventory []AnsiblePlaybookHost `json:"inventory"`

	// Vars describe the public value about Vars in AnsiblePlaybookTemplate.
	// +optional
	Vars map[string]IntOrStringOrYamlStore `json:"vars,omitempty"`

	ProxyVars []AnsiblePlaybookProxyVar `json:"proxyVars,omitempty"`

	ResourceSpecBase `json:",inline"`
}

type AnsiblePlaybookProxyVar struct {
	Name    string `json:"name,omitempty"`
	Service string `json:"service,omitempty"`
}

// AnsiblePlaybookStatus defines the observed state of AnsiblePlaybook
type AnsiblePlaybookStatus struct {
	// Important: Run "make" to regenerate code after modifying this file
	ResourceStatusBase `json:",inline"`
	// +optional
	ExternalInfo AnsiblePlaybookInfo `json:"externalInfo,omitempty"`

	DevtoolSshInfos []DevtoolSshInfo    `json:"devtoolSshInfos,omitempty"`
	ServiceUrls     []DevtoolServiceUrl `json:"serviceUrls,omitempty"`
}

type DevtoolSshInfo struct {
	VMName     string `json:"vmName,omitempty"`
	Id         string `json:"id,omitempty"`
	User       string `json:"user,omitempty"`
	Host       string `json:"host,omitempty"`
	Port       int    `json:"port,omitempty"`
	ServerName string `json:"serverName,omitempty"`
}

type DevtoolServiceUrl struct {
	Id      string `json:"id,omitempty"`
	Service string `json:"service,omitempty"`
	Url     string `json:"url,omitempty"`
}

type AnsiblePlaybookHost struct {
	VirtualMachine ObjectReference `json:"virtualMachine"`
	// Vars describes the unique values ​​of the VirtualMachine
	// corresponding to the variables in the AnsiblePlaybookTemplate.
	// +optional
	Vars map[string]IntOrStringOrYamlStore `json:"vars,omitempty"`
}

type AnsiblePlaybookInfo struct {
	ExternalInfoBase `json:",inline"`
	// +optional
	Output string `json:"output,omitempty"`
}

// +kubebuilder:object:root=true

// AnsiblePlaybook is the Schema for the ansibleplaybooks API
// +kubebuilder:subresource:status
type AnsiblePlaybook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AnsiblePlaybookSpec   `json:"spec,omitempty"`
	Status AnsiblePlaybookStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AnsiblePlaybookList contains a list of AnsiblePlaybook
type AnsiblePlaybookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AnsiblePlaybook `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AnsiblePlaybook{}, &AnsiblePlaybookList{})
}
