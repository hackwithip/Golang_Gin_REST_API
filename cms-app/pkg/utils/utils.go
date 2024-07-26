package utils

import (
	"time"

	"golang.org/x/exp/rand"
)

func GenerateRandomPassword() string {
	const passwordLength = 10
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()_+[]{}|;:,.<>?"

	var password []byte
	rand.Seed(uint64(time.Now().UnixNano())) // Seed the random number generator

	for i := 0; i < passwordLength; i++ {
		randomIndex := rand.Intn(len(charset))
		password = append(password, charset[randomIndex])
	}

	return string(password)
}
