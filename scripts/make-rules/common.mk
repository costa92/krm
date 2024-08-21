#
# These are the common rules for all the makefiles in the project.
#

# include the common make file
ifeq ($(origin KRM_ROOT),undefined)
KRM_ROOT :=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
endif


include $(KRM_ROOT)/scripts/make-rules/common-versions.mk


# It's necessary to set this because some environments don't link sh -> bash.
# zh 有些环境不会将sh链接到bash，所以需要设置这个变量
SHELL := /usr/bin/env bash -o errexit -o pipefail +o nounset
.SHELLFLAGS = -ec

# It's necessary to set the errexit flags for the bash shell.
# zh: 为bash shell设置errexit标志是必要的
export SHELLOPTS := errexit


# ====================================================
# Build Option
# zh: 构建选项
PRJ_SRC_PATH :=github.com/costa92/krm

# zh: 用于将逗号分隔的字符串转换为空格分隔的字符串
COMMA := ,
SPACE :=
SPACE +=


# ===================
# Output directory
ifeq ($(origin OUTPUT_DIR),undefined)
OUTPUT_DIR := $(KRM_ROOT)/_output
$(shell mkdir -p $(OUTPUT_DIR))
endif

#  ===================
# Local binary directory
ifeq ($(origin TMP_DIR),undefined)
TMP_DIR := $(OUTPUT_DIR)/tmp
$(shell mkdir -p $(TMP_DIR))
endif



# set the version number. you should not need to do this
# for the majority of scenarios.
# zh: 设置版本号，大多数情况下不需要设置
ifeq ($(origin VERSION), undefined)
# Current version of the project.
  VERSION := $(shell git describe --tags --always --match='v*')
  ifneq (,$(shell git status --porcelain 2>/dev/null))
    VERSION := $(VERSION)-dirty
  endif
endif

# ===========
# golang

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
GOPATH ?= $(shell go env GOPATH)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

# 构建选项
# The OS must be linux when building docker images
# PLATFORMS ?= linux_amd64 linux_arm64
# The OS can be linux/windows/darwin when building binaries
PLATFORMS ?= darwin_amd64 windows_amd64 linux_amd64 linux_arm64

# Set a specific PLATFORM
ifeq ($(origin PLATFORM), undefined)
	#	打印 PLATFORM
	# Use the default platform when building images
	ifeq ($(origin GOOS), undefined)
		GOOS := $(shell go env GOOS)
	endif
	ifeq ($(origin GOARCH), undefined)
		GOARCH := $(shell go env GOARCH)
	endif
	PLATFORM := $(GOOS)_$(GOARCH)

	# 判断 GOOS 是否是 darwin
	ifeq ($(GOOS),darwin)
		# Use windows as the default OS when building images
		IMAGE_PLAT := darwin_$(GOARCH)
	else
		# Use linux as the default OS when building images
    	IMAGE_PLAT := linux_$(GOARCH)
	endif
else
	GOOS := $(word 1, $(subst _, ,$(PLATFORM)))
	GOARCH := $(word 2, $(subst _, ,$(PLATFORM)))
	IMAGE_PLAT := $(PLATFORM)
endif
#$(info IMAGE_PLAT: $(IMAGE_PLAT))

# =====================================================
# Makefile settings
#
# We don't need make's built-in rules.
# zh: 我们不需要make的内置规则
# 执行命令： V=1 make go.build  可以打印出所有的命令
MAKEFLAGS += --no-builtin-rules
ifeq ($(V),1)
  # 使用 `$(MAKECMDGOALS)` 打印警告消息，显示 Makefile 的目标。`$(MAKECMDGOALS)` 是一个特殊变量，包含在命令行上指定的目标。
  $(warning ***** starting Makefile for goal(s) "$(MAKECMDGOALS)")
  # 打印当前日期和时间的警告消息。
  $(warning ***** $(shell date))
else
  # If we're not debugging the Makefile, don't echo recipes.]
  MAKEFLAGS += -s --no-print-directory
endif

# =====================================================
# Linux command settings for the Makefile
# zh: Makefile 的 Linux 命令设置
FIND := find . ! -path './third_party/*' ! -path './vendor/*' ! -path './.git/*' ! -path './.idea/*' ! -path './_output/*'
XARGS := xargs --no-run-if-empty


# Helper function to get dependency version from go.mod
get_go_version = $(shell go list -m $1 | awk '{print $$2}')
define go_install
$(info ===========> Installing $(1)@$(2))
$(GO) install $(1)@$(2)
endef





SCRIPTS_DIR=$(KRM_ROOT)/scripts
GENERATED_DOCKERFILE_DIR=$(KRM_ROOT)/build/docker

# Kubernetes releated variables.
## Metadata for driving the build lives here.
META_DIR := $(KRM_ROOT)/.make
GENERATED_FILE_PREFIX := zz_generated.
EXTRA_GENERATE_PKG := k8s.io/api/core/v1
# This controls the verbosity of the build. Higher numbers mean more output.
# 这控制了构建的详细程度。数字越高，输出越多。
KUBE_VERBOSE ?= 1

APIROOT ?= $(KRM_ROOT)/pkg/api
APISROOT ?= $(KRM_ROOT)/pkg/apis
