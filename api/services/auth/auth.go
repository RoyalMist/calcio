package auth

import (
	"errors"
	"fmt"

	"calcio/api/settings/config"
	"github.com/vk-rv/pvx"
	"go.uber.org/fx"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	pv4       *pvx.ProtoV4Public
	publicKey *pvx.AsymPublicKey
	secretKey *pvx.AsymSecretKey
}

type Claims struct {
	pvx.RegisteredClaims
	UserId string
}

func (c Claims) Valid() error {
	validationErr := &pvx.ValidationError{}
	if err := c.RegisteredClaims.Valid(); err != nil {
		errors.As(err, &validationErr)
	}

	if c.UserId == "" {
		return fmt.Errorf("unvalid claims")
	}

	return nil
}

// Module makes the injectable available for FX.
var Module = fx.Provide(New)

// New creates a new injectable.
func New(config *config.Config) *Auth {
	return &Auth{
		publicKey: config.PublicKey(),
		secretKey: config.SecretKey(),
		pv4:       pvx.NewPV4Public(),
	}
}

const HashCost = 15

func (Auth) HashPassword(password string) (string, error) {
	if len(password) < 1 {
		return "", fmt.Errorf("password should not be empty")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), HashCost)
	return string(hash), err
}

func (Auth) CheckPassword(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func (a Auth) SignToken(claims Claims) (string, error) {
	return a.pv4.Sign(a.secretKey, claims)
}

func (a Auth) VerifyToken(token string) (Claims, error) {
	var claims Claims
	if err := a.pv4.Verify(token, a.publicKey).ScanClaims(&claims); err != nil {
		return Claims{}, err
	}

	return claims, nil
}
