package router

import (
	"github.com/16francs/examin_go/interface/handler"
	"github.com/16francs/examin_go/interface/middleware"
	"github.com/16francs/examin_go/registry"
	"github.com/gin-gonic/gin"
)

// Router - ルーティングの定義
func Router() *gin.Engine {
	registry := registry.NewRegistry()

	// ルーティング
	router := gin.Default()
	router.Use(middleware.Logger())

	router.GET("/health", registry.HealthHandler.GetHealth)

	router.GET("/sample", registry.SampleHandler.GetSample)
	router.POST("/sample", registry.SampleHandler.PostSample)

	router.POST("/api/users", registry.UserHandler.CreateUser)

	router.POST("/api/auth", registry.AuthHandler.SignIn)

	router.NoRoute(handler.NotFound)

	return router
}
