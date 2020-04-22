TEST?=$$(go list ./... |grep -v 'vendor'|grep -v 'examples')
GO111MODULE=on
GOFLAGS=-mod=vendor

default: build install

build: fmtcheck
	go build
	
build-docker:
	mkdir -p bin
	docker run --rm -v $$(pwd)/bin:/go/bin -v $$(pwd):/go/src/github.com/samjegal/terraform-provider-fincloud -w /go/src/github.com/samjegal/terraform-provider-fincloud -e GOOS golang:1.14 make build

install: fmtcheck
	cp terraform-provider-fincloud /usr/local/bin
	rm terraform-provider-fincloud

fmt:
	@echo "==> Fixing source code with gofmt..."
	find . -name '*.go' | grep -v vendor | xargs gofmt -s -w

fmtcheck:
	@sh "$(CURDIR)/scripts/gofmtcheck.sh"

testacc:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 180m -ldflags="-X=./version.ProviderVersion=acc"
