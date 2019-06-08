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
	teacher := &model.User{
		LoginID:  loginID,
		Name:     name,
		School:   school,
		Password: loginID,
		Role:     0,
	}

	createdTeacher, err := u.service.CreateTeacher(teacher)
	if err != nil {
		return nil, err
	}

	return createdTeacher, nil
}
