package Handlers

import (
	"github.com/code-challenge/app/Entities"
)

type FilterHandler struct {
	FilterEntity Entities.FilterEntity
}

func NewFilterHandler(filterEntity Entities.FilterEntity) *FilterHandler {
	return &FilterHandler{
		FilterEntity: filterEntity,
	}
}

func (handler *FilterHandler) FilterTransactions(transactions []*Entities.TransactionEntity) []*Entities.TransactionEntity {
	return handler.Filter(transactions, handler.IsItemValid)
}

func (handler *FilterHandler) IsItemValid(transaction *Entities.TransactionEntity) bool {
	isItemValid := true
	filters := handler.generateAllFilters(transaction)
	for _, filter := range filters {
		if filter() == false {
			isItemValid = false
		}
	}
	return isItemValid
}

func (handler *FilterHandler) generateAllFilters(transaction *Entities.TransactionEntity) map[string]func() bool {
	filters := map[string]func() bool{
		"currency": func() bool {
			return handler.IsMatchCurrency(transaction.Currency)
		},
		"Amount": func() bool {
			return handler.IsMatchAmount(transaction.Amount)
		},
		"statusCode": func() bool {
			return handler.IsMatchStatusCode(transaction.StatusCode)
		},
	}
	return filters
}

func (handler *FilterHandler) IsMatchCurrency(currency string) bool {
	if handler.FilterEntity.Currency == "" {
		return true
	}
	return currency == handler.FilterEntity.Currency
}

func (handler *FilterHandler) IsMatchAmount(amount int) bool {
	if handler.FilterEntity.MaxAmount == 0 && handler.FilterEntity.MinAmount == 0 {
		return true
	}
	validateMin := true
	validateMax := true
	if handler.FilterEntity.MinAmount != 0 {
		validateMin = amount >= handler.FilterEntity.MinAmount
	}

	if handler.FilterEntity.MaxAmount != 0 {
		validateMax = amount <= handler.FilterEntity.MaxAmount
	}

	return validateMax && validateMin
}

func (handler *FilterHandler) IsMatchStatusCode(statusCode string) bool {
	if handler.FilterEntity.StatusCode == "" {
		return true
	}
	return statusCode == handler.FilterEntity.StatusCode
}

func (handler *FilterHandler) Filter(transactions []*Entities.TransactionEntity, criteria func(*Entities.TransactionEntity) bool) []*Entities.TransactionEntity {
	var result []*Entities.TransactionEntity
	for _, transaction := range transactions {
		if criteria(transaction) {
			result = append(result, transaction)
		}
	}
	return result
}
