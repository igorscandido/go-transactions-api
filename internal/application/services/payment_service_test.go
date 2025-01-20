package services

import (
	"context"
	"errors"
	"testing"

	"github.com/igorscandido/go-transactions-api/internal/domain"
	"github.com/igorscandido/go-transactions-api/internal/ports"
	"github.com/igorscandido/go-transactions-api/mocks"
	"github.com/stretchr/testify/assert"
)

func TestPaymentService_ProcessPaymentOnGateway(t *testing.T) {
	mockGateway := new(mocks.Gateway)
	mockRepository := new(mocks.PaymentRepository)
	service := NewPaymentService(map[domain.Gateway]ports.Gateway{"gateway": mockGateway}, mockRepository)

	ctx := context.Background()
	payment := &domain.Payment{}
	mockPaymentID := "payment-id"
	mockGateway.On("ProcessPayment", ctx, payment).Return(&mockPaymentID, nil)

	result, err := service.ProcessPaymentOnGateway(ctx, "gateway", payment)
	assert.NoError(t, err)
	assert.Equal(t, &mockPaymentID, result)
	mockGateway.AssertExpectations(t)
}

func TestPaymentService_ProcessPaymentOnGateway_ErrorGatewayNotFound(t *testing.T) {
	mockRepository := new(mocks.PaymentRepository)
	service := NewPaymentService(map[domain.Gateway]ports.Gateway{}, mockRepository)

	ctx := context.Background()
	payment := &domain.Payment{}

	result, err := service.ProcessPaymentOnGateway(ctx, "gateway", payment)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, ErrorGatewayNotFound, err)
}

func TestPaymentService_GetPaymentStatusFromGateway(t *testing.T) {
	mockGateway := new(mocks.Gateway)
	mockRepository := new(mocks.PaymentRepository)
	service := NewPaymentService(map[domain.Gateway]ports.Gateway{"gateway": mockGateway}, mockRepository)

	ctx := context.Background()
	paymentID := "payment-id"
	status := "success"
	mockGateway.On("GetPaymentStatus", ctx, &paymentID).Return(&status, nil)

	result, err := service.GetPaymentStatusFromGateway(ctx, "gateway", &paymentID)
	assert.NoError(t, err)
	assert.Equal(t, &status, result)
	mockGateway.AssertExpectations(t)
}

func TestPaymentService_GetPaymentStatusFromGateway_ErrorGatewayNotFound(t *testing.T) {
	mockRepository := new(mocks.PaymentRepository)
	service := NewPaymentService(map[domain.Gateway]ports.Gateway{}, mockRepository)

	ctx := context.Background()
	paymentID := "payment-id"

	result, err := service.GetPaymentStatusFromGateway(ctx, "gateway", &paymentID)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, ErrorGatewayNotFound, err)
}

func TestPaymentService_GetPaymentStatusFromGateway_ErrorInGateway(t *testing.T) {
	mockGateway := new(mocks.Gateway)
	mockRepository := new(mocks.PaymentRepository)
	service := NewPaymentService(map[domain.Gateway]ports.Gateway{"gateway": mockGateway}, mockRepository)

	ctx := context.Background()
	paymentID := "payment-id"
	mockError := errors.New("gateway error")
	mockGateway.On("GetPaymentStatus", ctx, &paymentID).Return(nil, mockError)

	result, err := service.GetPaymentStatusFromGateway(ctx, "gateway", &paymentID)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, mockError, err)
}

func TestPaymentService_FetchPayment(t *testing.T) {
	mockRepository := new(mocks.PaymentRepository)
	service := NewPaymentService(nil, mockRepository)

	ctx := context.Background()
	paymentID := "payment-id"
	expectedPayment := &domain.Payment{ID: &paymentID}
	mockRepository.On("GetByID", ctx, paymentID).Return(expectedPayment, nil)

	result, err := service.FetchPayment(ctx, &paymentID)
	assert.NoError(t, err)
	assert.Equal(t, expectedPayment, result)
	mockRepository.AssertExpectations(t)
}

func TestPaymentService_SavePayment(t *testing.T) {
	mockRepository := new(mocks.PaymentRepository)
	service := NewPaymentService(nil, mockRepository)

	ctx := context.Background()
	payment := &domain.Payment{}
	mockRepository.On("Create", ctx, payment).Return(nil)

	err := service.SavePayment(ctx, payment)
	assert.NoError(t, err)
	mockRepository.AssertExpectations(t)
}
