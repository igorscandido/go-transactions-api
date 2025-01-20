package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/igorscandido/go-transactions-api/internal/infra/http/dto"
	"github.com/igorscandido/go-transactions-api/internal/ports"
)

type CurrencyHandler struct {
	currencyService ports.CurrencyService
}

func NewCurrencyHandler(currencyService ports.CurrencyService) *CurrencyHandler {
	return &CurrencyHandler{currencyService}
}

func (h *CurrencyHandler) ConvertCurrency(c *gin.Context) {
	var conversionRequest dto.CurrencyConversionRequest
	if err := c.ShouldBindJSON(&conversionRequest); err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultResponse{
			Status:  "error",
			Message: fmt.Sprintf("Invalid request body: %s", err.Error()),
		})
		return
	}

	conversionRate, err := h.currencyService.GetConversionRate(
		c.Request.Context(),
		conversionRequest.FromCurrency,
		conversionRequest.ToCurrency,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultResponse{
			Status:  "error",
			Message: fmt.Sprintf("Error converting currency: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, dto.CurrencyConversionResponse{
		FromCurrency:    conversionRequest.FromCurrency,
		ToCurrency:      conversionRequest.ToCurrency,
		OriginalAmount:  conversionRequest.Amount,
		ConvertedAmount: conversionRequest.Amount * (*conversionRate),
	})
}
