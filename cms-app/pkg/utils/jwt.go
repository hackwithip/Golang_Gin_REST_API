package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


func GenerateToken(email string, userId uint) (string, error) {
	
	var JWT_SECRET_KEY = os.Getenv("JWT_SECRET_KEY")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"email": email,
		"exp": time.Now().Add(time.Minute * 5).Unix(),
	})

	return token.SignedString([]byte(JWT_SECRET_KEY))
}