package service

import (
	"github.com/16francs/examin_go/domain/model"
	"github.com/16francs/examin_go/domain/repository"
)

// TTagService - 講師向け タグ モデルの操作
type TTagService interface {
	CreateTag(tag *model.Tag) (*model.Tag, error)
}

type tTagService struct {
	repository repository.TTagRepository
}

// NewTTagService - tTagService の生成
func NewTTagService(r repository.TTagRepository) TTagService {
	return &tTagService{r}
}

func (s *tTagService) CreateTag(tag *model.Tag) (*model.Tag, error) {
	createdTag, err := s.repository.CreateTag(tag)
	if err != nil {
		return nil, err
	}

	return createdTag, nil
}
