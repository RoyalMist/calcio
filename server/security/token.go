package security

import (
	ed255192 "crypto/ed25519"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/vk-rv/pvx"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/ed25519"
)

const HashCost = 15

var (
	pvX       = pvx.NewPV2Public()
	publicKey ed255192.PublicKey
	secretKey ed25519.PrivateKey
)

func init() {
	var err error
	publicKey, secretKey, err = ed25519.GenerateKey(nil)
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

func SignToken(claims Claims, validity time.Duration) (string, error) {
	expiration := time.Now().Add(validity)
	claims.Issuer = "calcio"
	claims.KeyID = uuid.NewString()
	claims.Expiration = &expiration
	return pvX.Sign(secretKey, claims, nil)
}

func verifyToken(token string) (Claims, error) {
	var claims Claims
	err := pvX.Verify(token, publicKey).ScanClaims(&claims)
	return claims, err
}

func IsAuthenticated(ctx *fiber.Ctx) error {
	authHeader := ctx.Get(fiber.HeaderAuthorization)
	if !strings.Contains(authHeader, "Bearer ") {
		return fiber.ErrBadRequest
	}

	token := strings.Split(authHeader, " ")[1]
	claims, err := verifyToken(token)
	if err != nil {
		return fiber.ErrUnauthorized
	}

	ctx.SetUserContext(NewContext(ctx.UserContext(), claims))
	return ctx.Next()
}
