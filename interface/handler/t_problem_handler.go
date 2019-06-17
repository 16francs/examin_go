package handler

import (
	"net/http"

	"github.com/16francs/examin_go/application/usecase"
	"github.com/16francs/examin_go/interface/request"
	"github.com/16francs/examin_go/interface/response"
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

	userID, ok := c.MustGet("UserID").(int)
	if !ok {
		Unauthorized(c)
		return
	}

	problem, tags, err := h.usecase.CreateProblem(request.Title, request.Content, uint(userID), request.Tags)
	if err != nil {
		ServerError(c)
		return
	}

	// レスポンスの整形
	tagContents := make([]string, 0, 4)
	for _, v := range tags {
		tagContents = append(tagContents, v.Content)
	}

	response := &response.TCreateProblem{
		ID:          problem.Base.ID,
		Title:       problem.Title,
		Content:     problem.Content,
		Tags:        tagContents,
		TeacherName: "taecher", // TODO: ログイン情報から取得
		CreatedAt:   problem.Base.CreatedAt,
		UpdatedAt:   problem.Base.UpdatedAt,
	}
	c.JSON(http.StatusCreated, response)
}
