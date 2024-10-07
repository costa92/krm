#!/usr/bin/env bash

# Copyright 2024 costa <costa9293@gmail.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file. The original repo for
# this file is https://github.com/costa92/krm.

# echo 密码 | sudo -S shell命令
function krm::sudo() {
    echo "${LINUX_PASSWORD}" | sudo -S $1
}

# host_os 获取当前系统的操作系统
function krm::util::host_os() {
  local host_os
  case "$(uname -s)" in
    Darwin)
      host_os=darwin
      ;;
    Linux)
      host_os=linux
      ;;
    *)
      onex::log::error "Unsupported host OS.  Must be Linux or Mac OS X."
      exit 1
      ;;
  esac
  echo "${host_os}"
}

function krm::util::host_os() {
  local host_os
  case "$(uname -s)" in
    Darwin)
      host_os=darwin
      ;;
    Linux)
      host_os=linux
      ;;
    *)
      onex::log::error "Unsupported host OS.  Must be Linux or Mac OS X."
      exit 1
      ;;
  esac
  echo "${host_os}"
}

function krm::util::host_arch() {
  local host_arch
  case "$(uname -m)" in
    x86_64*)
      host_arch=amd64
      ;;
    i?86_64*)
      host_arch=amd64
      ;;
    amd64*)
      host_arch=amd64
      ;;
    aarch64*)
      host_arch=arm64
      ;;
    arm64*)
      host_arch=arm64
      ;;
    arm*)
      host_arch=arm
      ;;
    i?86*)
      host_arch=x86
      ;;
    s390x*)
      host_arch=s390x
      ;;
    ppc64le*)
      host_arch=ppc64le
      ;;
    *)
      krm::log::error "Unsupported host arch. Must be x86_64, 386, arm, arm64, s390x or ppc64le."
      exit 1
      ;;
  esac
  echo "${host_arch}"
}

function krm::util::download_file() {
  local -r url=$1
  local -r destination_file=$2

  rm "${destination_file}" 2&> /dev/null || true

  for i in $(seq 5)
  do
    if ! curl -fsSL --retry 3 --keepalive-time 2 "${url}" -o "${destination_file}"; then
      echo "Downloading ${url} failed. $((5-i)) retries left."
      sleep 1
    else
      echo "Downloading ${url} succeed"
      return 0
    fi
  done
  return 1
}


# looks for $1 in well-known output locations for the host platform
# $KRM_ROOT must be set
function krm::util::find-binary() {
  krm::util::find-binary-for-platform "$1" "$(krm::util::host_platform)"
}


function krm::util::host_platform() {
  echo "$(krm::util::host_os)/$(krm::util::host_arch)"
}

# looks for $1 in well-known output locations for the platform ($2)
# $KRM_ROOT must be set
function krm::util::find-binary-for-platform() {
  local -r lookfor="$1"
  local -r platform="$2"
  local locations=(
    "${KRM_ROOT}/_output/bin/${lookfor}"
    "${KRM_ROOT}/_output/dockerized/bin/${platform}/${lookfor}"
    "${KRM_ROOT}/_output/local/bin/${platform}/${lookfor}"
    "${KRM_ROOT}/platforms/${platform}/${lookfor}"
  )

  # if we're looking for the host platform, add local non-platform-qualified search paths
  if [[ "${platform}" = "$(krm::util::host_platform)" ]]; then
    locations+=(
      "${KRM_ROOT}/_output/local/go/bin/${lookfor}"
      "${KRM_ROOT}/_output/dockerized/go/bin/${lookfor}"
    );
  fi

  # looks for $1 in the $PATH
  if which "${lookfor}" >/dev/null; then
    local -r local_bin="$(which "${lookfor}")"
    locations+=( "${local_bin}"  );
  fi

  # List most recently-updated location.
  local -r bin=$( (ls -t "${locations[@]}" 2>/dev/null || true) | head -1 )

  if [[ -z "${bin}" ]]; then
    krm::log::error "Failed to find binary ${lookfor} for platform ${platform}"
    return 1
  fi

  echo -n "${bin}"
}

# krm::util::read-array
# Reads in stdin and adds it line by line to the array provided. This can be
# used instead of "mapfile -t", and is bash 3 compatible.  If the named array
# exists and is an array, it will be used.  Otherwise it will be unset and
# recreated.
#
# Assumed vars:
#   $1 (name of array to create/modify)
#
# Example usage:
#   krm::util::read-array files < <(ls -1)
#
# When in doubt:
#  $ W=abc         # a string
#  $ X=(a b c)     # an array
#  $ declare -A Y  # an associative array
#  $ unset Z       # not set at all
#  $ declare -p W X Y Z
#  declare -- W="abc"
#  declare -a X=([0]="a" [1]="b" [2]="c")
#  declare -A Y
#  bash: line 26: declare: Z: not found
#  $ onex::util::read-array W < <(echo -ne "1 1\n2 2\n3 3\n")
#  bash: W is defined but isn't an array
#  $ onex::util::read-array X < <(echo -ne "1 1\n2 2\n3 3\n")
#  $ onex::util::read-array Y < <(echo -ne "1 1\n2 2\n3 3\n")
#  bash: Y is defined but isn't an array
#  $ onex::util::read-array Z < <(echo -ne "1 1\n2 2\n3 3\n")
#  $ declare -p W X Y Z
#  declare -- W="abc"
#  declare -a X=([0]="1 1" [1]="2 2" [2]="3 3")
#  declare -A Y
#  declare -a Z=([0]="1 1" [1]="2 2" [2]="3 3")
function krm::util::read-array {
  if [[ -z "$1" ]]; then
    echo "usage: ${FUNCNAME[0]} <varname>" >&2
    return 1
  fi
  if [[ -n $(declare -p "$1" 2>/dev/null) ]]; then
    if ! declare -p "$1" 2>/dev/null | grep -q '^declare -a'; then
      echo "${FUNCNAME[0]}: $1 is defined but isn't an array" >&2
      return 2
    fi
  fi
  # shellcheck disable=SC2034 # this variable _is_ used
  local __read_array_i=0
  while IFS= read -r "$1[__read_array_i++]"; do :; done
  if ! eval "[[ \${$1[--__read_array_i]} ]]"; then
    unset "$1[__read_array_i]" # ensures last element isn't empty
  fi
}

# Takes a group/version and returns the path to its location on disk, sans
# "pkg". E.g.:
# * default behavior: extensions/v1beta1 -> apis/extensions/v1beta1
# * default behavior for only a group: experimental -> apis/experimental
# * Special handling for empty group: v1 -> api/v1, unversioned -> api/unversioned
# * Special handling for groups suffixed with ".k8s.io": foo.k8s.io/v1 -> apis/foo/v1
# * Very special handling for when both group and version are "": / -> api
#
# $KRM_ROOT must be set.
# UPDATEME: When add new api group.
function krm::util::group-version-to-pkg-path() {
  local group_version="$1"

  # Special cases first.
  # TODO(lavalamp): Simplify this by moving pkg/api/v1 and splitting pkg/api,
  # moving the results to pkg/apis/api.
  case "${group_version}" in
    # both group and version are "", this occurs when we generate deep copies for internal objects of the legacy v1 API.
    __internal)
      echo "pkg/apis/core"
      ;;
    #core/v1)
      #echo "${KRM_ROOT}/pkg/apis/core/v1"
      #;;
    apps/v1beta1)
      echo "${KRM_ROOT}/pkg/apis/apps/v1beta1"
      ;;
    #coordination/v1)
      #echo "${KRM_ROOT}/pkg/apis/coordination/v1"
      #;;
    #apiextensions/v1)
      #echo "${KRM_ROOT}/pkg/apis/apiextensions/v1"
      #;;
    *)
      echo "pkg/apis/${group_version%__internal}"
      ;;
  esac
}



# Downloads cfssl/cfssljson/cfssl-certinfo into $1 directory if they do not already exist in PATH
#
# Assumed vars:
#   $1 (cfssl directory) (optional)
#
# Sets:
#  CFSSL_BIN: The path of the installed cfssl binary
#  CFSSLJSON_BIN: The path of the installed cfssljson binary
#  CFSSLCERTINFO_BIN: The path of the installed cfssl-certinfo binary
#

function krm::util::ensure-cfssl {
  host_os=$(krm::util::host_os)
  if [[ "${host_os}" == "linux" ]]; then
      krm::util::linux::ensure-cfssl "$1"
    exit 1
  fi
  if [[ "${host_os}" == "darwin" ]]; then
      krm::util::darwin::ensure-cfssl "$1"
      return 0
  fi

}

function krm::util::darwin::ensure-cfssl {
  if command -v cfssl &>/dev/null && command -v cfssljson &>/dev/null; then
    CFSSL_BIN=$(command -v cfssl)
    CFSSLJSON_BIN=$(command -v cfssljson)
    return 0
  fi
  popd > /dev/null || return 1
}

function krm::util::linux::ensure-cfssl {
  if command -v cfssl &>/dev/null && command -v cfssljson &>/dev/null && command -v cfssl-certinfo &>/dev/null; then
    CFSSL_BIN=$(command -v cfssl)
    CFSSLJSON_BIN=$(command -v cfssljson)
    CFSSLCERTINFO_BIN=$(command -v cfssl-certinfo)
    return 0
  fi

  host_arch=$(krm::util::host_arch)
  if [[ "${host_arch}" != "amd64" ]]; then
    echo "Cannot download cfssl on non-amd64 hosts and cfssl does not appear to be installed."
    echo "Please install cfssl, cfssljson and cfssl-certinfo and verify they are in \$PATH."
    echo "Hint: export PATH=\$PATH:\$GOPATH/bin; go get -u github.com/cloudflare/cfssl/cmd/..."
    exit 1
  fi

  # Create a temp dir for cfssl if no directory was given
  local cfssldir=${1:-}
  if [[ -z "${cfssldir}" ]]; then
    cfssldir="$HOME/bin"
  fi

  mkdir -p "${cfssldir}"
  pushd "${cfssldir}" > /dev/null || return 1

  echo "Unable to successfully run 'cfssl' from ${PATH}; downloading instead..."
  kernel=$(uname -s)
  case "${kernel}" in
    Linux)
      curl --retry 10 -L -o cfssl https://pkg.cfssl.org/R1.2/cfssl_linux-amd64
      curl --retry 10 -L -o cfssljson https://pkg.cfssl.org/R1.2/cfssljson_linux-amd64
      curl --retry 10 -L -o cfssl-certinfo https://pkg.cfssl.org/R1.2/cfssl-certinfo_linux-amd64
      ;;
    Darwin)
      curl --retry 10 -L -o cfssl https://pkg.cfssl.org/R1.2/cfssl_darwin-amd64
      curl --retry 10 -L -o cfssljson https://pkg.cfssl.org/R1.2/cfssljson_darwin-amd64
      curl --retry 10 -L -o cfssl-certinfo https://pkg.cfssl.org/R1.2/cfssl-certinfo_darwin-amd64
      ;;
    *)
      echo "Unknown, unsupported platform: ${kernel}." >&2
      echo "Supported platforms: Linux, Darwin." >&2
      exit 2
  esac

  chmod +x cfssl || true
  chmod +x cfssljson || true
  chmod +x cfssl-certinfo || true

  CFSSL_BIN="${cfssldir}/cfssl"
  CFSSLJSON_BIN="${cfssldir}/cfssljson"
  CFSSLCERTINFO_BIN="${cfssldir}/cfssl-certinfo"
  if [[ ! -x ${CFSSL_BIN} || ! -x ${CFSSLJSON_BIN} || ! -x ${CFSSLCERTINFO_BIN} ]]; then
    echo "Failed to download 'cfssl'."
    echo "Please install cfssl, cfssljson and cfssl-certinfo and verify they are in \$PATH."
    echo "Hint: export PATH=\$PATH:\$GOPATH/bin; go get -u github.com/cloudflare/cfssl/cmd/..."
    exit 1
  fi
  popd > /dev/null || return 1
}
