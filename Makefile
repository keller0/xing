GOFMT ?= gofmt "-s"
PACKAGES ?= $(shell go list ./... | grep -v /vendor/)
GOFILES := $(shell find . -name "*.go" -type f -not -path "./vendor/*")

help:
	@grep -P '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: fmt
fmt: ## format all go files (use go)
	$(GOFMT) -w $(GOFILES)

vet: ## vat all go files (use go)
	go vet $(PACKAGES)

build:fmt vet ## format vet and compile server (use go)
	go build -ldflags '-w -s' -o xingServer server/main.go

ctl: fmt vet ## format vet and compile ctl (use go)
	go build -ldflags '-w -s' -o xingCtl tools/main.go

clean: ## clean local build
	rm -f ./xingCtl
	rm -f ./xingServer
