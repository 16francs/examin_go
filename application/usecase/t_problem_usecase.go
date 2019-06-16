package usecase

import (
	"github.com/16francs/examin_go/domain/model"
	"github.com/16francs/examin_go/domain/service"
)

// TProblemUsecase - 講師用の 問題集ユースケース
type TProblemUsecase interface {
	CreateProblem(title, content string, userID uint, tags []string) (*model.Problem, error)
}

type tProblemUsecase struct {
	service service.TProblemService
}

// NewTProblemUsecase - tProblemUsecase の生成
func NewTProblemUsecase(s service.TProblemService) TProblemUsecase {
	return &tProblemUsecase{s}
}

func (u *tProblemUsecase) CreateProblem(title, content string, userID uint, tags []string) (*model.Problem, error) {
	problem := &model.Problem{
		Title:   title,
		Content: content,
		UserID:  userID,
	}

	createdProblem, err := u.service.CreateProblem(problem)
	if err != nil {
		return nil, err
	}

	return createdProblem, nil
}
