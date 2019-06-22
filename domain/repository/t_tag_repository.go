package repository

import "github.com/16francs/examin_go/domain/model"

// TTagRepository - 講師向け タグ モデルのデータベース操作
type TTagRepository interface {
	CreateTag(tag *model.Tag) (*model.Tag, error)
}
