#!/usr/bin/env bash

# Copyright 2024 costa <costa9293@gmail.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file. The original repo for
# this file is https://github.com/costa92/krm.


# shellcheck disable=SC2046
# shellcheck disable=SC2034
KRM_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../../" && pwd -P)"
source "${KRM_ROOT}/hack/lib/util.sh"

krm::sudo pwd


# shellcheck disable=SC2046

