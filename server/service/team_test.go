package service

import (
	"context"
	"fmt"
	"reflect"
	"sort"
	"testing"

	"calcio/ent"
	"calcio/ent/enttest"
	"calcio/server/helpers"
	"go.uber.org/zap/zaptest"
)

func TestTeam_Create(t1 *testing.T) {
	logger := (zaptest.NewLogger(t1)).Sugar()
	client := enttest.Open(t1, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	_ = client.Schema.Create(context.Background())
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {

		}
	}(client)

	playerOne := client.User.Create().SetName("player-one").SetPassword("password").SaveX(helpers.LoggedInCtx(true))
	playerTwo := client.User.Create().SetName("player-two").SetPassword("password").SaveX(helpers.LoggedInCtx(true))
	type args struct {
		otherPlayer string
		ctx         context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    *ent.Team
		wantErr bool
	}{
		{
			name: "an unauthenticated user should not be able to create a team",
			args: args{
				ctx:         context.Background(),
				otherPlayer: playerTwo.ID.String(),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "an authenticated user should be able to create a solo team",
			args: args{
				ctx: helpers.LoggedInCtxWithIdAndName(playerOne.ID, "player-one", false),
			},
			want: &ent.Team{
				Name: playerOne.Name,
			},
			wantErr: false,
		},
		{
			name: "an authenticated user should be able to create a duo team",
			args: args{
				ctx:         helpers.LoggedInCtxWithIdAndName(playerOne.ID, "player-one", false),
				otherPlayer: playerTwo.ID.String(),
			},
			want: &ent.Team{
				Name: fmt.Sprintf("%s & %s", playerOne.Name, playerTwo.Name),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Team{
				log:    logger,
				client: client,
			}
			got, err := t.Create(tt.args.otherPlayer, tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t1.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil {
				tt.want.ID = got.ID
				if got.String() != tt.want.String() {
					t1.Errorf("Create() got = %v, wanted with name %s", got, tt.want)
				}
			}
		})
	}
}

func TestTeam_List(t1 *testing.T) {
	logger := (zaptest.NewLogger(t1)).Sugar()
	client := enttest.Open(t1, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	_ = client.Schema.Create(context.Background())
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {

		}
	}(client)

	playerOne := client.User.Create().SetName("player-one").SetPassword("password").SaveX(helpers.LoggedInCtx(true))
	playerTwo := client.User.Create().SetName("player-two").SetPassword("password").SaveX(helpers.LoggedInCtx(true))

	service := Team{
		log:    logger,
		client: client,
	}

	teamPlayerOne, _ := service.Create("", helpers.LoggedInCtxWithIdAndName(playerOne.ID, playerOne.Name, false))
	teamPlayerTwo, _ := service.Create("", helpers.LoggedInCtxWithIdAndName(playerTwo.ID, playerTwo.Name, false))
	teamPlayerOneTwo, _ := service.Create(playerTwo.ID.String(), helpers.LoggedInCtxWithIdAndName(playerOne.ID, playerOne.Name, false))
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    ent.Teams
		wantErr bool
	}{
		{
			name:    "an unauthenticated user should not be able to list teams",
			args:    args{ctx: context.Background()},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "an authenticated (playerOne) user should be able to list teams where he is added",
			args:    args{ctx: helpers.LoggedInCtxWithIdAndName(playerOne.ID, playerOne.Name, false)},
			want:    ent.Teams{teamPlayerOne, teamPlayerOneTwo},
			wantErr: false,
		},
		{
			name:    "an authenticated (playerTwo) user should be able to list teams where he is added",
			args:    args{ctx: helpers.LoggedInCtxWithIdAndName(playerTwo.ID, playerTwo.Name, false)},
			want:    ent.Teams{teamPlayerTwo, teamPlayerOneTwo},
			wantErr: false,
		},
		{
			name:    "an admin user should be able to list all teams",
			args:    args{ctx: helpers.LoggedInCtx(true)},
			want:    ent.Teams{teamPlayerOne, teamPlayerTwo, teamPlayerOneTwo},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Team{
				log:    logger,
				client: client,
			}
			got, err := t.List(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t1.Errorf("ListAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			sort.Slice(got, func(i, j int) bool {
				return got[i].ID.String() < got[j].ID.String()
			})
			sort.Slice(tt.want, func(i, j int) bool {
				return tt.want[i].ID.String() < tt.want[j].ID.String()
			})
			if !reflect.DeepEqual(fmt.Sprintf("%v", got), fmt.Sprintf("%v", tt.want)) {
				t1.Errorf("ListAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}
