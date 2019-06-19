package repository

import (
	"github.com/16francs/examin_go/domain/model"
)

// TProblemsTagRepository - 講師向け 問題集タグ 中間モデルのデータベース操作
type TProblemsTagRepository interface {
	CreateProblemsTag(problemsTag *model.ProblemsTag) (*model.ProblemsTag, error)
}
