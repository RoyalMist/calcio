package settings

import (
	"calcio/api/settings/config"
	"calcio/api/settings/log"
	"calcio/api/settings/server"
	"go.uber.org/fx"
)

// Module makes the collection of injectables available for FX.
var Module = fx.Options(
	config.Module,
	log.Module,
	server.Module,
)
