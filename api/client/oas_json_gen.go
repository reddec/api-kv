// Code generated by ogen, DO NOT EDIT.

package client

import (
	"math/bits"
	"strconv"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/validate"
)

// Encode implements json.Marshaler.
func (s *Batch) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Batch) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("keys")
		e.ArrStart()
		for _, elem := range s.Keys {
			e.Str(elem)
		}
		e.ArrEnd()
	}
	{
		e.FieldStart("cursor")
		e.Int64(s.Cursor)
	}
}

var jsonFieldsNameOfBatch = [2]string{
	0: "keys",
	1: "cursor",
}

// Decode decodes Batch from json.
func (s *Batch) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Batch to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "keys":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				s.Keys = make([]string, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem string
					v, err := d.Str()
					elem = string(v)
					if err != nil {
						return err
					}
					s.Keys = append(s.Keys, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"keys\"")
			}
		case "cursor":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Int64()
				s.Cursor = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"cursor\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Batch")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfBatch) {
					name = jsonFieldsNameOfBatch[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Batch) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Batch) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}