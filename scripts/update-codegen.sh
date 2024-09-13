#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# This tool wants a different default than usual.
KUBE_VERBOSE="${KUBE_VERBOSE:-1}"

KRM_ROOT=$(dirname "${BASH_SOURCE[0]}")/..

source "${KRM_ROOT}/scripts/lib/init.sh"
source "${KRM_ROOT}/scripts/lib/protoc.sh"

krm::golang::setup_env

DBG_CODEGEN="${DBG_CODEGEN:-0}"
GENERATED_FILE_PREFIX="${GENERATED_FILE_PREFIX:-zz_generated.}"
UPDATE_API_KNOWN_VIOLATIONS="${UPDATE_API_KNOWN_VIOLATIONS:-}"
API_KNOWN_VIOLATIONS_DIR="${API_KNOWN_VIOLATIONS_DIR:-"${KRM_ROOT}/api/api-rules"}"


OUT_DIR="_output"
KRM_MODULE_NAME="github.com/costa92/krm"
BOILERPLATE_FILENAME="${KRM_ROOT}/scripts/boilerplate/boilerplate.generatego.txt"
PLURAL_EXCEPTIONS="Endpoints:Endpoints,ResourceClaimParameters:ResourceClaimParameters,ResourceClassParameters:ResourceClassParameters"
OUTPUT_PKG="github.com/costa92/krm/pkg/generated"
EXTRA_GENERATE_PKG=(k8s.io/api/core/v1 k8s.io/api/coordination/v1 k8s.io/api/flowcontrol/v1 k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1)
APPLYCONFIG_PKG="${OUTPUT_PKG}/applyconfigurations"


# Any time we call sort, we want it in the same locale.
export LC_ALL="C"
# Work around for older grep tools which might have options we don't want.
unset GREP_OPTIONS

if [[ "${DBG_CODEGEN}" == 1 ]]; then
    krm::log::status "DBG: starting generated_files"
fi

# Generate a list of directories we don't want to play in.
# shellcheck disable=SC2034
DIRS_TO_AVOID=()
krm::util::read-array DIRS_TO_AVOID < <(
    git ls-files -cmo --exclude-standard -- ':!:vendor/*' ':(glob)*/**/go.work' \
        | while read -r F; do \
            echo ':!:'"$(dirname "${F}")"; \
        done
    )


function git_find() {
    # Similar to find but faster and easier to understand.  We want to include
    # modified and untracked files because this might be running against code
    # which is not tracked by git yet.
    git ls-files -cmo --exclude-standard ':!:vendor/*' "${DIRS_TO_AVOID[@]}" "$@"
}


function git_grep() {
    # We want to include modified and untracked files because this might be
    # running against code which is not tracked by git yet.
    # We need vendor exclusion added at the end since it has to be part of
    # the pathspecs which are specified last.
    git grep --untracked "$@" ':!:vendor/*' "${DIRS_TO_AVOID[@]}"
}

# Generate a list of all files that have a `+k8s:` comment-tag.  This will be
# used to derive lists of files/dirs for generation tools.
if [[ "${DBG_CODEGEN}" == 1 ]]; then
    krm::log::status "DBG: finding all +k8s: tags"
fi
ALL_K8S_TAG_FILES=()
krm::util::read-array ALL_K8S_TAG_FILES < <(
    git_grep -l \
        -e '^// *+k8s:'                `# match +k8s: tags` \
        -- \
        ':!:*/testdata/*'              `# not under any testdata` \
        ':(glob)**/*.go'               `# in any *.go file` \
    )
if [[ "${DBG_CODEGEN}" == 1 ]]; then
    krm::log::status "DBG: found ${#ALL_K8S_TAG_FILES[@]} +k8s: tagged files"
fi



#
# Code generation logic.
#

# protobuf generation
#
# Some of the later codegens depend on the results of this, so it needs to come
# first in the case of regenerating everything.
function codegen::protobuf() {
    # NOTE: All output from this script needs to be copied back to the calling
    # source tree.  This is managed in onex::build::copy_output in build/common.sh.
    # If the output set is changed update that function.
    echo 11
    local apis=()
    krm::util::read-array apis < <(
        git grep --untracked --null -l \
            -e '// +k8s:protobuf-gen=package' \
            -- \
            cmd pkg \
            | while read -r -d $'\0' F; do dirname "${F}"; done \
            | sed 's|^|github.com/costa92/krm/|;s|k8s.io/kubernetes/staging/src/||' \
            | sort -u)

    krm::log::status "Generating protobufs for ${#apis[@]} targets"
    if [[ "${DBG_CODEGEN}" == 1 ]]; then
        krm::log::status "DBG: generating protobufs for:"
        for dir in "${apis[@]}"; do
            krm::log::status "DBG:     $dir"
        done
    fi
    # NOTICE: must include k8s.io/api/core/v1, otherwise it will generate the message ObjectReference
    # in pkg/apis/apps/v1beta1/generated.proto, which will cause a compilation error when compiling
    # onex-apiserver: undefined: ObjectReference.
    apis+=("${EXTRA_GENERATE_PKG}")

    # Comment this out, otherwise it will delete some useful protobuf files.
    #git_find -z \
        #':(glob)**/generated.proto' \
        #':(glob)**/generated.pb.go' \
        #| xargs -0 rm -f

    if krm::protoc::check_protoc >/dev/null; then
      hack/_update-generated-protobuf-dockerized.sh "${apis[@]}"
    else
      krm::log::status "protoc ${PROTOC_VERSION} not found (can install with hack/install-protoc.sh); generating containerized..."
      build/run.sh hack/_update-generated-protobuf-dockerized.sh "${apis[@]}"
    fi

    # Fix `pkg/apis/apps/v1beta1/generated.pb.go:49:10: undefined: ObjectReference` compile errors
    # cp ${KRM_ROOT}/manifests/generated.pb.go.fix ${ONEX_ROOT}/pkg/apis/apps/v1beta1/generated.pb.go
}
