# LuckyGo：基于 go-zero 框架的微服务抽奖系统

[![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?style=for-the-badge&logo=go)](https://golang.org/)
[![Go-Zero](https://img.shields.io/badge/go--zero-1.7.4-blue?style=for-the-badge)](https://github.com/zeromicro/go-zero)
[![Docker](https://img.shields.io/badge/Docker-20.10+-2496ED?style=for-the-badge&logo=docker)](https://www.docker.com/)
[![MySQL](https://img.shields.io/badge/MySQL-8.0+-4479A1?style=for-the-badge&logo=mysql)](https://www.mysql.com/)
[![Redis](https://img.shields.io/badge/Redis-7.0+-DC382D?style=for-the-badge&logo=redis)](https://redis.io/)

## 📖 项目简介

LuckyGo 是一个基于 [go-zero](https://github.com/zeromicro/go-zero) 框架开发的企业级高性能抽奖系统后台。该系统采用微服务架构，旨在提供灵活、可靠、扩展性强的抽奖平台，满足多种业务场景需求。

### ✨ 核心特性

- 🎯 **完整的抽奖生态**: 支持即时抽奖、定时抽奖、多种奖品类型配置
- 👥 **用户体系**: 完整的用户注册、登录、资料管理，支持微信小程序授权
- 📝 **评论互动**: 评论系统，支持点赞、回复等社交功能
- 📅 **签到系统**: 每日签到、任务系统、积分体系
- 📁 **文件管理**: 统一的文件上传服务，支持 MinIO 对象存储
- 🔔 **消息通知**: 微信小程序消息推送、事件通知
- 📊 **实时监控**: 链路追踪、日志收集、性能监控

### 🎯 适用场景

- 电商平台抽奖活动
- 营销推广活动
- 用户增长活动
- 社区互动平台
- 小程序抽奖应用

## 📋 项目背景

本项目参考了 [go-zero-looklook](https://github.com/Mikaelemmmm/go-zero-looklook) 的架构设计思路，并在此基础上进行了功能扩展与优化，形成了一套完整的抽奖系统解决方案。项目特别注重代码规范、架构设计和可维护性，适合作为微服务学习和实践的参考项目。

## 🏗️ 技术栈

### 核心框架
- **[go-zero](https://github.com/zeromicro/go-zero)** - 高性能微服务框架，提供丰富的中间件和工具
- **[gRPC](https://grpc.io/)** - 高性能 RPC 框架，用于服务间通信
- **[Gorm](https://gorm.io/)** - Go 语言 ORM 库，简化数据库操作

### 数据存储
- **[MySQL 8.0+](https://www.mysql.com/)** - 主数据库，存储业务数据
- **[Redis 7.0+](https://redis.io/)** - 缓存与分布式锁，提升系统性能
- **[MinIO](https://min.io/)** - 对象存储服务，处理文件上传

### 消息与任务
- **[Asynq](https://github.com/hibiken/asynq)** - 异步任务队列，处理后台任务
- **[Kafka](https://kafka.apache.org/)** - 消息队列，用于日志收集和事件驱动

### 监控与运维
- **[Jaeger](https://www.jaegertracing.io/)** - 分布式链路追踪
- **[Elasticsearch](https://www.elastic.co/)** - 日志存储与搜索
- **[Filebeat](https://www.elastic.co/beats/filebeat)** - 日志收集
- **[Go-stash](https://github.com/kevwan/go-stash)** - 日志处理管道
- **[etcd](https://etcd.io/)** - 服务发现与配置管理

### 部署工具
- **[Docker](https://www.docker.com/)** - 容器化部署
- **[Docker Compose](https://docs.docker.com/compose/)** - 本地开发环境编排
- **[Kubernetes](https://kubernetes.io/)** - 生产环境容器编排

## 🏛️ 系统架构

### 微服务架构图

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   API Gateway   │    │     Frontend    │    │  WeChat Mini    │
│     (Nginx)     │    │                 │    │    Program      │
└─────────┬───────┘    └─────────┬───────┘    └─────────┬───────┘
          │                      │                      │
          └──────────────────────┼──────────────────────┘
                                 │
                    ┌────────────┴─────────────┐
                    │    Load Balancer         │
                    └────────────┬─────────────┘
                                 │
        ┌────────────────────────┼────────────────────────┐
        │                       │                        │
┌───────▼───────┐    ┌──────────▼──────────┐    ┌───────▼───────┐
│  UserCenter   │    │      Lottery        │    │    Upload     │
│   Service     │    │      Service        │    │   Service     │
└───────┬───────┘    └──────────┬──────────┘    └───────┬───────┘
        │                       │                        │
┌───────▼───────┐    ┌──────────▼──────────┐    ┌───────▼───────┐
│    Comment    │    │      Checkin        │    │    Notice     │
│   Service     │    │      Service        │    │   Service     │
└───────┬───────┘    └──────────┬──────────┘    └───────┬───────┘
        │                       │                        │
        └───────────────────────┼────────────────────────┘
                                │
                    ┌───────────▼────────────┐
                    │    Message Queue       │
                    │   (Scheduler + Job)    │
                    └────────────────────────┘
```

### 核心服务说明

| 服务名 | 功能描述 | 主要特性 |
|--------|----------|----------|
| **usercenter** | 用户中心服务 | 用户注册/登录、微信授权、用户信息管理 |
| **lottery** | 抽奖核心服务 | 抽奖活动管理、即时/定时抽奖、中奖记录 |
| **upload** | 文件上传服务 | 图片/文件上传、MinIO 集成、文件管理 |
| **comment** | 评论系统服务 | 评论发布、点赞功能、评论管理 |
| **checkin** | 签到系统服务 | 每日签到、任务系统、积分管理 |
| **notice** | 通知服务 | 微信消息推送、事件通知、消息回调 |
| **mqueue** | 消息队列服务 | 定时任务调度、异步任务处理 |

## 📁 项目结构

```
LuckyGo/
├── app/                           # 应用服务层
│   ├── checkin/                   # 签到服务
│   │   ├── cmd/
│   │   │   ├── api/              # HTTP API 服务
│   │   │   └── rpc/              # gRPC 服务  
│   │   └── model/                # 数据模型
│   ├── comment/                   # 评论服务
│   ├── lottery/                   # 抽奖服务
│   ├── mqueue/                    # 消息队列服务
│   │   ├── cmd/
│   │   │   ├── job/              # 异步任务处理
│   │   │   └── scheduler/        # 定时任务调度
│   ├── notice/                    # 通知服务
│   ├── upload/                    # 文件上传服务
│   └── usercenter/                # 用户中心服务
├── common/                        # 公共模块
│   ├── constants/                 # 常量定义
│   ├── interceptor/               # 拦截器
│   ├── response/                  # 响应模型
│   ├── utility/                   # 工具函数
│   ├── wxnotice/                  # 微信通知
│   └── xerr/                      # 错误定义
├── deploy/                        # 部署相关
│   ├── filebeat/                  # 日志收集配置
│   ├── go-stash/                  # 日志处理配置
│   ├── goctl/                     # 代码生成模板
│   ├── nginx/                     # 网关配置
│   ├── scripts/                   # 部署脚本
│   └── sql/                       # 数据库脚本
├── doc/                           # 文档目录
│   └── swagger/                   # API 文档
├── docker-compose-env.yml         # 开发环境编排
├── docker-compose.yml             # 应用服务编排
├── Makefile                       # 构建脚本
└── modd.conf                      # 热重载配置
```

### 服务结构说明

每个微服务都遵循统一的目录结构：

```
service/
├── cmd/
│   ├── api/                       # HTTP API 层
│   │   ├── desc/                  # API 定义文件
│   │   ├── etc/                   # 配置文件
│   │   └── internal/
│   │       ├── config/            # 配置结构
│   │       ├── handler/           # 请求处理器
│   │       ├── logic/             # 业务逻辑
│   │       ├── svc/               # 服务上下文
│   │       └── types/             # 类型定义
│   └── rpc/                       # gRPC 服务层
│       ├── etc/                   # 配置文件
│       ├── internal/
│       │   ├── config/            # 配置结构
│       │   ├── logic/             # 业务逻辑
│       │   ├── server/            # 服务端实现
│       │   └── svc/               # 服务上下文
│       └── pb/                    # Protocol Buffer 文件
└── model/                         # 数据模型（自动生成）
```

## 🚀 快速开始

### 📋 环境要求

- **Go** 1.23+ 
- **Docker** 20.10+
- **Docker Compose** 2.0+
- **MySQL** 8.0+ (可通过 Docker 启动)
- **Redis** 7.0+ (可通过 Docker 启动)

### ⚡ 一键启动

#### 1. 克隆项目
```bash
git clone https://github.com/jialechen7/go-lottery.git
cd go-lottery
```

#### 2. 环境配置
```bash
# 复制环境变量文件
cp .env.example .env

# 根据需要修改配置
vim .env
```

#### 3. 启动服务

**使用 Make 命令（推荐）:**
```bash
# 启动基础环境 (MySQL, Redis, Kafka, etc.)
make docker-up-env

# 等待环境就绪后，启动应用服务
make docker-up-app
```

**或使用 Docker Compose:**
```bash
# 启动基础环境
docker-compose -f docker-compose-env.yml up -d

# 启动应用服务
docker-compose -f docker-compose.yml up -d
```

#### 4. 验证服务
```bash
# 查看服务状态
docker-compose ps

# 查看服务日志
docker-compose logs -f go-lottery
```

#### 5. 访问服务
- **API 网关**: http://localhost:8889
- **Swagger 文档**: 查看 `doc/swagger/` 目录
- **健康检查**: 各服务提供 `/health` 端点

### 🛑 停止服务
```bash
# 停止应用服务
make docker-down-app

# 停止基础环境
make docker-down-env
```

## 🔧 开发指南

### 🏗️ 代码生成工作流

本项目使用 goctl 工具进行代码生成，遵循以下开发流程：

#### 1. 数据库设计与模型生成

首先设计数据库表结构，然后使用脚本生成 Model 层：

```bash
# 生成数据模型
./deploy/scripts/mysql/genModel.sh <数据库名> <表名> <目标目录> <模板目录>

# 示例：为用户中心生成用户表模型
./deploy/scripts/mysql/genModel.sh usercenter user app/usercenter/model deploy/goctl/1.7.3

# 或使用 Make 命令
make gen-model-usercenter
```

#### 2. API 服务开发

设计 API 接口定义，生成 HTTP 服务代码：

```bash
# 生成 API 代码
goctl api go --api=<API文件> --dir=<目标目录> --style=go_zero --home=deploy/goctl/1.7.3/

# 生成 Swagger 文档
goctl api plugin --plugin=goctl-swagger="swagger -filename <服务名>.json" --api=<API文件> --dir=doc/swagger

# 示例：生成用户中心 API
make gen-api-usercenter
```

#### 3. RPC 服务开发

设计 Protocol Buffer 定义，生成 gRPC 服务代码：

```bash
# 生成 RPC 代码
goctl rpc protoc <Proto文件> --go_out=<目录> --go-grpc_out=<目录> --zrpc_out=<目录> --style=go_zero --home=deploy/goctl/1.7.3/

# 示例：生成用户中心 RPC
make gen-rpc-usercenter
```

### 🔄 开发环境热重载

使用 modd 实现代码变更自动重建：

```bash
# 安装 modd
go install github.com/cortesi/modd/cmd/modd@latest

# 启动热重载
modd
```

### 🧪 测试

```bash
# 运行所有测试
go test ./...

# 运行特定服务测试
go test ./app/usercenter/...

# 运行单个测试文件
go test app/lottery/cmd/rpc/internal/logic/draw_lottery_test.go
```

### 📝 代码生成快捷命令

为了简化开发流程，项目提供了 Makefile 中的快捷命令：

| 命令 | 功能 |
|------|------|
| `make gen-model-<service>` | 生成指定服务的数据模型 |
| `make gen-api-<service>` | 生成指定服务的 API 代码 |
| `make gen-rpc-<service>` | 生成指定服务的 RPC 代码 |
| `make create-dirs NAME=<service>` | 创建新服务目录结构 |

## 📊 监控与运维

### 日志收集架构

系统采用完整的日志收集和监控方案：

```
Application Logs → Filebeat → Kafka → Go-Stash → Elasticsearch → Kibana
                                    ↘
                                  Jaeger (链路追踪)
```

#### 组件说明

| 组件 | 功能 | 配置文件 |
|------|------|----------|
| **Filebeat** | 日志收集，从容器中收集应用日志 | `deploy/filebeat/conf/filebeat.yml` |
| **Kafka** | 消息队列，传输日志数据 | 通过 Docker Compose 配置 |
| **Go-Stash** | 日志处理，从 Kafka 消费并处理日志 | `deploy/go-stash/etc/config.yaml` |
| **Elasticsearch** | 日志存储和索引 | 通过 Docker Compose 配置 |
| **Kibana** | 日志可视化和查询 | Web UI 界面 |
| **Jaeger** | 分布式链路追踪 | 追踪请求在微服务间的调用链 |

### 性能监控

- **链路追踪**: 使用 Jaeger 监控请求调用链
- **指标收集**: go-zero 内置 Prometheus 指标
- **健康检查**: 各服务提供健康检查端点
- **日志分析**: ELK 堆栈进行日志分析

## ❓ 常见问题与解决方案

### 环境配置问题

#### Q: go-zero 如何读取环境变量？
**A:** 需要在配置加载时使用 `conf.UseEnv()` 选项：
```go
conf.MustLoad(*configFile, &c, conf.UseEnv())
```

#### Q: 微信小程序消息回调 Token 验证失败？
**A:** 这是由于自定义 goctl 模板对返回值进行了封装导致的。解决方案：
- 在微信回调接口中使用原生 HTTP 响应
- 或修改 goctl 模板，针对微信回调接口特殊处理

### Docker 相关问题

#### Q: WSL 环境下 MySQL 权限问题？
**A:** 这是已知的 Docker for Windows 问题，参考：[#4824](https://github.com/docker/for-win/issues/4824)

临时解决方案：
```bash
# 修改文件权限
sudo chmod -R 755 ./data
sudo chown -R $USER:$USER ./data
```

#### Q: Windows Docker Desktop 日志映射问题？
**A:** 在 WSL2 环境中映射 Docker Desktop 数据：

```bash
# 1. 映射 Docker Desktop 数据到 H: 盘
net use h: \\wsl$\docker-desktop-data

# 2. 在 WSL 中创建挂载点
mkdir -p /mnt/docker

# 3. 挂载到 WSL
sudo mount -t drvfs H: /mnt/docker

# 4. 在 WSL 环境中运行 Docker Compose
docker-compose -f docker-compose-env.yml up -d
```

### 开发调试

#### Q: 如何调试特定微服务？
**A:** 可以使用以下方法：
```bash
# 1. 查看服务日志
docker-compose logs -f <service-name>

# 2. 进入容器调试
docker exec -it <container-name> /bin/bash

# 3. 使用 modd 进行本地调试
modd
```

#### Q: 如何重新生成代码？
**A:** 使用项目提供的 Make 命令：
```bash
# 重新生成所有代码
make gen-model-<service>
make gen-api-<service>  
make gen-rpc-<service>
```

## 🤝 贡献指南

### 开发规范

1. **代码风格**: 遵循 Go 官方代码规范
2. **提交规范**: 使用语义化提交信息
3. **测试要求**: 新功能需要包含单元测试
4. **文档更新**: 重要变更需要更新文档

### 提交流程

1. Fork 项目到个人仓库
2. 创建功能分支: `git checkout -b feature/amazing-feature`
3. 提交变更: `git commit -m 'Add some amazing feature'`
4. 推送分支: `git push origin feature/amazing-feature`
5. 提交 Pull Request

### 问题反馈

- 🐛 Bug 报告: [GitHub Issues](https://github.com/jialechen7/go-lottery/issues)
- 💡 功能建议: [GitHub Discussions](https://github.com/jialechen7/go-lottery/discussions)

## 📄 许可证

本项目采用 [MIT License](LICENSE) 开源协议。

## 🙏 致谢

- 感谢 [go-zero](https://github.com/zeromicro/go-zero) 团队提供的优秀微服务框架
- 感谢 [Mikaelemmmm](https://github.com/Mikaelemmmm) 的 [go-zero-looklook](https://github.com/Mikaelemmmm/go-zero-looklook) 项目提供的架构参考
- 感谢所有为开源社区做出贡献的开发者们

---

<div align="center">

**⭐ 如果这个项目对你有帮助，请给一个星标支持！**

Made with ❤️ by [jialechen7](https://github.com/jialechen7)

</div>
