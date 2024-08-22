
GOLANGCI_LINT_VERSION := v1.59.0
GOLANGCI_LINT_VERSION := v1.55.2
GO_APIDIFF_VERSION ?= v0.8.2
PROTOC_GEN_OPENAPI_VERSION ?= v0.7.0
BUF_VERSION ?= v1.28.1
GO_SWAGGER_VERSION ?= v0.31.0

# call get_go_version,github.com/golang/mock  # 获取golang/mock的版本
MOCKGEN_VERSION ?= $(call get_go_version,github.com/golang/mock)
CODE_GENERATOR_VERSION ?= $(call get_go_version,k8s.io/code-generator)
WIRE_VERSION ?= $(call get_go_version,github.com/google/wire)
#PROTOC_GEN_GO_GRPC_VERSION ?= $(call get_go_version,google.golang.org/grpc)
PROTOC_GEN_GO_GRPC_VERSION ?= v1.5.1
PROTOC_GEN_GO_VERSION ?= $(call get_go_version,google.golang.org/protobuf)
GRPC_GATEWAY_VERSION ?= $(call get_go_version,github.com/grpc-ecosystem/grpc-gateway/v2)
PROTOC_GEN_VALIDATE_VERSION ?= $(call get_go_version,github.com/envoyproxy/protoc-gen-validate)
#KRATOS_VERSION ?= $(call get_go_version,github.com/go-kratos/kratos/v2)
KRATOS_VERSION ?= latest