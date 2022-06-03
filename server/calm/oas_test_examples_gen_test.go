// Code generated by ogen, DO NOT EDIT.

package calm

import (
	"fmt"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/validate"

	std "encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestApiErrorsResponse_EncodeDecode(t *testing.T) {
	var typ ApiErrorsResponse
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 ApiErrorsResponse
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}

func TestApiErrorsResponse_Examples(t *testing.T) {

	for i, tc := range []struct {
		Input string
	}{
		{Input: "{\"errors\":[{\"code\":\"internal-server-error\",\"detail\":\"Something went wrong.\"}],\"trace_id\":\"ab28bdaad3d6e136e6eb9ce31cca9855\"}"},
	} {
		tc := tc
		t.Run(fmt.Sprintf("Test%d", i+1), func(t *testing.T) {
			var typ ApiErrorsResponse

			if err := typ.Decode(jx.DecodeStr(tc.Input)); err != nil {
				var validateErr *validate.Error
				if errors.As(err, &validateErr) {
					t.Skipf("Validation error: %v", validateErr)
					return
				}
				require.NoErrorf(t, err, "Input: %s", tc.Input)
			}

			e := jx.Encoder{}
			typ.Encode(&e)
			require.True(t, std.Valid(e.Bytes()), "Encoded: %s", e.Bytes())

			var typ2 ApiErrorsResponse
			require.NoError(t, typ2.Decode(jx.DecodeBytes(e.Bytes())))
		})
	}
}
func TestBillingAddress_EncodeDecode(t *testing.T) {
	var typ BillingAddress
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 BillingAddress
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestCard_EncodeDecode(t *testing.T) {
	var typ Card
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Card
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}

func TestCard_Examples(t *testing.T) {

	for i, tc := range []struct {
		Input string
	}{
		{Input: "{\"billing_address\":{\"address1\":\"555 Unblock Us St\",\"address2\":\"M13 9PL\",\"city\":\"New York\",\"company\":\"John Doe Company\",\"country\":\"US\",\"name\":\"John Doe\",\"phone\":\"+14842634673\",\"postal_code\":12301,\"region\":\"NY\"},\"capabilities\":[\"ACCOUNT_UPDATER\"],\"exp_month\":7,\"exp_year\":24,\"name\":\"John Doe\",\"number\":5573491171027312}"},
		{Input: "{\"capabilities\":[\"ACCOUNT_UPDATER\"],\"exp_month\":7,\"exp_year\":24,\"name\":\"John Doe\",\"number\":5573491171027312}"},
	} {
		tc := tc
		t.Run(fmt.Sprintf("Test%d", i+1), func(t *testing.T) {
			var typ Card

			if err := typ.Decode(jx.DecodeStr(tc.Input)); err != nil {
				var validateErr *validate.Error
				if errors.As(err, &validateErr) {
					t.Skipf("Validation error: %v", validateErr)
					return
				}
				require.NoErrorf(t, err, "Input: %s", tc.Input)
			}

			e := jx.Encoder{}
			typ.Encode(&e)
			require.True(t, std.Valid(e.Bytes()), "Encoded: %s", e.Bytes())

			var typ2 Card
			require.NoError(t, typ2.Decode(jx.DecodeBytes(e.Bytes())))
		})
	}
}
func TestCardCapabilitiesItem_EncodeDecode(t *testing.T) {
	var typ CardCapabilitiesItem
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 CardCapabilitiesItem
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestCardData_EncodeDecode(t *testing.T) {
	var typ CardData
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 CardData
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestCardEnrollApiError_EncodeDecode(t *testing.T) {
	var typ CardEnrollApiError
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 CardEnrollApiError
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestCardEnrollApiErrorsResponse_EncodeDecode(t *testing.T) {
	var typ CardEnrollApiErrorsResponse
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 CardEnrollApiErrorsResponse
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}

func TestCardEnrollApiErrorsResponse_Examples(t *testing.T) {

	for i, tc := range []struct {
		Input string
	}{
		{Input: "{\"errors\":[{\"code\":\"card-brand-not-supported\",\"detail\":\"Provided card brand is not supported.\"}],\"trace_id\":\"ab28bdaad3d6e136e6eb9ce31cca9855\"}"},
		{Input: "{\"errors\":[{\"code\":\"card-not-found\",\"detail\":\"Card not found.\"}],\"trace_id\":\"ab28bdaad3d6e136e6eb9ce31cca9855\"}"},
	} {
		tc := tc
		t.Run(fmt.Sprintf("Test%d", i+1), func(t *testing.T) {
			var typ CardEnrollApiErrorsResponse

			if err := typ.Decode(jx.DecodeStr(tc.Input)); err != nil {
				var validateErr *validate.Error
				if errors.As(err, &validateErr) {
					t.Skipf("Validation error: %v", validateErr)
					return
				}
				require.NoErrorf(t, err, "Input: %s", tc.Input)
			}

			e := jx.Encoder{}
			typ.Encode(&e)
			require.True(t, std.Valid(e.Bytes()), "Encoded: %s", e.Bytes())

			var typ2 CardEnrollApiErrorsResponse
			require.NoError(t, typ2.Decode(jx.DecodeBytes(e.Bytes())))
		})
	}
}
func TestCardObject_EncodeDecode(t *testing.T) {
	var typ CardObject
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 CardObject
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestCardObjectCapabilitiesItem_EncodeDecode(t *testing.T) {
	var typ CardObjectCapabilitiesItem
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 CardObjectCapabilitiesItem
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestEnrollCardApplicationJSONBadRequest_EncodeDecode(t *testing.T) {
	var typ EnrollCardApplicationJSONBadRequest
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 EnrollCardApplicationJSONBadRequest
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestEnrollCardApplicationJSONUnprocessableEntity_EncodeDecode(t *testing.T) {
	var typ EnrollCardApplicationJSONUnprocessableEntity
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 EnrollCardApplicationJSONUnprocessableEntity
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestError_EncodeDecode(t *testing.T) {
	var typ Error
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Error
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}