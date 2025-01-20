package main

import (
	"fmt"

	"github.com/igorscandido/go-transactions-api/internal/application/gateways"
	"github.com/igorscandido/go-transactions-api/internal/application/services"
	"github.com/igorscandido/go-transactions-api/internal/domain"
	"github.com/igorscandido/go-transactions-api/internal/infra/http"
	"github.com/igorscandido/go-transactions-api/internal/infra/http/router"
	"github.com/igorscandido/go-transactions-api/internal/infra/repositories"
	"github.com/igorscandido/go-transactions-api/internal/ports"
	"github.com/igorscandido/go-transactions-api/pkg/cache"
	"github.com/igorscandido/go-transactions-api/pkg/client"
	"github.com/igorscandido/go-transactions-api/pkg/configs"
	"github.com/igorscandido/go-transactions-api/pkg/database"
)

func main() {
	configs := configs.NewConfigs()

	postgresAdapter, err := database.NewPostgresAdapter(configs)
	if err != nil {
		panic(err)
	}

	redisClient, err := cache.NewRedisCache(configs)
	if err != nil {
		panic(err)
	}

	ratesClient := client.NewOpenExchangeRatesClient(configs)
	stripeClient := client.NewStripePaymentClient(configs)
	stripeGateway := gateways.NewStripeGateway(stripeClient)

	currencyRepository := repositories.NewCurrencyRepository(ratesClient, redisClient, configs)
	currencyService := services.NewCurrencyService(currencyRepository)
	paymentRepository := repositories.NewPaymentRepository(postgresAdapter, redisClient)
	paymentService := services.NewPaymentService(map[domain.Gateway]ports.Gateway{
		domain.Stripe: stripeGateway,
	}, paymentRepository)
	paymentHandler := http.NewPaymentHandler(paymentService, currencyService)
	currencyHandler := http.NewCurrencyHandler(currencyService)

	router := router.NewRouter(
		paymentHandler,
		currencyHandler,
	)
	router.Run(fmt.Sprintf(":%d", configs.Server.Port))
}
