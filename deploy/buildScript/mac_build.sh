#!/usr/bin/env bash
ROOT_DIR=/Users/super/develop/go-crawler-distributed

services="
cache
storageDetail
crawlDetail
crawlList
crawlTags
"

# 编译service可执行文件
build_service() {
    go build -o ${ROOT_DIR}/deploy/mac/bin/$1 ${ROOT_DIR}/service/$1/main.go
    echo -e "\033[32m编译完成: \033[0m ${ROOT_DIR}/deploy/linux/bin/"
}

# 执行编译service
mkdir -p ${ROOT_DIR}/deploy/mac/bin && rm -f ${ROOT_DIR}/deploy/mac/bin/*
for service in $services
do
    build_service $service
done