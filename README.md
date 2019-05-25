[![CircleCI](https://circleci.com/gh/16francs/examin_go.svg?style=shield)](https://circleci.com/gh/16francs/examin_go)
[![Coverage Status](https://coveralls.io/repos/github/16francs/examin_go/badge.svg?branch=master)](https://coveralls.io/github/16francs/examin_go?branch=master)

# Examin (バックエンド)

[仕様書: API Blueprint](https://github.com/16francs/examin_blueprint)     
[フロントエンド: Vue.js](https://github.com/16francs/examin_vue)     
[バックエンド: Golang API](https://github.com/16francs/examin_go)
[インフラ周り](https://github.com/16francs/examin)

## 開発環境

## 環境構築

* プロジェクトをクローンする

> $ mkdir -p $(GOPATH)/src/github.com/16francs  
> $ cd $(GOPATH)/src/github.com/16francs  
> $ git clone https://github.com/16francs/examin_go

* 以下のコマンドを実行

> $ make build

## 起動方法

* MySQLの起動

> $ sudo mysql.server start

* サーバを起動

> $ make run
