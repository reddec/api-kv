package client

import (
	"context"
)

// HeaderToken is header base authorization with pre-shared token.
type HeaderToken string

func (ta HeaderToken) HeaderAuth(_ context.Context, _ string) (HeaderAuth, error) {
	return HeaderAuth{
		APIKey: string(ta),
	}, nil
}
