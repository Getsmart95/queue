package middlewares

import (
	"time"
	"queue/tokens"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"queue/api/services"
	"strings"
)

func GuardRole(server *services.UserService) func(next httprouter.Handle) httprouter.Handle {
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
			if err != nil || ok == false {
				http.Error(writer, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			if claims.ExpiresAt < time.Now().Unix() {
				http.Error(writer, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			role, _ := server.GetRoleByUser(claims.Login)
			if role.Name == "admin" {
				next(writer, request, params)
				return
			}
			http.Error(writer, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}
	}
}
