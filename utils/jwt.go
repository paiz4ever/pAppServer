package utils

import (
	"pAppServer/global"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToke(username string) (string, error) {
	claims := Claims{username, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Duration(global.Config.MaxAge)).Unix(),
		Issuer:    username,
	},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(global.Config.Jwt.SignedKey))
}

func VerifyToken(tokenString string) bool {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.Config.Jwt.SignedKey), nil
	})
	return token.Valid && err != nil
}
