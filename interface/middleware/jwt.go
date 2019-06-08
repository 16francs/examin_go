package middleware

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"

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
