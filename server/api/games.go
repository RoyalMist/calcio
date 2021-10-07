package api

import (
	"calcio/server/service"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Games struct {
	log  *zap.SugaredLogger
	game *service.Game
}

// GamesModule makes the injectable available for FX.
var GamesModule = fx.Provide(NewGames)

// NewGames creates a new injectable.
func NewGames(logger *zap.SugaredLogger, game *service.Game) *Games {
	return &Games{
		log:  logger,
		game: game,
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
