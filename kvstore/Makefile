PROJECT_REPOSITORY=github.com/becosuke/golang-examples/kvstore
PROJECT_NAME=kvstore
COMMAND_DIR=internal/cmd
BUILD_DIR=build

CODE_DIRS ?= $(shell go list ./internal/... | grep -v pkg)
TEST_DIRS ?= $(shell go list ./tests/...)
LINT_DIRS ?= $(shell go list ./internal/... | grep -v pkg)

TIMESTAMP := $(shell date "+%Y%m%d%H%M%S")
GIT_REF := $(shell git rev-parse --short=7 HEAD)
VERSION ?= $(TIMESTAMP)-$(GIT_REF)

all: mod test build

mod:
	go mod vendor

mod-tidy:
	go mod tidy

vet:
	go vet $(CODE_DIRS)

lint:
	golangci-lint run $(args) ./internal/...
	go-consistent $(cons_args) $(LINT_DIRS)

fmt:
	@find ./internal -iname "*.go" -not -path "./pkg/**" | xargs gofmt -w

imports:
	@find ./internal -iname "*.go" -not -path "./pkg/**" | xargs goimports -w -local $(PROJECT_REPOSITORY)

run: run-backend run-gateway

run-backend:
	go run -ldflags "-X main.version=$(VERSION) -X main.serviceName=$(PROJECT_NAME)-backend" \
	$(PROJECT_REPOSITORY)/$(COMMAND_DIR)/backend

run-gateway:
	go run -ldflags "-X main.version=$(VERSION) -X main.serviceName=$(PROJECT_NAME)-gateway" \
	$(PROJECT_REPOSITORY)/$(COMMAND_DIR)/gateway

.PHONY: test
test: unit-test functional-test

unit-test:
	go test $(args) -race -cover $(CODE_DIRS)

functional-test:
	go test $(args) -race -cover $(TEST_DIRS)

.PHONY: clean
clean:
	go clean
	rm -f $(BUILD_DIR)/backend $(BUILD_DIR)/gateway

clean-test:
	go clean -testcache

.PHONY: build
build: update-build-dependencies build-backend build-gateway

update-build-dependencies:
	docker run --rm -v $(shell pwd):/root golang:1.18.0-alpine sh -c 'cp /usr/local/go/lib/time/zoneinfo.zip /root/build/zoneinfo.zip && cp /etc/ssl/certs/ca-certificates.crt /root/build/ca-certificates.crt'

.PHONY: build-backend
build-backend: build-backend-binary build-backend-container

.PHONY: build-gateway
build-gateway: build-gateway-binary build-gateway-container

build-backend-binary:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo \
	-ldflags "-w -X main.version=$(VERSION) -X main.serviceName=$(PROJECT_NAME)-backend" \
	-o $(BUILD_DIR)/$(PROJECT_NAME)-backend $(COMMAND_DIR)/backend/main.go

build-gateway-binary:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo \
	-ldflags "-w -X main.version=$(VERSION) -X main.serviceName=$(PROJECT_NAME)-gateway" \
	-o $(BUILD_DIR)/$(PROJECT_NAME)-gateway $(COMMAND_DIR)/gateway/main.go

build-backend-container:
	cd $(BUILD_DIR) && \
	docker build . --no-cache --build-arg name=$(PROJECT_NAME)-backend -t $(PROJECT_NAME)-backend:latest -t $(PROJECT_NAME)-backend:$(VERSION)

build-gateway-container:
	cd $(BUILD_DIR) && \
	docker build . --no-cache --build-arg name=$(PROJECT_NAME)-gateway -t $(PROJECT_NAME)-gateway:latest -t $(PROJECT_NAME)-gateway:$(VERSION)

protoc:
	protoc -I proto -I $(shell brew --prefix)/opt/protobuf/include \
	-I modules/github.com/googleapis/googleapis \
	-I modules/github.com/envoyproxy/protoc-gen-validate \
	--go_out pb --go_opt paths=source_relative \
	--go-grpc_out pb --go-grpc_opt paths=source_relative \
	--grpc-gateway_out pb --grpc-gateway_opt paths=source_relative \
	--validate_out "lang=go,paths=source_relative:pb" \
	proto/kvstore.proto

generate-mocks:
	mockgen -source internal/domain/pack/usecase.go      -destination mocks/domain/pack/usecase.go      -package pack
	mockgen -source internal/domain/pack/repository.go   -destination mocks/domain/pack/repository.go   -package pack
	mockgen -source internal/drivers/syncmap/syncmap.go  -destination mocks/drivers/syncmap/syncmap.go  -package syncmap

