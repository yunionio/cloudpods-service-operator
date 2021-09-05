package resources

import (
	"context"

	"yunion.io/x/jsonutils"
	"yunion.io/x/onecloud/pkg/mcclient/modules"
	"yunion.io/x/pkg/errors"
)

var (
	RequestDevtoolSshInfo    = Request.Resource(ResourceDevtoolSshInfo)
	RequestDevtoolServiceUrl = Request.Resource(ResourceDevtoolServiceUrl)
)

func init() {
	Register(ResourceDevtoolSshInfo, modules.DevToolSshInfos)
	Register(ResourceDevtoolServiceUrl, modules.DevToolServiceUrls)
}

type DevtoolSshInfo struct {
	Id           string
	User         string
	Host         string
	Port         int
	ServerName   string
	Status       string
	FailedReason string
}

type DevtoolServiceUrl struct {
	Id           string
	Url          string
	Status       string
	FailedReason string
}

type ServiceUrlCreateParam struct {
	ServerId          string
	Service           string
	ServerAnsibleInfo ServerAnsibleInfo
}

type ServerAnsibleInfo struct {
	User string `json:"user"`
	IP   string `json:"ip"`
	Port int    `json:"port"`
	Name string `json:"name"`
}

func CreateSshInfo(ctx context.Context, serverId string) (string, error) {
	params := jsonutils.NewDict()
	params.Set("server_id", jsonutils.NewString(serverId))
	ret, _, err := RequestDevtoolSshInfo.Operation(OperCreate).Apply(ctx, "", params)
	if err != nil {
		return "", errors.Wrap(err, "unable to create devtool ssh info")
	}
	id, _ := ret.GetString("id")
	return id, nil
}

func GetSshInfo(ctx context.Context, sshInfoId string) (DevtoolSshInfo, error) {
	ret, _, err := RequestDevtoolSshInfo.Operation(OperGet).Apply(ctx, sshInfoId, nil)
	if err != nil {
		return DevtoolSshInfo{}, errors.Wrap(err, "unable to get devtool ssh info")
	}
	sshInfo := DevtoolSshInfo{}
	err = ret.Unmarshal(&ret)
	if err != nil {
		return DevtoolSshInfo{}, errors.Wrap(err, "unable to unmarshal ret of getting devtool ssh info")
	}
	return sshInfo, nil
}

func DeleteSshInfo(ctx context.Context, sshInfoId string) error {
	_, _, err := RequestDevtoolSshInfo.Operation(OperDelete).Apply(ctx, sshInfoId, nil)
	if err != nil {
		return errors.Wrap(err, "unable to delete devtool ssh info")
	}
	return nil
}

func CreateServiceUrl(ctx context.Context, params ServiceUrlCreateParam) (string, error) {
	data := jsonutils.Marshal(params).(*jsonutils.JSONDict)
	ret, _, err := RequestDevtoolServiceUrl.Operation(OperCreate).Apply(ctx, "", data)
	if err != nil {
		return "", err
	}
	id, _ := ret.GetString("id")
	return id, nil
}

func GetServiceUrl(ctx context.Context, serviceUrlId string) (DevtoolServiceUrl, error) {
	ret, _, err := RequestDevtoolServiceUrl.Operation(OperGet).Apply(ctx, serviceUrlId, nil)
	if err != nil {
		return DevtoolServiceUrl{}, nil
	}
	serviceUrl := DevtoolServiceUrl{}
	ret.Unmarshal(&serviceUrl)
	return serviceUrl, nil
}

func DeleteServiceUrl(ctx context.Context, serviceUrlId string) error {
	_, _, err := RequestDevtoolServiceUrl.Operation(OperDelete).Apply(ctx, serviceUrlId, nil)
	if err != nil {
		return errors.Wrap(err, "unable to delete serviceUrl")
	}
	return nil
}
