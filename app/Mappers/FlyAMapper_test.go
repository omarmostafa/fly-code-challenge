package Mappers

import (
	"github.com/fly-code-challenge/app/Entities"
	"github.com/fly-code-challenge/app/Repositories/Mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlyAMapper_MapToTransactionEntity(t *testing.T) {

	flyARepository := new(Mocks.FlyARepository)
	var flyATransactions = []Entities.FlyPayATransactions{
		{
			TransactionId:  "fly a1",
			Amount:         1999,
			Currency:       "EGP",
			OrderReference: "reference",
			StatusCode:     1,
		},
		{
			TransactionId:  "fly a2",
			Amount:         3000,
			Currency:       "AUD",
			OrderReference: "reference",
			StatusCode:     2,
		},
	}
	var actualTransactions []*Entities.TransactionEntity
	var transactions = []*Entities.TransactionEntity{
		{
			Id:             "fly a1",
			Amount:         1999,
			Currency:       "EGP",
			OrderReference: "reference",
			StatusCode:     "authorised",
		},
		{
			Id:             "fly a2",
			Amount:         3000,
			Currency:       "AUD",
			OrderReference: "reference",
			StatusCode:     "decline",
		},
	}
	var flyAEntity = Entities.FlyPayAEntity{Transactions: flyATransactions}
	flyARepository.On("GetFlyAEntity").Return(&flyAEntity, nil)

	flyMapper := FlyAMapper{flyARepository}

	flyMapper.MapToTransactionEntity(&actualTransactions)

	assert.Equal(t, transactions, actualTransactions)
}

func BenchmarkFlyAMapper_MapToTransactionEntity(b *testing.B) {

	flyARepository := new(Mocks.FlyARepository)
	var flyATransactions = []Entities.FlyPayATransactions{
		{
			TransactionId:  "fly a1",
			Amount:         1999,
			Currency:       "EGP",
			OrderReference: "reference",
			StatusCode:     1,
		},
		{
			TransactionId:  "fly a2",
			Amount:         3000,
			Currency:       "AUD",
			OrderReference: "reference",
			StatusCode:     2,
		},
	}
	var actualTransactions []*Entities.TransactionEntity

	var flyAEntity = Entities.FlyPayAEntity{Transactions: flyATransactions}
	flyARepository.On("GetFlyAEntity").Return(&flyAEntity, nil)

	flyMapper := FlyAMapper{flyARepository}

	for i := 0; i < b.N; i++ {
		flyMapper.MapToTransactionEntity(&actualTransactions)
	}

}