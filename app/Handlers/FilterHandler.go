package Handlers

import (
	"fmt"
	"github.com/code-challenge/app/Entities"
)

type FilterHandler struct {
	filterEntity Entities.FilterEntity
}

func NewFilterHandler(filterEntity Entities.FilterEntity) *FilterHandler{
	return &FilterHandler{
		filterEntity:filterEntity,
	}
}

func (handler *FilterHandler) FilterTransactions(transactions []*Entities.TransactionEntity) []*Entities.TransactionEntity   {
	return handler.Filter(transactions,handler.IsItemValid)
}

func (handler *FilterHandler) IsItemValid(transaction *Entities.TransactionEntity) bool {
	isItemValid :=true
	filters := handler.generateAllFilters(transaction)
	for _, filter := range filters {
		 if filter() == false {
		 	isItemValid = false
		 }
	}
	return isItemValid
}

func (handler *FilterHandler) generateAllFilters(transaction *Entities.TransactionEntity) map[string] func() bool {
	filters := map[string] func() bool {
		"currency" : func() bool{
			return handler.IsMatchCurrency(transaction.Currency)
		},
		"Amount" : func() bool{
			return handler.IsMatchAmount(transaction.Amount)
		},
		"statusCode" : func() bool{
			return handler.IsMatchCurrency(transaction.StatusCode)
		},
	}
	return filters
}

func(handler *FilterHandler) IsMatchCurrency(currency string) bool  {
	fmt.Println(handler.filterEntity.Currency == "")
	if handler.filterEntity.Currency == "" {
		return true
	}
	return currency == handler.filterEntity.Currency
}

func(handler *FilterHandler) IsMatchAmount(amount int) bool  {
	if handler.filterEntity.MaxAmount != 0 && handler.filterEntity.MinAmount != 0 {
		return true
	}
	return amount >= handler.filterEntity.MinAmount && amount <= handler.filterEntity.MaxAmount
}

func(handler *FilterHandler) IsMatchStatusCode(statusCode string) bool  {
	if handler.filterEntity.StatusCode == "" {
		return true
	}
	return statusCode == handler.filterEntity.StatusCode
}

func (handler *FilterHandler) Filter(transactions []*Entities.TransactionEntity, criteria func(*Entities.TransactionEntity) bool) ([]*Entities.TransactionEntity) {
	var result []*Entities.TransactionEntity
	for _, transaction := range transactions {
		if criteria(transaction) {
			result = append(result, transaction)
		}
	}
	return result
}