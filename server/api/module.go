package api

import (
	"calcio/server/api/auth"
	"go.uber.org/fx"
)

// Module makes the collection of injectables available for FX.
var Module = fx.Options(
	auth.Module,
)
