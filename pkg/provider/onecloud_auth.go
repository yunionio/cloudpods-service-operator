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
	"log"

	"yunion.io/x/onecloud/pkg/mcclient"
	"yunion.io/x/onecloud/pkg/mcclient/auth"

	"yunion.io/x/onecloud-resource-operator/pkg/options"
)

var AuthConfig AdminAuthConfig

type AdminAuthConfig struct {
	Region             string
	AuthURL            string
	AdminUsername      string
	AdminPassword      string
	AdminDomain        string
	AdminProject       string
	AdminProjectDomain string
}

func (config AdminAuthConfig) ToAuthInfo() *auth.AuthInfo {
	log.Println("start to auth info")
	if len(config.AuthURL) == 0 {
		log.Fatalln("Missing AuthURL")
	}

	if len(config.AdminUsername) == 0 {
		log.Fatalln("Mising AdminUser")
	}

	if len(config.AdminPassword) == 0 {
		log.Fatalln("Missing AdminPasswd")
	}

	if len(config.AdminProject) == 0 {
		log.Fatalln("Missing AdminProject")
	}
	return &auth.AuthInfo{
		AuthUrl:       config.AuthURL,
		Domain:        config.AdminDomain,
		Username:      config.AdminUsername,
		Passwd:        config.AdminPassword,
		Project:       config.AdminProject,
		ProjectDomain: config.AdminProjectDomain,
	}
}

func (op OnecloudProvider) Init() {
	AuthConfig = AdminAuthConfig{
		Region:        options.Options.Region,
		AuthURL:       options.Options.AuthURL,
		AdminUsername: options.Options.AdminUsername,
		AdminPassword: options.Options.AdminPassword,
		AdminDomain:   options.Options.AdminDomain,
		AdminProject:  options.Options.AdminProject,
	}
	auth.Init(AuthConfig.ToAuthInfo(), false, true, "", "")
}

func AdminSession(ctx context.Context) *mcclient.ClientSession {
	return auth.GetAdminSessionWithPublic(ctx, AuthConfig.Region, "")
}
