SHELL=/bin/bash -e -o pipefail
PWD = $(shell pwd)

# constants
GOLANGCI_VERSION = 1.48.0
DOCKER_REPO = stackit-api-manager-cli
OPENAPI_GENERATOR_VERSION=v6.0.1
DOCKER_TAG = latest

all: git-hooks  tidy ## Initializes all tools

out:
	@mkdir -p out

git-hooks:
	@git config --local core.hooksPath .githooks/

download: ## Downloads the dependencies
	@go mod download

tidy: ## Cleans up go.mod and go.sum
	@go mod tidy

fmt: ## Formats all code with go fmt
	@go fmt ./...

run: fmt ## Run the app
	@go run ./cmd/stackit-api-manager-cli/main.go

test-build: ## Tests whether the code compiles
	@go build -o /dev/null ./...

build: out/bin ## Builds all binaries

GO_BUILD = mkdir -pv "$(@)" && go build -ldflags="-w -s" -o "$(@)" ./...
.PHONY: out/bin
out/bin:
	$(GO_BUILD)

GOLANGCI_LINT = bin/golangci-lint-$(GOLANGCI_VERSION)
$(GOLANGCI_LINT):
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | bash -s -- -b bin v$(GOLANGCI_VERSION)
	@mv bin/golangci-lint "$(@)" --fix

lint: fmt $(GOLANGCI_LINT) download ## Lints all code with golangci-lint
	@$(GOLANGCI_LINT) run

lint-reports: out/lint.xml

.PHONY: out/lint.xml
out/lint.xml: $(GOLANGCI_LINT) out download
	@$(GOLANGCI_LINT) run ./... --out-format checkstyle | tee "$(@)"

test: ## Runs all tests
	@go test $(ARGS) ./...

coverage: out/report.json ## Displays coverage per func on cli
	go tool cover -func=out/cover.out

html-coverage: out/report.json ## Displays the coverage results in the browser
	go tool cover -html=out/cover.out

test-reports: out/report.json

.PHONY: out/report.json
out/report.json: out
	@go test -count 1 ./... -coverprofile=out/cover.out --json | tee "$(@)"

clean: ## Cleans up everything
	@rm -rf bin out

docker: ## Builds docker image
	docker buildx build -t $(DOCKER_REPO):$(DOCKER_TAG) .


.PHONY: generate-client-code
generate-client-code: ## generate API client code
	docker run --rm \
		-v ${PWD}:/local openapitools/openapi-generator-cli:${OPENAPI_GENERATOR_VERSION} generate \
		-i /local/api/api_manager.openapi.json \
		-g go \
		--additional-properties=packageName=client \
		-o /local/pkg/stackit_api_manager/client

.PHONY: generate-client
generate-client: generate-client-code fmt tidy ## genarte API client & run go mod tidy

ci: lint-reports test-reports ## Executes lint and test and generates reports

help: ## Shows the help
	@echo 'Usage: make <OPTIONS> ... <TARGETS>'
	@echo ''
	@echo 'Available targets are:'
	@echo ''
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
        awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
	@echo ''
