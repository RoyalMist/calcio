package service

import (
	"context"

	"calcio/ent"
	"calcio/ent/game"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Game struct {
	log    *zap.SugaredLogger
	client *ent.Client
}

// GameModule makes the injectable available for FX.
var GameModule = fx.Provide(NewGame)

// NewGame creates a new injectable.
func NewGame(logger *zap.SugaredLogger, client *ent.Client) *Game {
	return &Game{
		log:    logger,
		client: client,
	}
}

func (g Game) List(ctx context.Context) (ent.Games, error) {
	return g.client.Game.Query().WithParticipations().Order(ent.Desc(game.FieldDate)).All(ctx)
}

func (g Game) Create(teamA, teamB *ent.Team, ctx context.Context) (*ent.Game, error) {
	return nil, nil
}
