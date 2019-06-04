package service

import (
	"github.com/16francs/examin_go/domain/model"
	"github.com/16francs/examin_go/domain/repository"
)

// TTeacherService - 講師向け TeacherService
type TTeacherService interface {
	CreateTeacher(loginID, name, school string) (*model.User, error)
}

type tTeacherService struct {
	repository repository.TTeacherRepository
}

// NewTTeacherService - TTeacherService の生成
func NewTTeacherService(r repository.TTeacherRepository) TTeacherService {
	return &tTeacherService{repository: r}
}

func (s *tTeacherService) CreateTeacher(loginID, name, school string) (*model.User, error) {
	// TODO: パスワードのハッシュ化処理

	teacher := &model.User{
		LoginID:  loginID,
		Name:     name,
		School:   school,
		Password: loginID,
		Role:			0,
	}

	createdTeacher, err := s.repository.CreateTeacher(teacher)
	if err != nil {
		return nil, err
	}

	return createdTeacher, nil
}
