## 开发环境搭建

本项目是基于[kubebuilder](https://github.com/kubernetes-sigs/kubebuilder)构建的，所以在开发之前请安装`kubebuilder`相关工具。

## 代码简介

`main.go`: 服务的入口，主要用来初始化各个 controller 以及 controllerManager。

`api/v1/`: v1 版本的 api。

`config/`: 主要是 crd 以及 controller manager 相关的 Kubernetes manifests 文件。

`controllers/`: 各个 controller 的实现代码。

`pkg/options/`: 参数相关的代码。

`pkg/provider/`: 对 OneCloud 各个 Service 以及资源操作的逻辑代码。

`pkg/util/`: 工具代码。

## 创建一个新的 CRD 以及 controller

```shell
kubebuilder create api --group onecloud --version v1 --kind AnsiblePlaybook
```

更加详细的操作请参考[kubebuilder book](https://kubebuilder.io/)。

## 生成代码

修改完代码之后，请执行`make manifests`以更新 manifests，执行`make generate`以更新自动生成的代码。

## 生成文档

修改完代码之后，请执行`make generate-doc`以更新文档。
