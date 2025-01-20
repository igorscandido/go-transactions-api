package domain

type CurrencyRates struct {
	Base  string             `json:"base"`
	Rates map[string]float64 `json:"rates"`
}
