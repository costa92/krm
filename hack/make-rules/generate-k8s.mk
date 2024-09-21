# This file contains rules for generating Kubernetes code.

.PHONY: generated.files
generated.files: tools.verify.code-generator gen_prerelease_lifecycle gen_deepcopy gen_defaulter gen_conversion gen_openapi gen_client ## Generate all necessary kubernetes files.


# This variable holds a list of every directory that contains Go files in this
# project.  Other rules and variables can use this as a starting point to
# reduce filesystem accesses.
ifeq ($(V),1)
    $(warning ***** finding all *.go dirs)
endif
ALL_GO_DIRS := $(shell                                                   \
    $(SCRIPTS_DIR)/cache_go_dirs.sh $(META_DIR)/all_go_dirs.mk  \
)
ifeq ($(V),1)
    $(warning ***** found $(shell echo $(ALL_GO_DIRS) | wc -w) *.go dirs)
endif

# The result file, in each pkg, of deep-copy generation.
DEEPCOPY_BASENAME := $(GENERATED_FILE_PREFIX)deepcopy
DEEPCOPY_FILENAME := $(DEEPCOPY_BASENAME).go

# The tool used to generate deep copies.
DEEPCOPY_GEN := $(GOPATH)/bin/deepcopy-gen
DEEPCOPY_GEN_TODO := $(DEEPCOPY_GEN).todo

# This rule aggregates the set of files to generate and then generates them all
# in a single run of the tool.
.PHONY: gen_deepcopy
gen_deepcopy: $(DEEPCOPY_GEN) $(META_DIR)/$(DEEPCOPY_GEN_TODO) ## Generate deepcopy generated files.
	echo "ALL_GO_DIRS: $(ALL_GO_DIRS)"
	if [[ -s $(META_DIR)/$(DEEPCOPY_GEN_TODO) ]]; then                  \
	    pkgs=$$(cat $(META_DIR)/$(DEEPCOPY_GEN_TODO) | paste -sd, -);   \
	    if [[ "$(V)" == 1 ]]; then                                      \
	        echo "DBG: running $(DEEPCOPY_GEN) for $$pkgs";             \
	    fi;                                                             \
	    N=$$(cat $(META_DIR)/$(DEEPCOPY_GEN_TODO) | wc -l);             \
	    echo "Generating deepcopy code for $$N targets";                \
	    $(SCRIPTS_DIR)/run-in-gopath.sh $(DEEPCOPY_GEN)                 \
	        --v $(KUBE_VERBOSE)                                         \
	        --logtostderr                                               \
	        -i "$$pkgs"                                                 \
	        --bounding-dirs $(PRJ_SRC_PATH),"k8s.io/api"                \
	        -O $(DEEPCOPY_BASENAME)                                     \
			--output-base "${GOPATH}/src"                               \
			--go-header-file ${SCRIPTS_DIR}/boilerplate.go.txt          \
	        "$$@";                                                      \
	fi


# The result file, in each pkg, of defaulter generation.
DEFAULTER_BASENAME := $(GENERATED_FILE_PREFIX)defaulter
DEFAULTER_FILENAME := $(DEFAULTER_BASENAME).go

# The tool used to generate defaulters.
DEFAULTER_GEN := defaulter-gen
