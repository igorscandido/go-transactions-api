package repositories

import (
	"context"
	"fmt"

	"github.com/igorscandido/go-transactions-api/internal/domain"
	"github.com/igorscandido/go-transactions-api/internal/ports"
	"github.com/igorscandido/go-transactions-api/pkg/configs"
)

const (
	currencyRateInBasedCacheKey      = "based_%s_to_%s_rate"
	latestsRatesBaseCurrencyCacheKey = "latests_rates_based_%s"
)

type currencyRepository struct {
	requestClient ports.RatesClient
	cacheClient   ports.Cache
	configs       *configs.Configs
}

func NewCurrencyRepository(requestClient ports.RatesClient, cacheClient ports.Cache, configs *configs.Configs) ports.CurrencyRepository {
	return &currencyRepository{requestClient, cacheClient, configs}
}

func (r *currencyRepository) GetConversionRateForCurrency(ctx context.Context, baseCurrency string, destCurrency string) (*float64, error) {
	var rate *float64
	configCurrency := r.configs.Currency

	if configCurrency.CacheRates {
		rate = r.getCachedRateFactor(ctx, baseCurrency, destCurrency)
	}

	if rate == nil {
		currencyRates := r.getCachedLatestsRatesForBaseCurrency(ctx, baseCurrency)
		if currencyRates == nil {
			var err error
			currencyRates, err = r.requestClient.GetLatestRatesBasedOn(ctx, baseCurrency)
			if err != nil {
				return nil, fmt.Errorf("failed to get currency rates from client: %w", err)
			}

			r.setLatestsRatesForBaseCurrencyCache(ctx, baseCurrency, currencyRates)
		}

		cRate, ok := currencyRates.Rates[destCurrency]
		if !ok {
			return nil, fmt.Errorf("destination currency [%s] not supported by API", destCurrency)
		}
		r.setRatesFactorCache(ctx, baseCurrency, destCurrency, cRate)

		rate = &cRate
	}

	return rate, nil
}

func (r *currencyRepository) getCachedRateFactor(ctx context.Context, baseCurrency string, destCurrency string) *float64 {
	if !r.configs.Currency.CacheRates {
		return nil
	}

	cacheVal, ok := r.cacheClient.Get(ctx, fmt.Sprintf(currencyRateInBasedCacheKey, baseCurrency, destCurrency))
	if ok {
		if val, ok := cacheVal.(float64); ok {
			return &val
		}
	}
	return nil
}

func (r *currencyRepository) setRatesFactorCache(ctx context.Context, baseCurrency string, destCurrency string, rate float64) {
	if r.configs.Currency.CacheRates {
		r.cacheClient.Set(
			ctx,
			fmt.Sprintf(currencyRateInBasedCacheKey, baseCurrency, destCurrency),
			rate,
			r.configs.Currency.CacheTTLSeconds)
	}
}

func (r *currencyRepository) getCachedLatestsRatesForBaseCurrency(ctx context.Context, baseCurrency string) *domain.CurrencyRates {
	if !r.configs.Currency.CacheRates {
		return nil
	}

	cacheVal, ok := r.cacheClient.Get(ctx, fmt.Sprintf(latestsRatesBaseCurrencyCacheKey, baseCurrency))
	if ok {
		if val, ok := cacheVal.(domain.CurrencyRates); ok {
			return &val
		}
	}
	return nil
}

func (r *currencyRepository) setLatestsRatesForBaseCurrencyCache(ctx context.Context, baseCurrency string, currencyRates *domain.CurrencyRates) {
	if r.configs.Currency.CacheRates {
		r.cacheClient.Set(
			ctx,
			fmt.Sprintf(latestsRatesBaseCurrencyCacheKey, baseCurrency),
			*currencyRates,
			r.configs.Currency.CacheTTLSeconds,
		)
	}
}
