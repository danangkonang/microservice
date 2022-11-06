package helper

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

type MyClaims struct {
	jwt.StandardClaims
	UserId string
}

func GenerateToken(id string) (string, error) {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	var APP_KEY = os.Getenv("APP_KEY")
	var APP_NAME = os.Getenv("APP_NAME")

	var LOGIN_EXPIRATION_DURATION = time.Duration(24) * time.Hour
	var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
	var JWT_SIGNATURE_KEY = []byte(APP_KEY)
	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    APP_NAME,
			ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
		},
		UserId: id,
	}
	token := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		claims,
	)
	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
