package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWT_SECRET_KEY = os.Getenv("JWT_SECRET_KEY")

func GenerateToken(email string, userId uint) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"email": email,
		"exp": time.Now().Add(time.Minute * 5).Unix(),
	})

	return token.SignedString([]byte(JWT_SECRET_KEY))
}

func VerifyToken(token string) ( uint, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		// Check if the signing method is correct
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		// Return byte stringed secretKey
		return []byte(JWT_SECRET_KEY), nil
	})

	if err != nil {
		return 0, errors.New("failed to parse token")
	}

	if !parsedToken.Valid {
        return 0, errors.New("invalid token")
    }

    claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("invalid token claims")
	}

	userId := claims["userId"].(float64)

	return uint(userId), nil
}