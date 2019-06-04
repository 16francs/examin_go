package handler

import (
	"net/http"

	"github.com/16francs/examin_go/application/usecase"
	"github.com/16francs/examin_go/interface/request"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	CreateUser(c *gin.Context)
}

type userHandler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(u usecase.UserUsecase) UserHandler {
	return &userHandler{usecase: u}
}

func (h *userHandler) CreateUser(c *gin.Context) {
	var request request.CreateUser
	if err := c.BindJSON(&request); err != nil {
		BadRequest(c)
		return
	}

	c.JSON(http.StatusOK, request)
}
