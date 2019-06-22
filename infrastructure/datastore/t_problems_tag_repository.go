package datastore

import (
	"github.com/16francs/examin_go/domain/model"
	"github.com/16francs/examin_go/domain/repository"
)

type tProblemsTagRepository struct {
}

// NewTProblemsTagRepository - tProblemsTagRepository の生成
func NewTProblemsTagRepository() repository.TProblemsTagRepository {
	return &tProblemsTagRepository{}
}

func (r *tProblemsTagRepository) CreateProblemsTag(problemsTag *model.ProblemsTag) (*model.ProblemsTag, error) {
	db := Connect()
	defer db.Close()

	err := db.Create(&problemsTag).Error
	return problemsTag, err
}
