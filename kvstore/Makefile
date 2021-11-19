PROJECT_REPOSITORY=github.com/becosuke/golang-examples/kvstore
PROJECT_NAME=kvstore
COMMAND_DIR=internal/cmd
COMMAND_NAME=api
COMMAND_PATH=${COMMAND_DIR}/${COMMAND_NAME}
BUILD_DIR=build

CODE_DIRS ?= $(shell go list ./internal/... | grep -v tests | grep -v mocks)
TEST_DIRS ?= $(shell go list ./internal/tests/...)
LINT_DIRS ?= $(shell go list ./internal/... | grep -v mocks)

all: mod test build

mod:
	go mod vendor

mod-tidy:
	go mod tidy

vet:
	go vet ${CODE_DIRS}

lint:
	golangci-lint run $(args) ./...
	go-consistent $(cons_args) ${LINT_DIRS}

fmt:
	@find . -iname "*.go" -not -path "./vendor/**" -not -path "./mocks/**" | xargs gofmt -w

imports:
	@find . -iname "*.go" -not -path "./vendor/**" -not -path "./mocks/**" | xargs goimports -w -local ${PROJECT_REPOSITORY}

run:
	go run ${PROJECT_REPOSITORY}/${COMMAND_PATH}

.PHONY: test
test: utest ftest

utest:
	go test $(args) -race -cover ${CODE_DIRS}

ftest:
	go test $(args) -race -cover $(TEST_DIRS)

.PHONY: clean
clean:
	go clean
	rm -f ${BUILD_DIR}/${COMMAND_NAME}

.PHONY: build
build: build-grpc build-rest

build-grpc:
	${PRECOMMAND} go build -a -installsuffix cgo -ldflags '-w' -o ${OUTPUT}/api ${GRPC_PATH}/main/main.go
	cd ${OUTPUT} && docker build . --no-cache --build-arg name=kvstore-api -t kvstore-api:latest -t kvstore-api:${TIMESTAMP}

build-rest:
	${PRECOMMAND} go build -a -installsuffix cgo -ldflags '-w' -o ${OUTPUT}/gateway ${REST_PATH}/main/main.go
	cd ${OUTPUT} && docker build . --no-cache --build-arg name=kvstore-gateway -t kvstore-gateway:latest -t kvstore-gateway:${TIMESTAMP}

build-container:
	cd ${BUILD_DIR} && docker build . --no-cache --build-arg name=${COMMAND_NAME} -t ${PROJECT_NAME}-${COMMAND_NAME}:latest -t ${PROJECT_NAME}-${COMMAND_NAME}:$(shell date "+%Y%m%d%H%M%S")

protoc:
	protoc -I internal/protobuf -I $(shell brew --prefix)/opt/protobuf/include \
	-I modules/github.com/googleapis/googleapis \
	-I modules/github.com/envoyproxy/protoc-gen-validate \
	--go_out internal/protogen --go_opt paths=source_relative \
	--go-grpc_out internal/protogen --go-grpc_opt paths=source_relative \
	--grpc-gateway_out internal/protogen --grpc-gateway_opt paths=source_relative \
	--validate_out "lang=go,paths=source_relative:internal/protogen" \
	internal/protobuf/kvstore.proto

mock:
	mockgen -source internal/usecase/interactor.go -destination internal/mock/usecase/interactor.go -package usecase
	mockgen -source internal/adapter/repository/repository.go -destination internal/mock/adapter/repository/repository.go -package repository
	mockgen -source internal/infrastructure/syncmap/syncmap.go -destination internal/mock/infrastructure/syncmap/syncmap.go -package syncmap
