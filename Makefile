# Variables
DOCKER_COMPOSE_APP=docker-compose.yml
DOCKER_COMPOSE_ENV=docker-compose-env.yml

# Default target
.DEFAULT_GOAL := help

# Help information
help:
	@echo "Usage:"
	@echo "  make create-dirs NAME=<micro_service_name> - Create directory structure for specified microservice name"
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
	@echo "  make gen-model-lottery - Generate model code for lottery"
	@echo "  make gen-api-lottery - Generate api code for lottery"
	@echo "  make gen-rpc-lottery - Generate rpc code for lottery"
	@echo "  make gen-api-notice - Generate api code for notice"
	@echo "  make gen-model-comment - Generate model code for comment"
	@echo "  make gen-api-comment - Generate api code for comment"
	@echo "  make gen-rpc-comment - Generate rpc code for comment"
	@echo "  make gen-model-checkin - Generate model code for checkin"
	@echo "  make gen-api-checkin - Generate api code for checkin"
	@echo "  make gen-rpc-checkin - Generate rpc code for checkin"

# Target: Create directory structure for specified project name
create-dirs:
	@echo "Creating directory structure for $(NAME)..."
	mkdir -p app/$(NAME)/cmd/api/desc
	mkdir -p app/$(NAME)/cmd/rpc/pb
	mkdir -p app/$(NAME)/model
	@echo "Directory structure created for $(NAME)."

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

# Generate model code for usercenter
gen-model-usercenter:
	./deploy/scripts/mysql/genModel.sh usercenter user app/usercenter/model deploy/goctl/1.7.3 && \
	./deploy/scripts/mysql/genModel.sh usercenter user_auth app/usercenter/model deploy/goctl/1.7.3 && \
	./deploy/scripts/mysql/genModel.sh usercenter user_sponsor app/usercenter/model deploy/goctl/1.7.3

# Generate API code for usercenter
gen-api-usercenter:
	goctl api go --api=app/usercenter/cmd/api/desc/main.api --dir=app/usercenter/cmd/api/ --style=go_zero --home=deploy/goctl/1.7.3/ && \
	goctl api plugin --plugin=goctl-swagger="swagger -filename usercenter.json" --api=app/usercenter/cmd/api/desc/main.api --dir=doc/swagger

# Generate RPC code for usercenter
gen-rpc-usercenter:
	goctl rpc protoc app/usercenter/cmd/rpc/pb/usercenter.proto --go_out=app/usercenter/cmd/rpc/ --go-grpc_out=app/usercenter/cmd/rpc/ --zrpc_out=app/usercenter/cmd/rpc/ --style=go_zero --home=deploy/goctl/1.7.3/

# Generate model code for upload
gen-model-upload:
	./deploy/scripts/mysql/genModel.sh upload upload_file app/upload/model deploy/goctl/1.7.3

# Generate API code for upload
gen-api-upload:
	goctl api go --api=app/upload/cmd/api/desc/upload.api --dir=app/upload/cmd/api/ --style=go_zero --home=deploy/goctl/1.7.3/ && \
	goctl api plugin --plugin=goctl-swagger="swagger -filename upload.json" --api=app/upload/cmd/api/desc/upload.api --dir=doc/swagger

# Generate RPC code for upload
gen-rpc-upload:
	goctl rpc protoc app/upload/cmd/rpc/pb/upload.proto --go_out=app/upload/cmd/rpc/ --go-grpc_out=app/upload/cmd/rpc/ --zrpc_out=app/upload/cmd/rpc/ --style=go_zero --home=deploy/goctl/1.7.3/

# Generate model code for lottery
gen-model-lottery:
	./deploy/scripts/mysql/genModel.sh lottery lottery app/lottery/model deploy/goctl/1.7.3 && \
	./deploy/scripts/mysql/genModel.sh lottery lottery_participation app/lottery/model deploy/goctl/1.7.3 && \
	./deploy/scripts/mysql/genModel.sh lottery prize app/lottery/model deploy/goctl/1.7.3

# Generate API code for lottery
gen-api-lottery:
	goctl api go --api=app/lottery/cmd/api/desc/main.api --dir=app/lottery/cmd/api/ --style=go_zero --home=deploy/goctl/1.7.3/ && \
	goctl api plugin --plugin=goctl-swagger="swagger -filename lottery.json" --api=app/lottery/cmd/api/desc/main.api --dir=doc/swagger

# Generate RPC code for lottery
gen-rpc-lottery:
	goctl rpc protoc app/lottery/cmd/rpc/pb/lottery.proto --go_out=app/lottery/cmd/rpc/ --go-grpc_out=app/lottery/cmd/rpc/ --zrpc_out=app/lottery/cmd/rpc/ --style=go_zero --home=deploy/goctl/1.7.3/

# Generate model code for notice
gen-api-notice:
	goctl api go --api=app/notice/cmd/api/desc/main.api --dir=app/notice/cmd/api/ --style=go_zero --home=deploy/goctl/1.7.3/ && \
	goctl api plugin --plugin=goctl-swagger="swagger -filename notice.json" --api=app/notice/cmd/api/desc/main.api --dir=doc/swagger

# Generate model code for comment
gen-model-comment:
	./deploy/scripts/mysql/genModel.sh comment comment app/comment/model deploy/goctl/1.7.3 && \
	./deploy/scripts/mysql/genModel.sh comment praise app/comment/model deploy/goctl/1.7.3

# Generate API code for comment
gen-api-comment:
	goctl api go --api=app/comment/cmd/api/desc/main.api --dir=app/comment/cmd/api/ --style=go_zero --home=deploy/goctl/1.7.3/ && \
	goctl api plugin --plugin=goctl-swagger="swagger -filename comment.json" --api=app/comment/cmd/api/desc/main.api --dir=doc/swagger

# Generate RPC code for comment
gen-rpc-comment:
	goctl rpc protoc app/comment/cmd/rpc/pb/comment.proto --go_out=app/comment/cmd/rpc/ --go-grpc_out=app/comment/cmd/rpc/ --zrpc_out=app/comment/cmd/rpc/ --style=go_zero --home=deploy/goctl/1.7.3/

gen-model-checkin:
	./deploy/scripts/mysql/genModel.sh checkin checkin_record app/checkin/model deploy/goctl/1.7.3 && \
	./deploy/scripts/mysql/genModel.sh checkin integral app/checkin/model deploy/goctl/1.7.3 && \
	./deploy/scripts/mysql/genModel.sh checkin integral_record app/checkin/model deploy/goctl/1.7.3 && \
	./deploy/scripts/mysql/genModel.sh checkin tasks app/checkin/model deploy/goctl/1.7.3 && \
	./deploy/scripts/mysql/genModel.sh checkin task_progress app/checkin/model deploy/goctl/1.7.3 && \
	./deploy/scripts/mysql/genModel.sh checkin task_record app/checkin/model deploy/goctl/1.7.3


gen-api-checkin:
	goctl api go --api=app/checkin/cmd/api/desc/main.api --dir=app/checkin/cmd/api/ --style=go_zero --home=deploy/goctl/1.7.3/ && \
	goctl api plugin --plugin=goctl-swagger="swagger -filename checkin.json" --api=app/checkin/cmd/api/desc/main.api --dir=doc/swagger

gen-rpc-checkin:
	goctl rpc protoc app/checkin/cmd/rpc/pb/checkin.proto --go_out=app/checkin/cmd/rpc/ --go-grpc_out=app/checkin/cmd/rpc/ --zrpc_out=app/checkin/cmd/rpc/ --style=go_zero --home=deploy/goctl/1.7.3/

# Default target
.PHONY: help create-dirs docker-up-env docker-up-app docker-down-env docker-down-app gen-model-usercenter gen-api-usercenter gen-rpc-usercenter gen-model-upload gen-api-upload gen-rpc-upload gen-model-lottery gen-api-lottery gen-rpc-lottery gen-api-notice gen-model-comment gen-api-comment gen-rpc-comment gen-model-checkin gen-api-checkin gen-rpc-checkin
