package ports

import (
	"context"

	"github.com/igorscandido/go-transactions-api/internal/domain"
)

type Gateway interface {
	ProcessPayment(ctx context.Context, payment *domain.Payment) (*string, error)
	GetPaymentStatus(ctx context.Context, paymentId *string) (*string, error)
}
