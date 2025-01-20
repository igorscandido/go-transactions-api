package dto

import (
	"strconv"
	"time"

	"github.com/igorscandido/go-transactions-api/internal/domain"
)

type CardDetails struct {
	Number string `json:"number"`
	Expiry string `json:"expiry"`
	CVV    string `json:"cvv"`
}

type CreatePaymentRequest struct {
	Gateway       string      `json:"gateway"`
	Amount        float64     `json:"amount"`
	Currency      string      `json:"currency"`
	PaymentMethod string      `json:"payment_method"`
	CardDetails   CardDetails `json:"card_details"`
}

func (p *CreatePaymentRequest) ToDomain(conversionRate *float64, baseCurrency string) (*domain.Payment, error) {
	expiryYear, err := strconv.Atoi(p.CardDetails.Expiry[3:])
	if err != nil {
		return nil, err
	}

	expiryMonth, err := strconv.Atoi(p.CardDetails.Expiry[:2])
	if err != nil {
		return nil, err
	}

	convertedAmount := p.Amount / (*conversionRate)
	creationRequestTime := time.Now()
	return &domain.Payment{
		PaymentBaseCurrency:            p.Currency,
		PaymentBaseAmount:              &p.Amount,
		Gateway:                        p.Gateway,
		GatewayTransactionatedCurrency: baseCurrency,
		GatewayTransactionatedAmount:   &convertedAmount,
		CurrencyConversionRate:         conversionRate,
		CardDetails: &domain.CardDetails{
			Number:       p.CardDetails.Number,
			CVV:          p.CardDetails.CVV,
			ExpiryNumber: expiryMonth,
			ExpiryYear:   expiryYear,
		},
		CreationRequestTime: &creationRequestTime,
	}, nil
}

type GetPaymentStatusResponse struct {
	ID                             string    `json:"id"`
	Status                         string    `json:"status"`
	PaymentBaseCurrency            string    `json:"payment_base_currency"`
	PaymentBaseAmount              float64   `json:"payment_base_amount"`
	Gateway                        string    `json:"gateway"`
	GatewayTransactionatedCurrency string    `json:"gateway_transactionated_currency"`
	GatewayTransactionatedAmount   float64   `json:"gateway_transactionated_amount"`
	CurrencyConversionRate         float64   `json:"currency_conversion_rate"`
	CreationRequestTime            time.Time `json:"creation_request_time"`
}

func (d *GetPaymentStatusResponse) FromDomain(paymentDomain *domain.Payment, status *string) {
	d.ID = *paymentDomain.ID
	d.Status = *status
	d.PaymentBaseCurrency = paymentDomain.PaymentBaseCurrency
	d.PaymentBaseAmount = *paymentDomain.PaymentBaseAmount
	d.Gateway = paymentDomain.Gateway
	d.GatewayTransactionatedCurrency = paymentDomain.GatewayTransactionatedCurrency
	d.GatewayTransactionatedAmount = *paymentDomain.GatewayTransactionatedAmount
	d.CurrencyConversionRate = *paymentDomain.CurrencyConversionRate
	d.CreationRequestTime = *paymentDomain.CreationRequestTime
}
