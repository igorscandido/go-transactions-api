package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/igorscandido/go-transactions-api/internal/domain"
	"github.com/igorscandido/go-transactions-api/internal/infra/repositories/queries"
	"github.com/igorscandido/go-transactions-api/internal/ports"
)

const (
	getByIdCacheKey = "payment_%s"
)

type paymentRepository struct {
	database ports.Database
	cache    ports.Cache
}

func NewPaymentRepository(database ports.Database, cache ports.Cache) ports.PaymentRepository {
	return &paymentRepository{database, cache}
}

func (p *paymentRepository) Create(ctx context.Context, payment *domain.Payment) error {
	_, err := p.database.Exec(ctx, queries.InsertPayment,
		*payment.ID,
		payment.PaymentBaseCurrency,
		*payment.PaymentBaseAmount,
		payment.Gateway,
		payment.GatewayTransactionatedCurrency,
		*payment.GatewayTransactionatedAmount,
		*payment.CurrencyConversionRate,
		*payment.CreationRequestTime,
	)
	return err
}

func (p *paymentRepository) GetByID(ctx context.Context, id string) (*domain.Payment, error) {
	var payment domain.Payment
	if val, ok := p.cache.Get(ctx, fmt.Sprintf(getByIdCacheKey, id)); ok {
		if payment, ok = val.(domain.Payment); ok {
			return &payment, nil
		}
	}

	row := p.database.QueryRow(ctx, queries.GetPaymentByID, id)
	if err := row.Scan(
		&payment.ID,
		&payment.PaymentBaseCurrency,
		&payment.PaymentBaseAmount,
		&payment.Gateway,
		&payment.GatewayTransactionatedCurrency,
		&payment.GatewayTransactionatedAmount,
		&payment.CurrencyConversionRate,
		&payment.CreationRequestTime,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("payment not found: %w", err)
		}
		return nil, fmt.Errorf("failed to scan payment: %w", err)
	}

	return &payment, nil
}
