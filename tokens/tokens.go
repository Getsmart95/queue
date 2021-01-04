package tokens

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	UserID int `json:"user_id"`
	Login string `json:"login"`
	jwt.StandardClaims
}

func SetToken(UserID int, Login string) (Token string) {
	mySigningKey := []byte("AllYourBase")

	// Create the Claims
	lifetime := 1500
	expirationTime := time.Now().Add(time.Duration(lifetime) * time.Second)
	claims := &Claims{
		UserID,
		Login,
		jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	Token, _ = token.SignedString(mySigningKey)
	return Token
}

func ParseToken(tokenString string) (claims *Claims, ok bool, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})
	if err != nil {
		return nil, false, err
	}
	if Claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return Claims, true,nil
	} else {
		return nil,false, nil
	}
}
