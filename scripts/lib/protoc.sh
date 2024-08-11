#!/usr/bin/env bash


set -o errexit
set -o nounset
set -o pipefail


# The root of the build/dist directory
KRM_ROOT=$(dirname "${BASH_SOURCE[0]}")/../..
source "${KRM_ROOT}/scripts/lib/init.sh"

# protocVersion is the version of protoc to install.
# shellcheck disable=SC2034
PROTOC_VERSION="3.19.4"

# Short-circuit if protoc.sh has already been sourced
[[ $(type -t krm::protoc::loaded) == function ]] && return 0



# Generates $1/api.pb.go from the protobuf file $1/api.proto
# and formats it correctly
# $1: Full path to the directory where the api.proto file is
function krm::protoc::generate_proto() {
  krm::golang::setup_env
  GO111MODULE=on GOPROXY=off go install k8s.io/code-generator/cmd/go-to-protobuf/protoc-gen-gogo

  krm::protoc::check_protoc

  local package=${1}
  krm::protoc::protoc "${package}"
  krm::protoc::format "${package}"
}


# Generates $1/api.pb.go from the protobuf file $1/api.proto
# $1: Full path to the directory where the api.proto file is
function krm::protoc::protoc() {
  local package=${1}
  gogopath=$(dirname "$(krm::util::find-binary "protoc-gen-gogo")")
  (
    cd "${package}"
    # This invocation of --gogo_out produces its output in the current
    # directory (despite gogo docs saying it would be source-relative, it
    # isn't).  The inputs to this function do not all have a common root, so
    # this works best for all inputs.
    PATH="${gogopath}:${PATH}" protoc \
      --proto_path="$(pwd -P)" \
      --proto_path="${KRM_ROOT}/vendor" \
      --proto_path="${KRM_ROOT}/third_party/google/protobuf" \
      --gogo_out=paths=source_relative,plugins=grpc:. \
      api.proto
  )
}



# Formats $1/api.pb.go, adds the boilerplate comments and run gofmt on it
# $1: Full path to the directory where the api.proto file is
function krm::protoc::format() {
  local package=${1}

  # Update boilerplate for the generated file.
  cat hack/boilerplate/boilerplate.generatego.txt "${package}/api.pb.go" > tmpfile && mv tmpfile "${package}/api.pb.go"

  # Run gofmt to clean up the generated code.
  krm::golang::verify_go_version
  gofmt -s -w "${package}/api.pb.go"
}


# Checks that the current protoc version matches the required version and
# exit 1 if it's not the case
function krm::protoc::check_protoc() {
  if [[ -z "$(/usr/bin/which protoc)" || "$(protoc --version)" != "libprotoc ${PROTOC_VERSION}"* ]]; then
    echo "Generating protobuf requires protoc ${PROTOC_VERSION}."
    echo "Run scripts/install-protoc.sh or download and install the"
    echo "platform-appropriate Protobuf package for your OS from"
    echo "https://github.com/protocolbuffers/protobuf/releases"
    return 1
  else
    echo "protoc v${PROTOC_VERSION} is already installed."
  fi
}

# Install protoc binary in $HOME/bin
function krm::protoc::install() {
  # run in a subshell to isolate caller from directory changes
  (
    local os
    local arch
    local download_folder
    local download_file
    echo "Installing protoc v${PROTOC_VERSION}..."
    os=$(krm::util::host_os)
    arch=$(krm::util::host_arch)
    download_folder="protoc-v${PROTOC_VERSION}-${os}-${arch}"
    download_file="${download_folder}.zip"
    cd "${HOME}/bin" || return 1
    if [[ ! -f protoc  ]]; then
      local url
      if [[ ${os} == "darwin" ]]; then
        # TODO: switch to universal binary when updating to 3.20+
        url="https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-osx-x86_64.zip"
      elif [[ ${os} == "linux" && ${arch} == "amd64" ]]; then
        url="https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip"
      elif [[ ${os} == "linux" && ${arch} == "arm64" ]]; then
        url="https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-aarch_64.zip"
      else
        krm::log::info "This install script does not support ${os}/${arch}"
        return 1
      fi
      krm::util::download_file "${url}" "${download_file}"
      unzip -o "${download_file}" -d "${download_folder}"
      mv "${download_folder}/bin/protoc" $HOME/bin/protoc
      chmod +rX $HOME/bin/protoc
      rm -rvf "${download_file}" "${download_folder}"
    fi
    krm::log::info "$HOME/bin/protoc v${PROTOC_VERSION} installed."
  )
}

# Marker function to indicate protoc.sh has been fully sourced
krm::protoc::loaded() {
  return 0
}