package main

import (
	"log"
	"net/http"
)

func main() {
	// config, middleware のロード

	// ポートの設定 ( default: 8080 )

	// 起動コマンド
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
