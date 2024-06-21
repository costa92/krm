
GOLANGCI_LINT_VERSION := v1.59.0
GOLANGCI_LINT_VERSION := v1.55.2

# call get_go_version,github.com/golang/mock  # 获取golang/mock的版本
MOCKGEN_VERSION ?= $(call get_go_version,github.com/golang/mock)
CODE_GENERATOR_VERSION ?= $(call get_go_version,k8s.io/code-generator)
WIRE_VERSION ?= $(call get_go_version,github.com/google/wire)
PROTOC_GEN_GO_GRPC_VERSION ?= $(call get_go_version,google.golang.org/grpc)
