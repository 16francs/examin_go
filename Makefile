GOCMD   := go
GOBUILD := $(GOCMD) build
GOGET   := $(GOCMD) get
GORUN   := $(GOCMD) run
GOTEST  := $(GOCMD) test

#　環境構築
setup:
	export GO111MODULE=on
	$(GOGET) -u golang.org/x/lint/golint

build:
	$(GOBUILD)

# サーバの起動
run:
	$(GORUN) ./main.go

# リントの実行
lint:
	golint ./...

# テストの実行
test:
	$(GOTEST) -v -cover ./...
