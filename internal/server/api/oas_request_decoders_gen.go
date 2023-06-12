// Code generated by ogen, DO NOT EDIT.

package api

import (
	"mime"
	"net/http"

	"github.com/go-faster/errors"
	"go.uber.org/multierr"

	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/validate"
)

func (s *Server) decodeSetRequest(r *http.Request) (
	req *SetReqWithContentType,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = multierr.Append(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ht.MatchContentType("*/*", ct):
		reader := r.Body
		request := SetReq{Data: reader}
		wrapped := SetReqWithContentType{
			ContentType: ct,
			Content:     request,
		}
		return &wrapped, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}