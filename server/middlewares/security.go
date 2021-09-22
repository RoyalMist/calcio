package middlewares

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/o1egl/paseto"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/ed25519"
)

const HashCost = 15

var (
	p          *paseto.V2
	publicKey  ed25519.PublicKey
	privateKey ed25519.PrivateKey
)

func init() {
	p = paseto.NewV2()
	var err error
	publicKey, privateKey, err = ed25519.GenerateKey(nil)
	if err != nil {
		log.Fatal("impossible to create a valid ed255519 keys pair")
	}
}

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

func SignToken(userId string, claims map[string]string, validity time.Duration) (string, error) {
	now := time.Now()
	jsonToken := paseto.JSONToken{
		Issuer:     "calcio",
		Jti:        uuid.New().String(),
		Subject:    userId,
		Expiration: now.Add(validity),
		IssuedAt:   now,
		NotBefore:  now,
	}

	for k, v := range claims {
		jsonToken.Set(k, v)
	}

	return p.Sign(privateKey, jsonToken, nil)
}

func VerifyToken(token string) (paseto.JSONToken, error) {
	var jsonToken paseto.JSONToken
	err := p.Verify(token, publicKey, &jsonToken, nil)
	if time.Now().After(jsonToken.Expiration) {
		return paseto.JSONToken{}, fmt.Errorf("expired token")
	}

	return jsonToken, err
}
