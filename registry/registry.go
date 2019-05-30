package registry

import (
	"github.com/16francs/examin_go/application/usecase"
	"github.com/16francs/examin_go/domain/service"
	"github.com/16francs/examin_go/interface/handler"
)

// Registry - 簡易の DI コンテナ
// router の実装が依存する handler を定義
type Registry struct {
	handler.HealthHandler
	handler.SampleHandler
}

// NewRegistry - registry の生成
func NewRegistry() Registry {
	healthHandler := handler.NewHealthHandler()

	sampleService := service.NewSampleService()
	sampleUsecase := usecase.NewSampleUsecase(sampleService)
	sampleHandler := handler.NewSampleHandler(sampleUsecase)

	return Registry{
		healthHandler,
		sampleHandler,
	}
}
