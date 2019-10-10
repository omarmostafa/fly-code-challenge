package Repositories

import "github.com/fly-code-challenge/app/Entities"

type IFlyARepository interface {
	GetFlyAEntity() *Entities.FlyPayAEntity
	GetSourceDestination() string
}
