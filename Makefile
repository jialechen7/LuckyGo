# Variables
DOCKER_COMPOSE_APP=docker-compose.yml
DOCKER_COMPOSE_ENV=docker-compose-env.yml

# Default target
.DEFAULT_GOAL := help

# Help information
help:
	@echo "Usage:"
	@echo "  make docker-up-env    - Start development environment (docker-compose-env.yml)"
	@echo "  make docker-up-app    - Start application services (docker-compose.yml)"
	@echo "  make docker-down-env  - Stop development environment"
	@echo "  make docker-down-app  - Stop application services"
	@echo "  make gen-model-usercenter - Generate model code for usercenter"
	@echo "  make gen-api-usercenter - Generate api code for usercenter"
	@echo "  make gen-rpc-usercenter - Generate rpc code for usercenter"
	@echo "  make gen-model-upload - Generate model code for upload"
	@echo "  make gen-api-upload - Generate api code for upload"
	@echo "  make gen-rpc-upload - Generate rpc code for upload"

# Target: Start development environment
docker-up-env:
	docker-compose -f $(DOCKER_COMPOSE_ENV) up -d

# Target: Start application services
docker-up-app:
	docker-compose -f $(DOCKER_COMPOSE_APP) up -d

# Target: Stop development environment
docker-down-env:
	docker-compose -f $(DOCKER_COMPOSE_ENV) down

# Target: Stop application services
docker-down-app:
	docker-compose -f $(DOCKER_COMPOSE_APP) down

gen-model-usercenter:
	./deploy/scripts/mysql/genModel.sh usercenter user app/usercenter/model deploy/goctl/1.7.3 && \
	./deploy/scripts/mysql/genModel.sh usercenter user_auth app/usercenter/model deploy/goctl/1.7.3

gen-api-usercenter:
	goctl api go --api=app/usercenter/cmd/api/desc/main.api --dir=app/usercenter/cmd/api/ --style=go_zero --home=deploy/goctl/1.7.3/ && \
	goctl api plugin --plugin=goctl-swagger="swagger -filename usercenter.json" --api=app/usercenter/cmd/api/desc/main.api --dir=doc/swagger

gen-rpc-usercenter:
	goctl rpc protoc app/usercenter/cmd/rpc/pb/usercenter.proto --go_out=app/usercenter/cmd/rpc/ --go-grpc_out=app/usercenter/cmd/rpc/ --zrpc_out=app/usercenter/cmd/rpc/ --style=go_zero --home=deploy/goctl/1.7.3/

gen-model-upload:
	./deploy/scripts/mysql/genModel.sh upload upload_file app/upload/model deploy/goctl/1.7.3

gen-api-upload:
	goctl api go --api=app/upload/cmd/api/desc/upload.api --dir=app/upload/cmd/api/ --style=go_zero --home=deploy/goctl/1.7.3/ && \
	goctl api plugin --plugin=goctl-swagger="swagger -filename upload.json" --api=app/upload/cmd/api/desc/upload.api --dir=doc/swagger

gen-rpc-upload:
	goctl rpc protoc app/upload/cmd/rpc/pb/upload.proto --go_out=app/upload/cmd/rpc/ --go-grpc_out=app/upload/cmd/rpc/ --zrpc_out=app/upload/cmd/rpc/ --style=go_zero --home=deploy/goctl/1.7.3/