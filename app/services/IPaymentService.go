package services

import (
	"github.com/fly-code-challenge/app/Entities"
	"github.com/fly-code-challenge/app/Handlers"
)

type IPaymentService interface {
	GetPaymentTransactions(filter *Handlers.FilterHandler) ([]*Entities.TransactionEntity, error)
}
