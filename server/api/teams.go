package api

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Teams struct {
	log *zap.SugaredLogger
}

// TeamsModule makes the injectable available for FX.
var TeamsModule = fx.Provide(NewTeams)

// NewTeams creates a new injectable.
func NewTeams(logger *zap.SugaredLogger) *Teams {
	return &Teams{
		log: logger,
	}
}

func (t Teams) Start(router fiber.Router, middlewares ...fiber.Handler) {
	for _, middleware := range middlewares {
		if middleware != nil {
			router.Use(middleware)
		}
	}
}
