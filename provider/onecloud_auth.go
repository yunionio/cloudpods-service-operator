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
	"flag"
	"log"

	"yunion.io/x/onecloud/pkg/mcclient"
	"yunion.io/x/onecloud/pkg/mcclient/auth"
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

func (op OnecloudProvider) InitConfig() {
	flag.StringVar(&AuthConfig.Region, "region", "", "Region name or ID")
	flag.StringVar(&AuthConfig.AuthURL, "auth-url", "", "Keystone auth URL")
	flag.StringVar(&AuthConfig.AdminUsername, "admin-username", "", "Admin username")
	flag.StringVar(&AuthConfig.AdminPassword, "admin-password", "", "Admin password")
	flag.StringVar(&AuthConfig.AdminDomain, "admin-domain", "", "Admin domain")
	flag.StringVar(&AuthConfig.AdminProject, "admin-project", "", "Admin project")
	flag.StringVar(&AuthConfig.AdminProjectDomain, "admin-project-domain", "", "Admin project domain")
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
	auth.Init(AuthConfig.ToAuthInfo(), false, true, "", "")
}

func AdminSession(ctx context.Context) *mcclient.ClientSession {
	return auth.GetAdminSessionWithPublic(ctx, AuthConfig.Region, "")
}
