PROJECT_REPOSITORY=github.com/becosuke/golang-examples/rest
PROJECT_NAME=rest
COMMAND_DIR=internal/cmd
COMMAND_NAME=server
COMMAND_PATH=${COMMAND_DIR}/${COMMAND_NAME}
BUILD_DIR=build
TEST_DIR=tests/functional

all: src/test build

mod:
	go mod vendor

mod-tidy:
	go mod tidy

fmt:
	sh -c "find . -type d \\( -name .git -o -name vendor -o -name protogen -o -name output \\) -prune -o -type f -name *.go -print | xargs -n1 goimports -w -local ${PROJECT_REPOSITORY}"

vet:
	sh -c "find . -type d \\( -name .git -o -name vendor -o -name protogen -o -name output \\) -prune -o -type d -print | xargs -IXXX sh -c 'find XXX -maxdepth 1 -type f -name *.go -print | xargs --no-run-if-empty go vet' || :"

run:
	go run ${PROJECT_REPOSITORY}/${COMMAND_PATH}

test: test-unit test-functional

test-unit:
	go test -v ${PROJECT_REPOSITORY}/internal/...

test-functional:
	go test -v ${PROJECT_REPOSITORY}/${TEST_DIR}

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
	mockgen -source internal/usecases/interactor.go -destination tests/mocks/usercases/interactor.go -package usecases
	mockgen -source internal/interfaces/repository/repository.go -destination tests/mocks/interfaces/repository/repository.go -package repository
	mockgen -source internal/infrastructure/syncmap/syncmap.go -destination tests/mocks/infrastructure/syncmap/syncmap.go -package syncmap
