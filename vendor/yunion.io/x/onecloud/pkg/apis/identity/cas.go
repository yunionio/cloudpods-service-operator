// Copyright 2019 Yunion
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

package identity

import "yunion.io/x/pkg/tristate"

type SCASIdpConfigOptions struct {
	// https://cas.example.org/cas/
	CASServerURL string `json:"cas_server_url"`

	// Deprecated
	CasProjectAttribute string `json:"cas_project_attribute" "deprecated-by":"project_attribute"`
	// Deprecated
	AutoCreateCasProject tristate.TriState `json:"auto_create_cas_project"`
	// Deprecated
	DefaultCasProjectId string `json:"default_cas_project_id" "deprecated-by":"default_project_id"`
	// Deprecated
	CasRoleAttribute string `json:"cas_role_attribute" "deprected-by":"role_attribute"`
	// Deprecated
	DefaultCasRoleId string `json:"default_cas_role_id" "deprecated-by":"default_role_id"`

	SIdpAttributeOptions
}
