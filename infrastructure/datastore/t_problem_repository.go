package datastore

import "github.com/16francs/examin_go/domain/repository"

type tProblemRepository struct {
}

// NewTProblemRepository - tProblemRepository の生成
func NewTProblemRepository() repository.TProblemRepository {
	return &tProblemRepository{}
}
