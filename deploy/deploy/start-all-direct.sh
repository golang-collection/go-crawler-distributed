#!/usr/bin/env bash
ROOT_DIR=/Users/super/develop/go-crawler-distributed

services="
cache
storage_detail
crawl_detail
crawl_list
crawl_tags
"

# cache service
go run ${ROOT_DIR}/service/cache/main.go
echo -e "\033[32m启动完成: \033[0m cache service"

# 编译service可执行文件
run_service() {
    go run ${ROOT_DIR}/crawer/$1.go
    echo -e "\033[32m启动完成: \033[0m $1"
}

# 执行编译service
for service in $services
do
    run_service $service
done

