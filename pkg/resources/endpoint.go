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

	"yunion.io/x/jsonutils"
	onecloudv1 "yunion.io/x/onecloud-service-operator/api/v1"
	"yunion.io/x/onecloud/pkg/mcclient/modules"
)

var (
	RequestEndpoint = Request.Resource(ResourceEndpoint)
	RequestService  = Request.Resource(ResourceSevice)
)

func init() {
	Register(ResourceEndpoint, modules.EndpointsV3)
	Register(ResourceSevice, modules.ServicesV3)
}

var (
	Service   = "external-service"
	ServiceID = ""
	Interface = "console"
)

type EndpointCreateParams struct {
	Interface          string `json:"interface"`
	ServiceId          string `json:"service_id"`
	URL                string `json:"url"`
	RegionId           string `json:"region_id"`
	Name               string `json:"name"`
	Enabled            bool   `json:"enabled"`
	ServiceCertificate string `json:"service_certificate"`
}

type Endpoint struct {
	Endpoint *onecloudv1.Endpoint
}

func NewEndpoint(ep *onecloudv1.Endpoint) Endpoint {
	return Endpoint{ep}
}

func (ep Endpoint) GetResourceName() Resource {
	return ResourceEndpoint
}

func (ep Endpoint) GetIResource() onecloudv1.IResource {
	return ep.Endpoint
}

func (ep Endpoint) Create(ctx context.Context, params interface{}) (onecloudv1.ExternalInfoBase, error) {
	cp, ok := params.(EndpointCreateParams)
	if !ok {
		return onecloudv1.ExternalInfoBase{}, fmt.Errorf("Invalid create params")
	}
	var err error
	// find serveice id
	if len(ServiceID) == 0 {
		ServiceID, err = RequestService.GetId(ctx, Service)
		if err != nil {
			return onecloudv1.ExternalInfoBase{}, err
		}
	}
	cp.ServiceId = ServiceID
	cp.Interface = Interface
	_, extInfo, err := RequestEndpoint.Operation(OperCreate).Apply(ctx, "", jsonutils.Marshal(cp).(*jsonutils.JSONDict))
	return extInfo, err
}

func (ep Endpoint) Delete(ctx context.Context) (onecloudv1.ExternalInfoBase, error) {
	rep := ep.Endpoint
	// disable delete first
	params := jsonutils.NewDict()
	params.Set("enabled", jsonutils.JSONFalse)
	_, s, e := RequestEndpoint.Operation(OperPatch).Apply(ctx, rep.Status.ExternalInfo.Id, params)
	if e != nil {
		return s, e
	}
	// delete
	_, extInfo, err := RequestEndpoint.Operation(OperDelete).Apply(ctx, rep.Status.ExternalInfo.Id, nil)
	return extInfo, err
}

func (ep Endpoint) GetStatus(ctx context.Context) (onecloudv1.IResourceStatus, error) {
	rep := ep.Endpoint
	_, extInfo, err := RequestEndpoint.Operation(OperGetStatus).Apply(ctx, rep.Status.ExternalInfo.Id, nil)
	if err != nil {
		return nil, err
	}
	return &onecloudv1.EndpointStatus{
		ResourceStatusBase: onecloudv1.ResourceStatusBase{
			Phase:        onecloudv1.ResourceReady,
			Reason:       "",
		},
		ExternalInfo: extInfo,
	}, nil
}

type EndpointUpdateField struct {
	Url                string `json:"url"`
	Name               string `json:"name"`
	Enabled            bool `json:"enabled"`
	ServiceCertificate string `json:"service_certificate"`
}

func (ep Endpoint) Reconcile(ctx context.Context) (*onecloudv1.EndpointStatus, error) {
	rep := ep.Endpoint
	ret, extInfo, err := RequestEndpoint.Operation(OperGet).Apply(ctx, rep.Status.ExternalInfo.Id, nil)
	if err != nil {
		return nil, err
	}
	epStatus := rep.Status.DeepCopy()
	epStatus.ExternalInfo = extInfo
	var uField EndpointUpdateField
	err = ret.Unmarshal(&uField)
	if err != nil {
		return nil, err
	}
	var change bool
	if  len(rep.Spec.Name) > 0 && uField.Name != rep.Spec.Name {
		uField.Name = rep.Spec.Name
		change = true
	}
	endabled := true
	if rep.Spec.Disabled != nil && *rep.Spec.Disabled {
		endabled = false
	}
	if uField.Enabled != endabled {
		uField.Enabled = endabled
		change = true
	}
	if uField.ServiceCertificate != rep.Spec.ServiceCertificate {
		uField.ServiceCertificate = rep.Spec.ServiceCertificate
		change = true
	}
	url, err := ep.Url(ctx)
	if err != nil {
		return nil, err
	}
	if len(url) != 0 && uField.Url != url {
		uField.Url = url
		change = true
	}
	if !change {
		return epStatus, nil
	}
	_, extInfo, err = RequestEndpoint.Operation(OperPatch).Apply(ctx, rep.Status.ExternalInfo.Id, jsonutils.Marshal(uField).(*jsonutils.JSONDict))
	if err != nil {
		return nil, err
	}
	epStatus.ExternalInfo = extInfo
	return epStatus, nil
}

func (ep Endpoint) Url(ctx context.Context) (string, error) {
	// fetch Host
	url := ep.Endpoint.Spec.URL
	host := url.Host
	v, err := host.GetValue(ctx)
	if err != nil || v == nil || v.IsZero() {
		return "", err
	}
	s := v.(onecloudv1.String)
	protocol := "http"
	if len(url.Protocol) != 0 {
		protocol = url.Protocol
	}
	if url.Port != nil && *url.Port > 0 {
		return fmt.Sprintf("%s://%s:%d/%s", protocol, s, *url.Port, url.Prefix), nil
	}
	return fmt.Sprintf("%s://%s/%s", protocol, s, url.Prefix), nil
}
