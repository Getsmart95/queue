package middlewares

import (
	"fmt"
	"queue/tokens"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
	"time"
)

func JWT() func(next httprouter.Handle) httprouter.Handle {
	return func(next httprouter.Handle) httprouter.Handle {
		return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
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
			fmt.Println(err)

			if err != nil || ok == false {
				http.Error(writer, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			fmt.Println(claims)
			if claims.ExpiresAt < time.Now().Unix() {
				http.Error(writer, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			next(writer, request, params)
		}
	}
}

