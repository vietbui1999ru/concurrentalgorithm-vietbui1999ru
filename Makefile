all: test

# Tool
GOLANGCI_LINT_VERSION = v1.61.0

.PHONY: test
test:

.PHONY: test-go
test: test-go
test-go:
	go test -v ./...

.PHONY: lint
test: lint
lint: tool/golangci-lint
	tool/golangci-lint run

.PHONY: fix
fix: tool/golangci-lint
	tool/golangci-lint run --fix

.PHONY:
test: complete
complete:
	grep -qF 'func TestInsertionSortConcurrent(t *testing.T)' algo/algo_test.go

tool/golangci-lint: tool/.golangci-lint.$(GOLANGCI_LINT_VERSION)
	GOBIN="$(PWD)/tool" go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)

tool/.golangci-lint.$(GOLANGCI_LINT_VERSION):
	@rm -f tool/.golangci-lint.*
	@mkdir -p tool
	touch $@

.PHONY: tool
tool: tool/golangci-lint
