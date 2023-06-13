package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"time"
)

func Dial(url string, namespace string, token string) (*KV, error) {
	c, err := NewClient(url, HeaderToken(token))
	if err != nil {
		return nil, fmt.Errorf("create client: %w", err)
	}
	return New(c, namespace), nil
}

func New(apiClient *Client, namespace string) *KV {
	return &KV{
		namespace: namespace,
		client:    apiClient,
	}
}

// KV is simplified wrapper around client.
type KV struct {
	namespace string
	client    *Client
}

// Get item from API-KV. Returns nil if key not found.
func (app *KV) Get(ctx context.Context, key string) ([]byte, error) {
	storedRef, err := app.client.Get(ctx, GetParams{
		Namespace: app.namespace,
		Key:       key,
	})

	if err != nil {
		return nil, fmt.Errorf("get key %s: %w", key, err)
	}
	if v, ok := storedRef.(*GetOKHeaders); ok {
		return io.ReadAll(v.Response.Data)
	}
	return nil, nil
}

// Set new or replace old value by key.
func (app *KV) Set(ctx context.Context, key string, data []byte) error {
	return app.SetExpire(ctx, key, data, 0)
}

// SetExpire create new or replace old value by key with desired TTL.
// TTL <= 0 means disabled expiration.
func (app *KV) SetExpire(ctx context.Context, key string, data []byte, ttl time.Duration) error {
	var params = SetParams{
		Namespace: app.namespace,
		Key:       key,
	}

	var req = &SetReqWithContentType{
		ContentType: "application/octet-stream",
		Content: SetReq{
			Data: bytes.NewReader(data),
		},
	}
	if ttl > 0 {
		params.TTL.SetTo(ttl.Seconds())
	}
	return app.client.Set(ctx, req, params)
}

// Keys in namespace.
func (app *KV) Keys() *Iterator {
	return &Iterator{
		app: app,
	}
}

// Delete single key.
func (app *KV) Delete(ctx context.Context, key string) error {
	return app.client.Delete(ctx, DeleteParams{
		Namespace: app.namespace,
		Key:       key,
	})
}

// Destroy the whole namespace.
func (app *KV) Destroy(ctx context.Context) error {
	return app.client.DeleteNamespace(ctx, DeleteNamespaceParams{Namespace: app.namespace})
}

type Iterator struct {
	app    *KV
	cursor int64
	err    error
	page   []string
	done   bool
}

func (it *Iterator) Next(ctx context.Context) bool {
	if it.err != nil || it.done {
		return false
	}
	res, err := it.app.client.Keys(ctx, KeysParams{
		Namespace: it.app.namespace,
		Cursor:    NewOptInt64(it.cursor),
	})
	if err != nil {
		it.err = err
		return false
	}
	it.cursor = res.Cursor
	it.page = res.Keys
	it.done = len(res.Keys) == 0
	return it.done
}

func (it *Iterator) Keys() []string {
	return it.page
}

func (it *Iterator) Error() error {
	return it.err
}
