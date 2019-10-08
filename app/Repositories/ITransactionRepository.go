package Repositories

import (
	"github.com/code-challenge/app/Entities"
)

type ITransactionRepository interface {
	ConvertToTransactionEntity(transactions *[]*Entities.TransactionEntity) error
	GetSourceDestination() string
}
