package datastore

import (
	"github.com/16francs/examin_go/domain/model"
	"github.com/16francs/examin_go/domain/repository"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysqlのドライバー
)

type userRepository struct {
}

func NewUserRepository() repository.UserRepository {
	return &userRepository{}
}

func (u *userRepository) Create(user model.User) error {
	db := Connect()
	defer db.Close()

	err := db.Create(&user).Error
	return err
}
