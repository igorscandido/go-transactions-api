package repositories

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/igorscandido/go-transactions-api/internal/domain"
	"github.com/igorscandido/go-transactions-api/mocks"
	"github.com/igorscandido/go-transactions-api/pkg/configs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCurrencyRepository_GetConversionRateForCurrency_CacheHit(t *testing.T) {
	mockRatesClient := new(mocks.RatesClient)
	mockCache := new(mocks.Cache)
	configs := &configs.Configs{
		Currency: configs.Currency{
			CacheRates:      true,
			CacheTTLSeconds: 60,
		},
	}
	repo := NewCurrencyRepository(mockRatesClient, mockCache, configs)

	baseCurrency := "USD"
	destCurrency := "EUR"
	expectedRate := 0.85

	mockCache.On("Get", mock.Anything, fmt.Sprintf(currencyRateInBasedCacheKey, baseCurrency, destCurrency)).
		Return(expectedRate, true)

	rate, err := repo.GetConversionRateForCurrency(context.Background(), baseCurrency, destCurrency)

	assert.NoError(t, err)
	assert.Equal(t, expectedRate, *rate)
	mockCache.AssertExpectations(t)
	mockRatesClient.AssertExpectations(t)
}

func TestCurrencyRepository_GetConversionRateForCurrency_CacheMiss_FetchFromAPI(t *testing.T) {
	mockRatesClient := new(mocks.RatesClient)
	mockCache := new(mocks.Cache)
	configs := &configs.Configs{
		Currency: configs.Currency{
			CacheRates:      true,
			CacheTTLSeconds: 60,
		},
	}
	repo := NewCurrencyRepository(mockRatesClient, mockCache, configs)

	baseCurrency := "USD"
	destCurrency := "EUR"
	expectedRate := 0.85
	currencyRates := &domain.CurrencyRates{
		Rates: map[string]float64{
			destCurrency: expectedRate,
		},
	}

	mockCache.On("Get", mock.Anything, fmt.Sprintf(currencyRateInBasedCacheKey, baseCurrency, destCurrency)).
		Return(nil, false)

	mockCache.On("Get", mock.Anything, fmt.Sprintf(latestsRatesBaseCurrencyCacheKey, baseCurrency)).
		Return(nil, false)

	mockRatesClient.On("GetLatestRatesBasedOn", mock.Anything, baseCurrency).
		Return(currencyRates, nil)

	mockCache.On("Set", mock.Anything, fmt.Sprintf(currencyRateInBasedCacheKey, baseCurrency, destCurrency), expectedRate, configs.Currency.CacheTTLSeconds).
		Return(nil)

	mockCache.On("Set", mock.Anything, fmt.Sprintf(latestsRatesBaseCurrencyCacheKey, baseCurrency), *currencyRates, configs.Currency.CacheTTLSeconds).
		Return(nil)

	rate, err := repo.GetConversionRateForCurrency(context.Background(), baseCurrency, destCurrency)

	assert.NoError(t, err)
	assert.Equal(t, expectedRate, *rate)
	mockCache.AssertExpectations(t)
	mockRatesClient.AssertExpectations(t)
}

func TestCurrencyRepository_GetConversionRateForCurrency_APIError(t *testing.T) {
	mockRatesClient := new(mocks.RatesClient)
	mockCache := new(mocks.Cache)
	configs := &configs.Configs{
		Currency: configs.Currency{
			CacheRates:      true,
			CacheTTLSeconds: 60,
		},
	}
	repo := NewCurrencyRepository(mockRatesClient, mockCache, configs)

	baseCurrency := "USD"
	destCurrency := "EUR"

	mockCache.On("Get", mock.Anything, fmt.Sprintf(currencyRateInBasedCacheKey, baseCurrency, destCurrency)).
		Return(nil, false)

	mockCache.On("Get", mock.Anything, fmt.Sprintf(latestsRatesBaseCurrencyCacheKey, baseCurrency)).
		Return(nil, false)

	mockRatesClient.On("GetLatestRatesBasedOn", mock.Anything, baseCurrency).
		Return(nil, errors.New("failed to fetch rates"))

	rate, err := repo.GetConversionRateForCurrency(context.Background(), baseCurrency, destCurrency)

	assert.Error(t, err)
	assert.Nil(t, rate)
	mockCache.AssertExpectations(t)
	mockRatesClient.AssertExpectations(t)
}

func TestCurrencyRepository_GetConversionRateForCurrency_DestinationCurrencyNotFound(t *testing.T) {
	mockRatesClient := new(mocks.RatesClient)
	mockCache := new(mocks.Cache)
	configs := &configs.Configs{
		Currency: configs.Currency{
			CacheRates:      true,
			CacheTTLSeconds: 60,
		},
	}
	repo := NewCurrencyRepository(mockRatesClient, mockCache, configs)

	baseCurrency := "USD"
	destCurrency := "AUD"

	currencyRates := &domain.CurrencyRates{
		Rates: map[string]float64{
			"EUR": 0.85,
		},
	}

	mockCache.On("Get", mock.Anything, fmt.Sprintf(currencyRateInBasedCacheKey, baseCurrency, destCurrency)).
		Return(nil, false)

	mockCache.On("Get", mock.Anything, fmt.Sprintf(latestsRatesBaseCurrencyCacheKey, baseCurrency)).
		Return(nil, false)

	mockRatesClient.On("GetLatestRatesBasedOn", mock.Anything, baseCurrency).
		Return(currencyRates, nil)

	mockCache.On("Set", mock.Anything, fmt.Sprintf(latestsRatesBaseCurrencyCacheKey, baseCurrency), *currencyRates, configs.Currency.CacheTTLSeconds).
		Return(nil)

	rate, err := repo.GetConversionRateForCurrency(context.Background(), baseCurrency, destCurrency)

	assert.Error(t, err)
	assert.Nil(t, rate)
	mockCache.AssertExpectations(t)
	mockRatesClient.AssertExpectations(t)
}
