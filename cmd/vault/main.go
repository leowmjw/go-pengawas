package main

import (
	"app/api/vault"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"net/http"
)

// FakerVaultServer might be how as Faker service is imeplemented
type FakerVaultServer struct {
	VaultServer       vault.ServerInterface // Implements Vault ServerInterface
	SetupScenarioFunc func() string         // Function to set how the Server will react ..
}

func (f FakerVaultServer) RevealMultipleAliases(w http.ResponseWriter, r *http.Request, params vault.RevealMultipleAliasesParams) {
	spew.Dump(params)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte("RES: RevealMultipleAliases"))
	return
}

func (f FakerVaultServer) CreateAliases(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (f FakerVaultServer) DeleteAlias(w http.ResponseWriter, r *http.Request, alias string) {
	//TODO implement me
	panic("implement me")
}

func (f FakerVaultServer) RevealAlias(w http.ResponseWriter, r *http.Request, alias string) {
	fmt.Println("IN RevealAlias: ")
	spew.Dump(alias)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte("RES: RevealAlias"))
	return
}

func (f FakerVaultServer) UpdateAlias(w http.ResponseWriter, r *http.Request, alias string) {
	//TODO implement me
	panic("implement me")
}

type MockVaultServer struct{} // This is how to do a simple Faker service?

func (mvs MockVaultServer) CreateAliases(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (mvs MockVaultServer) DeleteAlias(w http.ResponseWriter, r *http.Request, alias string) {
	//TODO implement me
	panic("implement me")
}

func (mvs MockVaultServer) RevealAlias(w http.ResponseWriter, r *http.Request, alias string) {
	spew.Dump(alias)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte("RES: RevealAlias"))
	return
}

func (mvs MockVaultServer) UpdateAlias(w http.ResponseWriter, r *http.Request, alias string) {
	//TODO implement me
	panic("implement me")
}

func (mvs MockVaultServer) RevealMultipleAliases(w http.ResponseWriter, r *http.Request, params vault.RevealMultipleAliasesParams) {

	w.WriteHeader(http.StatusOK)
	//render.Render(w, r, "bob")
	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte("RES: RevealMultipleAliases"))
	return
}

func main() {
	fmt.Println("Vault Implementation ..")

	// OpenAPI data ..
	swg, err := vault.GetSwagger()
	if err != nil {
		panic(err)
	}
	spew.Dump(swg.Tags)

	mvs := MockVaultServer{}
	fvs := FakerVaultServer{
		VaultServer: mvs,
		SetupScenarioFunc: func() string {
			return "boo"
		},
	}
	http.ListenAndServe(":8080", vault.Handler(fvs))
}
