PYTHON?=python3
ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
SAAS_DIR:=../../../saas

# Sync specs from SAAS
.PHONY: sync-specs
sync-specs:
	$(PYTHON) sync_api_client.py --saas-dir $(SAAS_DIR) --provider-dir .

# Generate the full spec
.PHONY: merge-specs
merge-specs: sync-specs
	$(PYTHON) specmerge.py -s pgedge.yaml -o merged.yaml

# Install all required dependencies
.PHONY: install-deps
install-deps:
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
	npm install -g @apidevtools/swagger-cli
	brew install daveshanley/vacuum/vacuum

# Validate the OpenAPI spec
.PHONY: validate-spec
validate-spec:
	swagger-cli validate merged.yaml

# Lint the OpenAPI spec
.PHONY: lint-spec
lint-spec:
	vacuum lint merged.yaml -d -r pgedge_vacuum_rules.yaml -n=warn

# Generate lint report
.PHONY: lint-report
lint-report:
	vacuum html-report merged.yaml -r pgedge_vacuum_rules.yaml

# Generate the API client
.PHONY: generate-client
generate-client: merge-specs validate-spec
	swagger-cli bundle -o pgedge.json merged.yaml
	oapi-codegen -package client -generate "types,client" ./merged.yaml > api.gen.go