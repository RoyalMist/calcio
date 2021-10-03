package service

import (
	"context"
	"fmt"

	"calcio/ent"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Team struct {
	log    *zap.SugaredLogger
	client *ent.Client
}

// TeamModule makes the injectable available for FX.
var TeamModule = fx.Provide(NewTeam)

// NewTeam creates a new injectable.
func NewTeam(logger *zap.SugaredLogger, client *ent.Client) *Team {
	return &Team{
		log:    logger,
		client: client,
	}
}

func (t Team) Create(playerOne string, playerTwo string, ctx context.Context) (*ent.Team, error) {
	var separator string
	if playerTwo != "" {
		separator = " & "
	}

	name := fmt.Sprintf("%s%s%s", playerOne, separator, playerTwo)
	operation := t.client.Team.Create().SetName(name)
	if id, err := uuid.Parse(playerOne); err != nil {
		return nil, errors.Wrapf(err, "impossible to parse uuid for playerOne got: %s", playerOne)
	} else {
		operation.AddPlayerIDs(id)
	}

	if playerTwo != "" {
		if id, err := uuid.Parse(playerTwo); err != nil {
			return nil, errors.Wrapf(err, "impossible to parse uuid for playerTwo got: %s", playerTwo)
		} else {
			operation.AddPlayerIDs(id)
		}
	}

	return operation.Save(ctx)
}

func (t Team) ListAll(ctx context.Context) (ent.Teams, error) {
	return nil, nil
}

func (t Team) ListNotMember(ctx context.Context) (ent.Teams, error) {
	return nil, nil
}
