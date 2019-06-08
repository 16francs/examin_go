package repository

import (
	"github.com/16francs/examin_go/domain/model"
)

type UserRepository interface {
	Create(user *model.User) error
	FindByLoginID(LoginID string) (model.User, error)
}
