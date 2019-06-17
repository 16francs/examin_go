package datastore

import (
	"github.com/16francs/examin_go/domain/model"
	"github.com/16francs/examin_go/domain/repository"
)

type tTagRepository struct {
}

// NewTTagRepository - tTagRepository の生成
func NewTTagRepository() repository.TTagRepository {
	return &tTagRepository{}
}

func (r *tTagRepository) CreateTag(tag *model.Tag) (*modek.Tag, error) {
	db := Connect()
	defer db.Close()

	err := db.Create(&tag).Error
	return tag, err
}
