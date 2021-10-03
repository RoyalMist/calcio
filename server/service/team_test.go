package service

import (
	"context"
	"reflect"
	"testing"

	"calcio/ent"
	"calcio/ent/enttest"
	"calcio/server/helpers"
	"go.uber.org/zap"
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

	playerOne := client.User.Create().SetName("playerOne").SetPassword("password").SaveX(helpers.LoggedInCtx(true))
	playerTwo := client.User.Create().SetName("playerTwo").SetPassword("password").SaveX(helpers.LoggedInCtx(true))
	type args struct {
		playerOne string
		playerTwo string
		ctx       context.Context
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
				ctx:       context.Background(),
				playerOne: playerOne.ID.String(),
				playerTwo: playerTwo.ID.String(),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Team{
				log:    logger,
				client: client,
			}
			got, err := t.Create(tt.args.playerOne, tt.args.playerTwo, tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t1.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTeam_ListAll(t1 *testing.T) {
	type fields struct {
		log    *zap.SugaredLogger
		client *ent.Client
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    ent.Teams
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Team{
				log:    tt.fields.log,
				client: tt.fields.client,
			}
			got, err := t.ListAll(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t1.Errorf("ListAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("ListAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTeam_ListNotMember(t1 *testing.T) {
	type fields struct {
		log    *zap.SugaredLogger
		client *ent.Client
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    ent.Teams
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Team{
				log:    tt.fields.log,
				client: tt.fields.client,
			}
			got, err := t.ListNotMember(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t1.Errorf("ListNotMember() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("ListNotMember() got = %v, want %v", got, tt.want)
			}
		})
	}
}
