package controllers

import (
	"github.com/fly-code-challenge/app/Entities"
	"github.com/fly-code-challenge/app/Handlers"
	"github.com/fly-code-challenge/app/services"
	"github.com/gorilla/schema"
	"net/http"
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
	_ = schema.NewDecoder().Decode(&filter, req.URL.Query())
	filterHandler := Handlers.NewFilterHandler(filter)
	entities, err := controller.service.GetPaymentTransactions(filterHandler)
	if err != nil {
		RespondBadRequest(res, req, err.Error(), nil)
		return
	}
	RespondAccepted(res, req, "Transactions returned successfully", entities)
}
