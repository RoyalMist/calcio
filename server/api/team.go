package api

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Team struct {
	app *fiber.App
	log *zap.SugaredLogger
}

// TeamModule makes the injectable available for FX.
var TeamModule = fx.Provide(NewTeam)

// NewTeam creates a new injectable.
func NewTeam(app *fiber.App, logger *zap.SugaredLogger) *Team {
	return &Team{
		app: app,
		log: logger,
	}
}

func (t Team) Start(base string, middlewares ...fiber.Handler) {
	router := t.app.Group(base)
	for _, middleware := range middlewares {
		if middleware != nil {
			router.Use(middleware)
		}
	}

	router.Get("/list", list)
}

func list(ctx *fiber.Ctx) error {
	return ctx.SendString("Coucou")
}
