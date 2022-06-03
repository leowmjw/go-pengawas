// Code generated by ogen, DO NOT EDIT.

package vault

import (
	"net/http"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/uri"
)

type DeleteAliasParams struct {
	// ID of card to fetch.
	Alias string
}

func decodeDeleteAliasParams(args [1]string, r *http.Request) (params DeleteAliasParams, _ error) {
	// Decode path: alias.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "alias",
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

				params.Alias = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: alias: not specified")
		}
	}
	return params, nil
}

type RevealAliasParams struct {
	// ID of card to fetch.
	Alias string
}

func decodeRevealAliasParams(args [1]string, r *http.Request) (params RevealAliasParams, _ error) {
	// Decode path: alias.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "alias",
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

				params.Alias = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: alias: not specified")
		}
	}
	return params, nil
}

type RevealMultipleAliasesParams struct {
	// Comma-separated list of aliases to reveal.
	Q string
}

func decodeRevealMultipleAliasesParams(args [0]string, r *http.Request) (params RevealMultipleAliasesParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: q.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "q",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Q = c
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "query: q: parse")
			}
		} else {
			return params, errors.Wrap(err, "query")
		}
	}
	return params, nil
}

type UpdateAliasParams struct {
	// ID of card to fetch.
	Alias string
}

func decodeUpdateAliasParams(args [1]string, r *http.Request) (params UpdateAliasParams, _ error) {
	// Decode path: alias.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "alias",
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

				params.Alias = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: alias: not specified")
		}
	}
	return params, nil
}