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

	router.NoRoute(handler.NotFound)

	return router
}
