package service

import (
	"context"

	"calcio/ent"
	"calcio/ent/game"
	"github.com/pkg/errors"
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
	tx, err := g.client.Tx(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "impossible to create transaction while trying to create a game with teamA %p and teamB %p", teamA, teamB)
	}

	participationA, err := tx.Participation.Create().SetTeam(teamA).Save(ctx)
	if err != nil {
		_ = tx.Rollback()
		return nil, errors.Wrapf(err, "impossible to create a participation for teamA %p", teamA)
	}

	participationB, err := tx.Participation.Create().SetTeam(teamB).Save(ctx)
	if err != nil {
		_ = tx.Rollback()
		return nil, errors.Wrapf(err, "impossible to create a participation for teamB %p", teamB)
	}

	newGame, err := tx.Game.Create().AddParticipations(participationA, participationB).Save(ctx)
	if err != nil {
		_ = tx.Rollback()
		return nil, errors.Wrapf(err, "impossible to create a game with teamA %p and teamB %p", teamA, teamB)
	}

	return newGame, tx.Commit()
}
