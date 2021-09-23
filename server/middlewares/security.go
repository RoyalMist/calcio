package middlewares

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/vk-rv/pvx"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/ed25519"
)

const HashCost = 15

var (
	pv4       = pvx.NewPV4Public()
	publicKey *pvx.AsymPublicKey
	secretKey *pvx.AsymSecretKey
)

func init() {
	pk, sk, err := ed25519.GenerateKey(nil)
	if err != nil {
		log.Fatal("impossible to create a valid ed255519 keys pair")
	}

	secretKey = pvx.NewAsymmetricSecretKey(sk, pvx.Version4)
	publicKey = pvx.NewAsymmetricPublicKey(pk, pvx.Version4)
}

type Claims struct {
	pvx.RegisteredClaims
	UserId string
}

func (c Claims) Valid() error {
	if err := c.RegisteredClaims.Valid(); err != nil {
		return err
	}

	if len(c.UserId) < 1 {
		return fmt.Errorf("invalid userId %s", c.UserId)
	}

	return nil
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
	claims.KeyID = uuid.New().String()
	claims.Expiration = &expiration
	return pv4.Sign(secretKey, claims)
}

func VerifyToken(token string) (Claims, error) {
	var claims Claims
	err := pv4.Verify(token, publicKey).ScanClaims(&claims)
	return claims, err
}