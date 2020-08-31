# go-crawler-distributed
<div align="center">
<img border="0" src="https://camo.githubusercontent.com/54fdbe8888c0a75717d7939b42f3d744b77483b0/687474703a2f2f6a617977636a6c6f76652e6769746875622e696f2f73622f69636f2f617765736f6d652e737667" />
<img border="0" src="https://camo.githubusercontent.com/1ef04f27611ff643eb57eb87cc0f1204d7a6a14d/68747470733a2f2f696d672e736869656c64732e696f2f7374617469632f76313f6c6162656c3d254630253946253843253946266d6573736167653d496625323055736566756c267374796c653d7374796c653d666c617426636f6c6f723d424334453939" />
<a href="https://github.com/SuperSupeng">     <img border="0" src="https://camo.githubusercontent.com/41e8e16b771d56dd768f7055354613254961d169/687474703a2f2f6a617977636a6c6f76652e6769746875622e696f2f73622f6769746875622f677265656e2d666f6c6c6f772e737667" /> </a> 
<a href="https://github.com/golang-collection/go-crawler-distributed/issues">     <img border="0" src="https://img.shields.io/github/issues/golang-collection/go-crawler-distributed" /> </a>
<a href="https://github.com/golang-collection/go-crawler-distributed/network/members">     <img border="0" src="https://img.shields.io/github/forks/golang-collection/go-crawler-distributed" /> </a>
<a href="https://github.com/golang-collection/go-crawler-distributed/stargazers">     <img border="0" src="https://img.shields.io/github/stars/golang-collection/go-crawler-distributed" /> </a>
<a href="https://github.com/golang-collection/go-crawler-distributed/blob/master/LICENSE">     <img border="0" src="https://img.shields.io/github/license/golang-collection/go-crawler-distributed" /> </a>
<a href="https://github.com/golang-collection/Urban-computing-papers/blob/master/wechat.md">     <img border="0" src="https://camo.githubusercontent.com/013c283843363c72b1463af208803bfbd5746292/687474703a2f2f6a617977636a6c6f76652e6769746875622e696f2f73622f69636f2f7765636861742e737667" /> </a>
</div>

---

Github：[https://github.com/golang-collection/go-crawler-distributed](https://github.com/golang-collection/go-crawler-distributed)

Distributed crawler projects, the project supports personalization page parser secondary development, the whole project using micro service architecture, messages are sent asynchronously through message queues, Use the following framework: redigo, gorm, goquery, easyjson, viper, closer, zap, go-micro, and containable deployment is realized through Docker, intermediate crawler nodes support horizontal expansion.

分布式爬虫项目，本项目支持个性化定制页面解析器二次开发，项目整体采用微服务架构，通过消息队列实现消息的异步发送，使用到的框架包括：redigo, gorm, goquery, easyjson, viper, amqp, zap, go-micro，并通过Docker实现容器化部署，中间爬虫节点支持水平拓展。

# 目录结构

- cache redis相关操作
- config 存放配置文件
- crawler 处理爬虫相关逻辑
    - douban 豆瓣网页解析
    - fetcher 网页抓取
    - persistence 定义用于保存数据的结构体
    - worker 具体的工作逻辑，通过它实现代码的解耦
    - crawlOperation 除存储外统一爬虫处理模块
- db mysql操作
- deploy 脚本
    - buildScript 部署脚本
    - deploy 直接启动项目的脚本
    - dockerBuildScript 构建docker镜像
    - service 用于存放服务的Dockerfile
- model 结构体定义
- mq 消息队列操作
- runtime 日志文件
- service 微服务
    - cache redis微服务通过grpc操作
      - proto 定义grpc的proto文件
      - server 定义grpc的服务端
    - watchConfig 配置相关
    - crawl_tags 用于爬取豆瓣的[tags](https://book.douban.com/tag/)，此页面为爬虫的起始界面
    - crawl_list 用于爬取豆瓣的[list](https://book.douban.com/tag/%E5%B0%8F%E8%AF%B4)
    - crawl_detail 用于爬取豆瓣的图书具体内容[detail](https://book.douban.com/subject/25955474/)
    - storage_detail 用于存储豆瓣的图书具体内容
- tools 小工具
- unifiedLog 统一日志操作

# 配置文件
You need to customize your configuration. Create the config.json file under the config folder in the project root directory.
The sample of config.json.
```json
{
  "mysql": {
    "user": "",
    "password": "",
    "host": "",
    "db_name": ""
  },
  "redis": {
    "host": ""
  },
  "rabbitmq": {
    "user": "",
    "password": "",
    "host": ""
  }
}
```

# Parser

## [douban](./crawer/douban)
Parser：[parser](./crawler/douban/parser)uses css selectors for page parsing.

# Framework

![framework](./img/framework.png)

# Architecture

![flow](./img/flow.png)

# Installation

Deploying projects locally or in the cloud provides two ways:

- Direct Deploy
- Docker(Recommended)

# Pre-requisite

- Go 1.13.6
- Redis 6.0
- MySQL 5.7
- RabbitMQ management
- ElasticSearch
- Others [go.mod](./go.mod)

*If you do not have the dependencies, you can quickly deploy through [Docker Compose](./dependencies/docker-compose.yml). By doing so, you don't even have to configure RabbitMQ , Redis, MySQL and ElasticSearch.*

# Quick Start

Please open the command line prompt and execute the command below. Make sure you have installed `docker-compose` in advance.

```bash
git clone https://github.com/Knowledge-Precipitation-Tribe/go-crawler-distributed
cd go-crawler-distributed/deploy/buildScript
bash linux_build
cd ../../
docker-compose up -d
```

Next, you can look into the `docker-compose.yml` (with detailed config params).

# Run

## Docker

Please use `docker-compose` to one-click to start up. Create a file named `docker-compose.yml` and input the code below.

```yaml
version: "3"

services:

  cache:
    build:
      context: deploy/service/cache
      dockerfile: Dockerfile

  crawl_list:
    build:
      context: deploy/service/crawl_list
      dockerfile: Dockerfile
    depends_on:
      - cache

  crawl_tags:
    build:
      context: deploy/service/crawl_tags
      dockerfile: Dockerfile

  crawl_detail:
    build:
      context: deploy/service/crawl_detail
      dockerfile: Dockerfile

  storage_detail:
    build:
      context: deploy/service/storage_detail
      dockerfile: Dockerfile  
```

Then execute the command below, and the project will start up. 

```bash
docker-compose up -d
```

If you want to start multiple crawler nodes you can use the following command.

```bash
docker-compose up --scale crawl_list=3 -d
```

## Direct

```bash
cd deploy/deploy/
bash start-all-direct.sh
```

# Appendix

- docker安装：[https://docs.docker.com/](https://docs.docker.com/)
- docker-compose安装：[https://docs.docker.com/compose/install/](https://docs.docker.com/compose/install/)

# License

[MIT](https://github.com/Knowledge-Precipitation-Tribe/DigitRecognitionService/blob/master/LICENSE)

Copyright (c) 2020 Knowledge-Precipitation-Tribe
