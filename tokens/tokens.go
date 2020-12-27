package tokens

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyCustomClaims struct {
	Login string `json:"login"`
	Password string `json:"password"`
	jwt.StandardClaims
}
func SetToken(login string, password string) (Token string, expiredIn int) {
	mySigningKey := []byte("AllYourBase")

	// Create the Claims
	lifetime := 1500
	expirationTime := time.Now().Add(time.Duration(lifetime) * time.Second)
	claims := &MyCustomClaims{
		login,
		password,
		jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	Token, _ = token.SignedString(mySigningKey)
	fmt.Printf("I am a token = %v\n", Token)
	return Token, lifetime
}

func ParseToken(tokenString string) (claims *MyCustomClaims, ok bool, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})
	if err != nil {
		return nil, false, err
	}
	if Claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return Claims, true,nil
	} else {
		return nil,false, nil
	}
}