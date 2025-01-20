package services

import (
	"context"

	"github.com/igorscandido/go-transactions-api/internal/ports"
)

type currencyService struct {
	repository ports.CurrencyRepository
}

func NewCurrencyService(repository ports.CurrencyRepository) ports.CurrencyService {
	return &currencyService{repository}
}

func (s *currencyService) GetConversionRate(ctx context.Context, baseCurrency string, destCurrency string) (*float64, error) {
	return s.repository.GetConversionRateForCurrency(ctx, baseCurrency, destCurrency)
}
