# AnsiblePlaybookTemplate

AnsiblePlaybookTemplate 本质上代表 ansible 中的 playbook，它没有对应的 OneCloud 的资源。

它应该与 [AnsiblePlaybook](./ansibleplaybook.md) 搭配使用，之所以叫做 Template，是因为 AnsiblePlaybook 需要指定主机，而 AnsiblePlaybookTemplate 并不需要，所以可以通过 AnsiblePlaybook 将 AnsiblePlaybookTemplate 复用给若干主机。

## 创建 AnsiblePlaybookTemplate

在创建 AnisblePlaybookTemplate 之前，你需要掌握 [ansible playbook](https://ansible-tran.readthedocs.io/en/latest/docs/playbooks_intro.html) 的相关知识，尤其是关于其模块化`role`的知识。

下面是一个创建 AnsiblePlaybookTemplate 的例子:

```yaml
apiVersion: onecloud.yunion.io/v1
kind: AnsiblePlaybookTemplate
metadata:
  name: jenkins
spec:
  playbook: |
    - hosts: all
      become: true

      roles:
        - role: geerlingguy.java
        - role: ansible-role-jenkins
  requirements: |
    - src: geerlingguy.java
    - src: https://github.com/rainzm/ansible-role-jenkins
  vars:
    - name: jenkins_hostname
      default: localhost
    - name: jenkins_http_port
      default: 8080
    - name: jenkins_admin_username
      default: admin
    - name: jenkins_admin_password
      default: admin
```

`playbook`就是 ansible playbook 的内容，`requirements`主要是`playbook`中对其他`role`的引用信息。

注意到这两个 field 后面的`|`，这是`yaml`语法的一部分，这表明下面的内容是保留格式的字符串，所以实际上，`playbook`以及`requirements`是字符串。

`playbook`内容中的`hosts`的取值应该一直是"all"。

`vars`是此 AnsiblePlaybookTemplate 对外提供的参数。

详细文档请参考 [AnsiblePlaybookTemplateSpec](../api/docs.md#onecloud.yunion.io/v1.AnsiblePlaybookTemplateSpec)。

总之，如果将上面的 AnsiblePlaybookTemplate 通过 AnsiblePlaybook 应用到 VirtualMachine ，最终会在 OneCloud 上创建一个运行着 jenkins 应用的虚拟机实例。并且用户可以提前设置 hostname, port, admin username, admin password。

## 更新 AnsiblePlaybookTemplate

由于其没有对应的 OneCloud 资源，所以所有的字段均可以更改。如果有`Pending`或者`Waiting`状态的 AnsiblePlaybook 引用了它，也会马上生效。

## 删除 AnsiblePlaybookTemplate

注意，如果有处于`Waiting`的 AnsiblePlaybook 引用了它，将会一直`Waiting`。
