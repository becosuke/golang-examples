export PROJECT_OWNER=example
export PROJECT_NAME=hello
export PROJECT_PATH=${GOPATH}/src/${PROJECT_OWNER}/${PROJECT_NAME}
export CONTROLLER_PATH=application/controller
export OUTPUT=binary

all: deps test build

deps:
	dep ensure

run-hello:
	make -f ${CONTROLLER_PATH}/hello/Makefile run

test: test-hello

test-hello:
	make -f ${CONTROLLER_PATH}/hello/Makefile test

build: build-hello

build-hello:
	make -f ${CONTROLLER_PATH}/hello/Makefile build

hello: test-hello build-hello

clean:
	go clean
	rm -f binary/*
