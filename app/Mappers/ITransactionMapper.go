package Mappers

import (
	"github.com/fly-code-challenge/app/Entities"
)

type ITransactionMapper interface {
	MapToTransactionEntity(transactions *[]*Entities.TransactionEntity)
}
