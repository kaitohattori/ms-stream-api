APP_NAME = ms-stream-api
EXTERNAL_APPS = postgresql ms-api

init: ## Initialize app
	go mod download
	go mod tidy
	go get -u github.com/swaggo/swag/cmd/swag
	mkdir -p ./logs/

run: ## Run on local
	swag init
	go run main.go

docker-build: ## Build on docker
	docker build -t $(APP_NAME) .

docker-run: ## Run on docker
	docker run --rm \
		--add-host=localhost:host-gateway \
		-p 8081:8081 \
		-v assets:/go/src/$(APP_NAME)/assets \
		-v logs:/go/src/$(APP_NAME)/logs \
		--name $(APP_NAME) \
		$(APP_NAME):latest

external-init: ## Initialize external apps
	rm -rf ./external-apps/logs

external-up: ## Up external apps
	docker-compose up -d $(EXTERNAL_APPS)

external-down: ## Down external apps
	docker-compose down

service-in: ## Service in app
	kubectl apply -f deploy/deployment.yaml
	kubectl apply -f deploy/service.yaml

service-out: ## Service out app
	kubectl delete -f deploy/deployment.yaml
	kubectl delete -f deploy/service.yaml

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
