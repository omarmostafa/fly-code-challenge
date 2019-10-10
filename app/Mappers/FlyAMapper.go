package Mappers

import (
	"github.com/fly-code-challenge/app/Entities"
	"github.com/fly-code-challenge/app/Repositories"
)

type FlyAMapper struct {
	repo Repositories.IFlyARepository
}

func NewFlyAMapper(repo Repositories.IFlyARepository) *FlyAMapper {
	return &FlyAMapper{
		repo: repo,
	}
}

func (mapper *FlyAMapper) MapToTransactionEntity(transactions *[]*Entities.TransactionEntity) {
	flyAEntity := mapper.repo.GetFlyAEntity()
	for _, transaction := range flyAEntity.Transactions {
		paymentEntity := &Entities.TransactionEntity{
			transaction.TransactionId,
			transaction.Amount,
			transaction.GetStatusCode(),
			transaction.OrderReference,
			transaction.Currency,
		}
		*transactions = append(*transactions, paymentEntity)
	}
}
