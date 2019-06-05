package request

import (
	"github.com/16francs/examin_go/domain/model"
	"time"
)

type CreateUser struct {
	LoginId  string `json:"login_id" binding:"required"`
	Password string `json:"password"  binding:"required"`
	Name     string `json:"name"  binding:"required"`
	School   string `json:"school"`
}

func (c *CreateUser) ToUser() *model.User {
	now := time.Now()
	return &model.User{
		LoginId:  c.LoginId,
		Password: c.Password,
		Name:     c.Name,
		School:   c.School,
		Role:     0,
		Base: model.Base{
			CreatedAt: now,
			UpdatedAt: now,
		},
	}
}
