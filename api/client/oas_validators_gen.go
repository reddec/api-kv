// Code generated by ogen, DO NOT EDIT.

package client

import (
	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/validate"
)

func (s *Batch) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Keys == nil {
			return errors.New("nil is invalid value")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "keys",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *GetOKHeaders) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.XTTL.Set {
			if err := func() error {
				if err := (validate.Float{}).Validate(float64(s.XTTL.Value)); err != nil {
					return errors.Wrap(err, "float")
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "XTTL",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
