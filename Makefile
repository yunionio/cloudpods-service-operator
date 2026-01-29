# Env init
ROOT_DIR := $(CURDIR)
BUILD_DIR := $(ROOT_DIR)/_output
REGISTRY ?= registry.cn-beijing.aliyuncs.com/yunionio
VERSION ?= $(shell git describe --exact-match 2> /dev/null || \
                git describe --match=$(git rev-parse --short=8 HEAD) --always --dirty --abbrev=8)


# Produce CRDs that work back to Kubernetes 1.11 (no version conversion)
CRD_OPTIONS ?= "crd:trivialVersions=true"

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

GOPROXY ?= direct

all: manager

# Run tests
test: generate fmt vet manifests
	go test ./... -coverprofile cover.out

# Build manager binary, simple version
manager:
	go build -mod vendor -o _output/bin/manager main.go

# Build manager binary, professional version
manager-pro: generate fmt vet
	go build -mod vendor -o _output/bin/manager main.go

# Run against the configured Kubernetes cluster in ~/.kube/config
run: generate fmt vet manifests
	go run ./main.go

# Install CRDs into a cluster
install: manifests
	kustomize build config/crd | kubectl apply -f -

# Uninstall CRDs from a cluster
uninstall: manifests
	kustomize build config/crd | kubectl delete -f -

# Deploy controller in the configured Kubernetes cluster in ~/.kube/config
deploy: manifests
	cd config/manager && kustomize edit set image controller=$(REGISTRY)/onecloud-service-operator:$(VERSION)
	kustomize build config/default | kubectl apply -f -

# Generate manifests e.g. CRD, RBAC etc.
manifests: controller-gen
	$(CONTROLLER_GEN) $(CRD_OPTIONS) rbac:roleName=manager-role webhook paths="./..." output:crd:artifacts:config=config/crd/bases

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...

# Generate code
generate: controller-gen
	GOPROXY=$(GOPROXY) GONOSUMDB=yunion.io/x $(CONTROLLER_GEN) object:headerFile="hack/boilerplate.go.txt" paths="./..."

# Generate doc
generate-doc:
	./scripts/gen-doc.sh

# Build the docker image
docker-build: test
	docker build . -t $(REGISTRY)/onecloud-service-operator:$(VERSION)

# Push the docker image
docker-push:
	docker push $(REGISTRY)/onecloud-service-operator:$(VERSION)

base-image:
	docker buildx build --platform linux/arm64,linux/amd64,linux/riscv64 --push -t $(REGISTRY)/onecloud-service-operator-base:v0.0.2 -f ./Dockerfile.base .

# Simple operator for build and push image in auto build env
image: generate image-only

image-only:
	DOCKER_DIR=${CURDIR} PUSH=true DEBUG=${DEBUG} \
	REGISTRY=${REGISTRY} TAG=${VERSION} ARCH=${ARCH} \
	${CURDIR}/scripts/docker_push.sh manager

# find or download controller-gen
# download controller-gen if necessary
controller-gen:
ifeq (, $(shell which controller-gen))
	@{ \
	set -e ;\
	CONTROLLER_GEN_TMP_DIR=$$(mktemp -d) ;\
	cd $$CONTROLLER_GEN_TMP_DIR ;\
	GOPROXY=$(GOPROXY) GONOSUMDB=yunion.io/x go mod init tmp ;\
	GOPROXY=$(GOPROXY) GONOSUMDB=yunion.io/x go install sigs.k8s.io/controller-tools/cmd/controller-gen@v0.2.5 ;\
	rm -fr $$CONTROLLER_GEN_TMP_DIR ;\
	}
CONTROLLER_GEN=$(GOBIN)/controller-gen
else
CONTROLLER_GEN=$(shell which controller-gen)
endif
