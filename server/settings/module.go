package settings

import (
	"calcio/server/settings/config"
	"calcio/server/settings/db"
	"calcio/server/settings/log"
	"calcio/server/settings/server"
	"go.uber.org/fx"
)

// Module makes the collection of injectables available for FX.
var Module = fx.Options(
	config.Module,
	db.Module,
	log.Module,
	server.Module,
)
