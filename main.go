package main

import (
	"log"
	"net/http"
	"os"

	"github.com/16francs/examin_go/infrastructure/router"
)

func main() {
	// config, middleware のロード

	// ポートの設定 (default: 8080)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 起動コマンド
	router := router.Router()
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("alert: %v", err)
	}
}
