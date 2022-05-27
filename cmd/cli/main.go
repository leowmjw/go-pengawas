package main

import (
	"app/api/payments"
	"fmt"
	"github.com/labstack/echo/v4"
)

type PaymentsServer struct {
	payments.ServerInterface
}

func main() {
	fmt.Println("go-pengawas!!")

	//echo.HandlerFunc()
	//var ctx echo.Context
	//
	//ps := PaymentsServer{}
	//
	//err := ps.CreateGateway(ctx)
	//if err != nil {
	//	// Do somtjong ..
	//}
	//
	//c, cerr := payments.NewClient("")
	//if cerr != nil {
	//
	//}
	s := echo.New()
	err := s.StartAutoTLS("localhost:8080")
	if err != nil {
		panic("BOOM!!!")
	}
}
