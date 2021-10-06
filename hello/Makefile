PROJECT_REPOSITORY=github.com/becosuke/golang-examples/hello
PROJECT_NAME=hello
COMMAND_DIR=command
COMMAND_NAME=server
COMMAND_PATH=${COMMAND_DIR}/${COMMAND_NAME}

all: test build

mod:
	go mod vendor

mod-tidy:
	go mod tidy

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
