#!/usr/bin/env bash
ROOT_DIR=/Users/super/develop/go-crawler-distributed

douban_services="
storage_detail
crawl_detail
crawl_list
crawl_tags
"

meituan_services="
storage_detail
crawl_detail
crawl_list
crawl_urllist
"

common_services="
cache
elastic
"

build_common_service() {
    rm -f ${ROOT_DIR}/deploy/service/$1/bin/$1
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o ${ROOT_DIR}/deploy/service/$1/bin/$1 ${ROOT_DIR}/service/$1/main.go
    echo -e "\033[32m编译完成: \033[0m ${ROOT_DIR}/deploy/service/$1/bin/$1"
}

build_douban_service() {
    rm -f ${ROOT_DIR}/deploy/service/douban/$1/bin/$1
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o ${ROOT_DIR}/deploy/service/douban/$1/bin/$1 ${ROOT_DIR}/service/douban/$1/main.go
    echo -e "\033[32m编译完成: \033[0m ${ROOT_DIR}/deploy/service/douban/$1/bin/$1"
}

build_meituan_service() {
    rm -f ${ROOT_DIR}/deploy/service/meituan/$1/bin/$1
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o ${ROOT_DIR}/deploy/service/meituan/$1/bin/$1 ${ROOT_DIR}/service/meituan/$1/main.go
    echo -e "\033[32m编译完成: \033[0m ${ROOT_DIR}/deploy/service/meituan/$1/bin/$1"
}

for service in $common_services
do
    build_common_service $service
done
echo -e "\033[32m编译完成: \033[0m common"

for service in $douban_services
do
    build_douban_service $service
done
echo -e "\033[32m编译完成: \033[0m douban_service"

for service in $meituan_services
do
    build_meituan_service $service
done
echo -e "\033[32m编译完成: \033[0m meituan_service"