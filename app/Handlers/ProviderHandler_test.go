package Handlers

import (
	"github.com/fly-code-challenge/app/Errors"
	"github.com/fly-code-challenge/app/Mappers"
	"github.com/fly-code-challenge/app/Repositories"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProviderHandler_GetProvidersWithNoFilters(t *testing.T) {
	providerHandler := NewProviderHandler()

	actualResult, _ := providerHandler.GetProvidersByFilter("")
	assert.Equal(t, len(actualResult), len(providerHandler.GetAllProviders()))
}

func TestProviderHandler_GetProvidersWithFilter(t *testing.T) {
	providerHandler := NewProviderHandler()

	actualResult, _ := providerHandler.GetProvidersByFilter("flypayA")
	expected := map[string]func() Mappers.ITransactionMapper{
		"flypayA": func() Mappers.ITransactionMapper {
			return Mappers.NewFlyAMapper(Repositories.NewFlyARepository())
		},
	}
	assert.Equal(t, len(actualResult), len(expected))
}

func TestProviderHandler_GetProvidersWithInvalidProvider(t *testing.T) {
	providerHandler := NewProviderHandler()

	_, err := providerHandler.GetProvidersByFilter("FlyPayC")
	expected := Errors.ValidationError{"Invalid provider"}
	assert.Equal(t, err, expected)
}


// Benchmarks
func BenchmarkProviderHandler_GetProvidersByFilter(b *testing.B) {
	providerHandler := NewProviderHandler()

	for i := 0; i < b.N; i++ {
		_,_=providerHandler.GetProvidersByFilter("")
	}
}

func BenchmarkProviderHandler_GetAllProviders(b *testing.B) {
	providerHandler := NewProviderHandler()

	for i := 0; i < b.N; i++ {
		_=providerHandler.GetAllProviders()
	}
}