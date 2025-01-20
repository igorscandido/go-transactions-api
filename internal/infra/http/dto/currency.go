package dto

type CurrencyConversionRequest struct {
	Amount       float64 `json:"amount" binding:"required"`
	FromCurrency string  `json:"from_currency" binding:"required"`
	ToCurrency   string  `json:"to_currency" binding:"required"`
}

type CurrencyConversionResponse struct {
	FromCurrency    string  `json:"from_currency"`
	ToCurrency      string  `json:"to_currency"`
	OriginalAmount  float64 `json:"original_amount"`
	ConvertedAmount float64 `json:"converted_amount"`
}
