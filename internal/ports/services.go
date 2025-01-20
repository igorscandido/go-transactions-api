package ports

import (
	"context"

	"github.com/igorscandido/go-transactions-api/internal/domain"
)

type PaymentService interface {
	ProcessPaymentOnGateway(ctx context.Context, gateway domain.Gateway, payment *domain.Payment) (*string, error)
	GetPaymentStatusFromGateway(ctx context.Context, gateway domain.Gateway, paymentId *string) (*string, error)
}
