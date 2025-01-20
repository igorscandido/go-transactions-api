package repositories

import (
	"context"
	"testing"
	"time"

	"github.com/igorscandido/go-transactions-api/internal/domain"
	"github.com/igorscandido/go-transactions-api/internal/infra/repositories/queries"
	"github.com/igorscandido/go-transactions-api/mocks"
	"github.com/stretchr/testify/assert"
)

func TestPaymentRepository_Create(t *testing.T) {
	mockDB := new(mocks.Database)
	mockCache := new(mocks.Cache)
	repo := NewPaymentRepository(mockDB, mockCache)

	ctx := context.Background()
	paymentID := "payment-id"
	payment := &domain.Payment{
		ID:                             &paymentID,
		PaymentBaseCurrency:            "USD",
		PaymentBaseAmount:              float64Pointer(100.50),
		Gateway:                        "gateway",
		GatewayTransactionatedCurrency: "EUR",
		GatewayTransactionatedAmount:   float64Pointer(90.25),
		CurrencyConversionRate:         float64Pointer(1.11),
		CreationRequestTime:            timePointer(time.Now()),
	}

	mockDB.On("Exec", ctx, queries.InsertPayment,
		paymentID,
		payment.PaymentBaseCurrency,
		*payment.PaymentBaseAmount,
		payment.Gateway,
		payment.GatewayTransactionatedCurrency,
		*payment.GatewayTransactionatedAmount,
		*payment.CurrencyConversionRate,
		*payment.CreationRequestTime,
	).Return(nil, nil)

	err := repo.Create(ctx, payment)

	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func float64Pointer(f float64) *float64 {
	return &f
}

func timePointer(t time.Time) *time.Time {
	return &t
}
