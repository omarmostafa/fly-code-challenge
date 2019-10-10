package Handlers

import "github.com/fly-code-challenge/app/Mappers"

type IProviderHandler interface {
	GetProvidersByFilter(provider string) (map[string]func() Mappers.ITransactionMapper, error)
}
