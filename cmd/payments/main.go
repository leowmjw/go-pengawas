package main

import (
	"app/api/payments"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"net/http"
)

type MockPaymentsServer struct{}

func (mps MockPaymentsServer) CreateThreeDSAuthentication(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (mps MockPaymentsServer) ThreeDSAuthenticationChallenge(w http.ResponseWriter, r *http.Request, n3dsAuthenticationId payments.N3dsAuthenticationId) {
	//TODO implement me
	panic("implement me")
}

func (mps MockPaymentsServer) ThreeDSAuthenticationDeviceFingerprint(w http.ResponseWriter, r *http.Request, n3dsAuthenticationId payments.N3dsAuthenticationId) {
	//TODO implement me
	panic("implement me")
}

func (mps MockPaymentsServer) ThreeDSAuthenticationNotifications(w http.ResponseWriter, r *http.Request, n3dsAuthenticationId payments.N3dsAuthenticationId) {
	//TODO implement me
	panic("implement me")
}

func (mps MockPaymentsServer) CreateThreeDSProvider(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (mps MockPaymentsServer) GetFinancialInstruments(w http.ResponseWriter, r *http.Request, params payments.GetFinancialInstrumentsParams) {
	//TODO implement me
	panic("implement me")
}

func (mps MockPaymentsServer) CreateFinancialInstrument(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (mps MockPaymentsServer) DeleteFinancialInstrumentById(w http.ResponseWriter, r *http.Request, financialInstrumentId payments.FinancialInstrumentId) {
	//TODO implement me
	panic("implement me")
}

func (mps MockPaymentsServer) GetFinancialInstrumentById(w http.ResponseWriter, r *http.Request, financialInstrumentId payments.FinancialInstrumentId) {
	//TODO implement me
	panic("implement me")
}

func (mps MockPaymentsServer) GetGateways(w http.ResponseWriter, r *http.Request, params payments.GetGatewaysParams) {
	//TODO implement me
	panic("implement me")
}

func (mps MockPaymentsServer) CreateGateway(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (mps MockPaymentsServer) DeleteGateway(w http.ResponseWriter, r *http.Request, gatewayId payments.GatewayId) {
	//TODO implement me
	panic("implement me")
}

func (mps MockPaymentsServer) GetGatewayById(w http.ResponseWriter, r *http.Request, gatewayId payments.GatewayId) {
	//TODO implement me
	panic("implement me")
}

func (mps MockPaymentsServer) UpdateGatewayById(w http.ResponseWriter, r *http.Request, gatewayId payments.GatewayId) {
	//TODO implement me
	panic("implement me")
}

func (mps MockPaymentsServer) CreateOrder(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (mps MockPaymentsServer) GetOrdersOrderId(w http.ResponseWriter, r *http.Request, orderId payments.OrderId) {
	//TODO implement me
	panic("implement me")
}

func (mps MockPaymentsServer) GetRules(w http.ResponseWriter, r *http.Request, params payments.GetRulesParams) {
	//TODO implement me
	panic("implement me")
}

func (mps MockPaymentsServer) CreateRule(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (mps MockPaymentsServer) DeleteRule(w http.ResponseWriter, r *http.Request, ruleId payments.RuleId) {
	//TODO implement me
	panic("implement me")
}

func (mps MockPaymentsServer) GetRuleById(w http.ResponseWriter, r *http.Request, ruleId payments.RuleId) {
	//TODO implement me
	panic("implement me")
}

func (mps MockPaymentsServer) UpdateRuleById(w http.ResponseWriter, r *http.Request, ruleId payments.RuleId) {
	//TODO implement me
	panic("implement me")
}

func (mps MockPaymentsServer) GetTransfers(w http.ResponseWriter, r *http.Request, params payments.GetTransfersParams) {
	//TODO implement me
	panic("implement me")
}

func (mps MockPaymentsServer) CreateTransfer(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (mps MockPaymentsServer) GetTransferById(w http.ResponseWriter, r *http.Request, transferId payments.TransferId) {
	//TODO implement me
	panic("implement me")
}

func (mps MockPaymentsServer) FinalizeTransfer(w http.ResponseWriter, r *http.Request, transferId payments.TransferId) {
	//TODO implement me
	panic("implement me")
}

func (mps MockPaymentsServer) ReverseTransferById(w http.ResponseWriter, r *http.Request, transferId payments.TransferId) {
	//TODO implement me
	panic("implement me")
}

func (mps MockPaymentsServer) CreateVerification(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func main() {
	fmt.Println("Service Payments ...")

	// OpenAPI data ..
	swg, err := payments.GetSwagger()
	if err != nil {
		panic(err)
	}
	spew.Dump(swg.Tags)

	mps := MockPaymentsServer{}
	http.ListenAndServe(":8080", payments.Handler(mps))
}
