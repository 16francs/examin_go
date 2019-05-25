package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthHandler - ヘルスチェック用の handler
type HealthHandler interface {
	GetHealth(c *gin.Context)
}

type healthHandler struct {
}

// NewHealthHandler - healthHandler の生成
func NewHealthHandler() HealthHandler {
	return &healthHandler{}
}

func (h *healthHandler) GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "UP"})
}
