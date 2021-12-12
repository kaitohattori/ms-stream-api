APP_NAME = ms-stream-api
DOCKER_STORAGE_PATH = ~/ms-tv

build: ## Build on local
	swag init
	go build main.go

run: ## Run on local
	swag init
	go run main.go

docker-build: ## Build on docker
	docker build -t $(APP_NAME) .

docker-run: ## Run on docker
	docker run --rm \
		-p 8081:8081 \
		-v $(DOCKER_STORAGE_PATH)/assets:/go/src/$(APP_NAME)/assets \
		-v $(DOCKER_STORAGE_PATH)/logs:/go/src/$(APP_NAME)/logs \
		--name $(APP_NAME) \
		$(APP_NAME):latest

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
