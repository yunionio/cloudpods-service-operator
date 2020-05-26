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

package provider

import (
	comapi "yunion.io/x/onecloud/pkg/apis/compute"

	onecloudv1 "yunion.io/x/onecloud-service-operator/api/v1"
)

func (OnecloudProvider) convertInt64Ptr(n *int64) int {
	if n == nil {
		return 0
	}
	return int(*n)
}

func (OnecloudProvider) convertInt32Ptr(n *int32) int {
	if n == nil {
		return 0
	}
	return int(*n)
}

func (OnecloudProvider) convertBoolPrt(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

func (OnecloudProvider) ConvertVMDisk(config onecloudv1.VMDiskSpec) comapi.DiskConfig {
	return comapi.DiskConfig{
		ImageId: config.Image,
		SizeMb:  int(config.SizeGB) * 1024,
		Driver:  string(config.Driver),
		Storage: config.Storage.Storage,
		Backend: config.Storage.Backend,
		Medium:  string(config.Storage.Medium),
	}
}

func (oc OnecloudProvider) ConvertVMNetwork(config onecloudv1.VMNetworkSpec) comapi.NetworkConfig {
	return comapi.NetworkConfig{
		Network: config.Network,
		Address: config.Address,
	}
}

func (oc OnecloudProvider) ConvertVMConfig(config onecloudv1.VirtualMachineConfig) comapi.ServerConfigs {
	disks := make([]*comapi.DiskConfig, 1, len(config.DataDisks)+1)
	rootDisk := oc.ConvertVMDisk(config.RootDisk)
	rootDisk.DiskType = "sys"
	disks[0] = &rootDisk
	for i := range config.DataDisks {
		disk := oc.ConvertVMDisk(config.DataDisks[i])
		disk.DiskType = "data"
		disks = append(disks, &disk)
	}
	networks := make([]*comapi.NetworkConfig, 0, len(config.Networks))
	for i := range config.Networks {
		network := oc.ConvertVMNetwork(config.Networks[i])
		networks = append(networks, &network)
	}
	return comapi.ServerConfigs{
		PreferManager: config.PreferManger,
		PreferRegion:  config.PreferRegion,
		PreferZone:    config.PreferZone,
		PreferWire:    config.PreferWire,
		PreferHost:    config.PreferHost,
		Hypervisor:    config.Hypervisor,
		InstanceType:  config.InstanceType,
		Disks:         disks,
		Networks:      networks,
	}
}

func (oc OnecloudProvider) ConvertVM(config onecloudv1.VirtualMachineSpec) comapi.ServerCreateInput {
	serverConfig := oc.ConvertVMConfig(config.VmConfig)
	createInput := comapi.ServerCreateInput{
		ServerConfigs: &serverConfig,
		KeypairId:     config.KeyPairId,
		Password:      config.Password,
		ResetPassword: config.ResetPassword,
		Duration:      config.BillDuration,
		AutoRenew:     oc.convertBoolPrt(config.AutoRenew),
	}

	createInput.VcpuCount = oc.convertInt64Ptr(config.VmConfig.VcpuCount)
	createInput.VmemSize = oc.convertInt64Ptr(config.VmConfig.VmemSizeGB) * 1024
	createInput.Description = config.Desciption
	createInput.Project = config.Project.Project
	createInput.ProjectDomain = config.Project.PoejectDomain
	if config.NameCheck == nil || *config.NameCheck {
		createInput.Name = config.Name
	} else {
		createInput.GenerateName = config.Name
	}
	createInput.AutoStart = true

	if config.NewEip != nil {
		createInput.EipBw = oc.convertInt64Ptr(config.NewEip.Bw)
		createInput.EipChargeType = config.NewEip.ChargeType
		createInput.EipAutoDellocate = true
	} else {
		createInput.Eip = config.Eip
	}
	return createInput
}
