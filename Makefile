#@IgnoreInspection BashAddShebang
export ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))
export CGO_ENABLED=0
export GO111MODULE=on
export GOPROXY=https://proxy.golang.org,direct

export ENV=development

LDFLAGS+=-X k8s.io/client-go/pkg/version.gitVersion=v0.1.0
LDFLAGS+=-X k8s.io/client-go/pkg/version.gitCommit=4f75300

.PHONY: all format lint build

build:
	go build -v -o $(ROOT)/bin/arvan -ldflags="$(LDFLAGS)" $(ROOT)/cmd/arvan/*.go
