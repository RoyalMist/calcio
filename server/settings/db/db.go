package db

import (
	"context"
	"fmt"

	"calcio/ent"
	_ "calcio/ent/runtime"
	"calcio/server/settings/config"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Module permits create a FX module.
var Module = fx.Provide(New)

// New create a new instance.
func New(lifecycle fx.Lifecycle, config *config.Config, logger *zap.SugaredLogger) (client *ent.Client, err error) {
	lifecycle.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			logger.Info("Closing DB connection")
			return client.Close()
		},
	})

	if client, err = ent.Open(config.DbDriver(), config.DbUrl()); err != nil {
		err = fmt.Errorf("impossible to connect to %s with driver = %s, %v", config.DbUrl(), config.DbDriver(), err)
		logger.Fatal(err)
		return
	}

	if err = client.Schema.Create(context.Background()); err != nil {
		err = fmt.Errorf("impossible to migrate schema, %v", err)
		logger.Fatal(err)
	}

	return
}
