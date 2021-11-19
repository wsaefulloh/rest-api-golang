package helpers

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var mySigninKey = []byte("secrects")

type claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func NewToken(username string) *claims {
	return &claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 5).Unix(),
		},
	}
}

func (c *claims) Create() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return token.SignedString(mySigninKey)
}

func CheckToken(token string) (bool, string) {
	tokens, err := jwt.ParseWithClaims(token, &claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(mySigninKey), nil
	})

	if err != nil {
		return false, ""
	}

	claims := tokens.Claims.(*claims)
	return tokens.Valid, claims.Username
}
