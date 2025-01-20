package ports

import (
	"context"

	"github.com/igorscandido/go-transactions-api/internal/domain"
)

type PaymentService interface {
	ProcessPaymentOnGateway(ctx context.Context, gateway domain.Gateway, payment *domain.Payment) (*string, error)
	GetPaymentStatusFromGateway(ctx context.Context, gateway domain.Gateway, paymentId *string) (*string, error)
	FetchPayment(ctx context.Context, paymentId *string) (*domain.Payment, error)
	SavePayment(ctx context.Context, payment *domain.Payment) error
}

type CurrencyService interface {
	GetConversionRate(ctx context.Context, baseCurrency string, destCurrency string) (*float64, error)
}
