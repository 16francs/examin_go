package registry

import (
	"github.com/16francs/examin_go/application/usecase"
	"github.com/16francs/examin_go/domain/service"
	"github.com/16francs/examin_go/infrastructure/datastore"
	"github.com/16francs/examin_go/interface/handler"
)

// Registry - 簡易の DI コンテナ
// router の実装が依存する handler を定義
type Registry struct {
	handler.HealthHandler
	handler.SampleHandler
	handler.TTeacherHandler
}

// NewRegistry - registry の生成
func NewRegistry() Registry {
	healthHandler := handler.NewHealthHandler()

	sampleService := service.NewSampleService()
	sampleUsecase := usecase.NewSampleUsecase(sampleService)
	sampleHandler := handler.NewSampleHandler(sampleUsecase)

	tTeacherRepository := datastore.NewTTeacherRepository()
	tTeacherService := service.NewTTeacherService(tTeacherRepository)
	tTeacherUsecase := usecase.NewTTeacherUsecase(tTeacherService)
	tTeacherHandler := handler.NewTTeacherHandler(tTeacherUsecase)

	return Registry{
		healthHandler,
		sampleHandler,
		tTeacherHandler,
	}
}
