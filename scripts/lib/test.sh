#!/usr/bin/env bash

# shellcheck disable=SC2046
# shellcheck disable=SC2034
KRM_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../../" && pwd -P)"
source "${KRM_ROOT}/scripts/lib/util.sh"

krm::sudo pwd