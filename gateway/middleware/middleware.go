package middleware

import (
	"net/http"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader == "" {
			MakeRespon(w, 401, "Anauthorize", nil)
			return
		}
		next(w, r)
	}
}
