// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnsiblePlaybook) DeepCopyInto(out *AnsiblePlaybook) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnsiblePlaybook.
func (in *AnsiblePlaybook) DeepCopy() *AnsiblePlaybook {
	if in == nil {
		return nil
	}
	out := new(AnsiblePlaybook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AnsiblePlaybook) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnsiblePlaybookHost) DeepCopyInto(out *AnsiblePlaybookHost) {
	*out = *in
	out.VirtualMachine = in.VirtualMachine
	if in.Vars != nil {
		in, out := &in.Vars, &out.Vars
		*out = make(map[string]IntOrStringOrYamlStore, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnsiblePlaybookHost.
func (in *AnsiblePlaybookHost) DeepCopy() *AnsiblePlaybookHost {
	if in == nil {
		return nil
	}
	out := new(AnsiblePlaybookHost)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnsiblePlaybookInfo) DeepCopyInto(out *AnsiblePlaybookInfo) {
	*out = *in
	out.ExternalInfoBase = in.ExternalInfoBase
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnsiblePlaybookInfo.
func (in *AnsiblePlaybookInfo) DeepCopy() *AnsiblePlaybookInfo {
	if in == nil {
		return nil
	}
	out := new(AnsiblePlaybookInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnsiblePlaybookList) DeepCopyInto(out *AnsiblePlaybookList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AnsiblePlaybook, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnsiblePlaybookList.
func (in *AnsiblePlaybookList) DeepCopy() *AnsiblePlaybookList {
	if in == nil {
		return nil
	}
	out := new(AnsiblePlaybookList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AnsiblePlaybookList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnsiblePlaybookProxyVar) DeepCopyInto(out *AnsiblePlaybookProxyVar) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnsiblePlaybookProxyVar.
func (in *AnsiblePlaybookProxyVar) DeepCopy() *AnsiblePlaybookProxyVar {
	if in == nil {
		return nil
	}
	out := new(AnsiblePlaybookProxyVar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnsiblePlaybookSpec) DeepCopyInto(out *AnsiblePlaybookSpec) {
	*out = *in
	if in.PlaybookTemplateRef != nil {
		in, out := &in.PlaybookTemplateRef, &out.PlaybookTemplateRef
		*out = new(LocalObjectReference)
		**out = **in
	}
	if in.PlaybookTemplate != nil {
		in, out := &in.PlaybookTemplate, &out.PlaybookTemplate
		*out = new(AnsiblePlaybookTemplateSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Inventory != nil {
		in, out := &in.Inventory, &out.Inventory
		*out = make([]AnsiblePlaybookHost, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Vars != nil {
		in, out := &in.Vars, &out.Vars
		*out = make(map[string]IntOrStringOrYamlStore, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.ProxyVars != nil {
		in, out := &in.ProxyVars, &out.ProxyVars
		*out = make([]AnsiblePlaybookProxyVar, len(*in))
		copy(*out, *in)
	}
	in.ResourceSpecBase.DeepCopyInto(&out.ResourceSpecBase)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnsiblePlaybookSpec.
func (in *AnsiblePlaybookSpec) DeepCopy() *AnsiblePlaybookSpec {
	if in == nil {
		return nil
	}
	out := new(AnsiblePlaybookSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnsiblePlaybookStatus) DeepCopyInto(out *AnsiblePlaybookStatus) {
	*out = *in
	out.ResourceStatusBase = in.ResourceStatusBase
	out.ExternalInfo = in.ExternalInfo
	if in.DevtoolSshInfos != nil {
		in, out := &in.DevtoolSshInfos, &out.DevtoolSshInfos
		*out = make([]DevtoolSshInfo, len(*in))
		copy(*out, *in)
	}
	if in.ServiceUrls != nil {
		in, out := &in.ServiceUrls, &out.ServiceUrls
		*out = make([]DevtoolServiceUrl, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnsiblePlaybookStatus.
func (in *AnsiblePlaybookStatus) DeepCopy() *AnsiblePlaybookStatus {
	if in == nil {
		return nil
	}
	out := new(AnsiblePlaybookStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnsiblePlaybookTemplate) DeepCopyInto(out *AnsiblePlaybookTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnsiblePlaybookTemplate.
func (in *AnsiblePlaybookTemplate) DeepCopy() *AnsiblePlaybookTemplate {
	if in == nil {
		return nil
	}
	out := new(AnsiblePlaybookTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AnsiblePlaybookTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnsiblePlaybookTemplateList) DeepCopyInto(out *AnsiblePlaybookTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AnsiblePlaybookTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnsiblePlaybookTemplateList.
func (in *AnsiblePlaybookTemplateList) DeepCopy() *AnsiblePlaybookTemplateList {
	if in == nil {
		return nil
	}
	out := new(AnsiblePlaybookTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AnsiblePlaybookTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnsiblePlaybookTemplateSpec) DeepCopyInto(out *AnsiblePlaybookTemplateSpec) {
	*out = *in
	if in.Vars != nil {
		in, out := &in.Vars, &out.Vars
		*out = make([]AnsiblePlaybookTemplateVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnsiblePlaybookTemplateSpec.
func (in *AnsiblePlaybookTemplateSpec) DeepCopy() *AnsiblePlaybookTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(AnsiblePlaybookTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnsiblePlaybookTemplateStatus) DeepCopyInto(out *AnsiblePlaybookTemplateStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnsiblePlaybookTemplateStatus.
func (in *AnsiblePlaybookTemplateStatus) DeepCopy() *AnsiblePlaybookTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(AnsiblePlaybookTemplateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnsiblePlaybookTemplateVar) DeepCopyInto(out *AnsiblePlaybookTemplateVar) {
	*out = *in
	if in.Required != nil {
		in, out := &in.Required, &out.Required
		*out = new(bool)
		**out = **in
	}
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(IntOrString)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnsiblePlaybookTemplateVar.
func (in *AnsiblePlaybookTemplateVar) DeepCopy() *AnsiblePlaybookTemplateVar {
	if in == nil {
		return nil
	}
	out := new(AnsiblePlaybookTemplateVar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevtoolServiceUrl) DeepCopyInto(out *DevtoolServiceUrl) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevtoolServiceUrl.
func (in *DevtoolServiceUrl) DeepCopy() *DevtoolServiceUrl {
	if in == nil {
		return nil
	}
	out := new(DevtoolServiceUrl)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevtoolSshInfo) DeepCopyInto(out *DevtoolSshInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevtoolSshInfo.
func (in *DevtoolSshInfo) DeepCopy() *DevtoolSshInfo {
	if in == nil {
		return nil
	}
	out := new(DevtoolSshInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DisplayIdenSpec) DeepCopyInto(out *DisplayIdenSpec) {
	*out = *in
	if in.NameCheck != nil {
		in, out := &in.NameCheck, &out.NameCheck
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DisplayIdenSpec.
func (in *DisplayIdenSpec) DeepCopy() *DisplayIdenSpec {
	if in == nil {
		return nil
	}
	out := new(DisplayIdenSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Endpoint) DeepCopyInto(out *Endpoint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Endpoint.
func (in *Endpoint) DeepCopy() *Endpoint {
	if in == nil {
		return nil
	}
	out := new(Endpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Endpoint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointList) DeepCopyInto(out *EndpointList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Endpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointList.
func (in *EndpointList) DeepCopy() *EndpointList {
	if in == nil {
		return nil
	}
	out := new(EndpointList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EndpointList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointSpec) DeepCopyInto(out *EndpointSpec) {
	*out = *in
	in.URL.DeepCopyInto(&out.URL)
	if in.Disabled != nil {
		in, out := &in.Disabled, &out.Disabled
		*out = new(bool)
		**out = **in
	}
	in.ResourceSpecBase.DeepCopyInto(&out.ResourceSpecBase)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointSpec.
func (in *EndpointSpec) DeepCopy() *EndpointSpec {
	if in == nil {
		return nil
	}
	out := new(EndpointSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointStatus) DeepCopyInto(out *EndpointStatus) {
	*out = *in
	out.ResourceStatusBase = in.ResourceStatusBase
	out.ExternalInfo = in.ExternalInfo
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointStatus.
func (in *EndpointStatus) DeepCopy() *EndpointStatus {
	if in == nil {
		return nil
	}
	out := new(EndpointStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalInfoBase) DeepCopyInto(out *ExternalInfoBase) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalInfoBase.
func (in *ExternalInfoBase) DeepCopy() *ExternalInfoBase {
	if in == nil {
		return nil
	}
	out := new(ExternalInfoBase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IntOrString) DeepCopyInto(out *IntOrString) {
	*out = *in
	out.IntOrString = in.IntOrString
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IntOrString.
func (in *IntOrString) DeepCopy() *IntOrString {
	if in == nil {
		return nil
	}
	out := new(IntOrString)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IntOrStringOrYamlStore) DeepCopyInto(out *IntOrStringOrYamlStore) {
	*out = *in
	if in.IsYaml != nil {
		in, out := &in.IsYaml, &out.IsYaml
		*out = new(bool)
		**out = **in
	}
	in.IntOrStringStore.DeepCopyInto(&out.IntOrStringStore)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IntOrStringOrYamlStore.
func (in *IntOrStringOrYamlStore) DeepCopy() *IntOrStringOrYamlStore {
	if in == nil {
		return nil
	}
	out := new(IntOrStringOrYamlStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IntOrStringStore) DeepCopyInto(out *IntOrStringStore) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(IntOrString)
		**out = **in
	}
	if in.Reference != nil {
		in, out := &in.Reference, &out.Reference
		*out = new(ObjectFieldReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IntOrStringStore.
func (in *IntOrStringStore) DeepCopy() *IntOrStringStore {
	if in == nil {
		return nil
	}
	out := new(IntOrStringStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalObjectReference) DeepCopyInto(out *LocalObjectReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalObjectReference.
func (in *LocalObjectReference) DeepCopy() *LocalObjectReference {
	if in == nil {
		return nil
	}
	out := new(LocalObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectFieldReference) DeepCopyInto(out *ObjectFieldReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectFieldReference.
func (in *ObjectFieldReference) DeepCopy() *ObjectFieldReference {
	if in == nil {
		return nil
	}
	out := new(ObjectFieldReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectReference) DeepCopyInto(out *ObjectReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectReference.
func (in *ObjectReference) DeepCopy() *ObjectReference {
	if in == nil {
		return nil
	}
	out := new(ObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecreatePolicy) DeepCopyInto(out *RecreatePolicy) {
	*out = *in
	if in.MatchStatus != nil {
		in, out := &in.MatchStatus, &out.MatchStatus
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Never != nil {
		in, out := &in.Never, &out.Never
		*out = new(bool)
		**out = **in
	}
	if in.Allways != nil {
		in, out := &in.Allways, &out.Allways
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecreatePolicy.
func (in *RecreatePolicy) DeepCopy() *RecreatePolicy {
	if in == nil {
		return nil
	}
	out := new(RecreatePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSpecBase) DeepCopyInto(out *ResourceSpecBase) {
	*out = *in
	if in.MaxTryTimes != nil {
		in, out := &in.MaxTryTimes, &out.MaxTryTimes
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSpecBase.
func (in *ResourceSpecBase) DeepCopy() *ResourceSpecBase {
	if in == nil {
		return nil
	}
	out := new(ResourceSpecBase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceStatusBase) DeepCopyInto(out *ResourceStatusBase) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceStatusBase.
func (in *ResourceStatusBase) DeepCopy() *ResourceStatusBase {
	if in == nil {
		return nil
	}
	out := new(ResourceStatusBase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StringStore) DeepCopyInto(out *StringStore) {
	*out = *in
	if in.Reference != nil {
		in, out := &in.Reference, &out.Reference
		*out = new(ObjectFieldReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StringStore.
func (in *StringStore) DeepCopy() *StringStore {
	if in == nil {
		return nil
	}
	out := new(StringStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *URL) DeepCopyInto(out *URL) {
	*out = *in
	in.Host.DeepCopyInto(&out.Host)
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new URL.
func (in *URL) DeepCopy() *URL {
	if in == nil {
		return nil
	}
	out := new(URL)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMDiskSpec) DeepCopyInto(out *VMDiskSpec) {
	*out = *in
	out.Storage = in.Storage
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMDiskSpec.
func (in *VMDiskSpec) DeepCopy() *VMDiskSpec {
	if in == nil {
		return nil
	}
	out := new(VMDiskSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMInfo) DeepCopyInto(out *VMInfo) {
	*out = *in
	out.ExternalInfoBase = in.ExternalInfoBase
	if in.Ips != nil {
		in, out := &in.Ips, &out.Ips
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMInfo.
func (in *VMInfo) DeepCopy() *VMInfo {
	if in == nil {
		return nil
	}
	out := new(VMInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMNetworkSpec) DeepCopyInto(out *VMNetworkSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMNetworkSpec.
func (in *VMNetworkSpec) DeepCopy() *VMNetworkSpec {
	if in == nil {
		return nil
	}
	out := new(VMNetworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMNewEipSpec) DeepCopyInto(out *VMNewEipSpec) {
	*out = *in
	if in.Bw != nil {
		in, out := &in.Bw, &out.Bw
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMNewEipSpec.
func (in *VMNewEipSpec) DeepCopy() *VMNewEipSpec {
	if in == nil {
		return nil
	}
	out := new(VMNewEipSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMPasswordSpec) DeepCopyInto(out *VMPasswordSpec) {
	*out = *in
	if in.ResetPassword != nil {
		in, out := &in.ResetPassword, &out.ResetPassword
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMPasswordSpec.
func (in *VMPasswordSpec) DeepCopy() *VMPasswordSpec {
	if in == nil {
		return nil
	}
	out := new(VMPasswordSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMPreferSpec) DeepCopyInto(out *VMPreferSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMPreferSpec.
func (in *VMPreferSpec) DeepCopy() *VMPreferSpec {
	if in == nil {
		return nil
	}
	out := new(VMPreferSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMProjectSpec) DeepCopyInto(out *VMProjectSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMProjectSpec.
func (in *VMProjectSpec) DeepCopy() *VMProjectSpec {
	if in == nil {
		return nil
	}
	out := new(VMProjectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMStorageSpec) DeepCopyInto(out *VMStorageSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMStorageSpec.
func (in *VMStorageSpec) DeepCopy() *VMStorageSpec {
	if in == nil {
		return nil
	}
	out := new(VMStorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachine) DeepCopyInto(out *VirtualMachine) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachine.
func (in *VirtualMachine) DeepCopy() *VirtualMachine {
	if in == nil {
		return nil
	}
	out := new(VirtualMachine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineConfig) DeepCopyInto(out *VirtualMachineConfig) {
	*out = *in
	out.VMPreferSpec = in.VMPreferSpec
	if in.VcpuCount != nil {
		in, out := &in.VcpuCount, &out.VcpuCount
		*out = new(int64)
		**out = **in
	}
	if in.VmemSizeGB != nil {
		in, out := &in.VmemSizeGB, &out.VmemSizeGB
		*out = new(int64)
		**out = **in
	}
	out.RootDisk = in.RootDisk
	if in.DataDisks != nil {
		in, out := &in.DataDisks, &out.DataDisks
		*out = make([]VMDiskSpec, len(*in))
		copy(*out, *in)
	}
	if in.Networks != nil {
		in, out := &in.Networks, &out.Networks
		*out = make([]VMNetworkSpec, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineConfig.
func (in *VirtualMachineConfig) DeepCopy() *VirtualMachineConfig {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineList) DeepCopyInto(out *VirtualMachineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualMachine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineList.
func (in *VirtualMachineList) DeepCopy() *VirtualMachineList {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineSpec) DeepCopyInto(out *VirtualMachineSpec) {
	*out = *in
	in.DisplayIdenSpec.DeepCopyInto(&out.DisplayIdenSpec)
	in.VmConfig.DeepCopyInto(&out.VmConfig)
	out.Project = in.Project
	in.VMPasswordSpec.DeepCopyInto(&out.VMPasswordSpec)
	if in.Secgropus != nil {
		in, out := &in.Secgropus, &out.Secgropus
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NewEip != nil {
		in, out := &in.NewEip, &out.NewEip
		*out = new(VMNewEipSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.AutoRenew != nil {
		in, out := &in.AutoRenew, &out.AutoRenew
		*out = new(bool)
		**out = **in
	}
	if in.RecreatePolicy != nil {
		in, out := &in.RecreatePolicy, &out.RecreatePolicy
		*out = new(RecreatePolicy)
		(*in).DeepCopyInto(*out)
	}
	in.ResourceSpecBase.DeepCopyInto(&out.ResourceSpecBase)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineSpec.
func (in *VirtualMachineSpec) DeepCopy() *VirtualMachineSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineStatus) DeepCopyInto(out *VirtualMachineStatus) {
	*out = *in
	out.ResourceStatusBase = in.ResourceStatusBase
	in.ExternalInfo.DeepCopyInto(&out.ExternalInfo)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineStatus.
func (in *VirtualMachineStatus) DeepCopy() *VirtualMachineStatus {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Yaml) DeepCopyInto(out *Yaml) {
	{
		in := &in
		*out = make(Yaml, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Yaml.
func (in Yaml) DeepCopy() Yaml {
	if in == nil {
		return nil
	}
	out := new(Yaml)
	in.DeepCopyInto(out)
	return *out
}
