package main

import (
	"log"
	"net/http"

	"github.com/16francs/examin_go/config"
	"github.com/16francs/examin_go/infrastructure/router"
)

func main() {
	// 環境変数のロード
	env, err := config.LoadEnv()
	if err != nil {
		log.Fatalf("failed to load env: %s", err)
	}

	// 起動コマンド
	router := router.Router()
	if err := http.ListenAndServe(":"+env.Port, router); err != nil {
		log.Fatalf("failed to server start: %v", err)
	}
}
