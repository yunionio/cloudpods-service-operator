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

package controllers

import (
	"context"
	"fmt"
	"time"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"yunion.io/x/onecloud-service-operator/pkg/resources"
	onecloudv1 "yunion.io/x/onecloud-service-operator/api/v1"
)

// EndpointReconciler reconciles a Endpoint object
type EndpointReconciler struct {
	ReconcilerBase
}

// +kubebuilder:rbac:groups=onecloud.yunion.io,resources=endpoints,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=onecloud.yunion.io,resources=endpoints/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=onecloud.yunion.io,resources=virtualmachines,verbs=get;list;watch

func (r *EndpointReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()

	var endpoint onecloudv1.Endpoint
	if err := r.Get(ctx, req.NamespacedName, &endpoint); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	log := r.GetLog(&endpoint)
	remoteEP := resources.NewEndpoint(&endpoint)

	dealErr := func(err error) (ctrl.Result, error) {
		return r.dealErr(ctx, remoteEP, err)
	}

	has, ret, err := r.UseFinallizer(ctx, remoteEP)
	if !has {
		return ret, err
	}

	if endpoint.Status.Phase == onecloudv1.ResourceInvalid {
		return ctrl.Result{}, nil
	}

	if len(endpoint.Status.ExternalInfo.Id) == 0 {
		url, err := remoteEP.Url(ctx)
		if err != nil {
			log.Error(err, "String.GetValue")
			endpoint.GetResourceStatus().SetPhase(onecloudv1.ResourceInvalid, err.Error())
			return ctrl.Result{}, r.Status().Update(ctx, &endpoint)
		}
		if len(url) == 0 {
			return r.MarkWaiting(ctx, &endpoint, fmt.Sprintf("wait for Endpoint(%s)", req.NamespacedName), 0)
		}
		createParams := resources.EndpointCreateParams{
			URL:                url,
			RegionId:           endpoint.Spec.RegionId,
			Name:               endpoint.Spec.Name,
			Enabled:            true,
			ServiceCertificate: endpoint.Spec.ServiceCertificate,
		}
		if endpoint.Spec.Disabled != nil && *endpoint.Spec.Disabled {
			createParams.Enabled = false
		}
		return r.Create(ctx, remoteEP, createParams, false)
	}

	// Unkown
	if endpoint.Status.Phase == onecloudv1.ResourceUnkown {
		return ctrl.Result{Requeue: true, RequeueAfter: 2 * time.Second}, nil
	}

	status, err := remoteEP.Reconcile(ctx)
	if err != nil {
		return dealErr(err)
	}
	endpoint.Status = *status
	return ctrl.Result{}, r.Status().Update(ctx, &endpoint)
}

func (r *EndpointReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&onecloudv1.Endpoint{}).
		Complete(r)
}
