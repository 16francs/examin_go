package main

import (
	"log"
	"net/http"

	"github.com/16francs/examin_go/config"
	"github.com/16francs/examin_go/infrastructure/router"
)

func main() {
	// ログ出力設定
	config.Logger()

	// 環境変数
	env, err := config.LoadEnv()
	if err != nil {
		log.Fatalf("alert: %s", err)
	}

	// 起動コマンド
	router := router.Router()
	if err := http.ListenAndServe(":"+env.Port, router); err != nil {
		log.Fatalf("alert: %v", err)
	}
}
