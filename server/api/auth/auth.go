package auth

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Auth struct {
	app *fiber.App
	log *zap.SugaredLogger
}

// Module makes the injectable available for FX.
var Module = fx.Provide(New)

// New creates a new injectable.
func New(app *fiber.App, logger *zap.SugaredLogger) *Auth {
	return &Auth{
		app: app,
		log: logger,
	}
}
