package api

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Auth struct {
	app *fiber.App
	log *zap.SugaredLogger
}

// AuthModule makes the injectable available for FX.
var AuthModule = fx.Provide(NewAuth)

// NewAuth creates a new injectable.
func NewAuth(app *fiber.App, logger *zap.SugaredLogger) *Auth {
	return &Auth{
		app: app,
		log: logger,
	}
}

func (a Auth) Start(base string, middlewares ...fiber.Handler) {
	router := a.app.Group(base)
	for _, middleware := range middlewares {
		if middleware != nil {
			router.Use(middleware)
		}
	}

	router.Get("/login", login)
}

func login(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello")
}
