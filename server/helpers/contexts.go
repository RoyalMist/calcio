package helpers

import (
	"context"

	"calcio/server/security"
	"github.com/google/uuid"
)

func LoggedInCtx(admin bool) context.Context {
	return security.NewContext(context.Background(), security.Claims{
		UserId:   uuid.NewString(),
		UserName: "user",
		IsAdmin:  admin,
	})
}
