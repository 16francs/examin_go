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
	handler.UserHandler
	handler.AuthHandler
	handler.TProblemHandler
	handler.TTeacherHandler
}

// NewRegistry - registry の生成
func NewRegistry() Registry {
	healthHandler := handler.NewHealthHandler()

	sampleService := service.NewSampleService()
	sampleUsecase := usecase.NewSampleUsecase(sampleService)
	sampleHandler := handler.NewSampleHandler(sampleUsecase)

	userRepository := datastore.NewUserRepository()
	userService := service.NewUserService(userRepository)
	userUsecase := usecase.NewUserUsecase(userService)
	userHandler := handler.NewUserHandler(userUsecase)

	authService := service.NewAuthService(userRepository)
	authUsecase := usecase.NewAuthUsecase(authService)
	authHandler := handler.NewAuthHandler(authUsecase)

	tTagRepository := datastore.NewTTagRepository()
	tTagService := service.NewTTagService(tTagRepository)

	tProblemRepository := datastore.NewTProblemRepository()
	tProblemService := service.NewTProblemService(tProblemRepository)
	tProblemUsecase := usecase.NewTProblemUsecase(tProblemService, tTagService)
	tProblemHandler := handler.NewTProblemHandler(tProblemUsecase)

	tTeacherRepository := datastore.NewTTeacherRepository()
	tTeacherService := service.NewTTeacherService(tTeacherRepository)
	tTeacherUsecase := usecase.NewTTeacherUsecase(tTeacherService)
	tTeacherHandler := handler.NewTTeacherHandler(tTeacherUsecase)

	return Registry{
		healthHandler,
		sampleHandler,
		userHandler,
		authHandler,
		tProblemHandler,
		tTeacherHandler,
	}
}
