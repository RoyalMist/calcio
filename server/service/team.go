package service

import (
	"context"
	"fmt"

	"calcio/ent"
	"calcio/ent/team"
	"calcio/ent/user"
	"calcio/server/security"
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

func (t Team) Create(otherPlayerId string, ctx context.Context) (*ent.Team, error) {
	operation := t.client.Team.Create()
	var otherPlayer *ent.User
	var name string
	initiator := security.FromContext(ctx)
	initiatorId, initiatorErr := uuid.Parse(initiator.UserId)
	if initiatorErr != nil {
		return nil, errors.Wrapf(initiatorErr, "impossible to parse uuid for initiator with id: %s", initiator.UserId)
	}

	if otherPlayerId != "" {
		uid, err := uuid.Parse(otherPlayerId)
		if err != nil {
			return nil, errors.Wrapf(err, "impossible to parse uuid for otherPlayer with id : %s", otherPlayerId)
		}

		otherPlayer, err = t.client.User.Query().Where(user.ID(uid)).First(ctx)
		if err != nil {
			return nil, errors.Wrapf(err, "impossible to find otherPlayer with uuid : %s", otherPlayerId)
		}

		name = fmt.Sprintf("%s & %s", initiator.UserName, otherPlayer.Name)
		operation = operation.SetName(name).AddUserIDs(initiatorId, otherPlayer.ID)
	} else {
		operation = operation.SetName(initiator.UserName).AddUserIDs(initiatorId)
	}

	return operation.Save(ctx)
}

func (t Team) List(ctx context.Context) (ent.Teams, error) {
	claims := security.FromContext(ctx)
	if claims.IsAdmin {
		return t.client.Team.Query().All(ctx)
	} else {
		id, err := uuid.Parse(claims.UserId)
		if err != nil {
			return nil, errors.Wrapf(err, "impossible to parse userId %s", claims.UserId)
		}

		return t.client.Team.Query().Where(team.HasUsersWith(user.ID(id))).All(ctx)
	}
}
