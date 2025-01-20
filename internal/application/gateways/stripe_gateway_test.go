package gateways

import (
	"context"
	"errors"
	"testing"

	"github.com/igorscandido/go-transactions-api/internal/domain"
	"github.com/igorscandido/go-transactions-api/mocks"
	"github.com/stretchr/testify/assert"
)

func TestStripeGateway_ProcessPayment(t *testing.T) {
	mockClient := new(mocks.StripeClient)
	gateway := NewStripeGateway(mockClient)
	ctx := context.Background()
	payment := &domain.Payment{
		CardDetails: &domain.CardDetails{
			Number:       "4242424242424242",
			CVV:          "123",
			ExpiryNumber: 12,
			ExpiryYear:   2030,
		},
		GatewayTransactionatedCurrency: "USD",
		GatewayTransactionatedAmount:   &[]float64{1000}[0],
	}
	mockMethodId := "mock-method-id"
	mockPaymentId := "mock-payment-id"
	mockClient.On("CreatePaymentMethod", ctx, payment.CardDetails).Return(&mockMethodId, nil)
	mockClient.On("ProcessPayment", ctx, payment, &mockMethodId).Return(&mockPaymentId, nil)
	result, err := gateway.ProcessPayment(ctx, payment)
	assert.NoError(t, err)
	assert.Equal(t, &mockPaymentId, result)
	mockClient.AssertExpectations(t)
}

func TestStripeGateway_ProcessPayment_ErrorInCreatePaymentMethod(t *testing.T) {
	mockClient := new(mocks.StripeClient)
	gateway := NewStripeGateway(mockClient)
	ctx := context.Background()
	payment := &domain.Payment{
		CardDetails: &domain.CardDetails{
			Number:       "4242424242424242",
			CVV:          "123",
			ExpiryNumber: 12,
			ExpiryYear:   2030,
		},
		GatewayTransactionatedCurrency: "USD",
		GatewayTransactionatedAmount:   &[]float64{1000}[0],
	}
	mockError := errors.New("failed to create payment method")
	mockClient.On("CreatePaymentMethod", ctx, payment.CardDetails).Return(&[]string{""}[0], mockError)
	result, err := gateway.ProcessPayment(ctx, payment)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, mockError, err)
	mockClient.AssertExpectations(t)
}

func TestStripeGateway_GetPaymentStatus(t *testing.T) {
	mockClient := new(mocks.StripeClient)
	gateway := NewStripeGateway(mockClient)
	ctx := context.Background()
	mockPaymentId := "mock-payment-id"
	mockStatus := "success"
	mockClient.On("GetPaymentStatus", ctx, &mockPaymentId).Return(&mockStatus, nil)
	result, err := gateway.GetPaymentStatus(ctx, &mockPaymentId)
	assert.NoError(t, err)
	assert.Equal(t, &mockStatus, result)
	mockClient.AssertExpectations(t)
}

func TestStripeGateway_GetPaymentStatus_Error(t *testing.T) {
	mockClient := new(mocks.StripeClient)
	gateway := NewStripeGateway(mockClient)
	ctx := context.Background()
	mockPaymentId := "mock-payment-id"
	mockError := errors.New("payment not found")
	mockClient.On("GetPaymentStatus", ctx, &mockPaymentId).Return(nil, mockError)
	result, err := gateway.GetPaymentStatus(ctx, &mockPaymentId)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, mockError, err)
	mockClient.AssertExpectations(t)
}
