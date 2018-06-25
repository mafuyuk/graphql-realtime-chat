HAVE_GOLINT:=$(shell which golint)
HAVE_VGO:=$(shell which vgo)
HAVE_GQLGEN:=$(shell which gqlgen)
HAVE_PECO:=$(shell which peco)

.PHONY: setup
setup: vgo gqlgen
	@echo "start setup"
	@vgo mod -vendor
	@go generate $(shell go list ./... | grep -v /vendor/)

## Go
.PHONY: lint vet test build run
lint: setup golint
	@echo "check lint"
	@golint $(shell go list ./...|grep -v vendor)
	@go vet ./...

test: setup
	@echo "go test"
	@go test -v $(shell go list ./... | grep -v /vendor/)

build: setup
	@echo "go build"
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 vgo build -o ./bin/graphql-realtime-chat ./main.go

run: setup
	@echo "go run"
	@export REDIS_ADDR=localhost:6379 && go run main.go

## Docker
CONTAINER_PREFIX:=graphql-realtime-chat

.PHONY: dstart dstop dstatus dlogin dclean dlog
dstart: setup
	@echo "docker start"
	@docker-compose up -d
	@export REDIS_ADDR=localhost:6379

dstop:
	@echo "docker stop"
	@docker-compose stop

dstatus:
	@echo "docker status"
	@docker ps --filter name=$(CONTAINER_PREFIX)

dlogin:
	@echo "docker login"
	@docker exec -it $(shell docker ps --all --format "{{.Names}}" | peco) /bin/bash

dclean:
	@echo "docker clean"
	@docker ps --all --filter name=$(CONTAINER_PREFIX) --quiet | xargs docker rm --force

dlog: peco
	@echo "docker log"
	@docker-compose logs -f $(shell docker ps --all --format "{{.Names}}" | peco | cut -d"_" -f2)


## Install package
.PHONY: vgo golint gqlgen peco
vgo:
ifndef HAVE_VGO
	@echo "Installing vgo"
	@go get -u golang.org/x/vgo
endif

golint:
ifndef HAVE_GOLINT
	@echo "Installing linter"
	@go get -u github.com/golang/lint/golint
endif

gqlgen:
ifndef HAVE_GQLGEN
	@echo "Installing gqlgen"
	@go get -u github.com/vektah/gqlgen
endif

peco:
ifndef HAVE_PECO
	@echo "Installing peco"
	@go get -u github.com/peco/peco/cmd/peco
endif
