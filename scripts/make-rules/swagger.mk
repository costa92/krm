# ==============================================================================
# Makefile helper functions for swagger
#

.PHONY: swagger.run
swagger.run: tools.verify.swagger
	@echo "===========> Generating swagger API docs"
    # 判断是否存在swagger文件夹，如果不存在则创建
	@if [ ! -d $(KRM_ROOT)/api/swagger ]; then mkdir -p $(KRM_ROOT)/api/swagger; fi
	#@swagger generate spec --scan-models -w $(KRM_ROOT)/cmd/gen-swagger-type-docs -o $(KRM_ROOT)/api/swagger/kubernetes.yaml
	@swagger mixin `find $(KRM_ROOT)/api/openapi -name "*.swagger.json"` \
		-q                                                    \
		--keep-spec-order                                     \
		--format=yaml                                         \
		--ignore-conflicts                                    \
		-o $(KRM_ROOT)/api/swagger/swagger.yaml
	@echo "Generated at: $(KRM_ROOT)/api/swagger/swagger.yaml"

.PHONY: swagger.serve
swagger.serve: tools.verify.swagger
	@swagger serve -F=redoc --no-open --port 65534 $(KRM_ROOT)/api/swagger/swagger.yaml
