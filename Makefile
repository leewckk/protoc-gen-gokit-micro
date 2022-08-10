
BUILD_TIME=`date`
BUILD_TIME_DATE=`date "+%Y-%m-%d %H:%M:%S"`
GO_VERSION=`go version`
GIT_HASH=`git rev-parse HEAD`
BIN_DIR=./

BIN_NAME=protoc-gen-gokit-micro
BUILD_NAME=${BIN_DIR}${BIN_NAME}
GIT_TAG=`git describe --always --tag`

build:
	GOOS=linux CGO_ENABLED=0 go build -ldflags "-X 'main.BuildTime=${BUILD_TIME_DATE}' \
		-X 'main.GitHash=${GIT_HASH}' \
		-X 'main.GitTag=${GIT_TAG}' \
		-X 'main.GoVersion=${GO_VERSION}'" -o ${BUILD_NAME} 
install:
	GOOS=linux CGO_ENABLED=0 go build -ldflags "-X 'main.BuildTime=${BUILD_TIME_DATE}' \
		-X 'main.GitHash=${GIT_HASH}' \
		-X 'main.GitTag=${GIT_TAG}' \
		-X 'main.GoVersion=${GO_VERSION}'" -o ${BUILD_NAME} 
	@cp ${BIN_NAME} ~/go/bin/

all:
	GOOS=linux CGO_ENABLED=0 go build -ldflags "-X 'main.BuildTime=${BUILD_TIME_DATE}' \
		-X 'main.GitHash=${GIT_HASH}' \
		-X 'main.GitTag=${GIT_TAG}' \
		-X 'main.GoVersion=${GO_VERSION}'" -o ${BUILD_NAME} 

default:
	GOOS=linux CGO_ENABLED=0 go build -ldflags "-X 'main.BuildTime=${BUILD_TIME_DATE}' \
		-X 'main.GitHash=${GIT_HASH}' \
		-X 'main.GitTag=${GIT_TAG}' \
		-X 'main.GoVersion=${GO_VERSION}'" -o ${BUILD_NAME} 

clean :
	@if [ -f ${BUILD_NAME} ]; then rm ${BUILD_NAME}; fi

