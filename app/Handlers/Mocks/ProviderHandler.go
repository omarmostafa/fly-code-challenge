package Mocks

import (
	"github.com/fly-code-challenge/app/Mappers"
	"github.com/stretchr/testify/mock"
)

type ProviderHandler struct {
	mock.Mock
}

func (_m ProviderHandler) GetProvidersByFilter(provider string) (map[string]func() Mappers.ITransactionMapper, error) {
	ret := _m.Called()

	var r1 map[string]func() Mappers.ITransactionMapper
	if rf, ok := ret.Get(0).(func() map[string]func() Mappers.ITransactionMapper); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(0).(map[string]func() Mappers.ITransactionMapper)
	}

	var r2 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(1)
	}
	return r1, r2
}
