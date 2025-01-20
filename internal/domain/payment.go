package domain

import "time"

type CardDetails struct {
	Number       string
	CVV          string
	ExpiryNumber int
	ExpiryYear   int
}

type Payment struct {
	ID                             *string
	PaymentBaseCurrency            string
	PaymentBaseAmount              *float64
	Gateway                        string
	GatewayTransactionatedCurrency string
	GatewayTransactionatedAmount   *float64
	CurrencyConversionRate         *float64
	CreationRequestTime            *time.Time
	CardDetails                    *CardDetails
}
