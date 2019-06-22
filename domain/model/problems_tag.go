package model

// ProblemsTag ー 問題集とタグの関連付け用 中間モデル
type ProblemsTag struct {
	Base
	ProblemID uint `json:"problem_id"`
	TagID     uint `json:"tag_id"`
}
