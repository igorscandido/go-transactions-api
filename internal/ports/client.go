package ports

import (
	"context"

	"github.com/igorscandido/go-transactions-api/internal/domain"
)

type Cache interface {
	Get(ctx context.Context, key string) (interface{}, bool)
	Set(ctx context.Context, key string, value interface{}, ttlSeconds int)
	Delete(ctx context.Context, key string)
}

type StripeClient interface {
	CreatePaymentMethod(ctx context.Context, cardDetails *domain.CardDetails) (*string, error)
	ProcessPayment(ctx context.Context, payment *domain.Payment, paymentMethodID *string) (*string, error)
	GetPaymentStatus(ctx context.Context, paymentID *string) (*string, error)
}

type RatesClient interface {
	GetLatestRatesBasedOn(ctx context.Context, baseCurrency string) (*domain.CurrencyRates, error)
}
