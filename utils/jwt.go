package utils

import (
	"pAppServer/global"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UUID string `json:"uuid"`
	jwt.StandardClaims
}

func GenerateToke(uuid string) (string, error) {
	claims := Claims{uuid, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Second * time.Duration(global.Config.MaxAge)).Unix(),
	},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(global.Config.Jwt.SignedKey))
}

func VerifyToken(tokenString string) (string, bool) {
	var c Claims
	token, err := jwt.ParseWithClaims(tokenString, &c, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.Config.Jwt.SignedKey), nil
	})
	return c.UUID, token.Valid && err == nil
}
