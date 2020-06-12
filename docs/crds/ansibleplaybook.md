# AnsiblePlaybook

AnsiblePlaybook 代表 OneCloud 中的 AnsiblePlaybook 实例。

创建一个 AnsiblePlaybook 就会在 OneCloud 中创建一个对应的 AnsiblePlaybook 实例，以下简称
 OneCloudAP。

不同于 [VirtualMachine](./virtualmachine.md) 的是，当 OneCloudAP 处于`succeeded`时，也就是对应的 ansible playbook 已经成功在对应的 OneCloudVM 中执行成功，OneCloudAP 就会被删除，AnsiblePlaybook 的`phase`会被设置为`Finished`。 

对 AnsiblePlaybook 执行的更新或者删除操作，也会应用到 OneCloudAp 上。

## 创建 AnsiblePlaybook

在创建之前，你最好了解以下 [AnsiblePlaybookTemplate](./ansibleplaybooktemplate.md)。

下面是一个创建 AnsiblePlaybook 的例子：

```yaml
apiVersion: onecloud.yunion.io/v1
kind: AnsiblePlaybook
metadata:
  name: jenkins-vm
spec:
  playbookTemplateRef:
    name: jenkins
  inventory:
    - virtualMachine:
        name: vm-example
  maxRetryTimes: 2
```

上面是一个简单的例子，其中`playbookTemplateRef`引用了 AnisblePlaybookTemplate 。

### 使用`playbookTemplate`代替`playbookTemplateRef`

还有一个类似的 field 叫做`playbookTemplate`，它允许你直接在 AnsiblePlaybook 中描述一个 AnsiblePlaybookTemplate 而不用去真正的创建一个 AnsiblePlaybookTemplate，当然这样的坏处是缺乏可复用性。

下面是一个类似的例子：

```yaml
apiVersion: onecloud.yunion.io/v1
kind: AnsiblePlaybook
metadata:
  name: jenkins-vm
spec:
  playbookTemplate:
    playbook: |
      - hosts: all
        become: true

        roles:
          - role: geerlingguy.java
          - role: ansible-role-jenkins
    requirements: |
      - src: geerlingguy.java
      - src: https://github.com/rainzm/ansible-role-jenkins
  inventory:
    - virtualMachine:
        name: vm-example
  maxRetryTimes: 2
```

`invertory`描述了 AnsiblePlaybook 操作的 VirtualMachine。

### 使用`vars`

还有一个重要的 field 叫做`vars`，`vars`分为全局的以及针对各个主机的，下面是一个例子：

```yaml
apiVersion: onecloud.yunion.io/v1
kind: AnsiblePlaybook
metadata:
  name: jenkins-slave-vm
spec:
  playbookTemplateRef:
    name: jenkins-slave-tem
  inventory:
    - virtualMachine:
        name: jenkins-slave1
      vars:
        jenkins_slave_name:
          value: slave1
    - virtualMachine:
        name: jenkins-slave2
      vars:
        jenkins_slave_name:
          value: slave2
  vars:
    jenkins_master_hostname:
      reference:
        kind: VirtualMachine
        namespace: onecloud
        name: jenkins-master
        fieldPath: Status.ExternalInfo.Ips[0]
  maxRetryTimes: 2
```

`vars`的类型是`map[string]IntOrStringStore`，[IntOrStringStore](../api/docs.md#onecloud.yunion.io/v1.IntOrStringStore)包含两个 field: `value`以及`reference`。

`IntOrStringStore`的意义是用来存储 int 或者 string 类型的值，你可以直接在`value`指定，就像上面的：

```yaml
vars:
  jenkins_slave_name:
    value: slave2
```

你也可以，通过`reference`来引用其他 CRD 的 field 的值，比如：

```yaml
vars:
  jenkins_master_hostname:
    reference:
      kind: VirtualMachine
      namespace: onecloud
      name: jenkins-master
      fieldPath: Status.ExternalInfo.Ips[0]
```

它的意思是，引用 namespace 为 onecloud，name 为 jenkins-master 的 VirtualMachine 的 Status.ExternalInfo.Ips[0]，也就是它的第一个IP。

这是一个强大的功能，很多的 CRD 都会有类似的用法，这可以让我们把不同类型或者相同类型的资源连接起来，描述依赖关系。

## 更新 AnsiblePlaybook

不建议更新 AnsiblePlaybook，因为不是很有必要，AnsiblePlaybook 只是一次性的工作，完全可以删除并新建。

## 删除 AnsiblePlaybook 

如果 OneCloudAP 已经存在，删除操作将会先尝试暂停正在运行的 OneCloudAP，然后才开始执行删除操作。

## AnsiblePlaybook Status

### Phase

`Phase`反映了 AnsiblePlaybook 当前所处的状态（阶段），`Reason`反应了其原因。

1. `Waiting`: 等待状态，正在等待其所依赖的资源，比如等待依赖的 VirtualMachine 处于`Running`中。
2. `Pending`: 中间状态，如 OneCloudAP 正在创建，正在执行等等。
3. `Failed`: 异常状态，如 OneCloudAp 创建失败，执行失败；这种状态下，会去重新尝试。
4. `Unkown`: 未知状态，会不断的尝试取同步 OneCloudVM。
5. `Finished`: 完成状态，表示 OneCloudAP 成功执行完毕。
6. `Invalid`: 无效，表明当前的 AnsiblePlaybook 没有也不会有 OneCloudVM，处于此种状态下的 AnsiblePlaybook 没有意义，应该被删除。一个典型的场景就是，尝试创建 OneCloudAP 若干次后失败。

### ExternalInfo

`ExternalInfo`存储了 OneCloudAP 的一些信息，包括通用的`Id`，`Status`。

此外还存储了 OneCloudAP 的输出`Output`。

更多请参考文档 [AnsiblePlaybookStatus](../api/docs.md#onecloud.yunion.io/v1.AnsiblePlaybookStatus)。
