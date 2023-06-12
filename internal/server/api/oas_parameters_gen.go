// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"
	"net/url"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// DeleteParams is parameters of delete operation.
type DeleteParams struct {
	// Key-Value namespace.
	Namespace string
	// Key name.
	Key string
}

func unpackDeleteParams(packed middleware.Parameters) (params DeleteParams) {
	{
		key := middleware.ParameterKey{
			Name: "namespace",
			In:   "path",
		}
		params.Namespace = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "key",
			In:   "path",
		}
		params.Key = packed[key].(string)
	}
	return params
}

func decodeDeleteParams(args [2]string, argsEscaped bool, r *http.Request) (params DeleteParams, _ error) {
	// Decode path: namespace.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "namespace",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Namespace = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "namespace",
			In:   "path",
			Err:  err,
		}
	}
	// Decode path: key.
	if err := func() error {
		param := args[1]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[1])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "key",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Key = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "key",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// DeleteNamespaceParams is parameters of deleteNamespace operation.
type DeleteNamespaceParams struct {
	// Key-Value namespace.
	Namespace string
}

func unpackDeleteNamespaceParams(packed middleware.Parameters) (params DeleteNamespaceParams) {
	{
		key := middleware.ParameterKey{
			Name: "namespace",
			In:   "path",
		}
		params.Namespace = packed[key].(string)
	}
	return params
}

func decodeDeleteNamespaceParams(args [1]string, argsEscaped bool, r *http.Request) (params DeleteNamespaceParams, _ error) {
	// Decode path: namespace.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "namespace",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Namespace = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "namespace",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// GetParams is parameters of get operation.
type GetParams struct {
	// Key-Value namespace.
	Namespace string
	// Key name.
	Key string
}

func unpackGetParams(packed middleware.Parameters) (params GetParams) {
	{
		key := middleware.ParameterKey{
			Name: "namespace",
			In:   "path",
		}
		params.Namespace = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "key",
			In:   "path",
		}
		params.Key = packed[key].(string)
	}
	return params
}

func decodeGetParams(args [2]string, argsEscaped bool, r *http.Request) (params GetParams, _ error) {
	// Decode path: namespace.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "namespace",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Namespace = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "namespace",
			In:   "path",
			Err:  err,
		}
	}
	// Decode path: key.
	if err := func() error {
		param := args[1]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[1])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "key",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Key = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "key",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// KeysParams is parameters of keys operation.
type KeysParams struct {
	// Key-Value namespace.
	Namespace string
	// Cursor of the previous page, should be passed in from the previous request or absent for new
	// request.
	Cursor OptInt64
}

func unpackKeysParams(packed middleware.Parameters) (params KeysParams) {
	{
		key := middleware.ParameterKey{
			Name: "namespace",
			In:   "path",
		}
		params.Namespace = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "cursor",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Cursor = v.(OptInt64)
		}
	}
	return params
}

func decodeKeysParams(args [1]string, argsEscaped bool, r *http.Request) (params KeysParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: namespace.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "namespace",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Namespace = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "namespace",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: cursor.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "cursor",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotCursorVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotCursorVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Cursor.SetTo(paramsDotCursorVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "cursor",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// SetParams is parameters of set operation.
type SetParams struct {
	// Key-Value namespace.
	Namespace string
	// Key name.
	Key string
	// Time-to-live for the key in seconds (floating).
	TTL OptFloat64
}

func unpackSetParams(packed middleware.Parameters) (params SetParams) {
	{
		key := middleware.ParameterKey{
			Name: "namespace",
			In:   "path",
		}
		params.Namespace = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "key",
			In:   "path",
		}
		params.Key = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "ttl",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.TTL = v.(OptFloat64)
		}
	}
	return params
}

func decodeSetParams(args [2]string, argsEscaped bool, r *http.Request) (params SetParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: namespace.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "namespace",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Namespace = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "namespace",
			In:   "path",
			Err:  err,
		}
	}
	// Decode path: key.
	if err := func() error {
		param := args[1]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[1])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "key",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Key = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "key",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: ttl.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "ttl",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotTTLVal float64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToFloat64(val)
					if err != nil {
						return err
					}

					paramsDotTTLVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.TTL.SetTo(paramsDotTTLVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if params.TTL.Set {
					if err := func() error {
						if err := (validate.Float{
							MinSet:        true,
							Min:           0,
							MaxSet:        false,
							Max:           0,
							MinExclusive:  false,
							MaxExclusive:  false,
							MultipleOfSet: false,
							MultipleOf:    nil,
						}).Validate(float64(params.TTL.Value)); err != nil {
							return errors.Wrap(err, "float")
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "ttl",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}