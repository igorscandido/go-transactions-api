package queries

const (
	GetPaymentByID = `
		SELECT id, paymentBaseCurrency, paymentBaseAmount, gateway, gatewayTransactionatedCurrency, 
			gatewayTransactionatedAmount, currencyConversionRate, creationRequestTime
		FROM payment
		WHERE id = $1
	`
	InsertPayment = `
		INSERT INTO payment (id, paymentBaseCurrency, paymentBaseAmount, gateway, gatewayTransactionatedCurrency,
			gatewayTransactionatedAmount, currencyConversionRate, creationRequestTime)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
)
