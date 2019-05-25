package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 起動確認用 -> helth check 用ちゃんと作る!!
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello you've requested: %s\n", r.URL.Path)
	})

	// config, middleware のロード

	// ポートの設定 ( default: 8080 )

	// 起動コマンド
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
