package repository

import "github.com/16francs/examin_go/domain/model"

// TTeacherRepository - 講師向け TeacherRepository
type TTeacherRepository interface {
	CreateTeacher(teacher *model.User) (*model.User, error)
}
