# Build all by default
.DEFAULT_GOAL := help

# ==================================
# include
include scripts/make-rules/common.mk
include scripts/make-rules/all.mk

# ==================================
# Usage
define USAGE_OPTIONS

\033[35mOptions:\033[0m
  DBG		Set to 1 enable debug build. Default is 0.
  V		Set to 1 enable verbose build. Default is 0.
endef
export USAGE_OPTIONS



# =================================
.PHONY: targets
targets: Makefile ## Show all Sub-makefile targets.
	@echo $(KRM_ROOT)
	@for mk in `echo $(MAKEFILE_LIST) | sed 's/Makefile //g'`; do echo -e \\n\\033[35m$$mk\\033[0m; awk -F':.*##' '/^[0-9A-Za-z._-]+:.*?##/ { printf "  \033[36m%-45s\033[0m %s\n", $$1, $$2 } /^\$$\([0-9A-Za-z_-]+\):.*?##/ { gsub("_","-", $$1); printf "  \033[36m%-45s\033[0m %s\n", tolower(substr($$1, 3, length($$1)-7)), $$2 }' $$mk;done;


# ==================================

.PHONY: help
help: Makefile ## @other Show this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<TARGETS> <OPTIONS>\033[0m\n\n\033[35mTargets:\033[0m\n"} /^[0-9A-Za-z._-]+:.*?##/ { printf "  \033[36m%-45s\033[0m %s\n", $$1, $$2 } /^\$$\([0-9A-Za-z_-]+\):.*?##/ { gsub("_","-", $$1); printf "  \033[36m%-45s\033[0m %s\n", tolower(substr($$1, 3, length($$1)-7)), $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' Makefile #$(MAKEFILE_LIST)
	@echo -e "$$USAGE_OPTIONS"