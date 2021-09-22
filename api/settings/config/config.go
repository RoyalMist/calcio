package config

import (
	"crypto/ed25519"

	"github.com/spf13/viper"
	"github.com/vk-rv/pvx"
	"go.uber.org/fx"
)

const (
	ApiHost  = "API_HOST"
	DbDriver = "DB_DRIVER"
	DbUrl    = "DB_URL"
)

type Config struct {
	apiHost   string
	dbDriver  string
	dbUrl     string
	publicKey *pvx.AsymPublicKey
	secretKey *pvx.AsymSecretKey
}

func (c Config) PublicKey() *pvx.AsymPublicKey {
	return c.publicKey
}

func (c Config) SecretKey() *pvx.AsymSecretKey {
	return c.secretKey
}

func (c Config) DbUrl() string {
	return c.dbUrl
}

func (c Config) DbDriver() string {
	return c.dbDriver
}

func (c Config) ApiHost() string {
	return c.apiHost
}

// Module makes the injectable available for FX.
var Module = fx.Provide(New)

// New creates a new injectable.
func New() *Config {
	viper.SetDefault(ApiHost, ":4000")
	viper.SetDefault(DbDriver, "postgres")
	viper.SetDefault(DbUrl, "host=localhost port=5432 user=postgres dbname=calcio password=postgres sslmode=disable")
	viper.AutomaticEnv()

	publicKey, privateKey, _ := ed25519.GenerateKey(nil)
	sk := pvx.NewAsymmetricSecretKey(privateKey, pvx.Version4)
	pk := pvx.NewAsymmetricPublicKey(publicKey, pvx.Version4)

	return &Config{
		apiHost:   viper.GetString(ApiHost),
		dbDriver:  viper.GetString(DbDriver),
		dbUrl:     viper.GetString(DbUrl),
		publicKey: pk,
		secretKey: sk,
	}
}
