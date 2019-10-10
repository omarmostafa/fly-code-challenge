package Mappers

import (
	"github.com/fly-code-challenge/app/Entities"
	"github.com/fly-code-challenge/app/Repositories"
)

type FlyBMapper struct {
	repo Repositories.IFlyBRepository
}

func NewFlyBMapper(repo Repositories.IFlyBRepository) *FlyBMapper {
	return &FlyBMapper{
		repo: repo,
	}
}

func (mapper *FlyBMapper) MapToTransactionEntity(transactions *[]*Entities.TransactionEntity) {
	flyAEntity := mapper.repo.GetFlyBEntity()
	for _, transaction := range flyAEntity.Transactions {
		paymentEntity := &Entities.TransactionEntity{
			transaction.PaymentId,
			transaction.Value,
			transaction.GetStatusCode(),
			transaction.OrderInfo,
			transaction.TransactionCurrency,
		}
		*transactions = append(*transactions, paymentEntity)
	}
}
