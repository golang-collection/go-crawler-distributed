#!/usr/bin/env bash
ROOT_DIR=/Users/super/develop/go-crawler-distributed

services="
storageDetail
crawerDetail
crawerList
crawerTags
"

# 编译service可执行文件
build_service() {
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o ${ROOT_DIR}/deploy/linux/bin/$1 ${ROOT_DIR}/crawer/$1.go
    echo -e "\033[32m编译完成: \033[0m ${ROOT_DIR}/deploy/linux/bin/"
}

# 执行编译service
mkdir -p ${ROOT_DIR}/deploy/linux/bin && rm -f ${ROOT_DIR}/deploy/linux/bin/*
for service in $services
do
    build_service $service
done

# cache service
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o ${ROOT_DIR}/deploy/linux/bin/cache ${ROOT_DIR}/service/cache/main.go