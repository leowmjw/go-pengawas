// Code generated by ogen, DO NOT EDIT.

package vault

import (
	"fmt"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/validate"
)

func (s Alias) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Format.Set {
			if err := func() error {
				if err := s.Format.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "format",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s AliasFormat) Validate() error {
	switch s {
	case "FPE_ACC_NUM_T_FOUR":
		return nil
	case "FPE_ALPHANUMERIC_ACC_NUM_T_FOUR":
		return nil
	case "FPE_SIX_T_FOUR":
		return nil
	case "FPE_SSN_T_FOUR":
		return nil
	case "FPE_T_FOUR":
		return nil
	case "NUM_LENGTH_PRESERVING":
		return nil
	case "PFPT":
		return nil
	case "RAW_UUID":
		return nil
	case "UUID":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s ApiErrorsResponse) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.Array{
			MinLength:    1,
			MinLengthSet: true,
			MaxLength:    0,
			MaxLengthSet: false,
		}).ValidateLength(len(s.Errors)); err != nil {
			return errors.Wrap(err, "array")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "errors",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s ApiErrorsResponseStatusCode) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.Response.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "Response",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s CreateAliasesCreated) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		var failures []validate.FieldError
		for i, elem := range s.Data {
			if err := func() error {
				if err := elem.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
		}
		if len(failures) > 0 {
			return &validate.Error{Fields: failures}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "data",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s CreateAliasesRequest) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Data == nil {
			return errors.New("nil is invalid value")
		}
		if err := (validate.Array{
			MinLength:    1,
			MinLengthSet: true,
			MaxLength:    20,
			MaxLengthSet: true,
		}).ValidateLength(len(s.Data)); err != nil {
			return errors.Wrap(err, "array")
		}
		var failures []validate.FieldError
		for i, elem := range s.Data {
			if err := func() error {
				if err := elem.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
		}
		if len(failures) > 0 {
			return &validate.Error{Fields: failures}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "data",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s CreateAliasesRequestDataItem) Validate() error {
	switch s.Type {
	case CreateAliasesRequestDataItem0CreateAliasesRequestDataItem:
		if err := s.CreateAliasesRequestDataItem0.Validate(); err != nil {
			return err
		}
		return nil
	case CreateAliasesRequestDataItem1CreateAliasesRequestDataItem:
		if err := s.CreateAliasesRequestDataItem1.Validate(); err != nil {
			return err
		}
		return nil
	default:
		return errors.Errorf("invalid type %q", s.Type)
	}
}

func (s CreateAliasesRequestDataItem0) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.Format.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "format",
			Error: err,
		})
	}
	if err := func() error {
		if s.Storage.Set {
			if err := func() error {
				if err := s.Storage.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "storage",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s CreateAliasesRequestDataItem0Storage) Validate() error {
	switch s {
	case "PERSISTENT":
		return nil
	case "VOLATILE":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s CreateAliasesRequestDataItem1) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.Format.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "format",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s RevealMultipleAliasesOK) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		var failures []validate.FieldError
		for i, elem := range s.Data {
			if err := func() error {
				if err := elem.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
		}
		if len(failures) > 0 {
			return &validate.Error{Fields: failures}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "data",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s RevealedData) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		var failures []validate.FieldError
		for i, elem := range s.Aliases {
			if err := func() error {
				if err := elem.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
		}
		if len(failures) > 0 {
			return &validate.Error{Fields: failures}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "aliases",
			Error: err,
		})
	}
	if err := func() error {
		if s.Storage.Set {
			if err := func() error {
				if err := s.Storage.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "storage",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s RevealedDataStorage) Validate() error {
	switch s {
	case "PERSISTENT":
		return nil
	case "VOLATILE":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s UpdateAliasRequest) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.Data.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "data",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s UpdateAliasRequestData) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Classifiers == nil {
			return errors.New("nil is invalid value")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "classifiers",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}