package service

import (
	"github.com/16francs/examin_go/domain/model"
	"github.com/16francs/examin_go/domain/repository"
	"github.com/16francs/examin_go/interface/middleware"
)

// TTeacherService - 講師向け TeacherService
type TTeacherService interface {
	CreateTeacher(teacher *model.User) (*model.User, error)
}

type tTeacherService struct {
	repository repository.TTeacherRepository
}

// NewTTeacherService - TTeacherService の生成
func NewTTeacherService(r repository.TTeacherRepository) TTeacherService {
	return &tTeacherService{repository: r}
}

func (s *tTeacherService) CreateTeacher(teacher *model.User) (*model.User, error) {
	// パスワードのハッシュ化処理
	hashPassword, err := middleware.GenerateHash(teacher.Password)
	if err != nil {
		return nil, err
	}
	teacher.Password = hashPassword

	createdTeacher, err := s.repository.CreateTeacher(teacher)
	if err != nil {
		return nil, err
	}

	return createdTeacher, nil
}
