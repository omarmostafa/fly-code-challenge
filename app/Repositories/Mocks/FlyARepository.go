package Mocks

import (
	"github.com/fly-code-challenge/app/Entities"
	"github.com/stretchr/testify/mock"
)

type FlyARepository struct {
	mock.Mock
}

func (_m FlyARepository) GetFlyAEntity() *Entities.FlyPayAEntity {
	ret := _m.Called()

	var r1 *Entities.FlyPayAEntity
	if rf, ok := ret.Get(0).(func() *Entities.FlyPayAEntity); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(0).(*Entities.FlyPayAEntity)
	}
	return r1
}

func (_m *FlyARepository) GetSourceDestination() string {
	ret := _m.Called()

	var r1 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r1 = rf()
	} else {
		r1 = ret.String(0)
	}
	return r1
}
