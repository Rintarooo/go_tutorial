yml := .devcontainer/docker-compose.yml
service := go-tutorial
port := 3000

build:
	docker-compose -f $(yml) build $(service)

up:
	docker-compose -f $(yml) up -d

# docker ps
curl:
	curl http://localhost:$(port)

logs:
	docker-compose -f $(yml) logs

down:
	docker-compose -f $(yml) down
