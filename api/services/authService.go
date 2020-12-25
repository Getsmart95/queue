package services

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type MyCustomClaims struct {
	Login string `json:"login"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func SetToken(login, password string) string {
	mySigningKey := []byte("AllYourBase")

	// Create the Claims
	claims := MyCustomClaims{
		login,
		password,
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(mySigningKey)
	fmt.Printf("I am a token = %v\n", ss)
	return ss
}