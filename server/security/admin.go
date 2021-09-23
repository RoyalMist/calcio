package security

import (
	"context"

	"calcio/ent"
)

func CreateAdmin(client *ent.Client) (err error) {
	ctx := newContext(context.Background(), Claims{
		UserId:  "admin",
		IsAdmin: true,
	})

	_, err = client.User.Create().SetAdmin(true).SetName("Admin").SetPassword("p@sSw0rd").Save(ctx)
	return
}
