

## Resources

### GroupName

`onecloud.yunion.io`

### Resources

- [VirtualMachine](./crds/virtualmachine.md)
- [AnsiblePlaybookTemplate](./crds/ansibleplaybooktemplate.md) 
- [AnsiblePlaybook](./crds/ansibleplaybook.md)

## CRUD

资源的全名，是上述提到的资源的小写负数再加上 GroupName，比如`virtualmachines.onecloud.yunion.io`

下面以 VirtualMachine 为例子介绍，假设已经有了名字为`newvm`的 VirtualMachine 的 yaml 文件：vm.yaml。

在正式操作之前，请确保你的 Kubernetes 集群运行正常，并且 kubectl 能够正确连接集群。

### Create


1. 使用以下命令创建 VirtualMachine `newvm`：

```shell
kubectl apply -f vm.yaml
```

2. 使用以下命令来检查 VirtualMachine `newvm`是否创建好：

```shell
kubectl get virtualmachines.onecloud.yunion.io newvm
```

3. 使用以下命令来查看 VirtualMachine `newvm`的详细信息：

```shell
kubectl get virtualmachines.onecloud.yunion.io newvm -o yaml
```
### Update

使用以下命令可以编辑更新 VirtualMachine `newvm`：

```shell
kubectl edit virtualmachines.onecloud.yunion.io newvm
```

### Delete

使用以下命令可以删除 VirtualMachine `newvm`：

```shell
kubectl delete virtualmachines.onecloud.yunion.io newvm
```
