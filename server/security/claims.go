package security

import (
	"context"
	"fmt"

	"github.com/vk-rv/pvx"
)

type Claims struct {
	pvx.RegisteredClaims
	UserId  string `json:"user_id"`
	IsAdmin bool   `json:"is_admin"`
}

func (c Claims) Valid() error {
	if err := c.RegisteredClaims.Valid(); err != nil {
		return err
	}

	if len(c.UserId) < 1 {
		return fmt.Errorf("invalid userId %s", c.UserId)
	}

	return nil
}

type ctxKey struct{}

// FromContext returns the Viewer stored in a context.
func FromContext(ctx context.Context) Claims {
	v, _ := ctx.Value(ctxKey{}).(Claims)
	return v
}

// NewContext returns a copy of parent context with the given Viewer attached with it.
func newContext(parent context.Context, c Claims) context.Context {
	return context.WithValue(parent, ctxKey{}, c)
}
