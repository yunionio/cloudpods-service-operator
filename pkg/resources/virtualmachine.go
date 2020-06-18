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

package resources

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-logr/logr"

	"yunion.io/x/jsonutils"
	comapi "yunion.io/x/onecloud/pkg/apis/compute"
	"yunion.io/x/onecloud/pkg/mcclient/modules"
	onecloudutils "yunion.io/x/pkg/utils"

	onecloudv1 "yunion.io/x/onecloud-service-operator/api/v1"
	"yunion.io/x/onecloud-service-operator/pkg/util"
)

type VirtualMachine struct {
	VirtualMachine *onecloudv1.VirtualMachine
}

func NewVirtualMachine(vm *onecloudv1.VirtualMachine) VirtualMachine {
	return VirtualMachine{vm}
}

func (vi VirtualMachine) GetIResource() onecloudv1.IResource {
	return vi.VirtualMachine
}

func (vi VirtualMachine) GetResourceName() Resource {
	return ResourceVM
}

func (vi VirtualMachine) Create(ctx context.Context, _ interface{}) (onecloudv1.ExternalInfoBase, error) {
	vm := vi.VirtualMachine
	serverCreateInput := ConvertVM(vm.Spec)
	if len(serverCreateInput.Name) == 0 && len(serverCreateInput.GenerateName) == 0 {
		serverCreateInput.GenerateName = vm.ObjectMeta.Name
	}
	params := serverCreateInput.JSON(serverCreateInput)
	_, s, e := RequestVM.Operation(OperCreate).Apply(ctx, "", params)
	return s, e
}

func (vi VirtualMachine) Delete(ctx context.Context) (onecloudv1.ExternalInfoBase, error) {
	vm := vi.VirtualMachine
	// disable delete first
	params := jsonutils.NewDict()
	params.Set("disable_delete", jsonutils.JSONFalse)
	_, s, e := RequestVM.Operation(OperUpdate).Apply(ctx, vm.Status.ExternalInfo.Id, params)
	if e != nil {
		return s, e
	}
	// delete
	_, s, e = RequestVMDelete.Apply(ctx, vm.Status.ExternalInfo.Id, nil)
	return s, e
}

func (vi VirtualMachine) GetStatus(ctx context.Context) (rs onecloudv1.IResourceStatus, err error) {
	vm := vi.VirtualMachine
	lastInfo := vm.Status.ExternalInfo.ExternalInfoBase
	_, extInfo, err := RequestVM.Operation(OperGetStatus).Apply(ctx, vm.Status.ExternalInfo.Id, nil)
	if err != nil {
		return nil, err
	}
	vmStatus := vm.Status.DeepCopy()
	rs = vmStatus
	var (
		phase  onecloudv1.ResourcePhase
		reason string
	)
	// The order about checking info.Status is Critical
	switch {
	case vi.isRunning(extInfo):
		phase = onecloudv1.ResourceRunning
	case vi.isStopped(extInfo):
		phase = onecloudv1.ResourceReady
	case vi.isFailed(extInfo, lastInfo):
		phase = onecloudv1.ResourceFailed
	case vi.needSync(extInfo):
		_, extInfo, err = RequestVMSyncstatus.Apply(ctx, vm.Status.ExternalInfo.Id, nil)
		if err != nil {
			return
		}
		phase = onecloudv1.ResourcePending
	case vi.isPendingWithoutFailed(extInfo):
		phase = onecloudv1.ResourcePending
	default:
		recreatePolicy := vm.Spec.RecreatePolicy
		if recreatePolicy == nil {
			recreatePolicy = recreatePolicyDefault
		}
		switch {
		case recreatePolicy.Never != nil && *recreatePolicy.Never:
			phase = onecloudv1.ResourcePending
			reason = "RecreatePolicy instruct to never recreate"
		case onecloudutils.IsInStringArray(extInfo.Status, recreatePolicy.MatchStatus):
			phase = onecloudv1.ResourceFailed
			reason = "RecreatePolicy instruct to recreate because vm'status matched"
		case recreatePolicy.Allways != nil && *recreatePolicy.Allways:
			reason = "RecreatePolicy instruct to allways recreate"
			phase = onecloudv1.ResourceFailed
		default:
			reason = "RecreatePolicy's MatchStatus doesn't contains this status"
			phase = onecloudv1.ResourceUnkown
		}
	}
	if len(reason) == 0 {
		reason = fmt.Sprintf("Exec '%s' successfully", extInfo.Action)
	}
	vmStatus.Phase = phase
	vmStatus.Reason = reason
	vmStatus.ExternalInfo.ExternalInfoBase = extInfo
	return
}

func (vi VirtualMachine) DefaultRecreatePolicy() *onecloudv1.RecreatePolicy {
	return recreatePolicyDefault
}

func (vi VirtualMachine) Start(ctx context.Context) (onecloudv1.ExternalInfoBase, error) {
	_, s, e := RequestVM.Operation(OperStart).Apply(ctx, vi.VirtualMachine.Status.ExternalInfo.Id, nil)
	return s, e
}

func (vi VirtualMachine) Stop(ctx context.Context) (onecloudv1.ExternalInfoBase, error) {
	_, s, e := RequestVM.Operation(OperStop).Apply(ctx, vi.VirtualMachine.Status.ExternalInfo.Id, nil)
	return s, e
}

func (vi VirtualMachine) Reconcile(ctx context.Context, logger logr.Logger) (oper *ReconcileOper, vmInfo *onecloudv1.VMInfo, err error) {
	vm := vi.VirtualMachine
	// fetch vm details
	vmJson, extInfo, err := RequestVMGetDetails.Apply(ctx, vm.Status.ExternalInfo.Id, nil)
	if err != nil {
		return
	}
	var serverDetail comapi.ServerDetails
	err = vmJson.Unmarshal(&serverDetail)
	if err != nil {
		return
	}

	// update ExternalInfo
	vmInfo = vm.Status.ExternalInfo.DeepCopy()
	vmInfo.ExternalInfoBase = extInfo
	if vm.Status.ExternalInfo.Eip != serverDetail.Eip {
		vmInfo.Eip = serverDetail.Eip
	}
	newIps := util.MapStringSlice(strings.TrimSpace, strings.Split(serverDetail.IPs, ","))
	if !util.EqualStringSlices(vm.Status.ExternalInfo.Ips, newIps) {
		vmInfo.Ips = newIps
	}

	// update operator which is a lightweight operation
	oper = vi.update(&serverDetail, &vm.Spec)
	if oper != nil {
		return
	}
	// change config
	oper = vi.changeConfig(logger, &serverDetail, &vm.Spec.VmConfig)
	if oper != nil {
		return
	}
	// disk resize
	oper = vi.diskResize(logger, &serverDetail, &vm.Spec.VmConfig)
	if oper != nil {
		return
	}
	// changebw
	//oper, err = vi.eipChangeBw(ctx, logger, &serverDetail, &vm.Spec)
	//if err != nil || oper != nil {
	//	return
	//}
	// set secgroups
	oper = vi.setSecGroups(&serverDetail, &vm.Spec)
	return
}

func (vi VirtualMachine) eipChangeBw(ctx context.Context, logger logr.Logger, serverDetail *comapi.ServerDetails, vmSpec *onecloudv1.VirtualMachineSpec) (*ReconcileOper, error) {
	if len(vmSpec.Eip) != 0 && vmSpec.Eip != serverDetail.Eip {
		logger.V(1).Info(fmt.Sprintf("The actual eip '%s' is different from this '%s' in spec", serverDetail.Eip,
			vmSpec.Eip))
		return nil, nil
	}
	if vmSpec.NewEip == nil {
		return nil, nil
	}
	// fetch eip's info
	params := jsonutils.NewDict()
	params.Set("ip_addr", jsonutils.NewString(serverDetail.Eip))
	r, _, e := Request.Resource(ResourceEIP).Operation(OperGet).Apply(ctx, "", params)
	if e != nil {
		return nil, e
	}
	eipId, _ := r.GetString("id")
	eipBw, _ := r.Int("bandwidth")
	if vmSpec.NewEip.Bw == nil || *vmSpec.NewEip.Bw == eipBw {
		return nil, nil
	}
	specEipBw := *vmSpec.NewEip.Bw
	ro := &ReconcileOper{}
	ro.OperDesc.Appendf(`change "EIP.Bandwidth" from (%d) to (%d)`, eipBw, specEipBw)
	ro.Operator = func(ctx context.Context) (onecloudv1.ExternalInfoBase, error) {
		params := jsonutils.NewDict()
		params.Set("bandwidth", jsonutils.NewInt(specEipBw))
		_, s, e := Request.Resource(ResourceEIP).Operation(OperChangeBw).Apply(ctx, eipId, params)
		return s, e
	}
	return ro, nil
}

func (vi VirtualMachine) setSecGroups(serverDetail *comapi.ServerDetails, vmSpec *onecloudv1.VirtualMachineSpec) *ReconcileOper {
	count := len(serverDetail.Secgroups)
	update := false
	for i := range vmSpec.Secgropus {
		for j := range serverDetail.Secgroups {
			if vmSpec.Secgropus[i] == serverDetail.Secgroups[j].Id || vmSpec.Secgropus[i] == serverDetail.
				Secgroups[i].Name {
				count--
				break
			}
		}
		update = true
	}
	if count > 1 {
		update = true
	}
	if !update {
		return nil
	}
	ro := &ReconcileOper{}
	ro.OperDesc.Appendf(`change "Secgroups" to (%s)`, vmSpec.Secgropus)
	ro.Operator = func(ctx context.Context) (onecloudv1.ExternalInfoBase, error) {
		params := jsonutils.NewDict()
		for i := range vmSpec.Secgropus {
			params.Set(fmt.Sprintf("secgrp.%d", i), jsonutils.NewString(vmSpec.Secgropus[i]))
		}
		_, es, e := RequestVM.Operation(OperSetSecgroups).Apply(ctx, serverDetail.Id, params)
		return es, e
	}
	return ro
}

func (vi VirtualMachine) update(serverDetail *comapi.ServerDetails, vmSpec *onecloudv1.VirtualMachineSpec) *ReconcileOper {
	operDesc := OperatorDesc{Name: "Update"}
	updateParams := jsonutils.NewDict()
	if serverDetail.Name != vmSpec.Name && vmSpec.NameCheck != nil && *vmSpec.NameCheck && len(vmSpec.Name) > 0 {
		updateParams.Set("name", jsonutils.NewString(vmSpec.Name))
		operDesc.Append("name", serverDetail.Name, vmSpec.Name)
	}
	if serverDetail.Description != vmSpec.Desciption {
		updateParams.Set("description", jsonutils.NewString(vmSpec.Desciption))
		operDesc.Append("description", serverDetail.Description, vmSpec.Desciption)
	}
	if updateParams.Length() > 0 {
		operator := func(ctx context.Context) (onecloudv1.ExternalInfoBase, error) {
			_, es, e := RequestVM.Operation(OperUpdate).Apply(ctx, serverDetail.Id, updateParams)
			return es, e
		}
		return &ReconcileOper{
			Operator: operator,
			OperDesc: operDesc,
			PrePhase: []onecloudv1.ResourcePhase{},
		}
	}
	return nil
}

func (vi VirtualMachine) diskResize(logger logr.Logger, serverDetails *comapi.ServerDetails,
	vmConfig *onecloudv1.VirtualMachineConfig) *ReconcileOper {
	diskSizeMap := make(map[string]int64)
	odesc := OperatorDesc{Name: "Disk Resize"}
	// check RootDisk
	rootDiskInfo := serverDetails.DisksInfo[0]
	if specSize, realSize := vmConfig.RootDisk.SizeGB*1024, int64(serverDetails.DisksInfo[0].
		SizeMb); realSize < specSize {
		diskSizeMap[rootDiskInfo.Id] = specSize
		odesc.Appendf(`resize "RootDisk.SizeGB" from (%d) to (%d)`, realSize, specSize)
	} else if realSize > specSize {
		logger.V(1).Info(fmt.Sprintf("The actual size '%d' of the vm system disk is greater than that '%d' stated in the spec",
			realSize, specSize))
	}
	// check DataDisk
	for i := range vmConfig.DataDisks {
		if specSize, realSize := vmConfig.DataDisks[i].SizeGB*1024, int64(serverDetails.DisksInfo[i+1].
			SizeMb); realSize < specSize {
			diskSizeMap[serverDetails.DisksInfo[i+1].Id] = specSize
			odesc.Appendf(`resize "DataDisks[%d].SizeGB" from (%d) to (%d)`, i, realSize, specSize)
		} else if realSize > specSize {
			logger.V(1).Info(fmt.Sprintf("The actual size '%d' of DataDisk whose index is '%d' is greater than that '%d"+
				"' stated in the spec", realSize, i+1, specSize))
		}
	}
	if len(diskSizeMap) == 0 {
		return nil
	}
	operator := func(ctx context.Context) (onecloudv1.ExternalInfoBase, error) {
		var (
			es  onecloudv1.ExternalInfoBase
			err error
		)
		for id, size := range diskSizeMap {
			dict := jsonutils.NewDict()
			dict.Set("size", jsonutils.NewInt(size))
			_, es, err = Request.Resource(ResourceDisk).Operation(OperResize).Apply(ctx, id, dict)
			if err != nil {
				return es, err
			}
		}
		return es, nil
	}
	return &ReconcileOper{
		Operator: operator,
		OperDesc: odesc,
		PrePhase: []onecloudv1.ResourcePhase{onecloudv1.ResourceReady},
	}
}

// support
func (vi VirtualMachine) changeConfig(logger logr.Logger, serverDetails *comapi.ServerDetails,
	vmConfig *onecloudv1.VirtualMachineConfig) *ReconcileOper {
	var (
		params   = jsonutils.NewDict()
		adesc    = OperatorDesc{Name: "Change Config"}
		needStop = false
	)
	if len(vmConfig.InstanceType) == 0 {
		if vcpuCount := convertInt64Ptr(vmConfig.VcpuCount); vcpuCount != serverDetails.VcpuCount {
			if vcpuCount < serverDetails.VcpuCount {
				needStop = true
			}
			params.Set("vcpu_count", jsonutils.NewInt(int64(vcpuCount)))
			adesc.Appendf(`change "%s" from (%d) to (%d)`, "vcpu_count", serverDetails.VcpuCount, vcpuCount)
		}
		if vmemSize := convertInt64Ptr(vmConfig.VmemSizeGB) * 1024; vmemSize != serverDetails.VmemSize {
			if vmemSize < serverDetails.VmemSize {
				needStop = true
			}
			params.Set("vmem_size", jsonutils.NewInt(int64(vmemSize)))
			adesc.Appendf(`change "%s" from (%d) to (%d)`, "vmem_size", serverDetails.VmemSize, vmemSize)
		}
	} else {
		if vmConfig.InstanceType != serverDetails.InstanceType {
			needStop = true
			params.Set("instance_type", jsonutils.NewString(vmConfig.InstanceType))
			adesc.Append("instance_type", serverDetails.InstanceType, vmConfig.InstanceType)
		}
	}
	// check if add DataDisk
	if specN, realN := len(vmConfig.DataDisks), serverDetails.DiskCount-1; specN > realN {
		disks := jsonutils.NewArray()
		// copy existed data disks' config
		for i := 0; i < realN; i++ {
			disks.Add(jsonutils.Marshal(serverDetails.DisksInfo[i+1]))
		}
		for i := realN; i < specN; i++ {
			diskSpec := ConvertVMDisk(vmConfig.DataDisks[i])
			diskSpec.Index = i + 1
			diskSpec.DiskType = "data"
			disks.Add(jsonutils.Marshal(diskSpec))
			adesc.Appendf("create Data Disk(%s)", vi.diskSpecString(vmConfig.DataDisks[i]))
		}
		params.Set("disks", disks)
	} else if specN < realN {
		logger.V(1).Info(fmt.Sprintf("The actual number '%d' of data disks is greater than that '%d stated in the spec",
			realN, specN))
	}
	// change config operator
	if params.Length() == 0 {
		return nil
	}

	var prePhase []onecloudv1.ResourcePhase
	if needStop {
		prePhase = append(prePhase, onecloudv1.ResourceReady)
	}

	operFunc := func(ctx context.Context) (onecloudv1.ExternalInfoBase, error) {
		_, es, e := RequestVM.Operation(OperChangeConfig).Apply(ctx, serverDetails.Id, params)
		return es, e
	}

	return &ReconcileOper{
		Operator: operFunc,
		OperDesc: adesc,
		PrePhase: prePhase,
	}
}

func (vi VirtualMachine) diskSpecString(spec onecloudv1.VMDiskSpec) string {
	var buf strings.Builder
	buf.WriteString(fmt.Sprintf("size: %dGB, ", spec.SizeGB))
	if len(spec.Image) != 0 {
		buf.WriteString(fmt.Sprintf("image: %s, ", spec.Image))
	}
	if len(spec.Driver) != 0 {
		buf.WriteString(fmt.Sprintf("driver: %s, ", spec.Driver))
	}
	if len(spec.Storage.Storage) != 0 {
		buf.WriteString(fmt.Sprintf("storage: %s, ", spec.Storage.Storage))
	}
	if len(spec.Storage.Backend) != 0 {
		buf.WriteString(fmt.Sprintf("storage's backend: %s", spec.Storage.Backend))
	}
	str := buf.String()
	return str[:len(str)-2]
}

type DiskInfo struct {
	Index   int    `json:"index"`
	SizeMb  int    `json:"size"`
	Driver  string `json:"driver"`
	ImageId string `json:"image_id"`
	Image   string `json:"image"`
}

func init() {
	Register(ResourceVM, modules.Servers.ResourceManager)
	Register(ResourceEIP, modules.Elasticips)
	Register(ResourceDisk, modules.Disks)

	deleteParams := jsonutils.NewDict()
	deleteParams.Set("override_pending_delete", jsonutils.JSONTrue)
	detailParams := jsonutils.NewDict()
	detailParams.Set("details", jsonutils.JSONTrue)

	RequestVM = Request.Resource(ResourceVM)
	RequestVMDelete = RequestVM.Operation(OperDelete).DefaultParams(deleteParams)
	RequestVMGetDetails = RequestVM.Operation(OperGet).DefaultParams(detailParams)
	RequestVMSyncstatus = RequestVM.Operation(OperSyncstatus)
}

var (
	RequestVM           SRequest
	RequestVMDelete     SRequest
	RequestVMGetDetails SRequest
	RequestVMSyncstatus SRequest
)

var (
	True                  = true
	recreatePolicyDefault = &onecloudv1.RecreatePolicy{
		Never:       &True,
		MatchStatus: []string{},
		Allways:     nil,
	}
	vmCreatingStatus = append(comapi.VM_CREATING_STATUS, comapi.VM_INIT, comapi.VM_SCHEDULE)
	vmStopStatus     = []string{
		comapi.VM_READY,
		comapi.VM_START_FAILED,
	}
	vmFailedStatus = []string{
		comapi.VM_SCHEDULE_FAILED,
		comapi.VM_NETWORK_FAILED,
		comapi.VM_DEVICE_FAILED,
		comapi.VM_CREATE_FAILED,
		comapi.VM_DISK_FAILED,
		comapi.VM_DEPLOY_FAILED,
		comapi.VM_DELETE_FAIL,
		comapi.VM_UNKNOWN,
	}
)

func (vi VirtualMachine) isPendingWithoutFailed(info onecloudv1.ExternalInfoBase) bool {
	if onecloudutils.IsInStringArray(info.Status, vmCreatingStatus) {
		return true
	}
	if onecloudutils.IsInStringArray(info.Status, []string{comapi.VM_ASSOCIATE_EIP}) {
		return true
	}
	if info.Status == comapi.VM_DELETING || strings.Contains(info.Status, "delete") {
		return true
	}
	if info.Status == comapi.VM_STARTING || strings.Contains(info.Status, "start") {
		return true
	}
	if strings.Contains(info.Status, "change") {
		return true
	}
	if strings.Contains(info.Status, "sync") {
		return true
	}
	return false
}

func (vi VirtualMachine) isRunning(info onecloudv1.ExternalInfoBase) bool {
	return info.Status == comapi.VM_RUNNING
}

func (vi VirtualMachine) isFailed(info, lastInfo onecloudv1.ExternalInfoBase) bool {
	// That action "RequestVMChangeConfig" result in status "deploy_fail" and "disk_fail" maybe not failed,
	// it's better to syncstatus first.
	if (info.Status == comapi.VM_DEPLOY_FAILED || info.Status == comapi.VM_DISK_FAILED) && lastInfo.Action != RequestVMSyncstatus.ResourceAction() {
		return false
	}
	return onecloudutils.IsInStringArray(info.Status, vmFailedStatus)
}

func (vi VirtualMachine) needSync(info onecloudv1.ExternalInfoBase) bool {
	return strings.Contains(info.Status, "fail") || strings.Contains(info.Status, "failed")
}

func (vi VirtualMachine) isStopped(info onecloudv1.ExternalInfoBase) bool {
	return onecloudutils.IsInStringArray(info.Status, vmStopStatus)
}
