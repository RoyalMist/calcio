package api

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Games struct {
	log *zap.SugaredLogger
}

// GamesModule makes the injectable available for FX.
var GamesModule = fx.Provide(NewGames)

// NewGames creates a new injectable.
func NewGames(logger *zap.SugaredLogger) *Games {
	return &Games{
		log: logger,
	}
}

func (g Games) Start(router fiber.Router, middlewares ...fiber.Handler) {
	for _, middleware := range middlewares {
		if middleware != nil {
			router.Use(middleware)
		}
	}

	// Routes go here
}
