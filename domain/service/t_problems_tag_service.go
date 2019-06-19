package service

import (
	"github.com/16francs/examin_go/domain/model"
	"github.com/16francs/examin_go/domain/repository"
)

// TProblemsTagService - 講師向け 問題集タグ 中間モデルの操作
type TProblemsTagService interface {
	CreateProblemsTag(problemsTag *model.ProblemsTag) (*model.ProblemsTag, error)
}

type tProblemsTagService struct {
	repository repository.TProblemsTagRepository
}

// NewTProblemsTagService - tProblemsTagService の生成
func NewTProblemsTagService(r repository.TProblemsTagRepository) TProblemsTagService {
	return &tProblemsTagService{r}
}

func (s *tProblemsTagService) CreateProblemsTag(problemsTag *model.ProblemsTag) (*model.ProblemsTag, error) {
	createdProblemsTag, err := s.repository.CreateProblemsTag(problemsTag)
	if err != nil {
		return nil, err
	}

	return createdProblemsTag, nil
}
