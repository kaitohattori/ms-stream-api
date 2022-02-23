# MS Stream API

## 準備

```
# GO111MODULEをオンにする
$ export GO111MODULE=on

# App初期設定
$ make init

# mod.modとgo.sumの差でエラーが出たら以下のコマンドを実行
$ go mod tidy
```

`config/config.ini` を配置してください。

## 動作

```
# docsを更新
$ swag init

# 実行
$ make run

# Dockerビルド
$ make docker-build

# Docker実行
$ make docker-run

# external appsを立ち上げる
$ make external-run

# external appsを終了する
$ make external-end

# ヘルプ
$ make help
```

## ドキュメント

http://localhost:8081/docs/api/v1/index.html#/

## Deploy

Deploy to minikube

```
# Start minikube
$ minikube start

# Use local image
$ eval $(minikube docker-env)

# Build docker image
$ docker build -t ms-stream-api .

# Deploy
$ kubectl apply -f deploy/deployment.yaml
$ kubectl apply -f deploy/service.yaml

# Get all status
$ kubectl get all

# Access to deployed app (Click the displayed url)
$ minikube service ms-stream-api --url
```

minikube common commands

```
$ minikube start
$ minikube status
$ minikube dashboard
$ minikube tunnel
$ minikube stop
```
