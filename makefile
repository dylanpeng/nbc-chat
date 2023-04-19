#!/bin/bash
export LANG=zh_CN.UTF-8

ifndef GIT_BRANCH
	GIT_BRANCH=`git symbolic-ref --short -q HEAD`
endif

ifndef GIT_HASH
	GIT_HASH=`git rev-parse --short HEAD`
endif

ifndef BUILD_TIME
	BUILD_TIME=`date '+%Y-%m-%dT%H:%M:%S'`
endif

ENVARG=CGO_ENABLED=0 GOPROXY=https://goproxy.cn,direct
LINUXARG=GOOS=linux GOARCH=amd64
BUILDARG=-mod=mod -ldflags " -s -X main.buildTime=${BUILD_TIME} -X main.gitHash=${GIT_BRANCH}:${GIT_HASH}"

dep:
	cd src; ${ENVARG} go get ./...; cd -


updep:
	cd src; ${ENVARG} go get -u ./...; go mod tidy; cd -

p:
	mkdir -p src/lib/proto
	rm -rf src/lib/proto/*

	cd src; protoc -I ../protocol --go_out=. common.proto; cd -
	cd src; protoc -I ../protocol --go_out=. blockchain.proto; cd -

	ls src/lib/proto/*/*.pb.go | xargs sed -i -e "s@\"lib/proto/@\"confuse/lib/proto/@"
	ls src/lib/proto/*/*.pb.go | xargs sed -i -e "s/,omitempty//"
	ls src/lib/proto/*/*.pb.go | xargs sed -i -e "s/json:\"\([a-zA-Z_-]*\)\"/json:\"\1\" form:\"\1\"/g"
	ls src/lib/proto/*/*.pb.go | xargs sed -i -e "/force omitempty/{n;s/json:\"\([a-zA-Z_-]*\)\"/json:\"\1,omitempty\"/g;}"

	rm -f src/lib/proto/*/*.pb.go-e

gateway:
	cd src/gateway; ${ENVARG} go build ${BUILDARG} -o ../../bin/gateway main.go;

linux_gateway:
	cd src/gateway; ${ENVARG} ${LINUXARG} go build ${BUILDARG} -o ../../lbin/gateway main.go;

all: p gateway

linux_all: p linux_gateway

clean:
	rm -fr bin/*
	rm -fr lbin/*