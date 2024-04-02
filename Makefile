GOLINT_VER = v1.55.2

##@ General

.PHONY: verify
verify: test checks go-lint ## verify simulates same behaviour as 'verify' GitHub Action which run on every PR

.PHONY: checks
checks: check-go-mod-tidy ## run different Go related checks

.PHONY: go-lint
go-lint: go-lint-install ## linter config in file at root of project -> '.golangci.yaml'
	golangci-lint run

go-lint-install: ## linter config in file at root of project -> '.golangci.yaml'
	@if [ "$(shell command golangci-lint version --format short)" != "$(GOLINT_VER)" ]; then \
  		echo golangci in version $(GOLINT_VER) not found. will be downloaded; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLINT_VER); \
		echo golangci installed with version: $(shell command golangci-lint version --format short); \
	fi;
	

##@ Tests

.PHONY: test 
test: ## run Go tests
	go test ./...

.PHONY: fix
fix: go-lint-install
	go mod tidy -v
	golangci-lint run --fix