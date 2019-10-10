package Mappers

import (
	"github.com/fly-code-challenge/app/Entities"
	"github.com/fly-code-challenge/app/Repositories/Mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlyBMapper_MapToTransactionEntity(t *testing.T) {

	flyBRepository := new(Mocks.FlyBRepository)
	var flyBTransactions = []Entities.FlyPayBTransactions{
		{
			PaymentId:           "fly b1",
			Value:               1999,
			TransactionCurrency: "EGP",
			OrderInfo:           "reference",
			StatusCode:          100,
		},
		{
			PaymentId:           "fly b2",
			Value:               3000,
			TransactionCurrency: "AUD",
			OrderInfo:           "reference",
			StatusCode:          200,
		},
	}
	var actualTransactions []*Entities.TransactionEntity
	var transactions = []*Entities.TransactionEntity{
		{
			Id:             "fly b1",
			Amount:         1999,
			Currency:       "EGP",
			OrderReference: "reference",
			StatusCode:     "authorised",
		},
		{
			Id:             "fly b2",
			Amount:         3000,
			Currency:       "AUD",
			OrderReference: "reference",
			StatusCode:     "decline",
		},
	}
	var flyBEntity = Entities.FlyPayBEntity{Transactions: flyBTransactions}
	flyBRepository.On("GetFlyBEntity").Return(&flyBEntity, nil)

	flyMapper := FlyBMapper{flyBRepository}

	flyMapper.MapToTransactionEntity(&actualTransactions)

	assert.Equal(t, transactions, actualTransactions)
}
