package api

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Games struct {
	app *fiber.App
	log *zap.SugaredLogger
}

// GamesModule makes the injectable available for FX.
var GamesModule = fx.Provide(NewGames)

// NewGames creates a new injectable.
func NewGames(app *fiber.App, logger *zap.SugaredLogger) *Games {
	return &Games{
		app: app,
		log: logger,
	}
}

func (g Games) Start(base string, middlewares ...fiber.Handler) {
	router := g.app.Group(base)
	for _, middleware := range middlewares {
		if middleware != nil {
			router.Use(middleware)
		}
	}

	// Routes go here
}
