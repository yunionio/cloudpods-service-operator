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
	_, extInfo, err := RequestEndpoint.Operation(OperCreate).Apply(ctx, "", jsonutils.Marshal(cp))
	return extInfo, err
}

func (ep Endpoint) Delete(ctx context.Context) (onecloudv1.ExternalInfoBase, error) {
	rep := ep.Endpoint
	_, extInfo, err := RequestEndpoint.Operation(OperDelete).Apply(ctx, rep.Status.ExternalInfo.Id, nil)
	return extInfo, err
}

func (ep Endpoint) GetStatus(ctx context.Context) (*onecloudv1.EndpointStatus, error) {
	rep := ep.Endpoint
	_, extInfo, err := RequestEndpoint.Operation(OperGetStatus).Apply(ctx, rep.Status.ExternalInfo.Id, nil)
	if err != nil {
		return nil, err
	}
	_ := extInfo
	return nil, nil
}
