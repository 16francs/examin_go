# interface/request

Client からのリクエストのバリデーションを定義
domain に置くか、ここに置くかは要検討...

## テンプレート

```go
package request

// PostSample - PostSample のリクエスト
type PostSample struct {
	Name string `json:"name" binding:"required"`
}
```
