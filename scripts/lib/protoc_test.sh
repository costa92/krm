#!/usr/bin/env bash


KRM_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../../" && pwd -P)"
source "${KRM_ROOT}/scripts/lib/init.sh"
source "${KRM_ROOT}/scripts/lib/protoc.sh"


#krm::protoc::install
#krm::protoc::check_protoc
krm::protoc::generate_proto "${KRM_ROOT}/pkg/api/usercenter/v1/"