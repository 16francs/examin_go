# domain/repository

インフラ層における repository のインターフェースを定義

```go
// UserRepository - User 関連のインターフェース
type UserRespository interface {
  FindUser(userID uint) (*model.User, error)
}
```
