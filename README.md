## OneCloud Service Operator

This repo contains the OneCloud Service Operator for Kubernetes(OSO).

An operator in Kubernetes is the combination of one or more [custom resources](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/) and [controllers](https://kubernetes.io/docs/reference/glossary/?fundamental=true#term-controller) managing said custom resources.

The OSO will allow containerized applications and Kubernetes users to create, update, delete and retrieve the status of objects in OneCloud services such as Server, Image, Network, Storage, etc. using the Kubernetes API, for example using Kubernetes manifests or `kubectl` plugins.

## Installation

```shell
# install all crds
make install

# build
make manager

# run locally
./manager --config xxx.conf

# build and push docker image
make image
```

## Documention

See [Documention](./docs/main.md)

## Contribution

See [Contribution](./CONTRIBUTING.md).

## License

Apache License 2.0. See [LICENSE](./LICENSE).
