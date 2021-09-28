package service

import (
	"context"
	"reflect"
	"testing"

	"calcio/ent"
	"calcio/ent/enttest"
	"calcio/server/security"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap/zaptest"
)

func TestUser_Login(t *testing.T) {
	logger := (zaptest.NewLogger(t)).Sugar()
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	_ = client.Schema.Create(context.Background())
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {

		}
	}(client)

	fake := client.User.Create().SetName("fake").SetPassword("password").SaveX(security.NewContext(context.Background(), security.Claims{
		UserId:  uuid.New().String(),
		IsAdmin: true,
	}))

	type args struct {
		name     string
		password string
	}
	tests := []struct {
		name                string
		args                args
		wantRetrievedUserId string
		wantErr             bool
	}{
		{
			name: "invalid user name should return an error",
			args: args{
				name:     "coucou",
				password: "password",
			},
			wantRetrievedUserId: "",
			wantErr:             true,
		},
		{
			name: "invalid password should return an error",
			args: args{
				name:     "fake",
				password: "coucou",
			},
			wantRetrievedUserId: "",
			wantErr:             true,
		},
		{
			name: "incorrect name and password should return an error",
			args: args{
				name:     "coucou",
				password: "coucou",
			},
			wantRetrievedUserId: "",
			wantErr:             true,
		},
		{
			name: "valid user and password should return a valid user",
			args: args{
				name:     "fake",
				password: "password",
			},
			wantRetrievedUserId: fake.ID.String(),
			wantErr:             false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := User{
				log:    logger,
				client: client,
			}
			gotRetrievedUser, err := u.Login(tt.args.name, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotRetrievedUser != nil && !reflect.DeepEqual(gotRetrievedUser.ID.String(), tt.wantRetrievedUserId) {
				t.Errorf("Login() gotRetrievedUser = %v, want %v", gotRetrievedUser.ID.String(), tt.wantRetrievedUserId)
			}
		})
	}
}

func TestUser_List(t *testing.T) {
	logger := (zaptest.NewLogger(t)).Sugar()
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	_ = client.Schema.Create(context.Background())
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {

		}
	}(client)

	loggedInCtx := security.NewContext(context.Background(), security.Claims{
		UserId:  uuid.New().String(),
		IsAdmin: true,
	})

	client.User.Create().SetName("fake1").SetPassword("password").SaveX(loggedInCtx)
	client.User.Create().SetName("fake2").SetPassword("password").SaveX(loggedInCtx)
	client.User.Create().SetName("fake3").SetPassword("password").SaveX(loggedInCtx)

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "An unauthenticated user should not be able to retrieve the list of users",
			args:    args{ctx: context.Background()},
			want:    0,
			wantErr: true,
		},
		{
			name:    "An authenticated user should be able to retrieve the list of users",
			args:    args{ctx: loggedInCtx},
			want:    3,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := User{
				log:    logger,
				client: client,
			}
			got, err := u.List(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(len(got), tt.want) {
				t.Errorf("List() got = %v, want a length of %v", got, tt.want)
			}
		})
	}
}
