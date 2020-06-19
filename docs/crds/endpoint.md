# Endpoint

Endpoint 代表 OneCloud 中特定属性（interface 为 'console'，service 为 'external-service'）的 Endpoint 实例。

创建一个 Endpoint 就会在 OneCloud 中创建一个符合 [EndpointSpec](../api/docs.md#onecloud.yunion.io/v1.EndpointSpec) 描述的 Endpoint 实例（以下简称 OneCloudEndpoint)。

对 Endpoint 执行的更新或者删除操作，也会应用到 OneCloudEndpoint 上。

## 创建 Endpoint 

下面是一个创建 Endpoint 的例子: 
```yaml
apiVersion: onecloud.yunion.io/v1
kind: Endpoint
metadata:
  name: endpoint-sample
spec:
  regionId: region0
  name: test
  url:
    host:
      reference:
        kind: VirtualMachine
        namespace: onecloud
        name: ng-nihao
        fieldPath: Status.ExternalInfo.Ips[0]
    port: 8081
    prefix: test
```

`url`具体可见 [URL](../api/docs.md#onecloud.yunion.io/v1.URL)

重点关注`url`中的`host`，可以直接通过`value`输入主机地址（字符串），或者通过`reference`来引用一个 VirtualMachine 的`ip`或者`eip`，具体用法可以参考上面的例子。

## 更新 Endpoint

所有的 field 都支持更新。
更新操作会触发 OneCloudEndpoint 的更新。

## EndpointStatus

### phase

`phase`反映了 Endpoint 当前所处的状态（阶段），`reason`反映了其原因。

1. `Ready`: 正常状态。

2. `Unkown`: 未知状态，会不断的尝试取同步 OneCloudEndpoint。

3. `Invalid`: 无效，表明当前的 Endpoint 没有也不会有 OneCloudEndpoint，处于此种状态下的 Endpoint 没有意义，应该被删除。一个典型的场景就是，尝试创建 OneCloudEndpoint 若干次后失败。

### tryTimes

尝试的次数。

### externalInfo

`externalInfo`存储了 OneCloudEndpoint 的一些信息，包括通用的`id`，`status`。

更多请参考文档 [EndpointStatus](../api/docs.md#onecloud.yunion.io/v1.EndpointStatus)。
