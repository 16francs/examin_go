package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

// Logger - ログの出力
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		buf, _ := ioutil.ReadAll(c.Request.Body)
		rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf)) // ログ出力で使用する用
		rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf)) // ログ出力以降で使用する用

		// リクエストの処理
		c.Request.Body = rdr2
		c.Next()

		// リクエスト値のログを出力
		requestLogger(rdr1)

		// レスポンス値のログを出力
	}
}

// requestLogger - リクエストボディの読み込み
func requestLogger(reader io.Reader) {
	// json ファイルの読み込み -> 値がなければ return
	data, _ := ioutil.ReadAll(reader)
	if len(data) == 0 {
		return
	}

	// json を データ型に変換する
	var params map[string]interface{}
	if err := json.Unmarshal([]byte(data), &params); err != nil {
		log.Printf("warn: JSONの整形に失敗しました")
		return
	}

	// password のみフィルターにかける
	if params["password"] != nil {
		params["password"] = "********"
	}

	// ログの整形・出力 -> リファクタ対象
	message := ""
	for key, value := range params {
		message += fmt.Sprintf("%s: '%s' ", key, value)
	}
	log.Printf("info: params: { %s}", message)
}
