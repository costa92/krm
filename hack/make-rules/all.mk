# ===================

# Include all make rules
include hack/make-rules/tools.mk
include hack/make-rules/golang.mk
include hack/make-rules/generate.mk
# include hack/make-rules/generate-k8s.mk
include hack/make-rules/copyright.mk
include hack/make-rules/lint.mk
include hack/make-rules/image.mk
include hack/make-rules/swagger.mk
