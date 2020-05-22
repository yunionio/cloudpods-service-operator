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

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	onecloudv1 "yunion.io/x/onecloud-service-operator/api/v1"
	"yunion.io/x/onecloud-service-operator/pkg/provider"
)

func dealErr(ctx context.Context, logger logr.Logger, sclient client.StatusClient, re onecloudv1.IResource,
	resource provider.Resource, err error) (ctrl.Result, error) {

	reErr, ok := err.(*provider.SRequestErr)
	if !ok {
		logger.Error(err, "")
		return ctrl.Result{}, err
	}

	if reErr.IsNotFound(resource) {
		re.SetExternalId("")
	}
	if reErr.IsClientErr() {
		re.SetResourcePhase(onecloudv1.ResourcePending, reErr.Error())
	}
	if reErr.IsServerErr() {
		re.SetResourcePhase(onecloudv1.ResourceUnkown, reErr.Error())
	}
	return ctrl.Result{}, sclient.Status().Update(ctx, re)
}
