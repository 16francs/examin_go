# aplication

ユースケースの記述をする
handler から受け取った値を service に送る
service からの結果を handler に返す

```
// UserUsecase - User 用のインターフェース
type UserUsecase interface {
  FindUser(userID uint) (*model.User, error)
}

type userUsecase struct {
  service service.UserService
}

// NewUserUsecase - userUsecaseの生成
func NewUserUsecase(s services.UserService) UserUsecase {
	return &userUsecase{service: s}
}

// FindUser - ユーザ情報の取得
func (u *userUsecase) FindUser(userID uint) (*model.User, error) {
  user, err := u.service.FindUser(userID)
  if err != nil {
    return nil, err
  }

  return user, nil
}
```
