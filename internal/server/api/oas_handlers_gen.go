// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/http"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.19.0"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/otelogen"
)

// handleDeleteRequest handles delete operation.
//
// Delete key. Does nothing if key not exists.
//
// DELETE /{namespace}/{key}
func (s *Server) handleDeleteRequest(args [2]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("delete"),
		semconv.HTTPMethodKey.String("DELETE"),
		semconv.HTTPRouteKey.String("/{namespace}/{key}"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "Delete",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "Delete",
			ID:   "delete",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityHeaderAuth(ctx, "Delete", r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "HeaderAuth",
					Err:              err,
				}
				recordError("Security:HeaderAuth", err)
				s.cfg.ErrorHandler(ctx, w, r, err)
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			recordError("Security", err)
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
	}
	params, err := decodeDeleteParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *DeleteNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "Delete",
			OperationID:   "delete",
			Body:          nil,
			Params: middleware.Parameters{
				{
					Name: "namespace",
					In:   "path",
				}: params.Namespace,
				{
					Name: "key",
					In:   "path",
				}: params.Key,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = DeleteParams
			Response = *DeleteNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackDeleteParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.Delete(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.Delete(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeDeleteResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleDeleteNamespaceRequest handles deleteNamespace operation.
//
// Delete namespace and all keys.
//
// DELETE /{namespace}
func (s *Server) handleDeleteNamespaceRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("deleteNamespace"),
		semconv.HTTPMethodKey.String("DELETE"),
		semconv.HTTPRouteKey.String("/{namespace}"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "DeleteNamespace",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "DeleteNamespace",
			ID:   "deleteNamespace",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityHeaderAuth(ctx, "DeleteNamespace", r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "HeaderAuth",
					Err:              err,
				}
				recordError("Security:HeaderAuth", err)
				s.cfg.ErrorHandler(ctx, w, r, err)
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			recordError("Security", err)
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
	}
	params, err := decodeDeleteNamespaceParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *DeleteNamespaceNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "DeleteNamespace",
			OperationID:   "deleteNamespace",
			Body:          nil,
			Params: middleware.Parameters{
				{
					Name: "namespace",
					In:   "path",
				}: params.Namespace,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = DeleteNamespaceParams
			Response = *DeleteNamespaceNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackDeleteNamespaceParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.DeleteNamespace(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.DeleteNamespace(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeDeleteNamespaceResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleGetRequest handles get operation.
//
// Get value by key.
//
// GET /{namespace}/{key}
func (s *Server) handleGetRequest(args [2]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("get"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/{namespace}/{key}"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "Get",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "Get",
			ID:   "get",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityHeaderAuth(ctx, "Get", r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "HeaderAuth",
					Err:              err,
				}
				recordError("Security:HeaderAuth", err)
				s.cfg.ErrorHandler(ctx, w, r, err)
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			recordError("Security", err)
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
	}
	params, err := decodeGetParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response GetRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "Get",
			OperationID:   "get",
			Body:          nil,
			Params: middleware.Parameters{
				{
					Name: "namespace",
					In:   "path",
				}: params.Namespace,
				{
					Name: "key",
					In:   "path",
				}: params.Key,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetParams
			Response = GetRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.Get(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.Get(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleKeysRequest handles keys operation.
//
// List keys in namespace page by page. Iteration may not be consistent if there are updates between
// pages.
// At the end of page, operation will return empty list.
// Iteration order is non-deterministic, but generally tends to be from oldest key to newest key.
//
// GET /{namespace}
func (s *Server) handleKeysRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("keys"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/{namespace}"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "Keys",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "Keys",
			ID:   "keys",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityHeaderAuth(ctx, "Keys", r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "HeaderAuth",
					Err:              err,
				}
				recordError("Security:HeaderAuth", err)
				s.cfg.ErrorHandler(ctx, w, r, err)
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			recordError("Security", err)
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
	}
	params, err := decodeKeysParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *Batch
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "Keys",
			OperationID:   "keys",
			Body:          nil,
			Params: middleware.Parameters{
				{
					Name: "namespace",
					In:   "path",
				}: params.Namespace,
				{
					Name: "cursor",
					In:   "query",
				}: params.Cursor,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = KeysParams
			Response = *Batch
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackKeysParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.Keys(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.Keys(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeKeysResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleSetRequest handles set operation.
//
// Create or replace value in namespace.
//
// POST /{namespace}/{key}
func (s *Server) handleSetRequest(args [2]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("set"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/{namespace}/{key}"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "Set",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "Set",
			ID:   "set",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityHeaderAuth(ctx, "Set", r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "HeaderAuth",
					Err:              err,
				}
				recordError("Security:HeaderAuth", err)
				s.cfg.ErrorHandler(ctx, w, r, err)
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			recordError("Security", err)
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
	}
	params, err := decodeSetParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	request, close, err := s.decodeSetRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *SetNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "Set",
			OperationID:   "set",
			Body:          request,
			Params: middleware.Parameters{
				{
					Name: "namespace",
					In:   "path",
				}: params.Namespace,
				{
					Name: "key",
					In:   "path",
				}: params.Key,
				{
					Name: "ttl",
					In:   "query",
				}: params.TTL,
			},
			Raw: r,
		}

		type (
			Request  = *SetReqWithContentType
			Params   = SetParams
			Response = *SetNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackSetParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.Set(ctx, request, params)
				return response, err
			},
		)
	} else {
		err = s.h.Set(ctx, request, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeSetResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}