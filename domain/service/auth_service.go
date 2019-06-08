package service

import (
	"github.com/16francs/examin_go/domain/repository"
	"github.com/16francs/examin_go/interface/middleware"
)

type AuthService interface {
	SignIn(loginID string, password string) (string, error)
}

type authService struct {
	repository repository.UserRepository
}

func NewAuthService(u repository.UserRepository) AuthService {
	return &authService{repository: u}
}

func (a authService) SignIn(loginID string, password string) (string, error) {
	user, err := a.repository.FindByLoginID(loginID)
	if err != nil {
		return "", err
	}
	err = middleware.Compare(user.Password, password)
	if err != nil {
		return "", err
	}
	return middleware.GenerateToken(user)
}
