package Handlers

import "github.com/code-challenge/app/Repositories"

func GetAllProviders() map[string] func() Repositories.ITransactionRepository {
	all_providers := map[string] func() Repositories.ITransactionRepository  {
		"fly_a" : func() Repositories.ITransactionRepository{
			return Repositories.NewFlyARepository()
		},
		"fly_b" : func() Repositories.ITransactionRepository{
			return Repositories.NewFlyBRepository()
		},
	}
	return all_providers
}
