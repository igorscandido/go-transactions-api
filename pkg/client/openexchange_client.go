package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/igorscandido/go-transactions-api/internal/domain"
	"github.com/igorscandido/go-transactions-api/internal/ports"
	"github.com/igorscandido/go-transactions-api/pkg/configs"
)

type RatesResult struct {
	Rates map[string]float64 `json:"rates"`
}

type openExchangeRatesClient struct {
	apiKey string
}

func NewOpenExchangeRatesClient(configs *configs.Configs) ports.RatesClient {
	return &openExchangeRatesClient{configs.ExternalKeys.OpenExchangeRates}
}

func (c *openExchangeRatesClient) GetLatestRatesBasedOn(ctx context.Context, baseCurrency string) (*domain.CurrencyRates, error) {
	url := fmt.Sprintf("https://openexchangerates.org/api/latest.json?app_id=%s&base=%s", c.apiKey, baseCurrency)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request to openexchangerates: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request to openexchangerates: %w", err)
	}
	defer resp.Body.Close()

	var data RatesResult
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to parse request result from openexchangerates: %w", err)
	}

	return &domain.CurrencyRates{
		Base:  baseCurrency,
		Rates: data.Rates,
	}, nil
}
