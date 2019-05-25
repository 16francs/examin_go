package registry

import "github.com/16francs/examin_go/interface/handler"

// Registry - 簡易の DI コンテナ
// router の実装が依存する handler を定義
type Registry struct {
	handler.HealthHandler
}

// NewRegistry - registry の生成
func NewRegistry() Registry {
	healthHandler := handler.NewHealthHandler()

	return Registry{
		healthHandler,
	}
}
