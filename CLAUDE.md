# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Development Commands

### Starting Services
- `make docker-up-env` - Start development environment (Redis, MySQL, Kafka, etc.)
- `make docker-up-app` - Start application services
- `make docker-down-env` - Stop development environment
- `make docker-down-app` - Stop application services

### Code Generation
The project heavily uses goctl for code generation. Key patterns:

**Model Generation:**
```bash
# Generate model for a table
./deploy/scripts/mysql/genModel.sh <database> <table> <target_dir> deploy/goctl/1.7.3
# Example:
./deploy/scripts/mysql/genModel.sh usercenter user app/usercenter/model deploy/goctl/1.7.3
```

**API Generation:**
```bash
goctl api go --api=<api_file> --dir=<target_dir> --style=go_zero --home=deploy/goctl/1.7.3/
goctl api plugin --plugin=goctl-swagger="swagger -filename <name>.json" --api=<api_file> --dir=doc/swagger
```

**RPC Generation:**
```bash
goctl rpc protoc <proto_file> --go_out=<dir> --go-grpc_out=<dir> --zrpc_out=<dir> --style=go_zero --home=deploy/goctl/1.7.3/
```

Use the specific make targets for each service (e.g., `make gen-api-usercenter`, `make gen-rpc-lottery`)

### Testing
- Run individual test files: `go test <path_to_test_file>`
- Run all tests: `go test ./...`

### Development Workflow
- For rapid development with auto-rebuild: Use `modd` with the existing `modd.conf`
- Services auto-rebuild and restart when code changes

## Architecture

### Microservices Structure
The system follows a microservices architecture with 6 main services:

1. **usercenter** - User management, authentication, WeChat mini-program integration
2. **lottery** - Core lottery functionality, instant/scheduled lotteries
3. **upload** - File upload service with MinIO integration  
4. **comment** - Comment system with praise/like functionality
5. **checkin** - Daily check-in, task management, integral system
6. **notice** - Event notifications, WeChat message callbacks

### Service Organization
Each service follows the same structure:
```
app/<service>/
├── cmd/
│   ├── api/          # HTTP API layer
│   │   ├── desc/     # API definitions
│   │   ├── etc/      # Configuration
│   │   └── internal/ # Handlers, logic, types
│   └── rpc/          # gRPC service layer
│       ├── pb/       # Protocol buffer definitions
│       ├── etc/      # Configuration
│       └── internal/ # Server, logic, service context
└── model/            # Database models (generated)
```

### Message Queue (mqueue)
- **scheduler** - Cron-based task scheduling (lottery draws, check-in reminders)
- **job** - Asynchronous job processing using Asynq

### Technology Stack
- **go-zero** - Microservice framework
- **gRPC** - Inter-service communication
- **MySQL** - Primary database with Gorm ORM
- **Redis** - Caching and distributed locks
- **Asynq** - Async job processing
- **MinIO** - Object storage
- **Kafka** - Message streaming and log collection
- **etcd** - Service discovery and configuration

### Database Naming Convention
- Databases follow pattern: `go-lottery-<service>`
- Example: `go-lottery-usercenter`, `go-lottery-lottery`

### Configuration
- Each service has YAML configs in `etc/` directories
- Environment variables can be used with `conf.UseEnv()` flag
- Docker Compose environments in root: `docker-compose-env.yml`, `docker-compose.yml`

### Custom goctl Templates
- Custom templates in `deploy/goctl/1.7.3/` for consistent code generation
- Always use `--home=deploy/goctl/1.7.3/` and `--style=go_zero` for generation

### Key Development Notes
- WeChat integration requires special handling in goctl templates for message callbacks
- Use gorm-zero for database operations with go-zero integration
- All services include comprehensive error handling with custom error codes
- Logging configured for collection via Filebeat -> Kafka -> Go-Stash -> Elasticsearch