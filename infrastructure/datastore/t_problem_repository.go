package datastore

import (
	"github.com/16francs/examin_go/domain/model"
	"github.com/16francs/examin_go/domain/repository"
)

type tProblemRepository struct {
}

// NewTProblemRepository - tProblemRepository の生成
func NewTProblemRepository() repository.TProblemRepository {
	return &tProblemRepository{}
}

func (r *tProblemRepository) CreateProblem(problem *model.Problem) (*model.Problem, error) {
	db := Connect()
	defer db.Close()

	err := db.Create(&problem).Error
	return problem, err
}
