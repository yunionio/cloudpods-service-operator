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

// VirtualMachineSpec defines the desired state of VirtualMachine
type VirtualMachineSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	DisplayIdenSpec `json:",inline"`

	VmConfig VirtualMachineConfig `json:"vmConfig"`

	Project VMProjectSpec `json:"projectConfig"`

	VMPasswordSpec `json:",inline"`

	// +optional
	// +kubebuilder:validation:MinItems=1
	Secgropus []string `json:"secgroups,omitempty"`

	// NewEip indicates that create a new EIP and bind it with VM.
	// It conflicts with Eip.
	// +optional
	NewEip *VMNewEipSpec `json:"newEip,omitempty"`
	// Eip indicates that bind the existing EIP with VM.
	// It conflicts with NewEip.
	// +optional
	Eip string `json:"eip,omitempty"`

	// BillDuration describes the duration of the annual and monthly billing type.
	// That length of BillDuration represents the billing type is 'prepaid'.
	// +optional
	BillDuration string `json:"billDuration,omitempty"`
	// AutoRenew indicates whether to automatically renewal.
	// +optional
	AutoRenew *bool `json:"autoRenew,omitempty"`

	// +optional
	RecreatePolicy   *RecreatePolicy `json:"recreatePolicy,omitempty"`
	ResourceSpecBase `json:",inline"`
}

// RecreatePolicy describe that when the virtual machine is abnormal, how to deal with it,
// specifically determine whether to delete and recreate.
type RecreatePolicy struct {
	MatchStatus []string `json:"matchStatus,omitempty"`
	// +optional
	Never *bool `json:"never,omitempty"`
	// +optional
	Allways *bool `json:"allways,omitempty"`
}

type VMNewEipSpec struct {
	// Bw indicates the bandwidth of the Elastic Public IP.
	// +kubebuilder:validation:Minimum=1
	Bw *int64 `json:"bw,omitempty"`
	// The charge type of Elastic Public IP
	// +optional
	ChargeType string `json:"chargeType,omitempty"`
}

type VMPreferSpec struct {
	// PreferManager is the parameter passed to the scheduler which makes
	// the virtual machine created in the preferred cloud provider.
	// +optional
	PreferManger string `json:"preferManager,omitempty"`
	// PreferRegion is the parameter passed to the scheduler which makes
	// the virtual machine created in the preferred cloud region.
	// +optional
	PreferRegion string `json:"preferRegion,omitempty"`
	// PreferZone is the parameter passed to the scheduler which makes
	// the virtual machine created in the preferred cloud zone.
	// +optional
	PreferZone string `json:"preferZone,omitempty"`
	// PreferWire is the parameter passed to the scheduler which makes
	// the machine created in the preferred wire.
	// +optional
	PreferWire string `json:"preferWire,omitempty"`
	// PreferHost is the parameter passed to the scheduler which makes
	// the machine created in the preferred host.
	// +optional
	PreferHost string `json:"preferHost,omitempty"`
}

type VMPasswordSpec struct {
	// +optional
	KeyPairId string `json:"keyPairId,omitempty"`
	// +optional
	Password string `json:"password,omitempty"`
	// +optional
	ResetPassword *bool `json:"resetPassword,omitempty"`
}

type VirtualMachineConfig struct {
	VMPreferSpec `json:",inline"`
	// +optional
	Hypervisor string `json:"hypervisor"`

	// VcpuCount represents the number of CPUs of the virtual machine.
	// It conflicts with InstanceType and it is It is required if InstanceType is not specified.
	// +kubebuilder:validation:Minimum=1
	// +optional
	VcpuCount *int64 `json:"vcpuCount,omitempty"`
	// VmemSizeGB represents the size of memory of the virtual machine.
	// It conflicts with InstanceType and it is It is required if InstanceType is not specified.
	// +optional
	VmemSizeGB *int64 `json:"vmemSizeGB,omitempty"`
	// InstanceType describes the specifications of the virtual machine,
	// which are predefined by the cloud provider.
	// It conflicts with VcpuCount and VmemSizeGB.
	// +optional
	InstanceType string `json:"instanceType,omitempty"`

	// RootDisk describes the configuration of the system disk
	RootDisk VMDiskSpec `json:"rootDisk"`

	// DataDisks describes the configuration of data disks
	// +kubebuilder:validation:MaxItems=7
	// +optional
	DataDisks []VMDiskSpec `json:"dataDisks,omitempty"`

	// +optional
	Networks []VMNetworkSpec `json:"networks,omitempty"`
}

type VMProjectSpec struct {
	Project string `json:"project"`
	// +optional
	PoejectDomain string `json:"projectDomain,omitempty"`
}

type VMDiskSpec struct {
	// The disk will be created from the image represented by ImageId.
	// +optional
	Image string `json:"image,omitempty"`

	// SizeGB represents the size(unit: GB) of disk.
	// +optional
	SizeGB int64 `json:"sizeGB"`

	// +optional
	Driver DiskDriver `json:"driver,omitempty"`

	// +optional
	Storage VMStorageSpec `json:"storageConfig,omitempty"`
}

type VMStorageSpec struct {
	// Storage represents specific storage
	// +optional
	Storage string `json:"storage,omitempty"`
	// Backend represents backend of storage
	// +optinal
	Backend string `json:"backend,omitempty"`
	// +optional
	Medium StorageMedium `json:"medium,omitempty"`
}

type VMNetworkSpec struct {
	Network string `json:"network"`

	// +optional
	Address string `json:"address,omitempty"`
}

// StorageMedium represents storage media type
// +kubebuilder:validation:Enum=rotate;ssd;hybrid
type StorageMedium string

const (
	// Mechanical disk
	StorageMediumRotate StorageMedium = "rotate"
	// Solid state disk
	StorageMediumSsd StorageMedium = "ssd"
	// Hybrid disk
	StorageMediumHybrid StorageMedium = "hybrid"
)

// Driver represents the drive method of the disk on the virtual machine.
// +kubebuilder:validation:Enum=virtio;ide;scsi;sata;pvscsi
type DiskDriver string

const (
	DiskDriverVirtio DiskDriver = "virtio"
	DiskDriverIde    DiskDriver = "ide"
	DiskDriverScsi   DiskDriver = "scsi"
	DiskDriverSata   DiskDriver = "sata"
	DiskDriverPvscsi DiskDriver = "pvscsi"
)

// VirtualMachineStatus defines the observed state of VirtualMachine
type VirtualMachineStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	ResourceStatusBase `json:",inline"`
	// +optional
	ExternalInfo VMInfo `json:"externalInfo,omitempty"`
}

type VMInfo struct {
	ExternalInfoBase `json:",inline"`
	// +optional
	Eip string `json:"eip,omitempty"`
	// +optional
	Ips []string `json:"ips,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualMachine is the Schema for the virtualmachines API
// +kubebuilder:subresource:status
type VirtualMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualMachineSpec   `json:"spec,omitempty"`
	Status VirtualMachineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualMachineList contains a list of VirtualMachine
type VirtualMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualMachine `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VirtualMachine{}, &VirtualMachineList{})
}
