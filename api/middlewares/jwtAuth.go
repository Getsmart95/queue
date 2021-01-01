package middlewares

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func JWTAuth() func(next httprouter.Handle) httprouter.Handle {
	return func(next httprouter.Handle) httprouter.Handle {
		return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

			claims := JWT(writer, request, params)
			if claims == nil {
				json.NewEncoder(writer).Encode(http.StatusUnauthorized)
				return
			}

			next(writer, request, params)
		}
	}
}


