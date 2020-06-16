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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// EndpointSpec defines the desired state of Endpoint
type EndpointSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Url of this Endpoint
	URL      URL    `json:"url"`
	RegionId string `json:"regionId"`
	// +optional
	Name string `json:"name,omitempty"`
	// +optional
	Disabled *bool `json:"disabled,omitempty"`
	// Service certificate id or name
	// +optional
	ServiceCertificate string `json:"serviceCertificate,omitempty"`
}

// URL is used to construct url string 'Protocol://Host:Port/Prefix'
type URL struct {
	Protocol string
	Host     StringStore
	Port     *int32
	Prefix   string
}

// EndpointStatus defines the observed state of Endpoint
type EndpointStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +optional
	Phase ResourcePhase `json:"phase,omitempty"`
	// A human readable message indicating details about why endpoint is in this phase.
	// +optional
	Reason string `json:"reason,omitempty"`
	// +optional
	ExternalInfo ExternalInfoBase `json:"externalInfo,omitempty"`
	// CreateTimes record the continuous creation times.
	CreateTimes int32 `json:"createTimes"`
}

// +kubebuilder:object:root=true

// Endpoint is the Schema for the endpoints API
type Endpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EndpointSpec   `json:"spec,omitempty"`
	Status EndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EndpointList contains a list of Endpoint
type EndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Endpoint `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Endpoint{}, &EndpointList{})
}
