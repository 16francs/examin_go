# domain/service

アプリケーション層から渡された値に対して処理を行う　　
いわゆる MVCS(Model View Controller Service) の S

```
// UserService - User の Service インターフェース
type UserService interface {
  FindUser(userID uint) (*model.User, error)
}

type userService struct {
  repository repository.UserRepository
}

// NewUserService - userService の生成
func NewUserService(r repository.UserRepository) UserService {
  return &userService{repository: r}
}

// ユーザ情報の取得
func (s *userService) FindUser(userID uint) (*model.User, error) {
  user, err := s.repository.FindUser(userID)
  if err != nil {
    return nil, err
  }

  return user, nil
}
```
