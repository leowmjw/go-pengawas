package main

import (
	"app/api/calm"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"net/http"
)

type MockCALMServer struct{}

func (mcs MockCALMServer) EnrollCard(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (mcs MockCALMServer) GetCard(w http.ResponseWriter, r *http.Request, cardId string) {
	//TODO implement me
	panic("implement me")
}

func main() {
	fmt.Println("Service CALM ...")

	// OpenAPI data ..
	swg, err := calm.GetSwagger()
	if err != nil {
		panic(err)
	}
	spew.Dump(swg.Tags)

	mcs := MockCALMServer{}
	http.ListenAndServe(":8080", calm.Handler(mcs))
}
