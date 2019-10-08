package services

import (
	"github.com/code-challenge/app/Entities"
	"github.com/code-challenge/app/Handlers"
)

type PaymentService struct {
}


func NewPaymentService()  *PaymentService{
	return &PaymentService{
	}
}

func (service *PaymentService) GetPaymentTransactions(filter *Handlers.FilterHandler) ([]*Entities.TransactionEntity,error) {
	var transactions []*Entities.TransactionEntity
	providers := Handlers.GetAllProviders()
	var err error
	for _, provider := range providers {
		err = provider().ConvertToTransactionEntity(&transactions)
	}
	transactions = filter.FilterTransactions(transactions)
	return  transactions , err
}


