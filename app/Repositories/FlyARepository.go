package Repositories

import (
	"github.com/code-challenge/app/Entities"
	"github.com/code-challenge/app/Infrastructures"
)

type FlyARepository struct {
	reader Infrastructures.JsonReader
}

func NewFlyARepository() *FlyARepository {
	return &FlyARepository{
		reader: Infrastructures.NewJsonReader(),
	}
}

func (repo *FlyARepository) GetSourceDestination() string {
	return "/app/data/flypayA.json"
}

func (repo *FlyARepository) ConvertToTransactionEntity(transactions *[]*Entities.TransactionEntity) {
	var flyAEntity *Entities.FlyPayAEntity
	err := repo.reader.ReadDataFromJsonFile(repo.GetSourceDestination(), &flyAEntity)

	if err != nil {
		panic(err)
	}

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
