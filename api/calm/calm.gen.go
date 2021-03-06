// Package calm provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.10.1 DO NOT EDIT.
package calm

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
)

const (
	OAuth2Scopes = "OAuth2.Scopes"
)

// Defines values for CardCapabilities.
const (
	CardCapabilitiesACCOUNTUPDATER CardCapabilities = "ACCOUNT_UPDATER"
)

// Defines values for CardEnrollApiErrorCode.
const (
	CardEnrollApiErrorCodeAccountUpdaterNotSupported CardEnrollApiErrorCode = "account-updater-not-supported"

	CardEnrollApiErrorCodeCardBrandNotSupported CardEnrollApiErrorCode = "card-brand-not-supported"

	CardEnrollApiErrorCodeCardNotFound CardEnrollApiErrorCode = "card-not-found"

	CardEnrollApiErrorCodeInternalServerError CardEnrollApiErrorCode = "internal-server-error"

	CardEnrollApiErrorCodeInvalidPayload CardEnrollApiErrorCode = "invalid-payload"

	CardEnrollApiErrorCodeMerchantNotFound CardEnrollApiErrorCode = "merchant-not-found"

	CardEnrollApiErrorCodeUnsupportedMediaType CardEnrollApiErrorCode = "unsupported-media-type"

	CardEnrollApiErrorCodeValidationFailed CardEnrollApiErrorCode = "validation-failed"
)

// Defines values for CardObjectFailureCode.
const (
	CardObjectFailureCodeAccountUpdaterNotSupported CardObjectFailureCode = "account-updater-not-supported"

	CardObjectFailureCodeMerchantNotFound CardObjectFailureCode = "merchant-not-found"
)

// Defines values for CardObjectState.
const (
	CardObjectStateClosed CardObjectState = "closed"

	CardObjectStateEnrolled CardObjectState = "enrolled"

	CardObjectStateEnrolling CardObjectState = "enrolling"

	CardObjectStateFailed CardObjectState = "failed"
)

// ApiErrorsResponse defines model for ApiErrorsResponse.
type ApiErrorsResponse struct {
	Errors []Error `json:"errors"`

	// A unique identifier of the failed request.
	TraceId *string `json:"trace_id,omitempty"`
}

// Customer's billing.
type BillingAddress struct {
	Address1 *string `json:"address1,omitempty"`
	Address2 *string `json:"address2,omitempty"`
	City     *string `json:"city,omitempty"`
	Company  *string `json:"company,omitempty"`

	// Country code in ISO 3166-1 alpha-3
	Country *string `json:"country,omitempty"`
	Name    *string `json:"name,omitempty"`

	// Telephone number in E.164
	Phone      *string `json:"phone,omitempty"`
	PostalCode *string `json:"postal_code,omitempty"`

	// Principal subdivision in ISO 3166-2
	Region *string `json:"region,omitempty"`
}

// Payment card object.
type Card struct {
	// Customer's billing.
	BillingAddress *BillingAddress `json:"billing_address,omitempty"`

	// CALM capabilities to enable. If blank all products are enrolled.
	Capabilities *[]CardCapabilities `json:"capabilities,omitempty"`

	// Card's expiration month.
	ExpMonth int `json:"exp_month"`

	// Card's expiration year.
	ExpYear int `json:"exp_year"`

	// Card owner's full name.
	Name string `json:"name"`

	// Customer's card number.
	Number string `json:"number"`
}

// CardCapabilities defines model for Card.Capabilities.
type CardCapabilities string

// CardData defines model for CardData.
type CardData struct {
	Data CardObject `json:"data"`
}

// CardEnrollApiError defines model for CardEnrollApiError.
type CardEnrollApiError struct {
	// Embedded struct due to allOf(#/components/schemas/Error)
	Error `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	// Application-specific error code:
	//   * `card-brand-not-supported` - Specified card brand is not supported, only Mastercard cards are supported
	//   * `account-updater-not-supported` - Account Updater is not implemented for your card issuer
	//   * `card-not-found` - Card is not enrolled in CALM
	//   * `merchant-not-found` - No merchant was registered for organization
	//   * `internal-server-error` - Something went wrong
	//   * `invalid-payload` - Invalid request payload
	//   * `validation-failed` - Request validation failed
	//   * `unsupported-media-type` - Request media type is not supported
	Code *CardEnrollApiErrorCode `json:"code,omitempty"`
}

// Application-specific error code:
//   * `card-brand-not-supported` - Specified card brand is not supported, only Mastercard cards are supported
//   * `account-updater-not-supported` - Account Updater is not implemented for your card issuer
//   * `card-not-found` - Card is not enrolled in CALM
//   * `merchant-not-found` - No merchant was registered for organization
//   * `internal-server-error` - Something went wrong
//   * `invalid-payload` - Invalid request payload
//   * `validation-failed` - Request validation failed
//   * `unsupported-media-type` - Request media type is not supported
type CardEnrollApiErrorCode string

// CardEnrollApiErrorsResponse defines model for CardEnrollApiErrorsResponse.
type CardEnrollApiErrorsResponse struct {
	Errors []CardEnrollApiError `json:"errors"`

	// A unique identifier of the failed request.
	TraceId *string `json:"trace_id,omitempty"`
}

// CardObject defines model for CardObject.
type CardObject struct {
	// Embedded struct due to allOf(#/components/schemas/Card)
	Card `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	// Creation time, in UTC.
	CreatedAt time.Time `json:"created_at"`

	// Application-specific failure code which can only occur with failed state:
	//   * `account-updater-not-supported` - Account Updater is not implemented for your card issuer
	//   * `merchant-not-found` - No merchant was registered for organization
	FailureCode *CardObjectFailureCode `json:"failure_code,omitempty"`

	// Unique reference identifier for this card
	Id string `json:"id"`

	// Current state of a card being processed:
	//   * `enrolling` - card is in enrolling state and is not yet available for updates
	//   * `enrolled` - card was successfully enrolled and is available for Push and Pull updates
	//   * `failed` - card enrollment failed, see `failure_code` for details
	//   * `closed` - card was closed, but can still be retrieved
	State CardObjectState `json:"state"`

	// Last time card was updated, in UTC.
	UpdatedAt time.Time `json:"updated_at"`
}

// Application-specific failure code which can only occur with failed state:
//   * `account-updater-not-supported` - Account Updater is not implemented for your card issuer
//   * `merchant-not-found` - No merchant was registered for organization
type CardObjectFailureCode string

// Current state of a card being processed:
//   * `enrolling` - card is in enrolling state and is not yet available for updates
//   * `enrolled` - card was successfully enrolled and is available for Push and Pull updates
//   * `failed` - card enrollment failed, see `failure_code` for details
//   * `closed` - card was closed, but can still be retrieved
type CardObjectState string

// An error object
type Error struct {
	// Explanation of what exactly went wrong.
	Detail *string `json:"detail,omitempty"`
}

// CardEnrollResponse defines model for CardEnrollResponse.
type CardEnrollResponse CardData

// Payment card object.
type CardEnrollRequest Card

// EnrollCardJSONRequestBody defines body for EnrollCard for application/json ContentType.
type EnrollCardJSONRequestBody CardEnrollRequest

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// EnrollCard request with any body
	EnrollCardWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	EnrollCard(ctx context.Context, body EnrollCardJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetCard request
	GetCard(ctx context.Context, cardId string, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) EnrollCardWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewEnrollCardRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) EnrollCard(ctx context.Context, body EnrollCardJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewEnrollCardRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetCard(ctx context.Context, cardId string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetCardRequest(c.Server, cardId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewEnrollCardRequest calls the generic EnrollCard builder with application/json body
func NewEnrollCardRequest(server string, body EnrollCardJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewEnrollCardRequestWithBody(server, "application/json", bodyReader)
}

// NewEnrollCardRequestWithBody generates requests for EnrollCard with any type of body
func NewEnrollCardRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/cards")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewGetCardRequest generates requests for GetCard
func NewGetCardRequest(server string, cardId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "cardId", runtime.ParamLocationPath, cardId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/cards/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// EnrollCard request with any body
	EnrollCardWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*EnrollCardResponse, error)

	EnrollCardWithResponse(ctx context.Context, body EnrollCardJSONRequestBody, reqEditors ...RequestEditorFn) (*EnrollCardResponse, error)

	// GetCard request
	GetCardWithResponse(ctx context.Context, cardId string, reqEditors ...RequestEditorFn) (*GetCardResponse, error)
}

type EnrollCardResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON202      *CardData
	JSON400      *CardEnrollApiErrorsResponse
	JSON422      *CardEnrollApiErrorsResponse
	JSONDefault  *ApiErrorsResponse
}

// Status returns HTTPResponse.Status
func (r EnrollCardResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r EnrollCardResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetCardResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *CardObject
	JSON404      *CardEnrollApiErrorsResponse
}

// Status returns HTTPResponse.Status
func (r GetCardResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetCardResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// EnrollCardWithBodyWithResponse request with arbitrary body returning *EnrollCardResponse
func (c *ClientWithResponses) EnrollCardWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*EnrollCardResponse, error) {
	rsp, err := c.EnrollCardWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseEnrollCardResponse(rsp)
}

func (c *ClientWithResponses) EnrollCardWithResponse(ctx context.Context, body EnrollCardJSONRequestBody, reqEditors ...RequestEditorFn) (*EnrollCardResponse, error) {
	rsp, err := c.EnrollCard(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseEnrollCardResponse(rsp)
}

// GetCardWithResponse request returning *GetCardResponse
func (c *ClientWithResponses) GetCardWithResponse(ctx context.Context, cardId string, reqEditors ...RequestEditorFn) (*GetCardResponse, error) {
	rsp, err := c.GetCard(ctx, cardId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetCardResponse(rsp)
}

// ParseEnrollCardResponse parses an HTTP response from a EnrollCardWithResponse call
func ParseEnrollCardResponse(rsp *http.Response) (*EnrollCardResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &EnrollCardResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 202:
		var dest CardData
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON202 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest CardEnrollApiErrorsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 422:
		var dest CardEnrollApiErrorsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON422 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest ApiErrorsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseGetCardResponse parses an HTTP response from a GetCardWithResponse call
func ParseGetCardResponse(rsp *http.Response) (*GetCardResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetCardResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest CardObject
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest CardEnrollApiErrorsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Enroll a new card in CALM
	// (POST /cards)
	EnrollCard(w http.ResponseWriter, r *http.Request)
	// Get information on a specific enrolled card
	// (GET /cards/{cardId})
	GetCard(w http.ResponseWriter, r *http.Request, cardId string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// EnrollCard operation middleware
func (siw *ServerInterfaceWrapper) EnrollCard(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, OAuth2Scopes, []string{"cards:write"})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.EnrollCard(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetCard operation middleware
func (siw *ServerInterfaceWrapper) GetCard(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "cardId" -------------
	var cardId string

	err = runtime.BindStyledParameter("simple", false, "cardId", chi.URLParam(r, "cardId"), &cardId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "cardId", Err: err})
		return
	}

	ctx = context.WithValue(ctx, OAuth2Scopes, []string{"cards:read"})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetCard(w, r, cardId)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/cards", wrapper.EnrollCard)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/cards/{cardId}", wrapper.GetCard)
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"ujE3mWDrs2r74+3jvwIAAP//paarLjgjAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
