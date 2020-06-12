# VirtualMachines

VirtualMachine 代表 OneCloud 中的虚拟机实例。

创建一个 VirtualMachine 就会在 OneCloud 中创建一个符合 [VirtualMachineSpec](../api/docs.md#onecloud.yunion.io/v1.VirtualMachineSpec) 描述的虚拟机实例。

以下称 VirtualMachine 对应的 OneCloud 虚拟机实例为 OneCloudVM。

对 VirtualMachine 执行的更新或者删除操作，也会应用到 OneCloudVM 上。

## 创建 VirtualMachine

下面是一个创建 VirtualMahine 的例子, 它将会创建一个2C4G，系统盘大小40G，系统镜像为`CentOS-7.6.1810-20190430.qcow2`，数据盘大小为20G，属于`Default`域`lizexi`项目的，名字为`vm-example`的 kvm 虚拟机。

```yaml
apiVersion: onecloud.yunion.io/v1
kind: VirtualMachine
metadata:
  name: vm-example
spec:
  description: k8s create
  vmConfig:
    hypervisor: kvm
    vcpuCount: 2
    vmemSizeGB: 4
    rootDisk:
      image: CentOS-7.6.1810-20190430.qcow2
      sizeGB: 40
    dataDisks:
      - sizeGB: 20
  projectConfig:
    project: lizexi
    projectDomain: Default
```
`spec`还支持更多的属性，比如`NewEip`（自动创建并绑定EIP），`secgroup`（安全组）。

详细文档请参考 [VirtualMachineSpec](../api/docs.md#onecloud.yunion.io/v1.VirtualMachineSpec)，其中标记`Optional`的 field 是可选的，标记`Required`的 field 是必需的。

许多 field 都有自己的验证（待做）。

## 更新 VirtualMachine

目前只有部分 field 是允许更新的（严格来说，都可以更新，只不过只有部分 field 的更新有实际作用）。

详细文档请参考 [VirtualMachineSpec](../api/docs.md#onecloud.yunion.io/v1.VirtualMachineSpec)，其中标记`AllowUpdate`的 field 是允许更新的。

## 删除 VirtualMachine

删除操作需要等待，因为 controller 需要做清理工作。

## VirtualMachine Status

### Phase 

`Phase`反映了 VirtualMachine 当前所处的状态（阶段），`Reason`反应了其原因。

1. `Pending`: 中间状态，如 OneCloudVM 此时正在创建中、调整配置中或者删除中等等。

2. `Ready`: 关机状态。

3. `Running`: 正常运行状态。

4. `Failed`: 异常状态，OneCloudVM 会被删除并重新创建。

5. `Unkown`: 未知状态，会不断的尝试取同步 OneCloudVM。

6. `Invalid`: 无效，表明当前的 VirtualMachine 没有也不会有 OneCloudVM，处于此种状态下的 VirtualMachine 没有意义，应该被删除。一个典型的场景就是，尝试创建 OneCloudVM 若干次后失败。

### ExternalInfo

`ExternalInfo`存储了 OneCloudVM 的一些信息，包括通用的`Id`，`Status`。

此外还存储了 OneCloudVM 想对于其他 OneCloud 资源特有的`Eip`，`Ips`。

更多请参考文档 [VirtualMachineStatus](../api/docs.md#onecloud.yunion.io/v1.VirtualMachineStatus)。




