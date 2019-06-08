package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		LoginID, err := Parse(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":      http.StatusText(http.StatusUnauthorized),
				"description": "認証エラーです．",
			})
			c.Abort()
			return
		}

		c.Set("LoginID", LoginID)
		c.Next()
	}
}
