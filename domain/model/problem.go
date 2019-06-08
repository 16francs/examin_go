package model

// Problem - 問題集 モデル
type Problem struct {
	Base
	Title   string `json:"title"`
	Conetnt string `json:"content"`
	UserID  uint   `json:"user_id"`
}
