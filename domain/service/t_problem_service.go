package service

import (
	"github.com/16francs/examin_go/domain/repository"
)

// TProblemService - 講師向け 問題集 モデルの操作
type TProblemService interface {
}

type tProblemService struct {
	repository repository.TProblemRepository
}

// NewTProblemService - tProblemService の生成
func NewTProblemService(r *repository.TProblemRepository) TProblemService {
	return &tProblemService{r}
}
