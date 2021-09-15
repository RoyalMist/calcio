package config

import (
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

const (
	ApiHost  = "API_HOST"
	DbDriver = "DB_DRIVER"
	DbUrl    = "DB_URL"
)

type Config struct {
	apiHost  string
	dbDriver string
	dbUrl    string
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
	viper.SetDefault(DbUrl, "host=localhost port=5432 user=postgres dbname=change password=postgres sslmode=disable")
	viper.AutomaticEnv()

	return &Config{
		apiHost:  viper.GetString(ApiHost),
		dbDriver: viper.GetString(DbDriver),
		dbUrl:    viper.GetString(DbUrl),
	}
}
