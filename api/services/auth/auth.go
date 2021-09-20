package auth

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const HashCost = 15

func HashPassword(password string) (string, error) {
	if len(password) < 1 {
		return "", fmt.Errorf("password should not be empty")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), HashCost)
	return string(hash), err
}

func CheckPassword(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
