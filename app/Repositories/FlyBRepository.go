package Repositories

import (
	"github.com/fly-code-challenge/app/Entities"
	"github.com/fly-code-challenge/app/Infrastructures"
)

type FlyBRepository struct {
	reader Infrastructures.IJsonReader
}

func NewFlyBRepository() *FlyBRepository {
	return &FlyBRepository{
		reader: Infrastructures.NewJsonReader(),
	}
}

func (repo *FlyBRepository) GetSourceDestination() string {
	return "/data/flypayB.json"
}

func (repo *FlyBRepository) GetFlyBEntity() *Entities.FlyPayBEntity {
	var flyBEntity *Entities.FlyPayBEntity
	err := repo.reader.ReadDataFromJsonFile(repo.GetSourceDestination(), &flyBEntity)

	if err != nil {
		panic(err)
	}
	return flyBEntity
}
