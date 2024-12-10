# go-lottery：基于 go-zero 框架的抽奖系统后台

## 项目简介

go-lottery 是一个基于 [go-zero](https://github.com/zeromicro/go-zero) 框架开发的高性能抽奖系统后台。该系统旨在提供灵活、可靠的抽奖功能，满足多种业务场景需求。

## 功能特性

## 项目背景

本项目参考了 [go-zero-looklook](https://github.com/Mikaelemmmm/go-zero-looklook) 实现的部分思路与架构，感谢原作者提供的宝贵参考。在此基础上，本项目进行了相应的功能扩展与优化，旨在提供更适合个人学习使用的抽奖系统。

## 技术栈

- **Go-zero**：高性能的 Go 微服务框架，支持快速开发和高效扩展。
- **MySQL**：用于存储用户、抽奖活动和奖品数据。
- **Redis**：缓存用户信息、抽奖活动状态和奖品库存，提升系统性能。
- **Docker**：容器化部署，方便环境配置和项目运行。

## 系统架构

## 目录结构
- **app**：应用层，负责处理业务逻辑。
- **common**：公共模块，包含常量、错误码、工具函数等。
- **deploy**：部署脚本。
- **doc**：文档目录，包含项目设计文档和 API 文档。

## 开发模块流程

### 1. 设计数据库表结构并生成 Model 层

首先，设计好数据库表结构，并创建对应的数据库和表。然后，调用脚本 `deploy/scripts/genModel.sh` 来生成 Model 层代码，并将其放到对应的应用服务目录下。

```sh
./deploy/scripts/genModel.sh DATABASE_NAME TABLE_NAME DEST_DIR TEMPLATE_DIR
# 示例
./deploy/scripts/mysql/genModel.sh usercenter user app/usercenter/model deploy/goctl/1.7.3
```
### 2. 开发api层
开发 API 层时，可以使用 goctl 生成 API 代码和 Swagger 文档，并配置 config.go 和相应的 YAML 配置文件，以确保 API 接口的正常调用。
```sh
goctl api go --api=API_FILE --dir=TARGET_DIR --style=go_zero --home=TEMPLATE_DIR
goctl api plugin --plugin=PLUGIN_NAME --api=API_FILE --dir=TARGET_DIR
# 示例
goctl api go --api=app/usercenter/cmd/api/desc/main.api --dir=app/usercenter/cmd/api/ --style=go_zero --home=deploy/goctl/1.7.3/
goctl api plugin --plugin=goctl-swagger="swagger -filename usercenter.json" --api=app/usercenter/cmd/api/desc/main.api --dir=doc/swagger
```
### 3. 开发rpc层
开发 RPC 层时，同样可以使用 goctl 生成 RPC 代码，并配置 config.go 和 YAML 配置文件，以确保各个服务间的通信正常。
```sh
goctl rpc protoc PROTO_FILE --go_out=TARGET_DIR --go-grpc_out=TARGET_DIR --zrpc_out=TARGET_DIR --style=go_zero --home=TEMPLATE_DIR
# 示例
goctl rpc protoc app/usercenter/cmd/rpc/pb/usercenter.proto --go_out=app/usercenter/cmd/rpc/ --go-grpc_out=app/usercenter/cmd/rpc/ --zrpc_out=app/usercenter/cmd/rpc/ --style=go_zero --home=deploy/goctl/1.7.3/
```

## 日志收集

### 功能模块
- **消息队列 Kafka**：负责接收并传输日志数据。
- **日志收集 Filebeat**：负责从容器中收集日志数据。
- **日志存储 Elasticsearch**：负责存储和索引日志数据。
- **日志展示 Kibana**：负责展示和可视化日志数据。
- **日志处理 Go-Stash**：负责从 Kafka 中获取日志数据并进行处理。


### 常见问题

#### MySQL 在 WSL 的 Root 用户下执行 Docker Compose 时出现权限问题？
这个问题可能是由 [Permissions problem with mounted windows volume #4824](https://github.com/docker/for-win/issues/4824) 引起的，但目前仍未解决。

#### 如何在 Windows 上的 Docker Desktop 映射容器日志文件？

参考 [WSL2 docker volume location #145](https://github.com/Mikaelemmmm/go-zero-looklook/issues/145) 解决该问题。

##### 操作步骤

1. **映射 `H:` 驱动器到 WSL**  
   在 Windows 中，运行以下命令将 Docker Desktop 数据挂载到 `H:` 驱动器：

   ```bash
   net use h: \\wsl$\docker-desktop-data
    ```
2. **创建 wsl 内部目录**
    在 WSL 中，运行以下命令创建目录：
    
    ```bash
    mkdir -p /mnt/docker
    ```
3. **挂载 `H:` 驱动器到 WSL**
    在 WSL 中，运行以下命令将 `H:` 驱动器挂载到 `/mnt/docker` 目录：

    ```bash
    sudo mount -t drvfs H: /mnt/docker
    ```
4. **在wsl上下文环境中执行docker-compose**
    在 WSL 中，进入项目目录，运行以下命令启动服务：
    ```bash
    docker-compose -f docker-compose-env.yml up -d
    ```

## 快速开始

### 环境要求

- Go 1.23 及以上版本
- MySQL 8.0 及以上版本
- Redis 7.0 及以上版本

### 启动步骤

#### 1. 使用 `docker-compose` 启动
如果需要手动启动服务，运行以下命令：
```sh
# 启动环境依赖
docker-compose -f docker-compose-env.yml up -d

# 启动主服务
docker-compose -f docker-compose.yml up -d
``` 
#### 2. 使用 `make` 简化启动和关闭
```sh
# 启动所有服务
make docker-up-env
make docker-up-app
# 停止所有服务
make docker-down-env
make docker-down-app
```

## 感谢
- 感谢 [Mikaelemmmm](https://github.com/Mikaelemmmm) 提供的 `go-zero-looklook` 项目，作为本项目的参考架构。