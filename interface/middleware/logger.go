package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

// RequestLogger - リクエスト値のログを出力
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		buf, _ := ioutil.ReadAll(c.Request.Body)
		rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf)) // ログ出力で使用する用
		rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf)) // ログ出力以降で使用する用

		// ログの出力
		log.Printf("info: params: %v", readBody(rdr1))

		c.Request.Body = rdr2
		c.Next()
	}
}

// readBody - リクエストボディの読み込み
func readBody(reader io.Reader) string {
	buf := new(bytes.Buffer)
	data, _ := ioutil.ReadAll(reader)

	// data に値がなければ return
	if len(data) == 0 {
		return buf.String()
	}

	if err := json.Compact(buf, []byte(data)); err != nil {
		log.Printf("warn: JSONの整形に失敗しました")
	}

	return buf.String()
}
