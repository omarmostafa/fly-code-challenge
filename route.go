package main

import (
	"github.com/gorilla/mux"
	"sync"
)

type IRoute interface {
	InitRoute() *mux.Router
}

type router struct {

}

func (router *router) InitRoute() *mux.Router  {
	r:= mux.NewRouter()
	paymentController := ServiceContainer().InjectPaymentController()
	r.HandleFunc("/payments/transactions", paymentController.GetPaymentTransactions).Methods("GET")
	return  r
}

var (
	route *router
	routerOnce sync.Once
)

func NewRouter() IRoute {
	if route == nil {
		routerOnce.Do(func() {
			route = &router{}
		})
	}
	return  route
}
