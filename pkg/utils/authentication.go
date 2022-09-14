package utils

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var mSK = []byte("key")

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodES256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = "User"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	tokenString, err := token.SignedString(mSK)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return tokenString, nil
}
