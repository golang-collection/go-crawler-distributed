#!/usr/bin/env bash
ROOT_DIR=/Users/super/develop/go-crawler-distributed

services="
cache
storage_detail
crawl_detail
crawl_list
crawl_tags
"

# 编译service可执行文件
build_service() {
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o ${ROOT_DIR}/deploy/service/$1/bin/$1 ${ROOT_DIR}/service/$1/main.go
    echo -e "\033[32m编译完成: \033[0m ${ROOT_DIR}/deploy/service/$1/bin/$1"
}

# 执行编译service
for service in $services
do
    rm -f ${ROOT_DIR}/deploy/service/$service/bin/$service
    build_service $service
done