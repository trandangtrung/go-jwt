// services/jwt_service.go
package services

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var SECRET_KEY = []byte("your_secret_key")

func CreateToken(userID int) (string, error) {
	claims := jwt.StandardClaims{
		Id:        strconv.Itoa(userID),
		ExpiresAt: time.Now().Add(time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SECRET_KEY)
}

func VerifyToken(tokenString string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return SECRET_KEY, nil
	})
	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
