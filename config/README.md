# config

ログの設定や環境変数などを記述する

## ログの設定

* ログのフォーマット

```go
msg := "Hello, World!"
log.Print("alert: %s", msg)
```

* ログの重要度
  - debug
  - info
  - trace
  - warn
  - error
  - alert

* ログの重要度の記述なし -> debug

## 環境変数の設定

* 環境変数を追加する場合、構造体に値を追加する

```go
// Env - 環境変数用の構造体
type Env struct {
	Port string `envconfig:"PORT" default:"8080"`
}
```
