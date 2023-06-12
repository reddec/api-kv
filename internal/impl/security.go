package impl

import (
	"context"
	"crypto/subtle"
	"errors"

	"github.com/reddec/api-kv/internal/server/api"
)

var ErrInvalidToken = errors.New("invalid token")

type StaticToken string

func (st StaticToken) HandleHeaderAuth(ctx context.Context, _ string, t api.HeaderAuth) (context.Context, error) {
	if subtle.ConstantTimeCompare([]byte(st), []byte(t.APIKey)) != 1 {
		return nil, ErrInvalidToken
	}
	return ctx, nil
}
