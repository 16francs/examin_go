package handler

import (
	"net/http"

	"github.com/16francs/examin_go/application/usecase"
	"github.com/16francs/examin_go/interface/request"
	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	SignIn(c *gin.Context)
	AuthCheck(c *gin.Context)
}

type authHandler struct {
	usecase usecase.AuthUsecase
}

func NewAuthHandler(u usecase.AuthUsecase) AuthHandler {
	return &authHandler{usecase: u}
}

func (h *authHandler) SignIn(c *gin.Context) {
	var request request.Auth
	if err := c.BindJSON(&request); err != nil {
		BadRequest(c)
		return
	}
	token, err := h.usecase.SignIn(request)
	if err != nil {
		Unauthorized(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":       "success",
		"Authorization": token,
	})
}

func (h *authHandler) AuthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "認証確認済み",
	})
}
