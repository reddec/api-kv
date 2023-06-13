package impl

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/reddec/api-kv/internal/dbo"
	"github.com/reddec/api-kv/internal/server/api"
)

type Server struct {
	BatchSize int64
	DB        *dbo.Queries
}

func (srv *Server) Delete(ctx context.Context, params api.DeleteParams) error {
	return srv.DB.Delete(ctx, dbo.DeleteParams{
		Namespace: params.Namespace,
		Key:       params.Key,
	})
}

func (srv *Server) DeleteNamespace(ctx context.Context, params api.DeleteNamespaceParams) error {
	return srv.DB.DeleteNamespace(ctx, params.Namespace)
}

func (srv *Server) Get(ctx context.Context, params api.GetParams) (api.GetRes, error) {
	value, err := srv.DB.Get(ctx, dbo.GetParams{
		Namespace: params.Namespace,
		Key:       params.Key,
	})
	if errors.Is(err, sql.ErrNoRows) {
		return &api.GetNotFound{}, nil
	}
	if err != nil {
		return nil, err
	}

	res := api.GetOKHeaders{
		XContentType: value.ContentType,
		Response: api.GetOK{
			Data: bytes.NewReader(value.Value),
		},
	}
	if value.ExpireAt.Valid {
		res.XTTL.SetTo(time.Until(time.UnixMilli(value.ExpireAt.Int64)).Seconds())
	}

	return &res, nil
}

func (srv *Server) Keys(ctx context.Context, params api.KeysParams) (*api.Batch, error) {
	if !params.Cursor.IsSet() {
		return mapKeys(srv.DB.ListStart(ctx, dbo.ListStartParams{
			Namespace: params.Namespace,
			Limit:     srv.BatchSize,
		}))
	}
	return mapKeys(srv.DB.ListNext(ctx, dbo.ListNextParams{
		Namespace: params.Namespace,
		ID:        params.Cursor.Value,
		Limit:     srv.BatchSize,
	}))
}

func (srv *Server) Set(ctx context.Context, req *api.SetReqWithContentType, params api.SetParams) error {
	data, err := io.ReadAll(req.Content.Data)
	if err != nil {
		return err
	}
	contentType := req.ContentType
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	var args = dbo.PutParams{
		Namespace:   params.Namespace,
		Key:         params.Key,
		Value:       data,
		ContentType: req.ContentType,
	}
	if params.TTL.IsSet() {
		args.ExpireAt = sql.NullInt64{
			Int64: time.Now().UnixMilli() + int64(1000*params.TTL.Value),
			Valid: true,
		}
	}

	return srv.DB.Put(ctx, args)
}

func mapKeys(keys []dbo.ValueKey, err error) (*api.Batch, error) {
	if err != nil {
		return nil, err
	}
	var ans = make([]string, 0, len(keys))
	var id int64
	for _, k := range keys {
		ans = append(ans, k.Key)
		id = k.ID
	}

	return &api.Batch{
		Keys:   ans,
		Cursor: id,
	}, nil
}

func FixContentType(handler http.Handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctm := &contentTypeMapper{real: writer}
		handler.ServeHTTP(ctm, request)
	}
}

type contentTypeMapper struct {
	real http.ResponseWriter
	sent bool
}

func (ctm *contentTypeMapper) Header() http.Header {
	return ctm.real.Header()
}

func (ctm *contentTypeMapper) Write(data []byte) (int, error) {
	ctm.setHeader()
	return ctm.real.Write(data)
}

func (ctm *contentTypeMapper) WriteHeader(statusCode int) {
	ctm.setHeader()
	ctm.real.WriteHeader(statusCode)
}

func (ctm *contentTypeMapper) setHeader() {
	if ctm.sent {
		return
	}
	ctm.sent = true
	ct := ctm.real.Header().Get("X-Content-Type")
	if ct == "" {
		return
	}
	ctm.real.Header().Set("Content-Type", ct)
}
