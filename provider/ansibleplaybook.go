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

package provider

import (
	"context"
	"fmt"

	"yunion.io/x/jsonutils"
	anapi "yunion.io/x/onecloud/pkg/apis/ansible"
	"yunion.io/x/onecloud/pkg/mcclient/modules"
	"yunion.io/x/onecloud/pkg/util/ansiblev2"

	onecloudv1 "onecloud-service-operator/api/v1"
)

var (
	RequestAP = Request.Resource(ResourceAP)
)

func init() {
	Register(ResourceAP, modules.AnsiblePlaybooksV2.ResourceManager)
}

func (oc OnecloudProvider) APCreate(ctx context.Context, ap *onecloudv1.AnsiblePlaybook,
	hosts []AnsiblePlaybookHost, apt *onecloudv1.AnsiblePlaybookTemplate,
	commonVars map[string]string) (onecloudv1.ExternalInfoBase,
	error) {
	// build inventory
	params := jsonutils.NewDict()
	params.Set("playbook", jsonutils.NewString(apt.Spec.Playbook))
	params.Set("files", jsonutils.NewString(apt.Spec.Files))
	params.Set("requirements", jsonutils.NewString(apt.Spec.Requirements))

	args := make([]interface{}, 0, len(commonVars)*2)
	for k, v := range commonVars {
		args = append(args, k, v)
	}
	inv := ansiblev2.NewInventory(args...)
	for i := range hosts {
		host := oc.apHosts(hosts[i])
		inv.SetHost(hosts[i].VM.Name, host)
	}
	params.Set("inventory", jsonutils.NewString(inv.String()))

	params.Set("generate_name", jsonutils.NewString(ap.Name))
	_, extInfo, err := RequestAP.Operation(OperCreate).Apply(ctx, "", params)
	return extInfo, err
}

func (oc OnecloudProvider) APDelete(ctx context.Context, ap *onecloudv1.AnsiblePlaybook) (bool, onecloudv1.ExternalInfoBase, error) {
	_, extInfo, err := RequestAP.Operation(OperGet).Apply(ctx, ap.Status.ExternalInfo.Id, nil)
	if err != nil {
		return false, extInfo, err
	}
	if extInfo.Status == anapi.AnsiblePlaybookStatusRunning {
		// cancel first
		_, extInfo, err := RequestAP.Operation(OperStop).Apply(ctx, ap.Status.ExternalInfo.Id, nil)
		return false, extInfo, err
	}
	_, extInfo, err = RequestAP.Operation(OperDelete).Apply(ctx, ap.Status.ExternalInfo.Id, nil)
	return true, extInfo, err
}

func (oc OnecloudProvider) APGetStatus(ctx context.Context, ap *onecloudv1.AnsiblePlaybook) (*onecloudv1.AnsiblePlaybookStatus, error) {
	_, extInfo, err := RequestAP.Operation(OperGetStatus).Apply(ctx, ap.Status.ExternalInfo.Id, nil)
	if err != nil {
		return nil, err
	}
	apStatus := ap.Status.DeepCopy()
	switch extInfo.Status {
	case anapi.AnsiblePlaybookStatusInit, anapi.AnsiblePlaybookStatusRunning:
		apStatus.Phase = onecloudv1.ResourcePending
	case anapi.AnsiblePlaybookStatusCanceled, anapi.AnsiblePlaybookStatusFailed, anapi.AnsiblePlaybookStatusUnknown:
		apStatus.Phase = onecloudv1.ResourceFailed
	case anapi.AnsiblePlaybookStatusSucceeded:
		apStatus.Phase = onecloudv1.ResourceFinished
	}
	apStatus.ExternalInfo.ExternalInfoBase = extInfo
	apStatus.Reason = fmt.Sprintf("Exec '%s' successfully", extInfo.Action)
	return apStatus, nil
}

func (oc OnecloudProvider) APReconcile(ctx context.Context, ap *onecloudv1.AnsiblePlaybook) (*onecloudv1.AnsiblePlaybookStatus, error) {
	ret, extInfo, err := RequestAP.Operation(OperGet).Apply(ctx, ap.Status.ExternalInfo.Id, nil)
	if err != nil {
		return nil, err
	}
	apStatus := ap.Status.DeepCopy()
	// update output
	output, _ := ret.GetString("output")
	if apStatus.ExternalInfo.Output != output {
		apStatus.ExternalInfo.Output = output
	}
	switch extInfo.Status {
	case anapi.AnsiblePlaybookStatusInit, anapi.AnsiblePlaybookStatusRunning:
		apStatus.Phase = onecloudv1.ResourcePending
	case anapi.AnsiblePlaybookStatusCanceled, anapi.AnsiblePlaybookStatusFailed, anapi.AnsiblePlaybookStatusUnknown:
		apStatus.Phase = onecloudv1.ResourceFailed
	case anapi.AnsiblePlaybookStatusSucceeded:
		apStatus.Phase = onecloudv1.ResourceFinished
	}
	apStatus.ExternalInfo.ExternalInfoBase = extInfo
	return apStatus, nil
}

func (oc OnecloudProvider) apHosts(host AnsiblePlaybookHost) *ansiblev2.Host {
	var ip string
	switch {
	case len(host.VM.Status.ExternalInfo.Eip) > 0:
		ip = host.VM.Status.ExternalInfo.Eip
	case len(host.VM.Status.ExternalInfo.Ips) > 0:
		ip = host.VM.Status.ExternalInfo.Ips[0]
	default:
		// noway
	}
	vars := map[string]interface{}{
		"ansible_user": "root",
		"ansible_host": ip,
	}
	for k, v := range host.Vars {
		vars[k] = v
	}
	h := ansiblev2.NewHost()
	h.Vars = vars
	return h
}
