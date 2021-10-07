package service

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"calcio/ent"
	"calcio/ent/enttest"
	"calcio/server/helpers"
	"go.uber.org/zap/zaptest"
)

func TestGame_List(t *testing.T) {
	logger := (zaptest.NewLogger(t)).Sugar()
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	_ = client.Schema.Create(context.Background())
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {

		}
	}(client)

	userOne := client.User.Create().SetName("userOne").SetPassword("password").SaveX(helpers.LoggedInCtx(true))
	userTwo := client.User.Create().SetName("userTwo").SetPassword("password").SaveX(helpers.LoggedInCtx(true))
	userThree := client.User.Create().SetName("userThree").SetPassword("password").SaveX(helpers.LoggedInCtx(true))
	userFour := client.User.Create().SetName("userFour").SetPassword("password").SaveX(helpers.LoggedInCtx(true))
	teamOne := client.Team.Create().SetName("TeamOne").AddUsers(userOne, userTwo).SaveX(helpers.LoggedInCtx(false))
	teamTwo := client.Team.Create().SetName("TeamTwo").AddUserIDs(userThree.ID, userFour.ID).SaveX(helpers.LoggedInCtx(false))
	participationGameOneFirstTeam := client.Participation.Create().SetTeam(teamOne).SaveX(helpers.LoggedInCtx(false))
	participationGameOneSecondTeam := client.Participation.Create().SetTeam(teamTwo).SaveX(helpers.LoggedInCtx(false))
	participationGameTwoFirstTeam := client.Participation.Create().SetTeam(teamOne).SaveX(helpers.LoggedInCtx(false))
	participationGameTwoSecondTeam := client.Participation.Create().SetTeam(teamTwo).SaveX(helpers.LoggedInCtx(false))
	gameOne := client.Game.Create().AddParticipations(participationGameOneFirstTeam, participationGameOneSecondTeam).SetDate(time.Now().Add(-10 * time.Minute)).SaveX(helpers.LoggedInCtx(false))
	gameTwo := client.Game.Create().AddParticipations(participationGameTwoFirstTeam, participationGameTwoSecondTeam).SaveX(helpers.LoggedInCtx(false))

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    ent.Games
		wantErr bool
	}{
		{
			name:    "an unauthenticated user should not be able to list games",
			args:    args{ctx: context.Background()},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "an authenticated user should be able to list all games",
			args:    args{ctx: helpers.LoggedInCtx(false)},
			want:    ent.Games{gameTwo, gameOne},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Game{
				log:    logger,
				client: client,
			}
			got, err := g.List(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(fmt.Sprint(got), fmt.Sprint(tt.want)) {
				t.Errorf("List() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGame_Create(t *testing.T) {
	logger := (zaptest.NewLogger(t)).Sugar()
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	_ = client.Schema.Create(context.Background())
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {

		}
	}(client)

	userOne := client.User.Create().SetName("userOne").SetPassword("password").SaveX(helpers.LoggedInCtx(true))
	userTwo := client.User.Create().SetName("userTwo").SetPassword("password").SaveX(helpers.LoggedInCtx(true))
	userThree := client.User.Create().SetName("userThree").SetPassword("password").SaveX(helpers.LoggedInCtx(true))
	userFour := client.User.Create().SetName("userFour").SetPassword("password").SaveX(helpers.LoggedInCtx(true))
	teamA := client.Team.Create().SetName("TeamOne").AddUsers(userOne, userTwo).SaveX(helpers.LoggedInCtx(false))
	teamB := client.Team.Create().SetName("TeamTwo").AddUserIDs(userThree.ID, userFour.ID).SaveX(helpers.LoggedInCtx(false))

	type args struct {
		teamA *ent.Team
		teamB *ent.Team
		ctx   context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    *ent.Game
		wantErr bool
	}{
		{
			name: "an unauthenticated user should not be able to create a game",
			args: args{
				teamA: teamA,
				teamB: teamB,
				ctx:   context.Background(),
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Game{
				log:    logger,
				client: client,
			}
			got, err := g.Create(tt.args.teamA, tt.args.teamB, tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(fmt.Sprint(got), fmt.Sprint(tt.want)) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}
