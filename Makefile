PROJECT_NAME = kuafu
PACKAGES ?= $(shell go list ./... | grep -v /vendor/)
GOFILES := $(shell find . -name "*.go" -type f -not -path "./vendor/*")
GOFMT ?= gofmt "-s"
VERSION := $(shell cat VERSION)

all: install fmt build

install:
	@echo ">> install dependence"
	@exec ./bin/dep.sh

release:
	@echo ">> release ${VERSION}"
	@git tag "${VERSION}"
	@git push origin "${VERSION}"

.PHONY: test
test:
	@echo ">> run test"
	@go test -cover -race -v ./...

fmt:
	@echo ">> formatting code"
	@$(GOFMT) -w $(GOFILES)

.PHONY: fmt-check
fmt-check:
	# get all go files and run go fmt on them
	@diff=$$($(GOFMT) -d $(GOFILES)); \
	if [ -n "$$diff" ]; then \
		echo "Please run 'make fmt' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi;

.PHONY: build
build: $(GOFILES)
	@echo ">> building binaries"
	@CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o kuafu main.go
