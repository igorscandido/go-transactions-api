package ports

import (
	"context"

	"github.com/igorscandido/go-transactions-api/internal/domain"
)

type PaymentRepository interface {
	Create(ctx context.Context, payment *domain.Payment) error
	GetByID(ctx context.Context, id string) (*domain.Payment, error)
}

type CurrencyRepository interface {
	GetConversionRateForCurrency(ctx context.Context, baseCurrency string, destCurrency string) (*float64, error)
}
