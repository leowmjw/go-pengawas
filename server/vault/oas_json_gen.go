// Code generated by ogen, DO NOT EDIT.

package vault

import (
	"math/bits"
	"strconv"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/validate"
)

// Encode implements json.Marshaler.
func (s Alias) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s Alias) encodeFields(e *jx.Encoder) {
	{
		if s.Alias.Set {
			e.FieldStart("alias")
			s.Alias.Encode(e)
		}
	}
	{
		if s.Format.Set {
			e.FieldStart("format")
			s.Format.Encode(e)
		}
	}
}

var jsonFieldsNameOfAlias = [2]string{
	0: "alias",
	1: "format",
}

// Decode decodes Alias from json.
func (s *Alias) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Alias to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "alias":
			if err := func() error {
				s.Alias.Reset()
				if err := s.Alias.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"alias\"")
			}
		case "format":
			if err := func() error {
				s.Format.Reset()
				if err := s.Format.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"format\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Alias")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s Alias) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Alias) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes AliasFormat as json.
func (s AliasFormat) Encode(e *jx.Encoder) {
	e.Str(string(s))
}

// Decode decodes AliasFormat from json.
func (s *AliasFormat) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode AliasFormat to nil")
	}
	v, err := d.StrBytes()
	if err != nil {
		return err
	}
	// Try to use constant string.
	switch AliasFormat(v) {
	case AliasFormatFPEACCNUMTFOUR:
		*s = AliasFormatFPEACCNUMTFOUR
	case AliasFormatFPEALPHANUMERICACCNUMTFOUR:
		*s = AliasFormatFPEALPHANUMERICACCNUMTFOUR
	case AliasFormatFPESIXTFOUR:
		*s = AliasFormatFPESIXTFOUR
	case AliasFormatFPESSNTFOUR:
		*s = AliasFormatFPESSNTFOUR
	case AliasFormatFPETFOUR:
		*s = AliasFormatFPETFOUR
	case AliasFormatNUMLENGTHPRESERVING:
		*s = AliasFormatNUMLENGTHPRESERVING
	case AliasFormatPFPT:
		*s = AliasFormatPFPT
	case AliasFormatRAWUUID:
		*s = AliasFormatRAWUUID
	case AliasFormatUUID:
		*s = AliasFormatUUID
	default:
		*s = AliasFormat(v)
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s AliasFormat) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *AliasFormat) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s ApiError) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s ApiError) encodeFields(e *jx.Encoder) {
	{
		if s.Detail.Set {
			e.FieldStart("detail")
			s.Detail.Encode(e)
		}
	}
	{
		if s.Href.Set {
			e.FieldStart("href")
			s.Href.Encode(e)
		}
	}
	{
		if s.Status.Set {
			e.FieldStart("status")
			s.Status.Encode(e)
		}
	}
	{
		if s.Title.Set {
			e.FieldStart("title")
			s.Title.Encode(e)
		}
	}
}

var jsonFieldsNameOfApiError = [4]string{
	0: "detail",
	1: "href",
	2: "status",
	3: "title",
}

// Decode decodes ApiError from json.
func (s *ApiError) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ApiError to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "detail":
			if err := func() error {
				s.Detail.Reset()
				if err := s.Detail.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"detail\"")
			}
		case "href":
			if err := func() error {
				s.Href.Reset()
				if err := s.Href.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"href\"")
			}
		case "status":
			if err := func() error {
				s.Status.Reset()
				if err := s.Status.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"status\"")
			}
		case "title":
			if err := func() error {
				s.Title.Reset()
				if err := s.Title.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"title\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ApiError")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s ApiError) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ApiError) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s ApiErrorsResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s ApiErrorsResponse) encodeFields(e *jx.Encoder) {
	{
		if s.Errors != nil {
			e.FieldStart("errors")
			e.ArrStart()
			for _, elem := range s.Errors {
				elem.Encode(e)
			}
			e.ArrEnd()
		}
	}
}

var jsonFieldsNameOfApiErrorsResponse = [1]string{
	0: "errors",
}

// Decode decodes ApiErrorsResponse from json.
func (s *ApiErrorsResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ApiErrorsResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "errors":
			if err := func() error {
				s.Errors = make([]ApiError, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem ApiError
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Errors = append(s.Errors, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"errors\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ApiErrorsResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s ApiErrorsResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ApiErrorsResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s CreateAliasesCreated) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s CreateAliasesCreated) encodeFields(e *jx.Encoder) {
	{
		if s.Data != nil {
			e.FieldStart("data")
			e.ArrStart()
			for _, elem := range s.Data {
				elem.Encode(e)
			}
			e.ArrEnd()
		}
	}
}

var jsonFieldsNameOfCreateAliasesCreated = [1]string{
	0: "data",
}

// Decode decodes CreateAliasesCreated from json.
func (s *CreateAliasesCreated) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode CreateAliasesCreated to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "data":
			if err := func() error {
				s.Data = make([]RevealedData, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem RevealedData
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Data = append(s.Data, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"data\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode CreateAliasesCreated")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s CreateAliasesCreated) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *CreateAliasesCreated) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s CreateAliasesRequest) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s CreateAliasesRequest) encodeFields(e *jx.Encoder) {
	{

		e.FieldStart("data")
		e.ArrStart()
		for _, elem := range s.Data {
			elem.Encode(e)
		}
		e.ArrEnd()
	}
}

var jsonFieldsNameOfCreateAliasesRequest = [1]string{
	0: "data",
}

// Decode decodes CreateAliasesRequest from json.
func (s *CreateAliasesRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode CreateAliasesRequest to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "data":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				s.Data = make([]CreateAliasesRequestDataItem, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem CreateAliasesRequestDataItem
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Data = append(s.Data, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"data\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode CreateAliasesRequest")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
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
				if fieldIdx < len(jsonFieldsNameOfCreateAliasesRequest) {
					name = jsonFieldsNameOfCreateAliasesRequest[fieldIdx]
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
func (s CreateAliasesRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *CreateAliasesRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes CreateAliasesRequestDataItem as json.
func (s CreateAliasesRequestDataItem) Encode(e *jx.Encoder) {
	switch s.Type {
	case CreateAliasesRequestDataItem0CreateAliasesRequestDataItem:
		s.CreateAliasesRequestDataItem0.Encode(e)
	case CreateAliasesRequestDataItem1CreateAliasesRequestDataItem:
		s.CreateAliasesRequestDataItem1.Encode(e)
	}
}

// Decode decodes CreateAliasesRequestDataItem from json.
func (s *CreateAliasesRequestDataItem) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode CreateAliasesRequestDataItem to nil")
	}
	// Sum type fields.
	if typ := d.Next(); typ != jx.Object {
		return errors.Errorf("unexpected json type %q", typ)
	}

	var found bool
	if err := d.Capture(func(d *jx.Decoder) error {
		return d.ObjBytes(func(d *jx.Decoder, key []byte) error {
			switch string(key) {
			case "classifiers":
				match := CreateAliasesRequestDataItem0CreateAliasesRequestDataItem
				if found && s.Type != match {
					s.Type = ""
					return errors.Errorf("multiple oneOf matches: (%v, %v)", s.Type, match)
				}
				found = true
				s.Type = match
			case "storage":
				match := CreateAliasesRequestDataItem0CreateAliasesRequestDataItem
				if found && s.Type != match {
					s.Type = ""
					return errors.Errorf("multiple oneOf matches: (%v, %v)", s.Type, match)
				}
				found = true
				s.Type = match
			case "value":
				match := CreateAliasesRequestDataItem0CreateAliasesRequestDataItem
				if found && s.Type != match {
					s.Type = ""
					return errors.Errorf("multiple oneOf matches: (%v, %v)", s.Type, match)
				}
				found = true
				s.Type = match
			case "alias":
				match := CreateAliasesRequestDataItem1CreateAliasesRequestDataItem
				if found && s.Type != match {
					s.Type = ""
					return errors.Errorf("multiple oneOf matches: (%v, %v)", s.Type, match)
				}
				found = true
				s.Type = match
			}
			return d.Skip()
		})
	}); err != nil {
		return errors.Wrap(err, "capture")
	}
	if !found {
		return errors.New("unable to detect sum type variant")
	}
	switch s.Type {
	case CreateAliasesRequestDataItem0CreateAliasesRequestDataItem:
		if err := s.CreateAliasesRequestDataItem0.Decode(d); err != nil {
			return err
		}
	case CreateAliasesRequestDataItem1CreateAliasesRequestDataItem:
		if err := s.CreateAliasesRequestDataItem1.Decode(d); err != nil {
			return err
		}
	default:
		return errors.Errorf("inferred invalid type: %s", s.Type)
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s CreateAliasesRequestDataItem) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *CreateAliasesRequestDataItem) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s CreateAliasesRequestDataItem0) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s CreateAliasesRequestDataItem0) encodeFields(e *jx.Encoder) {
	{
		if s.Classifiers != nil {
			e.FieldStart("classifiers")
			e.ArrStart()
			for _, elem := range s.Classifiers {
				e.Str(elem)
			}
			e.ArrEnd()
		}
	}
	{

		e.FieldStart("format")
		s.Format.Encode(e)
	}
	{
		if s.Storage.Set {
			e.FieldStart("storage")
			s.Storage.Encode(e)
		}
	}
	{

		e.FieldStart("value")
		e.Str(s.Value)
	}
}

var jsonFieldsNameOfCreateAliasesRequestDataItem0 = [4]string{
	0: "classifiers",
	1: "format",
	2: "storage",
	3: "value",
}

// Decode decodes CreateAliasesRequestDataItem0 from json.
func (s *CreateAliasesRequestDataItem0) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode CreateAliasesRequestDataItem0 to nil")
	}
	var requiredBitSet [1]uint8
	s.setDefaults()

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "classifiers":
			if err := func() error {
				s.Classifiers = make([]string, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem string
					v, err := d.Str()
					elem = string(v)
					if err != nil {
						return err
					}
					s.Classifiers = append(s.Classifiers, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"classifiers\"")
			}
		case "format":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				if err := s.Format.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"format\"")
			}
		case "storage":
			if err := func() error {
				s.Storage.Reset()
				if err := s.Storage.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"storage\"")
			}
		case "value":
			requiredBitSet[0] |= 1 << 3
			if err := func() error {
				v, err := d.Str()
				s.Value = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"value\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode CreateAliasesRequestDataItem0")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00001010,
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
				if fieldIdx < len(jsonFieldsNameOfCreateAliasesRequestDataItem0) {
					name = jsonFieldsNameOfCreateAliasesRequestDataItem0[fieldIdx]
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
func (s CreateAliasesRequestDataItem0) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *CreateAliasesRequestDataItem0) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes CreateAliasesRequestDataItem0Storage as json.
func (s CreateAliasesRequestDataItem0Storage) Encode(e *jx.Encoder) {
	e.Str(string(s))
}

// Decode decodes CreateAliasesRequestDataItem0Storage from json.
func (s *CreateAliasesRequestDataItem0Storage) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode CreateAliasesRequestDataItem0Storage to nil")
	}
	v, err := d.StrBytes()
	if err != nil {
		return err
	}
	// Try to use constant string.
	switch CreateAliasesRequestDataItem0Storage(v) {
	case CreateAliasesRequestDataItem0StoragePERSISTENT:
		*s = CreateAliasesRequestDataItem0StoragePERSISTENT
	case CreateAliasesRequestDataItem0StorageVOLATILE:
		*s = CreateAliasesRequestDataItem0StorageVOLATILE
	default:
		*s = CreateAliasesRequestDataItem0Storage(v)
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s CreateAliasesRequestDataItem0Storage) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *CreateAliasesRequestDataItem0Storage) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s CreateAliasesRequestDataItem1) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s CreateAliasesRequestDataItem1) encodeFields(e *jx.Encoder) {
	{

		e.FieldStart("alias")
		e.Str(s.Alias)
	}
	{

		e.FieldStart("format")
		s.Format.Encode(e)
	}
}

var jsonFieldsNameOfCreateAliasesRequestDataItem1 = [2]string{
	0: "alias",
	1: "format",
}

// Decode decodes CreateAliasesRequestDataItem1 from json.
func (s *CreateAliasesRequestDataItem1) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode CreateAliasesRequestDataItem1 to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "alias":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.Alias = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"alias\"")
			}
		case "format":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				if err := s.Format.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"format\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode CreateAliasesRequestDataItem1")
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
				if fieldIdx < len(jsonFieldsNameOfCreateAliasesRequestDataItem1) {
					name = jsonFieldsNameOfCreateAliasesRequestDataItem1[fieldIdx]
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
func (s CreateAliasesRequestDataItem1) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *CreateAliasesRequestDataItem1) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes AliasFormat as json.
func (o OptAliasFormat) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes AliasFormat from json.
func (o *OptAliasFormat) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptAliasFormat to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptAliasFormat) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptAliasFormat) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes CreateAliasesRequest as json.
func (o OptCreateAliasesRequest) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes CreateAliasesRequest from json.
func (o *OptCreateAliasesRequest) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptCreateAliasesRequest to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptCreateAliasesRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptCreateAliasesRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes CreateAliasesRequestDataItem0Storage as json.
func (o OptCreateAliasesRequestDataItem0Storage) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes CreateAliasesRequestDataItem0Storage from json.
func (o *OptCreateAliasesRequestDataItem0Storage) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptCreateAliasesRequestDataItem0Storage to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptCreateAliasesRequestDataItem0Storage) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptCreateAliasesRequestDataItem0Storage) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes time.Time as json.
func (o OptDateTime) Encode(e *jx.Encoder, format func(*jx.Encoder, time.Time)) {
	if !o.Set {
		return
	}
	format(e, o.Value)
}

// Decode decodes time.Time from json.
func (o *OptDateTime) Decode(d *jx.Decoder, format func(*jx.Decoder) (time.Time, error)) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptDateTime to nil")
	}
	o.Set = true
	v, err := format(d)
	if err != nil {
		return err
	}
	o.Value = v
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptDateTime) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e, json.EncodeDateTime)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptDateTime) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d, json.DecodeDateTime)
}

// Encode encodes int as json.
func (o OptInt) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Int(int(o.Value))
}

// Decode decodes int from json.
func (o *OptInt) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptInt to nil")
	}
	o.Set = true
	v, err := d.Int()
	if err != nil {
		return err
	}
	o.Value = int(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptInt) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptInt) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes RevealedDataStorage as json.
func (o OptRevealedDataStorage) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes RevealedDataStorage from json.
func (o *OptRevealedDataStorage) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptRevealedDataStorage to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptRevealedDataStorage) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptRevealedDataStorage) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes string as json.
func (o OptString) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes string from json.
func (o *OptString) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptString to nil")
	}
	o.Set = true
	v, err := d.Str()
	if err != nil {
		return err
	}
	o.Value = string(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptString) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptString) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes UpdateAliasRequest as json.
func (o OptUpdateAliasRequest) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes UpdateAliasRequest from json.
func (o *OptUpdateAliasRequest) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptUpdateAliasRequest to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptUpdateAliasRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptUpdateAliasRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s RevealMultipleAliasesOK) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s RevealMultipleAliasesOK) encodeFields(e *jx.Encoder) {
	{
		if s.Data != nil {
			e.FieldStart("data")
			e.ArrStart()
			for _, elem := range s.Data {
				elem.Encode(e)
			}
			e.ArrEnd()
		}
	}
}

var jsonFieldsNameOfRevealMultipleAliasesOK = [1]string{
	0: "data",
}

// Decode decodes RevealMultipleAliasesOK from json.
func (s *RevealMultipleAliasesOK) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode RevealMultipleAliasesOK to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "data":
			if err := func() error {
				s.Data = make([]RevealedData, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem RevealedData
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Data = append(s.Data, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"data\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode RevealMultipleAliasesOK")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s RevealMultipleAliasesOK) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *RevealMultipleAliasesOK) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s RevealedData) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s RevealedData) encodeFields(e *jx.Encoder) {
	{
		if s.Aliases != nil {
			e.FieldStart("aliases")
			e.ArrStart()
			for _, elem := range s.Aliases {
				elem.Encode(e)
			}
			e.ArrEnd()
		}
	}
	{
		if s.Classifiers != nil {
			e.FieldStart("classifiers")
			e.ArrStart()
			for _, elem := range s.Classifiers {
				e.Str(elem)
			}
			e.ArrEnd()
		}
	}
	{
		if s.CreatedAt.Set {
			e.FieldStart("created_at")
			s.CreatedAt.Encode(e, json.EncodeDateTime)
		}
	}
	{
		if s.Storage.Set {
			e.FieldStart("storage")
			s.Storage.Encode(e)
		}
	}
	{
		if s.Value.Set {
			e.FieldStart("value")
			s.Value.Encode(e)
		}
	}
}

var jsonFieldsNameOfRevealedData = [5]string{
	0: "aliases",
	1: "classifiers",
	2: "created_at",
	3: "storage",
	4: "value",
}

// Decode decodes RevealedData from json.
func (s *RevealedData) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode RevealedData to nil")
	}
	s.setDefaults()

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "aliases":
			if err := func() error {
				s.Aliases = make([]Alias, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem Alias
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Aliases = append(s.Aliases, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"aliases\"")
			}
		case "classifiers":
			if err := func() error {
				s.Classifiers = make([]string, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem string
					v, err := d.Str()
					elem = string(v)
					if err != nil {
						return err
					}
					s.Classifiers = append(s.Classifiers, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"classifiers\"")
			}
		case "created_at":
			if err := func() error {
				s.CreatedAt.Reset()
				if err := s.CreatedAt.Decode(d, json.DecodeDateTime); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"created_at\"")
			}
		case "storage":
			if err := func() error {
				s.Storage.Reset()
				if err := s.Storage.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"storage\"")
			}
		case "value":
			if err := func() error {
				s.Value.Reset()
				if err := s.Value.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"value\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode RevealedData")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s RevealedData) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *RevealedData) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes RevealedDataStorage as json.
func (s RevealedDataStorage) Encode(e *jx.Encoder) {
	e.Str(string(s))
}

// Decode decodes RevealedDataStorage from json.
func (s *RevealedDataStorage) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode RevealedDataStorage to nil")
	}
	v, err := d.StrBytes()
	if err != nil {
		return err
	}
	// Try to use constant string.
	switch RevealedDataStorage(v) {
	case RevealedDataStoragePERSISTENT:
		*s = RevealedDataStoragePERSISTENT
	case RevealedDataStorageVOLATILE:
		*s = RevealedDataStorageVOLATILE
	default:
		*s = RevealedDataStorage(v)
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s RevealedDataStorage) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *RevealedDataStorage) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s UpdateAliasRequest) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s UpdateAliasRequest) encodeFields(e *jx.Encoder) {
	{

		e.FieldStart("data")
		s.Data.Encode(e)
	}
}

var jsonFieldsNameOfUpdateAliasRequest = [1]string{
	0: "data",
}

// Decode decodes UpdateAliasRequest from json.
func (s *UpdateAliasRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UpdateAliasRequest to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "data":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				if err := s.Data.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"data\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode UpdateAliasRequest")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
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
				if fieldIdx < len(jsonFieldsNameOfUpdateAliasRequest) {
					name = jsonFieldsNameOfUpdateAliasRequest[fieldIdx]
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
func (s UpdateAliasRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *UpdateAliasRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s UpdateAliasRequestData) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s UpdateAliasRequestData) encodeFields(e *jx.Encoder) {
	{

		e.FieldStart("classifiers")
		e.ArrStart()
		for _, elem := range s.Classifiers {
			e.Str(elem)
		}
		e.ArrEnd()
	}
}

var jsonFieldsNameOfUpdateAliasRequestData = [1]string{
	0: "classifiers",
}

// Decode decodes UpdateAliasRequestData from json.
func (s *UpdateAliasRequestData) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UpdateAliasRequestData to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "classifiers":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				s.Classifiers = make([]string, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem string
					v, err := d.Str()
					elem = string(v)
					if err != nil {
						return err
					}
					s.Classifiers = append(s.Classifiers, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"classifiers\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode UpdateAliasRequestData")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
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
				if fieldIdx < len(jsonFieldsNameOfUpdateAliasRequestData) {
					name = jsonFieldsNameOfUpdateAliasRequestData[fieldIdx]
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
func (s UpdateAliasRequestData) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *UpdateAliasRequestData) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}