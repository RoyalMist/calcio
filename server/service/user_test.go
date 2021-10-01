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
		UserId:  uuid.NewString(),
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

	adminCtx := security.NewContext(context.Background(), security.Claims{
		UserId:  uuid.NewString(),
		IsAdmin: true,
	})

	client.User.Create().SetName("fake1").SetPassword("password").SaveX(adminCtx)
	client.User.Create().SetName("fake2").SetPassword("password").SaveX(adminCtx)
	client.User.Create().SetName("fake3").SetPassword("password").SaveX(adminCtx)

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
			name: "An authenticated user should be able to retrieve the list of users",
			args: args{ctx: security.NewContext(context.Background(), security.Claims{
				UserId:  uuid.NewString(),
				IsAdmin: false,
			})},
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

func TestUser_Create(t *testing.T) {
	logger := (zaptest.NewLogger(t)).Sugar()
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	_ = client.Schema.Create(context.Background())
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {

		}
	}(client)

	type args struct {
		usr ent.User
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    *ent.User
		wantErr bool
	}{
		{
			name: "an unauthenticated user should not be able to create a user",
			args: args{
				usr: ent.User{
					Name:     "user",
					Password: "password",
					Admin:    false,
				},
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "an authenticated user with no admin rights should not be able to create a user",
			args: args{
				usr: ent.User{
					Name:     "user",
					Password: "password",
					Admin:    false,
				},
				ctx: security.NewContext(context.Background(), security.Claims{
					UserId:  uuid.NewString(),
					IsAdmin: false,
				}),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "an admin should be able to create a user",
			args: args{
				usr: ent.User{
					Name:     "user",
					Password: "password",
					Admin:    false,
				},
				ctx: security.NewContext(context.Background(), security.Claims{
					UserId:  uuid.NewString(),
					IsAdmin: true,
				}),
			},
			want: &ent.User{
				Name:     "user",
				Password: "password",
				Admin:    false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := User{
				log:    logger,
				client: client,
			}
			got, err := u.Create(tt.args.usr, tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil {
				if got.Name != tt.want.Name {
					t.Errorf("Create() got = %v, want %v", got, tt.want)
				}
				if got.Admin != tt.want.Admin {
					t.Errorf("Create() got = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestUser_Update(t *testing.T) {
	logger := (zaptest.NewLogger(t)).Sugar()
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	_ = client.Schema.Create(context.Background())
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {

		}
	}(client)

	currentUser := client.User.Create().SetName("user").SetPassword("password").SetAdmin(false).SaveX(security.NewContext(context.Background(), security.Claims{
		UserId:  uuid.NewString(),
		IsAdmin: true,
	}))

	type args struct {
		usr ent.User
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    *ent.User
		wantErr bool
	}{
		{
			name: "an unauthenticated user should not be able to update a user",
			args: args{
				usr: ent.User{
					ID:       currentUser.ID,
					Name:     "name",
					Password: "mypassword",
					Admin:    false,
				},
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "an authenticated user with no admin rights should not be able to update a user",
			args: args{
				usr: ent.User{
					ID:       currentUser.ID,
					Name:     "name",
					Password: "mypassword",
					Admin:    false,
				},
				ctx: security.NewContext(context.Background(), security.Claims{
					UserId:  uuid.NewString(),
					IsAdmin: false,
				}),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "an authenticated user with admin rights should be able to update a user",
			args: args{
				usr: ent.User{
					ID:       currentUser.ID,
					Name:     "name",
					Password: "mypassword",
					Admin:    false,
				},
				ctx: security.NewContext(context.Background(), security.Claims{
					UserId:  uuid.NewString(),
					IsAdmin: true,
				}),
			},
			want:    &ent.User{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := User{
				log:    logger,
				client: client,
			}
			got, err := u.Update(tt.args.usr, tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil {
				if got == nil {
					t.Errorf("we should get a user back here but got %p", got)
				}
			}
		})
	}
}

func TestUser_Delete(t *testing.T) {
	logger := (zaptest.NewLogger(t)).Sugar()
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	_ = client.Schema.Create(context.Background())
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {

		}
	}(client)

	adminCtx := security.NewContext(context.Background(), security.Claims{
		UserId:   uuid.NewString(),
		UserName: "admin",
		IsAdmin:  true,
	})

	userInDb := client.User.Create().SetName("user").SetPassword("password").SaveX(adminCtx)

	type args struct {
		id  string
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "an unauthenticated user should not be able to delete a user",
			args: args{
				id:  userInDb.ID.String(),
				ctx: context.Background(),
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "an authenticated user with no admin rights should not be able to delete a user",
			args: args{
				id: userInDb.ID.String(),
				ctx: security.NewContext(context.Background(), security.Claims{
					UserId:   uuid.NewString(),
					UserName: "non-admin",
					IsAdmin:  false,
				}),
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "an admin should not be able to delete a user if it does not exist",
			args: args{
				id:  uuid.NewString(),
				ctx: adminCtx,
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "an admin should be able to delete a user if it exists",
			args: args{
				id:  userInDb.ID.String(),
				ctx: adminCtx,
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := User{
				log:    logger,
				client: client,
			}
			got, err := u.Delete(tt.args.id, tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Delete() got = %v, want %v", got, tt.want)
			}
		})
	}
}
