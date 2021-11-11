PROJECT_REPOSITORY=github.com/becosuke/golang-examples/rest
PROJECT_NAME=rest
COMMAND_DIR=internal/cmd
COMMAND_NAME=server
COMMAND_PATH=${COMMAND_DIR}/${COMMAND_NAME}
BUILD_DIR=build

CODE_DIRS ?= $(shell go list ./... | grep -v tests | grep -v mocks)
TEST_DIRS ?= $(shell go list ./tests/...)
LINT_DIRS ?= $(shell go list ./... | grep -v mocks)

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
build: build-binary build-container

build-binary:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o ${BUILD}/${COMMAND_NAME} ${PROJECT_REPOSITORY}/${COMMAND_PATH}

build-container:
	cd ${BUILD_DIR} && docker build . --no-cache --build-arg name=${COMMAND_NAME} -t ${PROJECT_NAME}-${COMMAND_NAME}:latest -t ${PROJECT_NAME}-${COMMAND_NAME}:$(shell date "+%Y%m%d%H%M%S")

.PHONY: mock
mock:
	mockgen -source internal/usecases/interactor.go -destination mocks/usercases/interactor.go -package usecases
	mockgen -source internal/interfaces/repository/repository.go -destination mocks/interfaces/repository/repository.go -package repository
	mockgen -source internal/infrastructure/syncmap/syncmap.go -destination mocks/infrastructure/syncmap/syncmap.go -package syncmap
