package middleware

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	request "github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"

	"github.com/16francs/examin_go/domain/model"
)

const (
	secret = "4CHtx3AAnOPrxWYvqsHou6w8b8aO3BF7Ey93/D4clbBsgMDZf9Zt+Q=="
	iss    = "16france"
)

func GenerateToken(user model.User) (string, error) {

	token := jwt.New(jwt.GetSigningMethod("HS256"))

	token.Claims = jwt.MapClaims{
		"iss":  iss,
		"sub":  user.LoginID,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
		"role": user.Role,
	}

	return token.SignedString([]byte(secret))
}

func Parse(ctx *gin.Context) (string, error) {

	token, err := request.ParseFromRequest(ctx.Request, request.OAuth2Extractor,
		func(token *jwt.Token) (interface{}, error) {
			b := []byte(secret)
			return b, nil
		})
	if err != nil {
		return "", err
	}
	claims := token.Claims.(jwt.MapClaims)

	return claims["sub"].(string), nil
}
