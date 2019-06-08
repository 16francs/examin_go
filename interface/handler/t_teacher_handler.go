package handler

import (
	"net/http"

	"github.com/16francs/examin_go/application/usecase"
	"github.com/16francs/examin_go/interface/request"
	"github.com/gin-gonic/gin"
)

// TTeacherHandler - 講師用の講師ハンドラ
type TTeacherHandler interface {
	CreateTeacher(c *gin.Context)
}

type tTeacherHandler struct {
	usecase usecase.TTeacherUsecase
}

// NewTTeacherHandler - tTeacherHandler の生成
func NewTTeacherHandler(u usecase.TTeacherUsecase) TTeacherHandler {
	return &tTeacherHandler{usecase: u}
}

func (h *tTeacherHandler) CreateTeacher(c *gin.Context) {
	var request request.TCreateTeacher
	if err := c.BindJSON(&request); err != nil {
		BadRequest(c)
		return
	}

	response, err := h.usecase.CreateTeacher(request.LoginID, request.Name, request.School)
	if err != nil {
		ServerError(c)
		return
	}

	c.JSON(http.StatusCreated, response)
}
