package middleware

import (
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	request "github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"

	"github.com/16francs/examin_go/config"
	"github.com/16francs/examin_go/domain/model"
)

/*
jwt tokenを生成する
	iss tokenの発行者
	sub tokenの利用者を一意に特定する識別子
	iat tokenの発行時間
	exp tokenの有効時間
	role userの役割
*/
func GenerateToken(user model.User) (string, error) {

	env, err := config.LoadEnv()
	if err != nil {
		log.Fatalf("alert: %s", err)
	}

	token := jwt.New(jwt.GetSigningMethod("HS256"))

	token.Claims = jwt.MapClaims{
		"iss":  env.Iss,
		"sub":  user.LoginID,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
		"role": user.Role,
	}

	return token.SignedString([]byte(env.Secret))
}

/*
	jwt tokenから情報を取り出す
	ここではlogin_idとroleを返却する
*/
func Parse(ctx *gin.Context) (string, int, error) {

	env, err := config.LoadEnv()
	if err != nil {
		log.Fatalf("alert: %s", err)
	}

	token, err := request.ParseFromRequest(ctx.Request, request.OAuth2Extractor,
		func(token *jwt.Token) (interface{}, error) {
			b := []byte(env.Secret)
			return b, nil
		})
	if err != nil {
		return "", -1, err
	}

	claims := token.Claims.(jwt.MapClaims)

	return claims["sub"].(string), int(claims["role"].(float64)), nil
}
