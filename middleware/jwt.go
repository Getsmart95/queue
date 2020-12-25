package middleware

import (
	"queue/tokens"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
	"time"
)

func JWT() func(next httprouter.Handle) httprouter.Handle {
	return func(next httprouter.Handle) httprouter.Handle {
		return func(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
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

			claims := tokens.ParseToken(Token)
			if claims.ExpiresAt < time.Now().Unix() {
				http.Error(writer, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			next(writer, request, param)
			//ok, err := jwtcore.Verify(token, secret)
			//if err != nil {
			//	http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			//	return
			//}
			//
			//		if !ok {
			//			http.Error(writer, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			//			return
			//		}
			//
			//		payload := reflect.New(payloadType).Interface()
			//
			//		err = jwtcore.Decode(token, payload)
			//		if err != nil {
			//			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			//			return
			//		}
			//
			//		ok, err = jwtcore.IsNotExpired(payload, time.Now())
			//		if err != nil {
			//			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			//			return
			//		}
			//
			//		if !ok {
			//			http.Error(writer, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			//			return
			//		}
			//
			//		log.Print(payload)
			//
			//		ctx := context.WithValue(request.Context(), payloadContextKey, payload)
			//		next(writer, request.WithContext(ctx), param)
		}
	}
}

