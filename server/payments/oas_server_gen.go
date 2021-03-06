// Code generated by ogen, DO NOT EDIT.

package payments

import (
	"context"

	"go.opentelemetry.io/otel/metric/instrument/syncint64"

	"github.com/ogen-go/ogen/otelogen"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// CreateFinancialInstrument implements CreateFinancialInstrument operation.
	//
	// Create a financial instrument for later use for sending or receiving money.
	//
	// POST /financial_instruments
	CreateFinancialInstrument(ctx context.Context, req OptPaymentCardCreate) (CreateFinancialInstrumentRes, error)
	// CreateGateway implements CreateGateway operation.
	//
	// Create a gateway with PSP's credentials to point transfers to based on certain criteria.
	//
	// POST /gateways
	CreateGateway(ctx context.Context, req OptGatewayCreate) (CreateGatewayRes, error)
	// CreateOrder implements CreateOrder operation.
	//
	// Place an order and return session id for it.
	//
	// POST /orders
	CreateOrder(ctx context.Context, req OptOrderCreate) (CreateOrderRes, error)
	// CreateRule implements CreateRule operation.
	//
	// Create a rule by which the API will determine to which gateway to point the transfers to based on
	// whether the rule conditions are met.
	//
	// POST /rules
	CreateRule(ctx context.Context, req OptRuleCreate) (CreateRuleRes, error)
	// CreateThreeDSAuthentication implements CreateThreeDSAuthentication operation.
	//
	// Create a 3DS authentication.
	// Mainly to be invoked by VGS checkout.js, not meant to be called by user directly.
	//
	// POST /3ds_authentications
	CreateThreeDSAuthentication(ctx context.Context, req OptThreeDSAuthenticationCreate) (CreateThreeDSAuthenticationRes, error)
	// CreateThreeDSProvider implements CreateThreeDSProvider operation.
	//
	// Create a 3DS provider with the credentials and configuration.
	// Currently only one 3DS provider creation is allowed. Please contact VGS support if you need to
	// replace the old 3DS provider.
	//
	// POST /3ds_providers
	CreateThreeDSProvider(ctx context.Context, req OptThreeDSProviderCreate) (CreateThreeDSProviderRes, error)
	// CreateTransfer implements CreateTransfer operation.
	//
	// Creates a transfer record in one of the ways:
	// 1. Capture the funds instantly
	// 2. Create an authorization to be captured manually later.
	// The authorization places the funds on hold with the customer's bank. This can be useful in
	// situations where you need to make a sale, but won't be able to ship merchandise for several days.
	// You can authorize the transfer to ensure funds availability, then capture the transfer to obtain
	// the funds upon shipment.
	// 3. Creates a transfer from order information.
	//
	// POST /transfers
	CreateTransfer(ctx context.Context, req *TransferCreate) (CreateTransferRes, error)
	// CreateVerification implements CreateVerification operation.
	//
	// This request verifies that the customer's card is valid without actually charging an amount on the
	// card. For example, you may want to verify a credit card before adding it to a customer profile for
	// future billing transfers.
	// The exact verification method varies from psp to psp, but usually this is either a $0 charge, or a
	// $1 authorization with immediate cancellation.
	//
	// POST /verifications
	CreateVerification(ctx context.Context, req *Verification) (CreateVerificationRes, error)
	// DeleteFinancialInstrumentById implements DeleteFinancialInstrumentById operation.
	//
	// Delete a particular financial instrument by id.
	//
	// DELETE /financial_instruments/{financial_instrument_id}
	DeleteFinancialInstrumentById(ctx context.Context, params DeleteFinancialInstrumentByIdParams) (DeleteFinancialInstrumentByIdRes, error)
	// DeleteGateway implements DeleteGateway operation.
	//
	// DELETE /gateways/{gateway_id}
	DeleteGateway(ctx context.Context, params DeleteGatewayParams) (DeleteGatewayRes, error)
	// DeleteRule implements DeleteRule operation.
	//
	// DELETE /rules/{rule_id}
	DeleteRule(ctx context.Context, params DeleteRuleParams) (DeleteRuleRes, error)
	// FinalizeTransfer implements FinalizeTransfer operation.
	//
	// Finalize the transfer by moving its state from `pending` to `successful` (captured) or `canceled`
	// via the `action` parameter.
	// Only transfers with `debit` type and `pending` state is allowed.
	// When `capture` is passed in the `action` parameter, the funds, previously held with the customer's
	// bank,
	// will be transferred from the customer to the gateway's account owner or merchant.
	// It will also update the transfer's `state` from `pending` to `successful` if the capture is
	// successful and
	// return the same transfer with `200` status. The `amount_captured` will be updated accordingly.
	// When `cancel` is passed in the `action` parameter, the funds, previously held with the customer's
	// bank,
	// will be released. Some cards may not support a hard release and the funds will be made available
	// after
	// a period of seven days from the original authorization.
	// It will create a new transfer with `reversal` type, return it with `201` status.
	// The `amount_reversed` of original transfer will be updated accordingly.
	//
	// PUT /transfers/{transfer_id}
	FinalizeTransfer(ctx context.Context, req OptTransferUpdate, params FinalizeTransferParams) (FinalizeTransferRes, error)
	// GetFinancialInstrumentById implements GetFinancialInstrumentById operation.
	//
	// Get a particular financial instrument by id.
	//
	// GET /financial_instruments/{financial_instrument_id}
	GetFinancialInstrumentById(ctx context.Context, params GetFinancialInstrumentByIdParams) (GetFinancialInstrumentByIdRes, error)
	// GetFinancialInstruments implements GetFinancialInstruments operation.
	//
	// List financial instrument records filtered by the query condition if provided.
	//
	// GET /financial_instruments
	GetFinancialInstruments(ctx context.Context, params GetFinancialInstrumentsParams) (GetFinancialInstrumentsRes, error)
	// GetGatewayById implements GetGatewayById operation.
	//
	// Get a particular gateway by id.
	//
	// GET /gateways/{gateway_id}
	GetGatewayById(ctx context.Context, params GetGatewayByIdParams) (GetGatewayByIdRes, error)
	// GetGateways implements GetGateways operation.
	//
	// List gateway records filtered by the query condition if provided.
	//
	// GET /gateways
	GetGateways(ctx context.Context, params GetGatewaysParams) (GetGatewaysRes, error)
	// GetRuleById implements GetRuleById operation.
	//
	// Get a particular rule by id.
	//
	// GET /rules/{rule_id}
	GetRuleById(ctx context.Context, params GetRuleByIdParams) (GetRuleByIdRes, error)
	// GetRules implements GetRules operation.
	//
	// List rule records filtered by the query condition if provided.
	//
	// GET /rules
	GetRules(ctx context.Context, params GetRulesParams) (GetRulesRes, error)
	// GetTransferById implements GetTransferById operation.
	//
	// Get a particular transfer by id.
	//
	// GET /transfers/{transfer_id}
	GetTransferById(ctx context.Context, params GetTransferByIdParams) (GetTransferByIdRes, error)
	// GetTransfers implements GetTransfers operation.
	//
	// List transfer records filtered by the query condition if provided.
	//
	// GET /transfers
	GetTransfers(ctx context.Context, params GetTransfersParams) (GetTransfersRes, error)
	// OrdersOrderIDGet implements  operation.
	//
	// Get a particular order by id.
	//
	// GET /orders/{order_id}
	OrdersOrderIDGet(ctx context.Context, params OrdersOrderIDGetParams) (OrdersOrderIDGetRes, error)
	// R3dsAuthentications3dsAuthenticationIDChallengesPost implements  operation.
	//
	// Finish the challenge step for the 3DS authentication session.
	// Mainly to be invoked by VGS checkout.js, not meant to be called by user directly.
	//
	// POST /3ds_authentications/{3ds_authentication_id}/challenges
	R3dsAuthentications3dsAuthenticationIDChallengesPost(ctx context.Context, req *R3dsAuthentications3dsAuthenticationIDChallengesPostReq, params R3dsAuthentications3dsAuthenticationIDChallengesPostParams) (R3dsAuthentications3dsAuthenticationIDChallengesPostRes, error)
	// R3dsAuthentications3dsAuthenticationIDFingerprintsPost implements  operation.
	//
	// Finish the device fingerprint step for the 3DS authentication session.
	// Mainly to be invoked by VGS checkout.js, not meant to be called by user directly.
	//
	// POST /3ds_authentications/{3ds_authentication_id}/fingerprints
	R3dsAuthentications3dsAuthenticationIDFingerprintsPost(ctx context.Context, req *R3dsAuthentications3dsAuthenticationIDFingerprintsPostReq, params R3dsAuthentications3dsAuthenticationIDFingerprintsPostParams) (R3dsAuthentications3dsAuthenticationIDFingerprintsPostRes, error)
	// ReverseTransferById implements ReverseTransferById operation.
	//
	// Reverse a specific transfer by id.
	//
	// POST /transfers/{transfer_id}/reversals
	ReverseTransferById(ctx context.Context, req OptReversal, params ReverseTransferByIdParams) (ReverseTransferByIdRes, error)
	// UpdateGatewayById implements UpdateGatewayById operation.
	//
	// Update a particular gateway by id.
	//
	// PUT /gateways/{gateway_id}
	UpdateGatewayById(ctx context.Context, req OptGatewayUpdate, params UpdateGatewayByIdParams) (UpdateGatewayByIdRes, error)
	// UpdateRuleById implements UpdateRuleById operation.
	//
	// Update a particular rule by id.
	//
	// PUT /rules/{rule_id}
	UpdateRuleById(ctx context.Context, req OptRuleUpdate, params UpdateRuleByIdParams) (UpdateRuleByIdRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	cfg config

	requests syncint64.Counter
	errors   syncint64.Counter
	duration syncint64.Histogram
}

func NewServer(h Handler, opts ...Option) (*Server, error) {
	s := &Server{
		h:   h,
		cfg: newConfig(opts...),
	}
	var err error
	if s.requests, err = s.cfg.Meter.SyncInt64().Counter(otelogen.ServerRequestCount); err != nil {
		return nil, err
	}
	if s.errors, err = s.cfg.Meter.SyncInt64().Counter(otelogen.ServerErrorsCount); err != nil {
		return nil, err
	}
	if s.duration, err = s.cfg.Meter.SyncInt64().Histogram(otelogen.ServerDuration); err != nil {
		return nil, err
	}
	return s, nil
}
