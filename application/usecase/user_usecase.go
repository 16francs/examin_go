package usecase

import (
	"github.com/16francs/examin_go/domain/model"
	"github.com/16francs/examin_go/domain/service"
)

type UserUsecase interface {
	Create(user model.User) error
}

type userUsecase struct {
	service service.UserService
}

func NewUserUsecase(s service.UserService) UserUsecase {
	return &userUsecase{service: s}
}

func (u *userUsecase) Create(user model.User) error {
	err := u.service.Create(user)
	if err != nil {
		return err
	}
	return nil
}