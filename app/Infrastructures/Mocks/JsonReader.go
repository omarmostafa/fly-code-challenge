package Mocks

import (
	"github.com/stretchr/testify/mock"
)

type JsonReader struct {
	mock.Mock
}

func (_m JsonReader) ReadDataFromJsonFile(path string, entity interface{}) error {
	ret := _m.Called()

	var r1 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(0)
	}
	return r1
}
