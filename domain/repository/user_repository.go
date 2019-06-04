package repository

import (
	"github.com/16francs/examin_go/domain/model"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	Create(user model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}


func (u *userRepository) Create(user model.User) error {
	err := u.db.Create(&user).Error
	return err
}