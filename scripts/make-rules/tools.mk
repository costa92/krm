#  Makefile helper functions for tools

# 安装 k8s 代码生成器
CODE_GENERATOR_TOOLS= client-gen

# code-generator is a makefile target not a real tool.
CI_WORKFLOW_TOOLS := code-generator golangci-lint goimports wire
# The following tools may need to be installed manually
MANUAL_INSTALL_TOOLS := kafka

# unused tools in this project: gentool
OTHER_TOOLS := mockgen addlicense

.PHONY: tools.install
tools.install: _install.ci _install.other tools.print-manual-tool ## Install all tools.

.PHONY: tools.print-manual-tool
tools.print-manual-tool:
	@echo "===========> The following tools may need to be installed manually:"
	@echo $(MANUAL_INSTALL_TOOLS) | awk 'BEGIN{RS=" "} {printf("%15s%s\n","- ",$$0)}'


.PHONY: tools.verify.%
tools.verify.%: ## Verify a specified tool.
	@if ! which $* &>/dev/null; then $(MAKE) tools.install.$*; fi


.PHONY: _install.ci
_install.ci: $(addprefix tools.install., $(CI_WORKFLOW_TOOLS)) ## Install necessary tools used by CI/CD workflow.

# Install code-generator
.PHONY: _install.other
_install.other: $(addprefix tools.install., $(OTHER_TOOLS))


.PHONY: tools.install.%
tools.install.%: ## Install a specified tool.
	@echo "===========> Installing $*"
	@$(MAKE) _install.$*

# =========================
# install
.PHONY: tools.verify.code-generator
tools.verify.code-generator: $(addprefix _verify.code-generator., $(CODE_GENERATOR_TOOLS)) ## Verify a specified tool.


# 严重代码生成器
.PHONY: _verify.code-generator.%
_verify.code-generator.%:
	@if ! which $* &>/dev/null; then $(MAKE) tools.install.code-generator.$*; fi


.PHONY: _install.mockgen
_install.mockgen: ## Install mockgen.
	# 判断是否安装了 mockgen
	$(eval MOCKGEN_VERSION := $(if $(strip $(MOCKGEN_VERSION)),$(strip $(MOCKGEN_VERSION)),latest))
	@$(GO) install github.com/golang/mock/mockgen@$(MOCKGEN_VERSION); \


.PHONY: _install.code-generator
_install.code-generator: $(addprefix tools.install.code-generator., $(CODE_GENERATOR_TOOLS)) ## Install all necessary code-generator tools.


# 安装代码生成器
.PHONY: _install.code-generator.%
_install.code-generator.%: ## Install specified code-generator tool.
	$(eval CODE_GENERATOR_VERSION := $(if $(strip $(CODE_GENERATOR_VERSION)),$(strip $(CODE_GENERATOR_VERSION)),latest))
	@$(GO) install k8s.io/code-generator/cmd/$*@$(CODE_GENERATOR_VERSION)


.PHONY: _install.wire
_install.wire: ## Install wire.
	# 判断是否安装了 wire
	$(eval WIRE_VERSION := $(if $(strip $(WIRE_VERSION)),$(strip $(WIRE_VERSION)),latest))
	@$(GO) install github.com/google/wire/cmd/wire@$(WIRE_VERSION)

.PHONY: _install.swagger
_install.swagger:
	@$(GO) install github.com/go-swagger/go-swagger/cmd/swagger@$(GO_SWAGGER_VERSION)

.PHONY: _install.golangci-lint
_install.golangci-lint: ## Install golangci-lint.
   	# 判断 GOLANGCI_LINT_VERSION 是否为空 如果为空则设置为latest
	$(eval GOLANGCI_LINT_VERSION := $(if $(strip $(GOLANGCI_LINT_VERSION)),$(strip $(GOLANGCI_LINT_VERSION)),latest))
	echo "===========> Installing golangci-lint $(GOLANGCI_LINT_VERSION)"
	@$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)


.PHONY: _install.grpc
_install.grpc:
	@$(GO) install google.golang.org/protobuf/cmd/protoc-gen-go@$(PROTOC_GEN_GO_VERSION)
	@$(GO) install google.golang.org/grpc/cmd/protoc-gen-go-grpc@$(PROTOC_GEN_GO_GRPC_VERSION)
	@$(GO) install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@$(GRPC_GATEWAY_VERSION)
	@$(GO) install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@$(GRPC_GATEWAY_VERSION)
	@#$(SCRIPTS_DIR)/install-protoc.sh


.PHONY: _install.goimports
_install.goimports: ## Install goimports.
	@$(GO) install golang.org/x/tools/cmd/goimports@$(GO_IMPORTS_VERSION)

.PHONY: _install.kratos
_install.kratos: _install.grpc ## Install kratos toolkit, includes multiple protoc plugins.
	@$(GO) install github.com/joelanford/go-apidiff@$(GO_APIDIFF_VERSION)
	@$(GO) install github.com/envoyproxy/protoc-gen-validate@$(PROTOC_GEN_VALIDATE_VERSION)
	@$(GO) install github.com/google/gnostic/cmd/protoc-gen-openapi@$(PROTOC_GEN_OPENAPI_VERSION)
	@$(GO) install github.com/go-kratos/kratos/cmd/kratos/v2@$(KRATOS_VERSION)
	@$(GO) install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@$(KRATOS_VERSION)
	@$(GO) install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@$(KRATOS_VERSION)

	@#$(SCRIPTS_DIR)/add-completion.sh kratos bash

.PHONY: _install.buf
_install.buf: ## Install buf command line tool.
	@$(GO) install github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION)

# make tools.install.kratos
.PHONY: _install.addlicense
_install.addlicense: ## Install addlicense.
	# 判断 ADDLICENSE_VERSION 是否为空 如果为空则设置为latest
	$(eval ADDLICENSE_VERSION := $(if $(strip $(ADDLICENSE_VERSION)),$(strip $(ADDLICENSE_VERSION)),latest))
	@$(GO) install github.com/superproj/addlicense@$(ADDLICENSE_VERSION)

.PHONY: _install.gofumpt
_install.gofumpt:
	@$(GO) install mvdan.cc/gofumpt@$(GO_FUMPT_VERSION)
