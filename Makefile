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

# Run swagger generate
.PHONY: swagger
swagger:
	swagger generate client -f swagger.yaml -c client