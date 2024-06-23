#!/usr/bin/env bash

set -o errexit
set +o nounset
set -o pipefail



# Short-circuit if init.sh has already been sourced
[[ $(type -t krm::init::loaded) == function ]] && return 0

# Unset CDPATH so that path interpolation can work correctly
# https://github.com/minerrnetes/minerrnetes/issues/52255
unset CDPATH

# Default use go modules
export GO111MODULE=on

# The root of the build/dist directory
KRM_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd -P)"

KRM_OUTPUT_SUBPATH="${KRM_OUTPUT_SUBPATH:-_output}"
KRM_OUTPUT="${KRM_ROOT}/${KRM__OUTPUT_SUBPATH}"

source "${KRM_ROOT}/scripts/lib/util.sh"
source "${KRM_ROOT}/scripts/lib/logging.sh"
source "${KRM_ROOT}/scripts/lib/color.sh"



# Marker function to indicate init.sh has been fully sourced
krm::init::loaded() {
  return 0
}