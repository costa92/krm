#!/usr/bin/env bash

# 设置脚本环境
set -o errexit
set -o nounset
set -o pipefail

# 定义常量
KRM_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
source "${KRM_ROOT}/hack/lib/init.sh"
source "${KRM_ROOT}/hack/lib/protoc.sh"
cd "${KRM_ROOT}"

krm::golang::setup_env

DBG_CODEGEN="${DBG_CODEGEN:-0}"
GENERATED_FILE_PREFIX="${GENERATED_FILE_PREFIX:-zz_generated.}"
API_KNOWN_VIOLATIONS_DIR="${API_KNOWN_VIOLATIONS_DIR:-"${KRM_ROOT}/api/api-rules"}"

OUT_DIR="_output"
BOILERPLATE_FILENAME="hack/boilerplate/boilerplate.go.txt"
KRM_MODULE_NAME="github.com/costa92/krm"
PLURAL_EXCEPTIONS="Endpoints:Endpoints,ResourceClaimParameters:ResourceClaimParameters,ResourceClassParameters:ResourceClassParameters"
OUTPUT_PKG="github.com/costa92/krm/pkg/generated"
APPLYCONFIG_PKG="${OUTPUT_PKG}/applyconfigurations"
# Any time we call sort, we want it in the same locale.
export LC_ALL="C"


# Work around for older grep tools which might have options we don't want.
unset GREP_OPTIONS

if [[ "${DBG_CODEGEN}" == 1 ]]; then
    krm::log::status "DBG: starting generated_files"
fi

# Generate a list of directories we don't want to play in.
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

# Deep-copy generation
#
# Any package that wants deep-copy functions generated must include a
# comment-tag in column 0 of one file of the form:
#     // +k8s:deepcopy-gen=<VALUE>
#
# The <VALUE> may be one of:
#     generate: generate deep-copy functions into the package
#     register: generate deep-copy functions and register them with a
#               scheme
function codegen::deepcopy() {
    # Build the tool.
    GOPROXY=off go install \
        k8s.io/code-generator/cmd/deepcopy-gen

    # The result file, in each pkg, of deep-copy generation.
    local output_file="${GENERATED_FILE_PREFIX}deepcopy.go"
    echo "output_file: ${output_file}"
    # Find all the directories that request deep-copy generation.
    if [[ "${DBG_CODEGEN}" == 1 ]]; then
        krm::log::status "DBG: finding all +k8s:deepcopy-gen tags"
    fi
    local tag_dirs=()
    krm::util::read-array tag_dirs < <( \
        grep -l --null '+k8s:deepcopy-gen=' "${ALL_K8S_TAG_FILES[@]}" \
            | while read -r -d $'\0' F; do dirname "${F}"; done \
            | sort -u)
    if [[ "${DBG_CODEGEN}" == 1 ]]; then
        krm::log::status "DBG: found ${#tag_dirs[@]} +k8s:deepcopy-gen tagged dirs"
    fi

    local tag_pkgs=()
    for dir in "${tag_dirs[@]}"; do
        tag_pkgs+=("./$dir")
    done

    krm::log::status "Generating deepcopy code for ${#tag_pkgs[@]} targets"
    if [[ "${DBG_CODEGEN}" == 1 ]]; then
        krm::log::status "DBG: running deepcopy-gen for:"
        for dir in "${tag_dirs[@]}"; do
            krm::log::status "DBG:     $dir"
        done
    fi

    git_find -z ':(glob)**'/"${output_file}" | xargs -0 rm -f

    deepcopy-gen \
        -v "${KUBE_VERBOSE}" \
        --go-header-file "${BOILERPLATE_FILENAME}" \
        --output-file "${output_file}" \
        --bounding-dirs "${KRM_MODULE_NAME},k8s.io/api" \
        "${tag_pkgs[@]}" \
        "$@"

    if [[ "${DBG_CODEGEN}" == 1 ]]; then
        krm::log::status "Generated deepcopy code"
    fi
}



# OpenAPI generation
#
# Any package that wants open-api functions generated must include a
# comment-tag in column 0 of one file of the form:
#     // +k8s:openapi-gen=true
function codegen::openapi() {
    # Build the tool.
    GOPROXY=off go install \
        k8s.io/kube-openapi/cmd/openapi-gen

    # The result file, in each pkg, of open-api generation.
    local output_file="${GENERATED_FILE_PREFIX}openapi.go"

    local output_dir="pkg/generated/openapi"
    local output_pkg="k8s.io/kubernetes/${output_dir}"
    local known_violations_file="${API_KNOWN_VIOLATIONS_DIR}/violation_exceptions.list"

    local report_file="${OUT_DIR}/api_violations.report"
    # When UPDATE_API_KNOWN_VIOLATIONS is set to be true, let the generator to write
    # updated API violations to the known API violation exceptions list.
    if [[ "${UPDATE_API_KNOWN_VIOLATIONS}" == true ]]; then
        report_file="${known_violations_file}"
    fi

    if [[ "${DBG_CODEGEN}" == 1 ]]; then
        krm::log::status "DBG: finding all +k8s:openapi-gen tags"
    fi

    local tag_files=()
    krm::util::read-array tag_files < <(
        k8s_tag_files_except \
            staging/src/k8s.io/code-generator \
            staging/src/k8s.io/sample-apiserver
        )

    local tag_dirs=()
    krm::util::read-array tag_dirs < <(
        grep -l --null '+k8s:openapi-gen=' "${tag_files[@]}" \
            | while read -r -d $'\0' F; do dirname "${F}"; done \
            | sort -u)

    if [[ "${DBG_CODEGEN}" == 1 ]]; then
        krm::log::status "DBG: found ${#tag_dirs[@]} +k8s:openapi-gen tagged dirs"
    fi

    local tag_pkgs=()
    for dir in "${tag_dirs[@]}"; do
        tag_pkgs+=("./$dir")
    done

    krm::log::status "Generating openapi code"
    if [[ "${DBG_CODEGEN}" == 1 ]]; then
        krm::log::status "DBG: running openapi-gen for:"
        for dir in "${tag_dirs[@]}"; do
            krm::log::status "DBG:     $dir"
        done
    fi

    git_find -z ':(glob)pkg/generated/**'/"${output_file}" | xargs -0 rm -f

    openapi-gen \
        -v "${KUBE_VERBOSE}" \
        --go-header-file "${BOILERPLATE_FILENAME}" \
        --output-file "${output_file}" \
        --output-dir "${output_dir}" \
        --output-pkg "${output_pkg}" \
        --report-filename "${report_file}" \
        "${tag_pkgs[@]}" \
        "$@"

    touch "${report_file}"
    local known_filename="${known_violations_file}"
    if ! diff -u "${known_filename}" "${report_file}"; then
        echo -e "ERROR:"
        echo -e "\tAPI rule check failed - reported violations differ from known violations"
        echo -e "\tPlease read api/api-rules/README.md to resolve the failure in ${known_filename}"
    fi

    if [[ "${DBG_CODEGEN}" == 1 ]]; then
        krm::log::status "Generated openapi code"
    fi
}


function codegen::clients() {
    GOPROXY=off go install \
        k8s.io/code-generator/cmd/client-gen

    IFS=" " read -r -a group_versions <<< "${ONEX_AVAILABLE_GROUP_VERSIONS}"
    local gv_dirs=()
    for gv in "${group_versions[@]}"; do
        # add items, but strip off any leading apis/ you find to match command expectations
        local api_dir
        api_dir=$(krm::log::group-version-to-pkg-path "${gv}")
        local nopkg_dir=${api_dir#pkg/}
        nopkg_dir=${nopkg_dir#staging/src/k8s.io/api/}
        local pkg_dir=${nopkg_dir#apis/}

        # skip groups that aren't being served, clients for these don't matter
        if [[ " ${KUBE_NONSERVER_GROUP_VERSIONS} " == *" ${gv} "* ]]; then
          continue
        fi

        gv_dirs+=("${pkg_dir}")
    done
    gv_dirs+=("${EXTRA_GENERATE_PKG[@]}")

    krm::log::status "Generating client code for ${#gv_dirs[@]} targets"
    if [[ "${DBG_CODEGEN}" == 1 ]]; then
        krm::log::status "DBG: running client-gen for:"
        for dir in "${gv_dirs[@]}"; do
            krm::log::status "DBG:     $dir"
        done
    fi

    (git_grep -l --null \
        -e '^// Code generated by client-gen. DO NOT EDIT.$' \
        -- \
        ':(glob)pkg/generated/**/*.go' \
        || true) \
        | xargs -0 rm -f

    # UPDATEME: When add new k8s resource.
    client-gen \
        -v "${KUBE_VERBOSE}" \
        --go-header-file "${BOILERPLATE_FILENAME}" \
        --included-types-overrides core/v1/Namespace,core/v1/ConfigMap,core/v1/Event,core/v1/Secret,apiextensions/v1/CustomResourceDefinition \
        --output-dir "${KRM_ROOT}/pkg/generated/clientset" \
        --output-pkg="${OUTPUT_PKG}/clientset" \
        --clientset-name="versioned" \
        --input-base="${KRM_MODULE_NAME}" \
        --plural-exceptions "${PLURAL_EXCEPTIONS}" \
        --apply-configuration-package "${APPLYCONFIG_PKG}" \
        "$@"

    # Fix generated namespace clients
    ${KRM_ROOT}/hack/fix-generated-client.sh

    if [[ "${DBG_CODEGEN}" == 1 ]]; then
        krm::log::status "Generated client code"
    fi
}

# OpenAPI generation
#
# Any package that wants open-api functions generated must include a
# comment-tag in column 0 of one file of the form:
#     // +k8s:openapi-gen=true
function todo::codegen::openapi() {
    # Build the tool.
    # Please make sure to use openapi-gen version v0.29.3 here.
    if ! command -v openapi-gen &> /dev/null ; then
        GOPROXY=off go install k8s.io/code-generator/cmd/openapi-gen@v0.29.3
    fi

    # The result file, in each pkg, of open-api generation.
    local output_file="${GENERATED_FILE_PREFIX}openapi"

    local output_dir="pkg/generated/openapi"
    local output_pkg="github.com/costa92/krm/${output_dir}"
    local known_violations_file="${API_KNOWN_VIOLATIONS_DIR}/violation_exceptions.list"

    local report_file="${OUT_DIR}/api_violations.report"
    # When UPDATE_API_KNOWN_VIOLATIONS is set to be true, let the generator to write
    # updated API violations to the known API violation exceptions list.
    if [[ "${UPDATE_API_KNOWN_VIOLATIONS}" == true ]]; then
        report_file="${known_violations_file}"
    fi

    if [[ "${DBG_CODEGEN}" == 1 ]]; then
        krm::log::status "DBG: finding all +k8s:openapi-gen tags"
    fi

    local tag_files=()
    krm::util::read-array tag_files < <(
        k8s_tag_files_except \
            staging/src/k8s.io/code-generator \
            staging/src/k8s.io/sample-apiserver
        )

    local tag_dirs=()
    krm::util::read-array tag_dirs < <(
        grep -l --null '+k8s:openapi-gen=' "${tag_files[@]}" \
            | while read -r -d $'\0' F; do dirname "${F}"; done \
            | sort -u)

    if [[ "${DBG_CODEGEN}" == 1 ]]; then
        krm::log::status "DBG: found ${#tag_dirs[@]} +k8s:openapi-gen tagged dirs"
    fi

    local tag_pkgs=()
    for dir in "${tag_dirs[@]}"; do
        tag_pkgs+=("./$dir")
    done

    krm::log::status "Generating openapi code"
    if [[ "${DBG_CODEGEN}" == 1 ]]; then
        krm::log::status "DBG: running openapi-gen for:"
        for dir in "${tag_dirs[@]}"; do
            krm::log::status "DBG:     $dir"
        done
    fi

    git_find -z ':(glob)pkg/generated/**'/"${output_file}" | xargs -0 rm -f

    # UPDATEME: When add new k8s resources/group.
    openapi-gen \
        -v "${KUBE_VERBOSE}" \
        --go-header-file "${BOILERPLATE_FILENAME}" \
        --output-file "${output_file}" \
        -i 'k8s.io/apimachinery/pkg/apis/meta/v1,k8s.io/apimachinery/pkg/runtime,k8s.io/apimachinery/pkg/version,k8s.io/apimachinery/pkg/util/intstr,k8s.io/kubernetes/pkg/apis/core,k8s.io/api/core/v1,k8s.io/api/autoscaling/v1,k8s.io/api/coordination/v1,k8s.io/kubernetes/pkg/apis/flowcontrol,k8s.io/api/flowcontrol/v1,k8s.io/apiextensions-apiserver/pkg/apis/apiextensions,k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1,k8s.io/kube-aggregator/pkg/apis/apiregistration,k8s.io/kube-aggregator/pkg/apis/apiregistration/v1,github.com/costa92/krm/pkg/apis/apps/v1beta1' \
        --output-base "${GOPATH}/src" \
        -p "${output_pkg}" \
        --report-filename "${report_file}" \
        "${tag_pkgs[@]}" \
        "$@"

    if [[ "${DBG_CODEGEN}" == 1 ]]; then
        krm::log::status "Generated openapi code"
    fi
}


function codegen::informers() {
    echo "start codegen::informers"
    if ! command -v gen-swaggertype-docs &> /dev/null ; then
        GOPROXY=off go install k8s.io/code-generator/cmd/informer-gen
    fi

    local ext_apis=()
    krm::util::read-array ext_apis < <(
        cd "${KRM_ROOT}"
        git_find -z ':(glob)pkg/apis/**/*types.go' \
            | while read -r -d $'\0' F; do dirname "github.com/costa92/krm/${F}"; done \
            | sort -u)
    ext_apis+=("${EXTRA_GENERATE_PKG[@]}")
    echo "ext_apis: ${ext_apis[@]}"
    krm::log::status "Generating informer code for ${#ext_apis[@]} targets"
    if [[ "${DBG_CODEGEN}" == 1 ]]; then
        krm::log::status "DBG: running informer-gen for:"
        for api in "${ext_apis[@]}"; do
            krm::log::status "DBG: $api"
        done
    fi

    (git_grep -l --null \
        -e '^// Code generated by informer-gen. DO NOT EDIT.$' \
        -- \
        ':(glob)pkg/generated/**/*.go' \
        || true) \
        | xargs -0 rm -f

    #  打印 informer-gen 执行命令
    echo "informer-gen \\
        -v \"${KUBE_VERBOSE}\" \\
        --go-header-file \"${BOILERPLATE_FILENAME}\" \\
        --output-dir \"${KRM_ROOT}/pkg/generated/informers\" \\
        --output-pkg \"${OUTPUT_PKG}/informers\" \\
        --single-directory \\
        --versioned-clientset-package \"${OUTPUT_PKG}/clientset/versioned\" \\
        --listers-package \"${OUTPUT_PKG}/listers\" \\
        --plural-exceptions \"${PLURAL_EXCEPTIONS}\" \\
        \"${ext_apis[@]}\" \\
        \"$@\""

    informer-gen \
        -v "${KUBE_VERBOSE}" \
        --go-header-file "${BOILERPLATE_FILENAME}" \
        --output-dir "${KRM_ROOT}/pkg/generated/informers" \
        --output-pkg "${OUTPUT_PKG}/informers" \
        --single-directory \
        --versioned-clientset-package "${OUTPUT_PKG}/clientset/versioned" \
        --listers-package "${OUTPUT_PKG}/listers" \
        --plural-exceptions "${PLURAL_EXCEPTIONS}" \
        "${ext_apis[@]}" \
        "$@"

    if [[ "${DBG_CODEGEN}" == 1 ]]; then
        krm::log::status "Generated informer code"
    fi
}


function indent() {
    while read -r X; do
        echo "    ${X}"
    done
}


function codegen::listers() {
    if ! command -v gen-swaggertype-docs &> /dev/null ; then
        GOPROXY=off go install k8s.io/code-generator/cmd/lister-gen
    fi

    local ext_apis=()
    krm::util::read-array ext_apis < <(
        cd "${KRM_ROOT}"
        git_find -z ':(glob)pkg/apis/**/*types.go' \
            | while read -r -d $'\0' F; do dirname "github.com/costa92/krm/${F}"; done \
            | sort -u)
    ext_apis+=("${EXTRA_GENERATE_PKG[@]}")

    echo "ext_apis: ${ext_apis[@]}"
  
    krm::log::status "Generating lister code for ${#ext_apis[@]} targets"
    if [[ "${DBG_CODEGEN}" == 1 ]]; then
        krm::log::status "DBG: running lister-gen for:"
        for api in "${ext_apis[@]}"; do
            krm::log::status "DBG:     $api"
        done
    fi

    (git_grep -l --null \
        -e '^// Code generated by lister-gen. DO NOT EDIT.$' \
        -- \
        ':(glob)pkg/generated/**/*.go' \
        || true) \
        | xargs -0 rm -f

        # --included-types-overrides core/v1/Namespace,core/v1/ConfigMap,core/v1/Event,core/v1/Secret \
    #  打印 lister-gen 执行命令
    echo "lister-gen \\
        -v \"${KUBE_VERBOSE}\" \\
        --go-header-file \"${BOILERPLATE_FILENAME}\" \\
        --output-dir \"${KRM_ROOT}/pkg/generated/listers\" \\
        --output-pkg \"${OUTPUT_PKG}/listers\" \\
        --plural-exceptions \"${PLURAL_EXCEPTIONS}\" \\
        \"${ext_apis[@]}\" \\
        \"$@\""
        
    lister-gen \
        -v "${KUBE_VERBOSE}" \
        --go-header-file "${BOILERPLATE_FILENAME}" \
        --output-dir "${KRM_ROOT}/pkg/generated/listers" \
        --output-pkg "${OUTPUT_PKG}/listers" \
        --plural-exceptions "${PLURAL_EXCEPTIONS}" \
        "${ext_apis[@]}" \
        "$@"

    if [[ "${DBG_CODEGEN}" == 1 ]]; then
        krm::log::status "Generated lister code"
    fi
}
# $@: directories to exclude
# example:
#    k8s_tag_files_except foo bat/qux
function k8s_tag_files_except() {
    for f in "${ALL_K8S_TAG_FILES[@]}"; do
        local excl=""
        for x in "$@"; do
            if [[ "$f" =~ "$x"/.* ]]; then
                excl="true"
                break
            fi
        done
        if [[ "${excl}" != true ]]; then
            echo "$f"
        fi
    done
}

#
# main
#
function list_codegens() {
    (
        shopt -s extdebug
        declare -F \
            | cut -f3 -d' ' \
            | grep "^codegen::" \
            | while read -r fn; do declare -F "$fn"; done \
            | sort -n -k2 \
            | cut -f1 -d' ' \
            | sed 's/^codegen:://'
    )
}

# shellcheck disable=SC2207 # safe, no functions have spaces
all_codegens=($(list_codegens))

function print_codegens() {
    echo "available codegens:"
    for g in "${all_codegens[@]}"; do
        echo "    $g"
    done
}

# Validate and accumulate flags to pass thru and codegens to run if args are
# specified.
flags_to_pass=()
codegens_to_run=()
for arg; do
    # Use -? to list known codegens.
    if [[ "${arg}" == "-?" ]]; then
        print_codegens
        exit 0
    fi
    # Accumulate flags to pass thru.
    if [[ "${arg}" =~ ^- ]]; then
        flags_to_pass+=("${arg}")
        continue
    fi
    # Make sure each non-flag arg matches at least one codegen.
    nmatches=0
    for t in "${all_codegens[@]}"; do
        if [[ "$t" =~ ${arg} ]]; then
            nmatches=$((nmatches+1))
            # Don't run codegens twice, just keep the first match.
            # shellcheck disable=SC2076 # we want literal matching
            if [[ " ${codegens_to_run[*]} " =~ " $t " ]]; then
                continue
            fi
            codegens_to_run+=("$t")
            continue
        fi
    done
    if [[ ${nmatches} == 0 ]]; then
        echo "ERROR: no codegens match pattern '${arg}'"
        echo
        print_codegens
        exit 1
    fi
    # The array-syntax abomination is to accommodate older bash.
    codegens_to_run+=("${matches[@]:+"${matches[@]}"}")
done

# If no codegens were specified, run them all.
if [[ "${#codegens_to_run[@]}" == 0 ]]; then
    codegens_to_run=("${all_codegens[@]}")
fi

for g in "${codegens_to_run[@]}"; do
    # The array-syntax abomination is to accommodate older bash.
    "codegen::${g}" "${flags_to_pass[@]:+"${flags_to_pass[@]}"}"
done
