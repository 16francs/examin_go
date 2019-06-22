package model

// Tag - タグ モデル
type Tag struct {
	Base
	Content string `json:"content" gorm:"unique;not null"`
}
