#!/usr/bin/env bash



KRM_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../../" && pwd -P)"
source "${KRM_ROOT}/hack/lib/init.sh"

krm::golang::setup_env