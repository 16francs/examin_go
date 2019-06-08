# Rails とのディレクトリ対応

## 対応

### Golang - application

* usecase

```
app/controller - serviceとconceptsへの値渡し
``` 

### Glaong - config

```
config - 設定ファイル関連
```

### Golang - domain

* domain/model

```
app/models - モデルの定義
db/schemas - データベースへ登録する用のテーブル情報の定義
```

* domain/repository

```
app/models - ActiveRecord 関連がここにあたるかなと...笑
```

* domain/service

```
app/concepts/**/operation - モデル関連のロジック処理 -> バリデーションとかデータベース関連とか
app/services - データベースからの値取得とデータの整形
```

### Golang - infrastructure

* datastore

```
app/models - モデルの定義
config/database.yml - データベース接続用の設定
db/schemas - データベースへ登録する用のテーブル情報の定義
```

* router

```
app/controller - 認証周り
config/routes.rb - ルーティングの管理
```

* その他

```
app/channels - リアルタイムな双方向通信とかのロジック管理
app/jobs - バッチ処理とか
app/mailers - メール送信用の設定
```

### Golang - interface

* handler

```
app/controller - リクエスト(params) の取得とそれの受け渡し
app/errors - エラーハンドリング
app/views - レスポンスの整形
config/locales - 多言語対応とかエラーメッセージの管理とか
```

* middleware

```
app/controllers - 認証, 権限周りの処理
app/helpers - 共通のロジックとか
app/services - jwt とか
```

* request

```
app/concepts/**/contract - バリデーションの定義とか
app/controllers - リクエスト値の受け取り制限とか
```

---

## Rails のディレクトリ構成

```
.
├── app
│   ├── channels
│   │   └── application_cable
│   ├── concepts
│   │   ├── students
│   │   │   └── achievements
│   │   │       ├── contract
│   │   │       └── operation
│   │   ├── teachers
│   │   │   ├── problems
│   │   │   │   ├── contract
│   │   │   │   └── operation
│   │   │   ├── questions
│   │   │   │   ├── contract
│   │   │   │   └── operation
│   │   │   ├── students
│   │   │   │   ├── contract
│   │   │   │   └── operation
│   │   │   └── teachers
│   │   │       ├── contract
│   │   │       └── operation
│   │   └── users
│   │       ├── contract
│   │       └── operation
│   ├── controllers
│   │   ├── api
│   │   │   ├── students
│   │   │   └── teachers
│   │   └── concerns
│   ├── errors
│   ├── helpers
│   ├── jobs
│   ├── mailers
│   ├── models
│   │   └── concerns
│   ├── services
│   │   ├── students
│   │   └── teachers
│   └── views
│       └── layouts
├── bin
├── config
│   ├── environments
│   ├── initializers
│   └── locales
├── coverage
├── db
│   └── schemas
├── docker
│   ├── bin
│   │   └── sync
│   └── containers
│       └── db
├── lib
│   ├── tasks
│   └── test
├── log
├── public
├── spec
│   ├── factories
│   ├── models
│   ├── requests
│   │   └── api
│   │       ├── students
│   │       └── teachers
│   ├── routing
│   │   └── api
│   │       ├── students
│   │       └── teachers
│   └── test_helpers
├── storage
├── tmp
│   ├── cache
│   ├── pids
│   ├── sockets
│   └── storage
└── vendor

343 directories
```

## Golang のディレクトリ構成

```
.
├── application
│   └── usecase
├── config
├── docker
│   ├── bin
│   └── containers
│       ├── api
│       └── db
├── document
│   └── wiki
├── domain
│   ├── model
│   ├── repository
│   └── service
├── infrastructure
│   ├── datastore
│   │   └── migrate
│   └── router
├── interface
│   ├── handler
│   ├── middleware
│   └── request
├── registry
└── tmp

24 directories
```
