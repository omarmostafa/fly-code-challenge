package controllers

import (
	"github.com/code-challenge/app/Entities"
	"github.com/code-challenge/app/Handlers"
	"github.com/code-challenge/app/services"
	"net/http"
	"github.com/gorilla/schema"
)

type PaymentController struct {
	service services.IPaymentService
}

func NewPaymentController(service services.IPaymentService) *PaymentController {
	return &PaymentController{
		service: service,
	}
}


func (controller *PaymentController) GetPaymentTransactions(res http.ResponseWriter, req *http.Request) {
	var filter Entities.FilterEntity
	err := schema.NewDecoder().Decode(&filter, req.URL.Query())
	filterHandler := Handlers.NewFilterHandler(filter)
	entities,err :=controller.service.GetPaymentTransactions(filterHandler)
	if(err != nil){

	}
	RespondAccepted(res,req,"succes",entities)
}
