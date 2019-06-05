package usecase

import (
	// "github.com/16francs/examin_go/domain/model"
	"github.com/16francs/examin_go/domain/service"
	"github.com/16francs/examin_go/interface/request"
)

type UserUsecase interface {
	Create(request request.CreateUser) error
}

type userUsecase struct {
	service service.UserService
}

func NewUserUsecase(s service.UserService) UserUsecase {
	return &userUsecase{service: s}
}

func (u *userUsecase) Create(request request.CreateUser) error {
	user := request.ToUser()
	err := u.service.Create(user)
	if err != nil {
		return err
	}
	return nil
}
