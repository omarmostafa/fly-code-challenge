package Handlers

import (
	"github.com/fly-code-challenge/app/Errors"
	"github.com/fly-code-challenge/app/Mappers"
	"github.com/fly-code-challenge/app/Repositories"
)

type ProviderHandler struct {
}

func NewProviderHandler() *ProviderHandler {
	return &ProviderHandler{}
}

func (handler *ProviderHandler) GetProvidersByFilter(provider string) (map[string]func() Mappers.ITransactionMapper, error) {
	allProviders := handler.GetAllProviders()
	if provider == "" {
		return allProviders, nil
	}
	filteredProvider := allProviders[provider]
	if filteredProvider == nil {
		return nil, Errors.ValidationError{Message: "Invalid provider"}
	}
	return map[string]func() Mappers.ITransactionMapper{
		provider: filteredProvider,
	}, nil
}

func (handler *ProviderHandler) GetAllProviders() map[string]func() Mappers.ITransactionMapper {
	allProviders := map[string]func() Mappers.ITransactionMapper{
		"flypayA": func() Mappers.ITransactionMapper {
			return Mappers.NewFlyAMapper(Repositories.NewFlyARepository())
		},
		"flypayB": func() Mappers.ITransactionMapper {
			return Mappers.NewFlyBMapper(Repositories.NewFlyBRepository())
		},
	}
	return allProviders
}
