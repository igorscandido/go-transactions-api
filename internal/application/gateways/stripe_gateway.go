package gateways

import (
	"context"

	"github.com/igorscandido/go-transactions-api/internal/domain"
	"github.com/igorscandido/go-transactions-api/internal/ports"
)

type stripeGateway struct {
	client ports.StripeClient
}

func NewStripeGateway(client ports.StripeClient) ports.Gateway {
	return &stripeGateway{client}
}

func (g *stripeGateway) ProcessPayment(ctx context.Context, payment *domain.Payment) (*string, error) {
	methodId, err := g.client.CreatePaymentMethod(ctx, payment.CardDetails)
	if err != nil {
		return nil, err
	}

	return g.client.ProcessPayment(ctx, payment, methodId)
}

func (g *stripeGateway) GetPaymentStatus(ctx context.Context, paymentId *string) (*string, error) {
	return g.client.GetPaymentStatus(ctx, paymentId)
}
