package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/igorscandido/go-transactions-api/internal/domain"
	"github.com/igorscandido/go-transactions-api/internal/ports"
)

var (
	ErrorGatewayNotFound = errors.New("gateway not found")
)

type paymentService struct {
	gateways   map[domain.Gateway]ports.Gateway
	repository ports.PaymentRepository
}

func NewPaymentService(gateways map[domain.Gateway]ports.Gateway, repository ports.PaymentRepository) ports.PaymentService {
	return &paymentService{gateways, repository}
}

func (s *paymentService) ProcessPaymentOnGateway(ctx context.Context, gateway domain.Gateway, payment *domain.Payment) (*string, error) {
	paymentGateway, ok := s.gateways[gateway]
	if !ok {
		return nil, ErrorGatewayNotFound
	}

	return paymentGateway.ProcessPayment(ctx, payment)
}

func (s *paymentService) GetPaymentStatusFromGateway(ctx context.Context, gateway domain.Gateway, paymentId *string) (*string, error) {
	paymentGateway, ok := s.gateways[gateway]
	if !ok {
		return nil, ErrorGatewayNotFound
	}

	paymentStatus, err := paymentGateway.GetPaymentStatus(ctx, paymentId)
	if err != nil {
		return nil, err
	}
	if paymentStatus == nil {
		return nil, fmt.Errorf("unable to get payment status for payment: [%s]", *paymentId)
	}

	return paymentStatus, nil
}
