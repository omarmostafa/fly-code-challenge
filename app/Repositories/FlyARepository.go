package Repositories

import (
	"github.com/fly-code-challenge/app/Entities"
	"github.com/fly-code-challenge/app/Infrastructures"
)

type FlyARepository struct {
	reader Infrastructures.IJsonReader
}

func NewFlyARepository() *FlyARepository {
	return &FlyARepository{
		reader: Infrastructures.NewJsonReader(),
	}
}

func (repo *FlyARepository) GetSourceDestination() string {
	return "/data/flypayA.json"
}

func (repo *FlyARepository) GetFlyAEntity() *Entities.FlyPayAEntity {
	var flyAEntity *Entities.FlyPayAEntity
	err := repo.reader.ReadDataFromJsonFile(repo.GetSourceDestination(), &flyAEntity)

	if err != nil {
		panic(err)
	}
	return flyAEntity
}
