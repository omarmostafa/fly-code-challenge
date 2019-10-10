package Handlers

import (
	"github.com/fly-code-challenge/app/Entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilterHandler_IsMatchAmountWithValidAmount(t *testing.T) {
	filter := Entities.FilterEntity{
		MinAmount:  100,
		MaxAmount:  200,
		Currency:   "AUD",
		Provider:   "flyPayB",
		StatusCode: "authorized",
	}

	handler := NewFilterHandler(filter)

	result := handler.IsMatchAmount(150)
	assert.True(t, result)
}

func TestFilterHandler_IsMatchAmountWithEmpty(t *testing.T) {
	filter := Entities.FilterEntity{
		Currency:   "AUD",
		Provider:   "flyPayB",
		StatusCode: "authorized",
	}

	handler := NewFilterHandler(filter)

	result := handler.IsMatchAmount(150)
	assert.True(t, result)
}

func TestFilterHandler_IsMatchAmountWithInvalidAmount(t *testing.T) {
	filter := Entities.FilterEntity{
		MinAmount:  100,
		MaxAmount:  200,
		Currency:   "AUD",
		Provider:   "flyPayB",
		StatusCode: "authorized",
	}

	handler := NewFilterHandler(filter)

	result := handler.IsMatchAmount(300)
	assert.False(t, result)
}

func TestFilterHandler_IsMatchCurrencyValid(t *testing.T) {
	filter := Entities.FilterEntity{
		MinAmount:  100,
		MaxAmount:  200,
		Currency:   "AUD",
		Provider:   "flyPayB",
		StatusCode: "authorized",
	}

	handler := NewFilterHandler(filter)

	result := handler.IsMatchCurrency("AUD")
	assert.True(t, result)
}

func TestFilterHandler_IsMatchCurrencyInvalid(t *testing.T) {
	filter := Entities.FilterEntity{
		MinAmount:  100,
		MaxAmount:  200,
		Currency:   "AUD",
		Provider:   "flyPayB",
		StatusCode: "authorized",
	}

	handler := NewFilterHandler(filter)

	result := handler.IsMatchCurrency("EGP")
	assert.False(t, result)
}

func TestFilterHandler_IsMatchCurrencyEmpty(t *testing.T) {
	filter := Entities.FilterEntity{
		MinAmount:  100,
		MaxAmount:  200,
		Provider:   "flyPayB",
		StatusCode: "authorized",
	}

	handler := NewFilterHandler(filter)

	result := handler.IsMatchCurrency("EGP")
	assert.True(t, result)
}

func TestFilterHandler_IsMatchStatusCodeValid(t *testing.T) {
	filter := Entities.FilterEntity{
		MinAmount:  100,
		MaxAmount:  200,
		Currency:   "AUD",
		Provider:   "flyPayB",
		StatusCode: "authorised",
	}

	handler := NewFilterHandler(filter)

	result := handler.IsMatchStatusCode("authorised")
	assert.True(t, result)
}

func TestFilterHandler_IsMatchStatusCodeInvalid(t *testing.T) {
	filter := Entities.FilterEntity{
		MinAmount:  100,
		MaxAmount:  200,
		Currency:   "AUD",
		Provider:   "flyPayB",
		StatusCode: "authorised",
	}

	handler := NewFilterHandler(filter)

	result := handler.IsMatchStatusCode("declined")
	assert.False(t, result)
}

func TestFilterHandler_IsItemValid(t *testing.T) {
	filter := Entities.FilterEntity{
		MinAmount:  100,
		MaxAmount:  200,
		Currency:   "AUD",
		Provider:   "flyPayB",
		StatusCode: "authorised",
	}

	handler := NewFilterHandler(filter)

	item := Entities.TransactionEntity{
		Id:             "fly 1",
		Amount:         150,
		StatusCode:     "authorised",
		OrderReference: "test reference",
		Currency:       "AUD",
	}

	result := handler.IsItemValid(&item)
	assert.True(t, result)
}

func TestFilterHandler_IsItemValidInvalid(t *testing.T) {
	filter := Entities.FilterEntity{
		MinAmount:  100,
		MaxAmount:  200,
		Currency:   "AUD",
		Provider:   "flyPayB",
		StatusCode: "authorised",
	}

	handler := NewFilterHandler(filter)

	item := Entities.TransactionEntity{
		Id:             "fly 1",
		Amount:         50,
		StatusCode:     "authorised",
		OrderReference: "test reference",
		Currency:       "AUD",
	}

	result := handler.IsItemValid(&item)
	assert.False(t, result)
}

func TestFilterHandler_FilterByAmount(t *testing.T) {
	filter := Entities.FilterEntity{
		MinAmount: 100,
		MaxAmount: 200,
	}

	handler := NewFilterHandler(filter)

	var transactions = []*Entities.TransactionEntity{
		{
			Id:             "fly a1",
			Amount:         150,
			Currency:       "AUD",
			OrderReference: "reference",
			StatusCode:     "authorised",
		},
		{
			Id:             "fly a2",
			Amount:         50,
			Currency:       "AUD",
			OrderReference: "reference",
			StatusCode:     "authorised",
		},
	}

	var filteredTransactions = []*Entities.TransactionEntity{
		{
			Id:             "fly a1",
			Amount:         150,
			Currency:       "AUD",
			OrderReference: "reference",
			StatusCode:     "authorised",
		},
	}

	result := handler.Filter(transactions, handler.IsItemValid)
	assert.Equal(t, result, filteredTransactions)
}

func TestFilterHandler_FilterByCurrency(t *testing.T) {
	filter := Entities.FilterEntity{
		Currency: "AUD",
	}

	handler := NewFilterHandler(filter)

	var transactions = []*Entities.TransactionEntity{
		{
			Id:             "fly a1",
			Amount:         150,
			Currency:       "AUD",
			OrderReference: "reference",
			StatusCode:     "authorised",
		},
		{
			Id:             "fly a2",
			Amount:         50,
			Currency:       "EGP",
			OrderReference: "reference",
			StatusCode:     "authorised",
		},
	}

	var filteredTransactions = []*Entities.TransactionEntity{
		{
			Id:             "fly a1",
			Amount:         150,
			Currency:       "AUD",
			OrderReference: "reference",
			StatusCode:     "authorised",
		},
	}

	result := handler.Filter(transactions, handler.IsItemValid)
	assert.Equal(t, result, filteredTransactions)
}

func TestFilterHandler_FilterByStatusCode(t *testing.T) {
	filter := Entities.FilterEntity{
		StatusCode: "authorised",
	}

	handler := NewFilterHandler(filter)

	var transactions = []*Entities.TransactionEntity{
		{
			Id:             "fly a1",
			Amount:         150,
			Currency:       "AUD",
			OrderReference: "reference",
			StatusCode:     "authorised",
		},
		{
			Id:             "fly a2",
			Amount:         50,
			Currency:       "EGP",
			OrderReference: "reference",
			StatusCode:     "decline",
		},
	}

	var filteredTransactions = []*Entities.TransactionEntity{
		{
			Id:             "fly a1",
			Amount:         150,
			Currency:       "AUD",
			OrderReference: "reference",
			StatusCode:     "authorised",
		},
	}

	result := handler.Filter(transactions, handler.IsItemValid)
	assert.Equal(t, result, filteredTransactions)
}

func TestFilterHandler_FilterCombined(t *testing.T) {
	filter := Entities.FilterEntity{
		MinAmount:  100,
		MaxAmount:  200,
		Currency:   "AUD",
		Provider:   "flyPayB",
		StatusCode: "authorised",
	}

	handler := NewFilterHandler(filter)

	var transactions = []*Entities.TransactionEntity{
		{
			Id:             "fly a1",
			Amount:         150,
			Currency:       "AUD",
			OrderReference: "reference",
			StatusCode:     "authorised",
		},
		{
			Id:             "fly a2",
			Amount:         170,
			Currency:       "EGP",
			OrderReference: "reference",
			StatusCode:     "authorized",
		},
		{
			Id:             "fly a2",
			Amount:         300,
			Currency:       "AUD",
			OrderReference: "reference",
			StatusCode:     "authorised",
		},
		{
			Id:             "fly a2",
			Amount:         150,
			Currency:       "AUD",
			OrderReference: "reference",
			StatusCode:     "declined",
		},
	}

	var filteredTransactions = []*Entities.TransactionEntity{
		{
			Id:             "fly a1",
			Amount:         150,
			Currency:       "AUD",
			OrderReference: "reference",
			StatusCode:     "authorised",
		},
	}

	result := handler.Filter(transactions, handler.IsItemValid)
	assert.Equal(t, result, filteredTransactions)
}
