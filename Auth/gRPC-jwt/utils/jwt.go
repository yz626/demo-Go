package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	jwt.StandardClaims
	Name string
}

var jwtSecret = []byte("secret")

func CreateToken(name string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
			Issuer:    "test",
			IssuedAt:  time.Now().Unix(),
		},
	})

	return token.SignedString(jwtSecret)
}

func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if Claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return Claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
