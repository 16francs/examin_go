package usecase

import (
	"github.com/16francs/examin_go/domain/service"
	"github.com/16francs/examin_go/interface/request"
)

type AuthUsecase interface {
	SignIn(request request.Auth) (string, error)
}

type authUsecase struct {
	service service.AuthService
}

func NewAuthUsecase(s service.AuthService) AuthUsecase {
	return &authUsecase{service: s}
}

func (a *authUsecase) SignIn(request request.Auth) (string, error) {
	token, err := a.service.SignIn(request.LoginID, request.Password)
	if err != nil {
		return "", err
	}
	return token, err
}
