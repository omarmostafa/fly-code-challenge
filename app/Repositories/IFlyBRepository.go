package Repositories

import "github.com/fly-code-challenge/app/Entities"

type IFlyBRepository interface {
	GetFlyBEntity() *Entities.FlyPayBEntity
	GetSourceDestination() string
}
