package router

import (
	"github.com/gin-gonic/gin"
	"github.com/igorscandido/go-transactions-api/internal/infra/http"
)

func NewRouter(paymentHandler *http.PaymentHandler, currencyHandler *http.CurrencyHandler) *gin.Engine {
	router := gin.Default()

	router.POST("/payments", paymentHandler.CreatePayment)
	router.GET("/payments/:id", paymentHandler.GetPaymentStatus)

	router.POST("/currency/convert", currencyHandler.ConvertCurrency)

	return router
}
