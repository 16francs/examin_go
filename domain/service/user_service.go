package service

import (
	"github.com/16francs/examin_go/domain/model"
	"github.com/16francs/examin_go/domain/repository"

	"github.com/16francs/examin_go/interface/middleware"
)

type UserService interface {
	Create(user model.User) error
}

type userService struct {
	repository.UserRepository
}

func NewUserService(u repository.UserRepository) UserService {
	return &userService{u}
}

func (u *userService) Create(user model.User) error {
	hash, err := middleware.GenerateHash(user.Password)
	if err != nil {
		return err
	}
	user.Password = hash
	return u.UserRepository.Create(user)
}