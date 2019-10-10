package services

import (
	"github.com/fly-code-challenge/app/Entities"
	"github.com/fly-code-challenge/app/Handlers"
)

type PaymentService struct {
	providerHandler Handlers.IProviderHandler
}

func NewPaymentService(handler Handlers.IProviderHandler) *PaymentService {
	return &PaymentService{
		handler,
	}
}

func (service *PaymentService) GetPaymentTransactions(filter *Handlers.FilterHandler) ([]*Entities.TransactionEntity, error) {
	var transactions []*Entities.TransactionEntity
	providers, err := service.providerHandler.GetProvidersByFilter(filter.FilterEntity.Provider)

	if err != nil {
		return nil, err
	}

	for _, provider := range providers {
		provider().MapToTransactionEntity(&transactions)
	}
	transactions = filter.FilterTransactions(transactions)
	return transactions, nil
}
