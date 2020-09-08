# Build the manager binary
# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM registry.cn-beijing.aliyuncs.com/yunionio/onecloud-base:v0.2
RUN echo http://dl-cdn.alpinelinux.org/alpine/edge/testing >>/etc/apk/repositories
RUN apk update && apk add kubectl && rm -rf /var/cache/apk/*

WORKDIR /
COPY config/crd/bases/ /etc/crds/
COPY _output/alpine-build/bin/manager .
