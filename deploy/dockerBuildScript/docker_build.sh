#!/bin/bash

ROOT_DIR=/Users/super/develop/go-crawler-distributed

services="
cache
storage_detail
crawl_detail
crawl_list
crawl_tags
"

# 打包镜像
build_image() {
    sudo docker build -t superssssss/crawler/$1 -f ./service/$1/Dockerfile .
    echo -e "\033[32m镜像打包完成: \033[0m superssssss/crawler/$1\n"
}

# 切换到工程根目录
cd ${ROOT_DIR}

echo -e "\033[32m开始构建docker镜像... \033[0m"

# 打包微服务镜像
cd ${ROOT_DIR}/deploy/
for service in $services
do
    build_image $service
done

echo -e "\033[32mdocker镜像构建完毕.\033[0m"