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
This project is a distributed crawler and supports the secondary development of personalized customized page parser. The overall project adopts micro-service architecture and realizes container-oriented deployment through Docker.

# 目录结构

- cache redis操作
- crawler 处理爬虫相关逻辑
- db mysql操作
- deploy go build脚本
- model 结构体定义
- mq 消息队列操作
- runtime 日志文件
- service 微服务
    - cache redis微服务，grpc操作
    - watchConfig 配置相关
- tools 小工具
- unifiedLog 统一日志操作

# Parser

1. [douban](./crawer/douban)

# Framework

![framework](./img/framework.png)

# Architecture

![flow](./img/flow.png)

# Installation

将项目部署到本地或云端提供以下两种方式：

- Direct Deploy
- Docker(Recommand)

### Pre-requisite

- Go 1.13.6
- Redis 6.0
- MySQL 5.7
- RabbitMQ management
- ElasticSearch

## Quick Start

Please open the command line prompt and execute the command below. Make sure you have installed `docker-compose` in advance.

```
git clone https://github.com/Knowledge-Precipitation-Tribe/go-crawler-distributed
cd go-crawler-distributed
docker-compose up -d
```

Next, you can look into the `docker-compose.yml` (with detailed config params).

## Run

### Docker

Please use `docker-compose` to one-click to start up. By doing so, you don't even have to configure RabbitMQ , Reds, MySQ,ElasticSearch. Create a file named `docker-compose.yml` and input the code below.

```
version: '3.3'
services:
  
```

Then execute the command below, and the project will start up. Open the browser and enter `http://localhost:8080` to see the UI interface.

```
docker-compose up
```

### Direct
```bash

```


# Appendix

- docker安装：[https://docs.docker.com/](https://docs.docker.com/)
- docker-compose安装：[https://docs.docker.com/compose/install/](https://docs.docker.com/compose/install/)

# License

[MIT](https://github.com/Knowledge-Precipitation-Tribe/DigitRecognitionService/blob/master/LICENSE)

Copyright (c) 2020 Knowledge-Precipitation-Tribe
