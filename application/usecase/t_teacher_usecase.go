package usecase

import (
	"github.com/16francs/examin_go/domain/model"
	"github.com/16francs/examin_go/domain/service"
)

// TTeacherUsecase - 講師向け TeacherUsecase
type TTeacherUsecase interface {
	CreateTeacher(loginID, name, school string) (*model.User, error)
}

type tTeacherUsecase struct {
	service service.TTeacherService
}

// NewTTeacherUsecase - tTeacherUsecase の生成
func NewTTeacherUsecase(s service.TTeacherService) TTeacherUsecase {
	return &tTeacherUsecase{service: s}
}

func (u *tTeacherUsecase) CreateTeacher(loginID, name, school string) (*model.User, error) {
	teacher, err := u.service.CreateTeacher(loginID, name, school)
	if err != nil {
		return nil, err
	}

	return teacher, nil
}
