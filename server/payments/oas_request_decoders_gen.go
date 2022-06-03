// Code generated by ogen, DO NOT EDIT.

package payments

import (
	"io"
	"mime"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/multierr"

	"github.com/ogen-go/ogen/validate"
)

func (s *Server) decodeCreateFinancialInstrumentRequest(r *http.Request, span trace.Span) (
	req OptPaymentCardCreate,
	close func() error,
	rerr error,
) {
	var closers []io.Closer
	close = func() error {
		var merr error
		for _, c := range closers {
			merr = multierr.Append(merr, c.Close())
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
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, close, nil
		}

		var request OptPaymentCardCreate
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, close, err
		}

		if written == 0 {
			return req, close, nil
		}

		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			request.Reset()
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "decode \"application/json\"")
		}

		return request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeCreateGatewayRequest(r *http.Request, span trace.Span) (
	req OptGatewayCreate,
	close func() error,
	rerr error,
) {
	var closers []io.Closer
	close = func() error {
		var merr error
		for _, c := range closers {
			merr = multierr.Append(merr, c.Close())
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
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, close, nil
		}

		var request OptGatewayCreate
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, close, err
		}

		if written == 0 {
			return req, close, nil
		}

		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			request.Reset()
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "decode \"application/json\"")
		}

		return request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeCreateOrderRequest(r *http.Request, span trace.Span) (
	req OptOrderCreate,
	close func() error,
	rerr error,
) {
	var closers []io.Closer
	close = func() error {
		var merr error
		for _, c := range closers {
			merr = multierr.Append(merr, c.Close())
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
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, close, nil
		}

		var request OptOrderCreate
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, close, err
		}

		if written == 0 {
			return req, close, nil
		}

		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			request.Reset()
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "decode \"application/json\"")
		}
		if err := func() error {
			if request.Set {
				if err := func() error {
					if err := request.Value.Validate(); err != nil {
						return err
					}
					return nil
				}(); err != nil {
					return err
				}
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "validate")
		}

		return request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeCreateRuleRequest(r *http.Request, span trace.Span) (
	req OptRuleCreate,
	close func() error,
	rerr error,
) {
	var closers []io.Closer
	close = func() error {
		var merr error
		for _, c := range closers {
			merr = multierr.Append(merr, c.Close())
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
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, close, nil
		}

		var request OptRuleCreate
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, close, err
		}

		if written == 0 {
			return req, close, nil
		}

		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			request.Reset()
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "decode \"application/json\"")
		}
		if err := func() error {
			if request.Set {
				if err := func() error {
					if err := request.Value.Validate(); err != nil {
						return err
					}
					return nil
				}(); err != nil {
					return err
				}
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "validate")
		}

		return request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeCreateThreeDSAuthenticationRequest(r *http.Request, span trace.Span) (
	req OptThreeDSAuthenticationCreate,
	close func() error,
	rerr error,
) {
	var closers []io.Closer
	close = func() error {
		var merr error
		for _, c := range closers {
			merr = multierr.Append(merr, c.Close())
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
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, close, nil
		}

		var request OptThreeDSAuthenticationCreate
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, close, err
		}

		if written == 0 {
			return req, close, nil
		}

		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			request.Reset()
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "decode \"application/json\"")
		}
		if err := func() error {
			if request.Set {
				if err := func() error {
					if err := request.Value.Validate(); err != nil {
						return err
					}
					return nil
				}(); err != nil {
					return err
				}
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "validate")
		}

		return request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeCreateThreeDSProviderRequest(r *http.Request, span trace.Span) (
	req OptThreeDSProviderCreate,
	close func() error,
	rerr error,
) {
	var closers []io.Closer
	close = func() error {
		var merr error
		for _, c := range closers {
			merr = multierr.Append(merr, c.Close())
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
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, close, nil
		}

		var request OptThreeDSProviderCreate
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, close, err
		}

		if written == 0 {
			return req, close, nil
		}

		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			request.Reset()
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "decode \"application/json\"")
		}

		return request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeCreateTransferRequest(r *http.Request, span trace.Span) (
	req *TransferCreate,
	close func() error,
	rerr error,
) {
	var closers []io.Closer
	close = func() error {
		var merr error
		for _, c := range closers {
			merr = multierr.Append(merr, c.Close())
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
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, close, nil
		}

		var request *TransferCreate
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, close, err
		}

		if written == 0 {
			return req, close, nil
		}

		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			request = nil
			var elem TransferCreate
			if err := elem.Decode(d); err != nil {
				return err
			}
			request = &elem
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "decode \"application/json\"")
		}

		return request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeCreateVerificationRequest(r *http.Request, span trace.Span) (
	req *Verification,
	close func() error,
	rerr error,
) {
	var closers []io.Closer
	close = func() error {
		var merr error
		for _, c := range closers {
			merr = multierr.Append(merr, c.Close())
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
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, close, nil
		}

		var request *Verification
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, close, err
		}

		if written == 0 {
			return req, close, nil
		}

		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			request = nil
			var elem Verification
			if err := elem.Decode(d); err != nil {
				return err
			}
			request = &elem
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "decode \"application/json\"")
		}

		return request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeFinalizeTransferRequest(r *http.Request, span trace.Span) (
	req OptTransferUpdate,
	close func() error,
	rerr error,
) {
	var closers []io.Closer
	close = func() error {
		var merr error
		for _, c := range closers {
			merr = multierr.Append(merr, c.Close())
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
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, close, nil
		}

		var request OptTransferUpdate
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, close, err
		}

		if written == 0 {
			return req, close, nil
		}

		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			request.Reset()
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "decode \"application/json\"")
		}
		if err := func() error {
			if request.Set {
				if err := func() error {
					if err := request.Value.Validate(); err != nil {
						return err
					}
					return nil
				}(); err != nil {
					return err
				}
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "validate")
		}

		return request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeReverseTransferByIdRequest(r *http.Request, span trace.Span) (
	req OptReversal,
	close func() error,
	rerr error,
) {
	var closers []io.Closer
	close = func() error {
		var merr error
		for _, c := range closers {
			merr = multierr.Append(merr, c.Close())
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
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, close, nil
		}

		var request OptReversal
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, close, err
		}

		if written == 0 {
			return req, close, nil
		}

		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			request.Reset()
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "decode \"application/json\"")
		}
		if err := func() error {
			if request.Set {
				if err := func() error {
					if err := request.Value.Validate(); err != nil {
						return err
					}
					return nil
				}(); err != nil {
					return err
				}
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "validate")
		}

		return request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeThreeDSAuthenticationChallengeRequest(r *http.Request, span trace.Span) (
	req *ThreeDSAuthenticationChallengeReq,
	close func() error,
	rerr error,
) {
	var closers []io.Closer
	close = func() error {
		var merr error
		for _, c := range closers {
			merr = multierr.Append(merr, c.Close())
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
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, close, nil
		}

		var request *ThreeDSAuthenticationChallengeReq
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, close, err
		}

		if written == 0 {
			return req, close, nil
		}

		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			request = nil
			var elem ThreeDSAuthenticationChallengeReq
			if err := elem.Decode(d); err != nil {
				return err
			}
			request = &elem
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "decode \"application/json\"")
		}

		return request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeThreeDSAuthenticationDeviceFingerprintRequest(r *http.Request, span trace.Span) (
	req *ThreeDSAuthenticationDeviceFingerprintReq,
	close func() error,
	rerr error,
) {
	var closers []io.Closer
	close = func() error {
		var merr error
		for _, c := range closers {
			merr = multierr.Append(merr, c.Close())
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
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, close, nil
		}

		var request *ThreeDSAuthenticationDeviceFingerprintReq
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, close, err
		}

		if written == 0 {
			return req, close, nil
		}

		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			request = nil
			var elem ThreeDSAuthenticationDeviceFingerprintReq
			if err := elem.Decode(d); err != nil {
				return err
			}
			request = &elem
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "decode \"application/json\"")
		}

		return request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeUpdateGatewayByIdRequest(r *http.Request, span trace.Span) (
	req OptGatewayUpdate,
	close func() error,
	rerr error,
) {
	var closers []io.Closer
	close = func() error {
		var merr error
		for _, c := range closers {
			merr = multierr.Append(merr, c.Close())
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
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, close, nil
		}

		var request OptGatewayUpdate
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, close, err
		}

		if written == 0 {
			return req, close, nil
		}

		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			request.Reset()
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "decode \"application/json\"")
		}

		return request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeUpdateRuleByIdRequest(r *http.Request, span trace.Span) (
	req OptRuleUpdate,
	close func() error,
	rerr error,
) {
	var closers []io.Closer
	close = func() error {
		var merr error
		for _, c := range closers {
			merr = multierr.Append(merr, c.Close())
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
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, close, nil
		}

		var request OptRuleUpdate
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, close, err
		}

		if written == 0 {
			return req, close, nil
		}

		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			request.Reset()
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "decode \"application/json\"")
		}
		if err := func() error {
			if request.Set {
				if err := func() error {
					if err := request.Value.Validate(); err != nil {
						return err
					}
					return nil
				}(); err != nil {
					return err
				}
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "validate")
		}

		return request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}
