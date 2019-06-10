package handler

import (
	"net/http"

	"github.com/16francs/examin_go/application/usecase"
	"github.com/16francs/examin_go/interface/request"
	"github.com/gin-gonic/gin"
)

// TProblemHandler - 講師用の問題集ハンドラ
type TProblemHandler interface {
	CreateProblem(c *gin.Context)
}

type tProblemHandler struct {
	usecase usecase.TProblemUsecase
}

// NewTProblemHandler - tProblemHandler の生成
func NewTProblemHandler(u usecase.TProblemUsecase) TProblemHandler {
	return &tProblemHandler{u}
}

func (h *tProblemHandler) CreateProblem(c *gin.Context) {
	var request request.TCreateProblem
	if err := c.BindJSON(&request); err != nil {
		BadRequest(c)
		return
	}

	loginID := c.MustGet("LoginID")
	strLoginID, ok := loginID.(string)
	if !ok {
		Unauthorized(c)
		return
	}

	response, err := h.usecase.CreateProblem(request.Title, request.Content, strLoginID)
	if err != nil {
		ServerError(c)
		return
	}

	c.JSON(http.StatusCreated, response)
}
