package helper

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

type Login struct {
	Id   int
	Name string
}

func CreateJwt(data Login) (string, error) {
	godotenv.Load()

	key := os.Getenv("JWT_KEY")

	claims := jwt.MapClaims{
		"id":   data.Id,
		"name": data.Name,
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
		"iat":  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
