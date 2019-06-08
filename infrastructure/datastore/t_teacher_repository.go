package datastore

import (
	"github.com/16francs/examin_go/domain/model"
	"github.com/16francs/examin_go/domain/repository"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysqlのドライバー
)

type tTeacherRepository struct {
}

// NewTTeacherRepository - tTeacherRepository の生成
func NewTTeacherRepository() repository.TTeacherRepository {
	return &tTeacherRepository{}
}

func (r *tTeacherRepository) CreateTeacher(teacher *model.User) (*model.User, error) {
	db := Connect()
	defer db.Close()

	err := db.Create(&teacher).Error
	return teacher, err
}
