package services

import (
	"github.com/code-challenge/app/Entities"
	"github.com/code-challenge/app/Handlers"
)

type IPaymentService interface {
	GetPaymentTransactions (filter *Handlers.FilterHandler) ([]*Entities.TransactionEntity,error)
}

