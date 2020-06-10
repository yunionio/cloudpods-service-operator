# Build the manager binary
FROM golang:1.13 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
# RUN go mod download

# Copy the go source
COPY main.go main.go
COPY api/ api/
COPY controllers/ controllers/
COPY pkg/ pkg/
COPY vendor/ vendor/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -mod vendor -a -o manager main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM registry.cn-beijing.aliyuncs.com/yunionio/onecloud-base:latest
WORKDIR /
COPY --from=builder /workspace/manager .

# Install CRDs yaml
COPY config/crd/bases/ /etc/crds/

# Install kubectl from Docker Hub.
COPY --from=lachlanevenson/k8s-kubectl:v1.16.9 /usr/local/bin/kubectl /usr/local/bin/kubectl

# USER nonroot:nonroot
# ENTRYPOINT ["kubectl"]
