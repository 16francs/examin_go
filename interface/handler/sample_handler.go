package handler

import (
	"net/http"

	"github.com/16francs/examin_go/application/usecase"
	"github.com/gin-gonic/gin"
)

// SampleHandler - handler のサンプル
type SampleHandler interface {
	GetSample(c *gin.Context)
}

type sampleHandler struct {
	usecase usecase.SampleUsecase
}

// NewSampleHandler - sampleHandlerの生成
func NewSampleHandler(u usecase.SampleUsecase) SampleHandler {
	return &sampleHandler{usecase: u}
}

func (h *sampleHandler) GetSample(c *gin.Context) {
	response, err := h.usecase.GetSample()
	if err != nil {
		c.JSON(500, nil)
	}
	c.JSON(http.StatusOK, response)
}
