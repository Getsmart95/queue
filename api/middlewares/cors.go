package middlewares

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func CORS(next httprouter.Handle) httprouter.Handle {
		return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

			writer.Header().Set("Content-Type", "application/json, text/html")
			   writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			//writer.Header().Set("Access-Control-Allow-Origin", "*")
			writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS")
			//writer.Header().Set("Access-Control-Allow-Credentials", "true")
			writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, RefreshToken")
			//writer.Header().Set("Accept", "*/*")

			next(writer, request, params)
		}
	}
