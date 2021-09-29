package security

import (
	"context"

	"calcio/ent"
	"github.com/google/uuid"
)

func CreateAdmin(client *ent.Client) (err error) {
	ctx := NewContext(context.Background(), Claims{
		UserId:  uuid.NewString(),
		IsAdmin: true,
	})

	_, err = client.User.Create().SetAdmin(true).SetName("admin").SetPassword("admin123").Save(ctx)
	return
}
