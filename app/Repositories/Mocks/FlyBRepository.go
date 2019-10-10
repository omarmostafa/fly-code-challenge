package Mocks

import (
	"github.com/fly-code-challenge/app/Entities"
	"github.com/stretchr/testify/mock"
)

type FlyBRepository struct {
	mock.Mock
}

func (_m FlyBRepository) GetFlyBEntity() *Entities.FlyPayBEntity {
	ret := _m.Called()

	var r1 *Entities.FlyPayBEntity
	if rf, ok := ret.Get(0).(func() *Entities.FlyPayBEntity); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(0).(*Entities.FlyPayBEntity)
	}
	return r1
}

func (_m *FlyBRepository) GetSourceDestination() string {
	ret := _m.Called()

	var r1 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r1 = rf()
	} else {
		r1 = ret.String(0)
	}
	return r1
}
