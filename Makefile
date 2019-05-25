GOCMD   := go
GOBUILD := $(GOCMD) build
GOGET   := $(GOCMD) get
GOLINT  := golint
GOTEST  := $(GOCMD) test
COMPOSE := docker-compose

#　環境構築
setup:
	export GO111MODULE=on

# 依存ライブラリをビルド
install:
	$(GOBUILD)

# コンテナのビルド
build:
	$(COMPOSE) build

# サーバの起動
run:
	$(COMPOSE) up

# リントの実行
lint:
	$(GOLINT) ./...

# テストの実行
test:
	$(GOTEST) -v -cover ./...
