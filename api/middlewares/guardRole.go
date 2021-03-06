package middlewares

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"queue/api/services"
)

func GuardRole(server *services.UserService) func(next httprouter.Handle) httprouter.Handle {
	return func(next httprouter.Handle) httprouter.Handle {
		return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

			claims := JWT(writer, request, params)
			if claims == nil {
				json.NewEncoder(writer).Encode(http.StatusUnauthorized)
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
