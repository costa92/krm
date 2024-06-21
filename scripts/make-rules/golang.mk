
GO:=go
# Minimum version of Go required
GO_MINIMUM_VERSION:=1.22

# Check if Go is installed
GOPATH := $(shell go env GOPATH)
ifeq ($(origin GOPATH), undefined)
  GOBIN := $(GOPATH)/bin
endif

# 打印 KRM_ROOT
# $(info KRM_ROOT: $(KRM_ROOT))
# zh: 获取所有的cmd目录
CMD_DIRS := $(wildcard $(KRM_ROOT)/cmd/*)
# 打印 CMD_DIRS
#@ $(info CMD_DIRS: $(CMD_DIRS))
# Filter out directories without Go files, as these directories cannot be compiled.
# zh: 过滤掉没有Go文件的目录，因为这些目录不能被编译
COMMANDS := $(filter-out $(wildcard %.md), $(foreach dir, $(CMD_DIRS), $(if $(wildcard $(dir)/*.go), $(dir),)))
BINS ?= $(foreach cmd,${COMMANDS},$(notdir ${cmd}))

ifeq (${COMMANDS},)  #  如果COMMANDS为空
  $(error Could not determine COMMANDS, set CMD_DIRS or run in source dir)
endif

ifeq (${BINS},) # 如果BINS为空
  $(error Could not determine BINS, set CMD_DIRS or run in source dir)
endif

# make go.build apiserver linux_amd64
.PHONY: go.build.verify
go.build.verify: ## Verify supported go versions.
ifneq ($(shell $(GO) version|awk -v min=$(GO_MINIMUM_VERSION) '{gsub(/go/,"",$$3);if($$3 >= min){print 0}else{print 1}}'), 0)
	$(error unsupported go version. Please install a go version which is greater than or equal to '$(GO_MINIMUM_VERSION)')
endif

#
.PHONY: go.build.%
go.build.%: ## Build specified applications with platform, os and arch.
	$(eval COMMAND := $(word 2,$(subst ., ,$*)))
	$(eval PLATFORM := $(word 1,$(subst ., ,$*)))
	$(eval OS := $(word 1,$(subst _, ,$(PLATFORM))))
	$(eval ARCH := $(word 2,$(subst _, ,$(PLATFORM))))
	#@KRM_GIT_VERSION=$(VERSION) $(SCRIPTS_DIR)/build.sh $(COMMAND) $(PLATFORM)
	@if grep -q "func main()" $(KRM_ROOT)/cmd/$(COMMAND)/*.go &>/dev/null; then \
		echo "===========> Building binary $(COMMAND) $(VERSION) for $(OS) $(ARCH)" ; \
		CGO_ENABLED=0 GOOS=$(OS) GOARCH=$(ARCH) $(GO) build $(GO_BUILD_FLAGS) \
		-o $(OUTPUT_DIR)/platforms/$(OS)/$(ARCH)/$(COMMAND)$(GO_OUT_EXT) $(PRJ_SRC_PATH)/cmd/$(COMMAND) ; \
	fi


.PHONY: go.version
go.version: ## Show go version.
	@$(GO) version

# 打印 GOBIN
#$(info GOBIN: $(GOBIN))
# 打印 PLATFORM
#$(info PLATFORM: $(PLATFORM))

# make go.build BINS=krm-apiserver  # 单独编译 krm-apiserver
# make go.build
.PHONY: go.build # Build all applications.
go.build: $(addprefix go.build., $(addprefix $(PLATFORM)., $(BINS))) ## Build all applications.

.PHONY: go.build.multiarch
go.build.multiarch: $(foreach p,$(PLATFORMS),$(addprefix go.build., $(addprefix $(p)., $(BINS)))) ## Build all applications with all supported arch.
