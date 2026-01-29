FROM registry.cn-beijing.aliyuncs.com/yunionio/onecloud-service-operator-base:v0.0.2
ENV TZ UTC
WORKDIR /
COPY config/crd/bases/ /etc/crds/
COPY _output/alpine-build/bin/manager .
