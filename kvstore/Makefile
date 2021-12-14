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

run: run-api run-gateway

run-api:
	go run -ldflags "-X main.version=$(VERSION) -X main.serviceName=$(PROJECT_NAME)-api" \
	$(PROJECT_REPOSITORY)/$(COMMAND_DIR)/api

run-gateway:
	go run -ldflags "-X main.version=$(VERSION) -X main.serviceName=$(PROJECT_NAME)-gateway" \
	$(PROJECT_REPOSITORY)/$(COMMAND_DIR)/gateway

.PHONY: test
test: utest ftest

utest:
	go test $(args) -race -cover $(CODE_DIRS)

ftest:
	go test $(args) -race -cover $(TEST_DIRS)

.PHONY: clean
clean:
	go clean
	rm -f $(BUILD_DIR)/api $(BUILD_DIR)/gateway

clean-test:
	go clean -testcache

.PHONY: build
build: build-binary

build-binary: build-binary-api build-binary-gateway

build-binary-api:
	go build -a -installsuffix cgo \
	-ldflags "-w -X main.version=$(VERSION) -X main.serviceName=$(PROJECT_NAME)-api" \
	-o $(BUILD_DIR)/api $(COMMAND_DIR)/api/main.go

build-binary-gateway:
	go build -a -installsuffix cgo \
	-ldflags "-w -X main.version=$(VERSION) -X main.serviceName=$(PROJECT_NAME)-gateway" \
	-o $(BUILD_DIR)/gateway $(COMMAND_DIR)/gateway/main.go

build-container-api:
	cd $(BUILD_DIR) && \
	docker build . --no-cache --build-arg name=$(PROJECT_NAME)-api -t $(PROJECT_NAME)-api:latest -t $(PROJECT_NAME)-api:$(VERSION)

build-container-gateway:
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
	internal/protobuf/kvstore.proto

mockgen:
	mockgen -source internal/usecases/interactor/interactor.go -destination mocks/usecases/interactor/interactor.go -package interactor
	mockgen -source internal/adapters/repository/repository.go -destination mocks/adapters/repository/repository.go -package repository
	mockgen -source internal/drivers/syncmap/syncmap.go        -destination mocks/drivers/syncmap/syncmap.go        -package syncmap
