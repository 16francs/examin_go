package repository

import "github.com/16francs/examin_go/domain/model"

// TProblemRepository - 講師用の 問題集 モデル関連のデータベース操作
type TProblemRepository interface {
	CreateProblem(problem *model.Problem) (*model.Problem, error)
}
