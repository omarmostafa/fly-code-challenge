package Handlers

import (
	"github.com/code-challenge/app/Errors"
	"github.com/code-challenge/app/Repositories"
)

type ProviderHandler struct {
}

func NewProviderHandler() *ProviderHandler {
	return &ProviderHandler{}
}

func (handler *ProviderHandler) GetProvidersByFilter(provider string) (map[string]func() Repositories.ITransactionRepository, error) {
	allProviders := handler.GetAllProviders()
	if provider == "" {
		return allProviders, nil
	}
	filteredProvider := allProviders[provider]
	if filteredProvider == nil {
		return nil, Errors.ValidationError{Message: "Invalid provider"}
	}
	return map[string]func() Repositories.ITransactionRepository{
		provider: filteredProvider,
	}, nil
}

func (handler *ProviderHandler) GetAllProviders() map[string]func() Repositories.ITransactionRepository {
	allProviders := map[string]func() Repositories.ITransactionRepository{
		"flypayA": func() Repositories.ITransactionRepository {
			return Repositories.NewFlyARepository()
		},
		"flypayB": func() Repositories.ITransactionRepository {
			return Repositories.NewFlyBRepository()
		},
	}
	return allProviders
}
