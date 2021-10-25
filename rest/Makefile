PROJECT_REPOSITORY=github.com/becosuke/golang-examples/rest
PROJECT_NAME=rest
COMMAND_DIR=command
COMMAND_NAME=server
COMMAND_PATH=${COMMAND_DIR}/${COMMAND_NAME}
BUILD_DIR=build

all: test build

mod:
	go mod vendor

mod-tidy:
	go mod tidy

fmt:
	sh -c "find . -type d \\( -name .git -o -name vendor -o -name protogen -o -name output \\) -prune -o -type f -name *.go -print | xargs -n1 goimports -w"

vet:
	sh -c "find . -type d \\( -name .git -o -name vendor -o -name protogen -o -name output \\) -prune -o -type d -print | xargs -IXXX sh -c 'find XXX -maxdepth 1 -type f -name *.go -print | xargs --no-run-if-empty go vet' || :"

run:
	go run ${PROJECT_REPOSITORY}/${COMMAND_PATH}

test:
	go test ${PROJECT_REPOSITORY}/${COMMAND_PATH}

.PHONY: build
build: build-binary build-container

build-binary:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o ${BUILD}/${COMMAND_NAME} ${PROJECT_REPOSITORY}/${COMMAND_PATH}

build-container:
	cd ${BUILD_DIR} && docker build . --no-cache --build-arg name=${COMMAND_NAME} -t ${PROJECT_NAME}-${COMMAND_NAME}:latest -t ${PROJECT_NAME}-${COMMAND_NAME}:$(shell date "+%Y%m%d%H%M%S")

clean:
	go clean
	rm -f ${BUILD_DIR}/${COMMAND_NAME}

.PHONY: mock
gen-mock:
	mockgen -source presentation/handler/handler.go -destination mock/handler/handler.go -package handler
	mockgen -source presentation/controller/controller.go -destination mock/controller/controller.go -package controller
	mockgen -source application/usecase/usecase.go -destination mock/usercase/usecase.go -package usecase
	mockgen -source domain/processor/processor.go -destination mock/processor/processor.go -package processor
	mockgen -source domain/repository/repository.go -destination mock/repository/repository.go -package repository
	mockgen -source infrastructure/memorymap/memorymap.go -destination mock/memorymap/memorymap.go -package memorymap
	mockgen -source infrastructure/syncmap/syncmap.go -destination mock/syncmap/syncmap.go -package syncmap
