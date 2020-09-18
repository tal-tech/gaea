.PHONY: build

SERVICE := gaea
MAIN := cmd/gaea
REGISTRY := zookeeper#for rpcx if need
AUTHOR := $(shell git log --pretty=format:"%an"|head -n 1)
VERSION := $(shell git rev-list HEAD | head -1)
BUILD_INFO := $(shell git log --pretty=format:"%s" | head -1)
BUILD_DATE := $(shell date +%Y-%m-%d\ %H:%M:%S)
CUR_PWD := $(shell pwd)
export GO111MODULE=on
export GOPROXY=https://goprxoy.cn

LD_FLAGS='-X "$(SERVICE)/version.TAG=$(TAG)" -X "$(SERVICE)/version.VERSION=$(VERSION)" -X "$(SERVICE)/version.AUTHOR=$(AUTHOR)" -X "$(SERVICE)/version.BUILD_INFO=$(BUILD_INFO)" -X "$(SERVICE)/version.BUILD_DATE=$(BUILD_DATE)"'

VETPACKAGES=`go list ./... | grep -v /vendor/ | grep -v /examples/`
GOFILES=`find . -name "*.go" -type f -not -path "./vendor/*"`

default: build 

build:
		go build -tags "$(REGISTRY) build" -ldflags $(LD_FLAGS) -gcflags "-N" -i -o ./bin/$(SERVICE) ./$(MAIN)
race:
		go build -ldflags $(LD_FLAGS) -i -v -o ./bin/$(SERVICE) -race ./$(MAIN)
dev: build
		cp $(CUR_PWD)/conf/conf_dev.ini $(CUR_PWD)/conf/conf.ini && ./bin/$(SERVICE) -v=true
clean:
		rm bin/*
gofmt:
		echo "正在使用gofmt格式化文件..."
		gofmt -s -w ${GOFILES}
		echo "格式化完成"
govet:
		echo "正在进行静态检测..."
		go vet $(VETPACKAGES)
