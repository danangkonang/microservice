package middleware

import (
	"net/http"
	"os"
	"strings"

	"html"

	"github.com/danangkonang/product/helper"
	"github.com/dgrijalva/jwt-go"
)

type MyClaims struct {
	jwt.StandardClaims
	UserId string
}

func Auth(next http.HandlerFunc) http.HandlerFunc {
	var key = os.Getenv("APP_KEY")
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		authorizationHeader := r.Header.Get("Authorization")

		if !strings.Contains(authorizationHeader, "Bearer") {
			helper.MakeRespon(w, http.StatusUnauthorized, "Unauthorizad ini", nil)
			return
		}
		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		})
		if err != nil {
			helper.MakeRespon(w, http.StatusUnauthorized, "invalid token", nil)
			return
		}

		if token.Valid {
			next(w, r)
		} else if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				helper.MakeRespon(w, http.StatusUnauthorized, "invalid token", nil)
				return
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				helper.MakeRespon(w, http.StatusUnauthorized, "token expired", nil)
				return
			} else {
				helper.MakeRespon(w, http.StatusUnauthorized, "invalid token", nil)
				return
			}
		} else {
			helper.MakeRespon(w, http.StatusUnauthorized, "invalid token", nil)
			return
		}
	}
}

func ExtractToken(r *http.Request) map[string]interface{} {
	authorizationHeader := r.Header.Get("Authorization")
	BearerToken := strings.Split(authorizationHeader, " ")
	tokenString := html.EscapeString(BearerToken[1])

	var key = os.Getenv("APP_KEY")

	claims := jwt.MapClaims{}
	token, _ := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if token.Valid {
		return claims
	}
	return claims
}
