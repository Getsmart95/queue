package middlewares

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"queue/tokens"
	"strings"
	"time"
)

func JWT(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) (claims *tokens.Claims) {
	bearerToken := request.Header.Get("Authorization")
	if bearerToken == "" {
		http.Error(writer, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	if !strings.HasPrefix(bearerToken, "Bearer ") {
		http.Error(writer, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	Token := bearerToken[len("Bearer "):]

	claims,	ok, err := tokens.ParseToken(Token)

	if err != nil || ok == false {
		http.Error(writer, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	if claims.ExpiresAt < time.Now().Unix() {
		http.Error(writer, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	return claims
}

