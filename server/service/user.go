package service

import (
	"context"
	"fmt"

	"calcio/ent"
	"calcio/ent/user"
	"calcio/server/security"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type User struct {
	log    *zap.SugaredLogger
	client *ent.Client
}

// UserModule makes the injectable available for FX.
var UserModule = fx.Provide(NewUser)

// NewUser creates a new injectable.
func NewUser(logger *zap.SugaredLogger, client *ent.Client) *User {
	return &User{
		log:    logger,
		client: client,
	}
}

func (u User) Login(name, password string) (*ent.User, error) {
	dummyUserContext := security.NewContext(context.Background(), security.Claims{UserId: uuid.NewString()})
	retrievedUser, err := u.client.User.Query().Where(user.Name(name)).First(dummyUserContext)
	if err != nil {
		return nil, errors.Wrap(err, "user not found for login")
	}

	if !security.CheckPassword(password, retrievedUser.Password) {
		return nil, fmt.Errorf("cannot match the password for user %v", retrievedUser.ID)
	}

	return retrievedUser, nil
}

func (u User) List(ctx context.Context) (ent.Users, error) {
	return u.client.User.Query().Order(ent.Asc(user.FieldName)).All(ctx)
}

func (u User) Create(usr ent.User, ctx context.Context) (*ent.User, error) {
	return u.client.User.Create().SetName(usr.Name).SetPassword(usr.Password).SetAdmin(usr.Admin).Save(ctx)
}

func (u User) CreateDefaultAdmin(password string) (err error) {
	ctx := security.NewContext(context.Background(), security.Claims{
		UserId:  uuid.NewString(),
		IsAdmin: true,
	})

	_, err = u.client.User.Create().SetAdmin(true).SetName("admin").SetPassword(password).Save(ctx)
	return
}

func (u User) Update(usr ent.User, ctx context.Context) (*ent.User, error) {
	current, err := u.client.User.Query().Where(user.ID(usr.ID)).First(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "impossible to retrieve a user with id: %v", usr.ID)
	}

	if usr.Password != "" {
		return current.Update().SetPassword(usr.Password).Save(ctx)
	} else {
		return current, nil
	}
}

func (u User) Delete(id string, ctx context.Context) (int, error) {
	uId, err := uuid.Parse(id)
	if err != nil {
		return 0, errors.Wrapf(err, "impossible to get uuid from %s", id)
	}

	return u.client.User.Delete().Where(user.ID(uId)).Exec(ctx)
}
