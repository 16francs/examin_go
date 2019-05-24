package model

import (
	"time"
)

// Base - モデル共通のカラム
type Base struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
