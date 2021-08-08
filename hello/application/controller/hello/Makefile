ENDPOINT_NAME=hello
ENDPOINT_PATH=application/controller/${ENDPOINT_NAME}

all: test build

run:
	go run ${ENDPOINT_PATH}/main.go ${ENDPOINT_PATH}/hello.go

test:
	go test ${PROJECT_OWNER}/${PROJECT_NAME}/${ENDPOINT_PATH}

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o ${OUTPUT}/${ENDPOINT_NAME} ${ENDPOINT_PATH}/main.go ${ENDPOINT_PATH}/hello.go
	cd ${OUTPUT} && docker build . --no-cache --build-arg name=${ENDPOINT_NAME} -t ${PROJECT_NAME}-api-${ENDPOINT_NAME}
