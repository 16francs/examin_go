# datastore

DBへアクセスする処理を記述
何かライブラリを用いる予定 -> たぶん gorm

```
type userRepository struct {
}

// NewUserRepository - userRepository の生成
func NewUserRepository() repository.UserRepository {
  return &userRepository{}
}

// FindUser - ユーザ情報の取得
func (r *userRepository) FindUser(userID uint) (*model.User, error) {
  user := new(model.User)
  if err := // DBへアクセスする処理; err != nil {
    return nil, err
  }

  return user, nil
}
```