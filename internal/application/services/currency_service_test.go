package services

import (
	"context"
	"errors"
	"testing"

	"github.com/igorscandido/go-transactions-api/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCurrencyService_GetConversionRate(t *testing.T) {
	mockRepository := new(mocks.CurrencyRepository)
	service := NewCurrencyService(mockRepository)

	ctx := context.Background()
	baseCurrency := "USD"
	destCurrency := "EUR"
	conversionRate := 0.85
	mockRepository.On("GetConversionRateForCurrency", ctx, baseCurrency, destCurrency).Return(&conversionRate, nil)

	result, err := service.GetConversionRate(ctx, baseCurrency, destCurrency)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, &conversionRate, result)
	mockRepository.AssertExpectations(t)
}

func TestCurrencyService_GetConversionRate_Error(t *testing.T) {
	mockRepository := new(mocks.CurrencyRepository)
	service := NewCurrencyService(mockRepository)

	ctx := context.Background()
	baseCurrency := "USD"
	destCurrency := "EUR"
	mockError := errors.New("conversion rate not found")
	mockRepository.On("GetConversionRateForCurrency", ctx, baseCurrency, destCurrency).Return(nil, mockError)

	result, err := service.GetConversionRate(ctx, baseCurrency, destCurrency)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, mockError, err)
	mockRepository.AssertExpectations(t)
}
