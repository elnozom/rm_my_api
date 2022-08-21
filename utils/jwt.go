package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JWTSecret = []byte("!!SECRET!!")

func GenerateJWT(id uint, dbIndex int) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["dbIndex"] = dbIndex
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, _ := token.SignedString(JWTSecret)
	return t
}
