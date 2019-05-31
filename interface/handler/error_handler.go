package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// BadRequest - 400
func BadRequest(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status":      http.StatusText(http.StatusBadRequest),
		"description": "不正なパラメータが入力されています．",
	})
	c.Abort()
}

// Unauthorized - 401
func Unauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"status":      http.StatusText(http.StatusUnauthorized),
		"description": "認証エラーです．",
	})
	c.Abort()
}

// Forbidden - 403
func Forbidden(c *gin.Context) {
	c.JSON(http.StatusForbidden, gin.H{
		"status":      http.StatusText(http.StatusForbidden),
		"description": "認証エラーです．",
	})
	c.Abort()
}

// NotFound - 404
func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":      http.StatusText(http.StatusNotFound),
		"description": "ページが存在しません．",
	})
	c.Abort()
}

// ServerError - 500
func ServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"status":      http.StatusText(http.StatusInternalServerError),
		"description": "サーバーエラーです．",
	})
	c.Abort()
}
