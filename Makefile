SHELL:=/bin/bash
APP:=birthdaytrivia
BUILD_DIR:=.build
BIN_DIR:=$(BUILD_DIR)/$(APP)/_bin

GO_OS?="linux"

DOCKER_REPO?="matache91mh"
IMAGE?=$(DOCKER_REPO)/$(APP)

VERSION ?= $(shell git describe --tags --dirty --always)
BUILD_DATE ?= $(shell date +%FT%T%z)
COMMIT_HASH ?= $(shell git rev-parse --short HEAD 2>/dev/null)

LDFLAGS += -X 'github.com/mimatache/birthdaytrivia/internal/info.appName=${APP}'
LDFLAGS += -X 'github.com/mimatache/birthdaytrivia/internal/info.version=${VERSION}'
LDFLAGS += -X 'github.com/mimatache/birthdaytrivia/internal/info.commitHash=${COMMIT_HASH}'
LDFLAGS += -X 'github.com/mimatache/birthdaytrivia/internal/info.buildDate=${BUILD_DATE}'


all: install-go-tools lint test build

build: build-ui
	CGO_ENABLED=0 GOOS=${GO_OS} go build -ldflags="$(LDFLAGS)" -v .

build-ui:
	pushd web/trivia-ui && \
	npm install --legacy-peer-dep && \
	npm run build && \
	popd

test-ci:
	go test -v  -race -json -coverprofile=coverage.out ./... > unit-test.json
	go tool cover -func=coverage.out

test:
	go test -v -race -cover ./...

install-go-tools:
	GO111MODULE=on CGO_ENABLED=0 go get github.com/golangci/golangci-lint/cmd/golangci-lint

lint:
	go vet ./...
	golangci-lint run

fmt:
	go mod tidy
	goimports -w .
	gofmt -s -w .

run: build-ui
	GO111MODULE=on CGO_ENABLED=0 go run .

build-full: build-ui build

docker-build:
	docker build -t $(IMAGE):$(VERSION) --build-arg VERSION=${VERSION} --build-arg BUILD_DATE=${BUILD_DATE} --build-arg COMMIT_HASH=${COMMIT_HASH} --build-arg APP=${APP} .

docker-push: app-image
	docker push $(IMAGE):$(VERSION)

