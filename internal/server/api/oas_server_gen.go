// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// Delete implements delete operation.
	//
	// Delete key. Does nothing if key not exists.
	//
	// DELETE /{namespace}/{key}
	Delete(ctx context.Context, params DeleteParams) error
	// DeleteNamespace implements deleteNamespace operation.
	//
	// Delete namespace and all keys.
	//
	// DELETE /{namespace}
	DeleteNamespace(ctx context.Context, params DeleteNamespaceParams) error
	// Get implements get operation.
	//
	// Get value by key.
	//
	// GET /{namespace}/{key}
	Get(ctx context.Context, params GetParams) (GetRes, error)
	// Keys implements keys operation.
	//
	// List keys in namespace page by page. Iteration may not be consistent if there are updates between
	// pages.
	// At the end of page, operation will return empty list.
	// Iteration order is non-deterministic, but generally tends to be from oldest key to newest key.
	//
	// GET /{namespace}
	Keys(ctx context.Context, params KeysParams) (*Batch, error)
	// Set implements set operation.
	//
	// Create or replace value in namespace.
	//
	// POST /{namespace}/{key}
	Set(ctx context.Context, req *SetReqWithContentType, params SetParams) error
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	sec SecurityHandler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, sec SecurityHandler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		sec:        sec,
		baseServer: s,
	}, nil
}
