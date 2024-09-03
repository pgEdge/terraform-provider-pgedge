# Run golangci-lint
.PHONY: lint
lint:
	golangci-lint run ./...

# Run all tests
.PHONY: test
test:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m

# Run doc generate
.PHONY: docs
docs:
	go generate ./...

PYTHON?=python3

# Install dependencies
.PHONY: install
install:
	go install github.com/go-swagger/go-swagger/cmd/swagger@latest

.PHONY: validate
validate:
	swagger validate client/swagger/swagger.yaml 

# Merge OpenAPI specification files
.PHONY: merge-openapi-specs
merge-openapi-specs:
	$(PYTHON) client/swagger/specmerge.py client/swagger/main.yaml > client/swagger/swagger.yaml

# Generate api client
.PHONY: generate-swagger-client
generate-client: merge-openapi-specs
	swagger generate client -f client/swagger/swagger.yaml -c client --skip-models
	swagger generate model -f client/swagger/swagger.yaml -t client
