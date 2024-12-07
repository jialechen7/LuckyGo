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
