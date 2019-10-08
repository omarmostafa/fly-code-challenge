package Repositories

import (
	"github.com/code-challenge/app/Entities"
	"github.com/code-challenge/app/Infrastructures"
)

type FlyBRepository struct {
	reader Infrastructures.JsonReader
}

func NewFlyBRepository() *FlyBRepository  {
	return &FlyBRepository{
		reader:Infrastructures.NewJsonReader(),
	}
}

func (repo *FlyBRepository) GetSourceDestination () string  {
	return "/app/data/flypayB.json"
}


func (repo *FlyBRepository) ConvertToTransactionEntity(transactions *[]*Entities.TransactionEntity) error {
	var flyBEntity *Entities.FlyPayBEntity
	err :=repo.reader.ReadDataFromJsonFile(repo.GetSourceDestination(),&flyBEntity)

	for _,transaction := range flyBEntity.Transactions {
		paymentEntity := &Entities.TransactionEntity{
			transaction.PaymentId,
			transaction.Value,
			transaction.GetStatusCode(),
			transaction.OrderInfo,
			transaction.TransactionCurrency,
		}
		*transactions = append(*transactions,paymentEntity)
	}
	return err
}
