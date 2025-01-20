package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/igorscandido/go-transactions-api/internal/domain"
	"github.com/igorscandido/go-transactions-api/internal/infra/http/dto"
	"github.com/igorscandido/go-transactions-api/internal/ports"
	"github.com/igorscandido/go-transactions-api/pkg/configs"
)

type PaymentHandler struct {
	paymentService  ports.PaymentService
	currencyService ports.CurrencyService
}

func NewPaymentHandler(paymentService ports.PaymentService, currencyService ports.CurrencyService) *PaymentHandler {
	return &PaymentHandler{paymentService, currencyService}
}

func (h *PaymentHandler) CreatePayment(c *gin.Context) {
	var paymentRequest dto.CreatePaymentRequest
	if err := c.ShouldBindJSON(&paymentRequest); err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultResponse{
			Status:  "error",
			Message: fmt.Sprintf("Invalid request body: %s", err.Error()),
		})
		return
	}

	conversionRate, err := h.currencyService.GetConversionRate(
		c.Request.Context(),
		configs.AppTransactionsBaseCurrency,
		paymentRequest.Currency,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultResponse{
			Status:  "error",
			Message: fmt.Sprintf("Error converting currency: %s", err.Error()),
		})
		return
	}

	paymentDomain, err := paymentRequest.ToDomain(conversionRate, configs.AppTransactionsBaseCurrency)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultResponse{
			Status:  "error",
			Message: fmt.Sprintf("Unable to parse request body: %s", err.Error()),
		})
		return
	}

	paymentId, err := h.paymentService.ProcessPaymentOnGateway(
		c.Request.Context(),
		domain.Gateway(paymentDomain.Gateway),
		paymentDomain,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultResponse{
			Status:  "error",
			Message: fmt.Sprintf("Error processing request: %s", err.Error()),
		})
		return
	}

	paymentDomain.ID = paymentId
	err = h.paymentService.SavePayment(c.Request.Context(), paymentDomain)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultResponse{
			Status:  "error",
			Message: fmt.Sprintf("Error trying to save payment: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, dto.DefaultResponse{
		Status:  "success",
		Message: fmt.Sprintf("Payment processed successfully. Payment ID: %s", *paymentId),
	})
}

func (h *PaymentHandler) GetPaymentStatus(c *gin.Context) {
	paymentID := c.Param("id")
	if paymentID == "" {
		c.JSON(http.StatusBadRequest, dto.DefaultResponse{
			Status:  "error",
			Message: "The id string parameter was not provided",
		})
		return
	}

	paymentDomain, err := h.paymentService.FetchPayment(c.Request.Context(), &paymentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultResponse{
			Status:  "error",
			Message: fmt.Sprintf("Error trying to fetch payment: %s", err.Error()),
		})
		return
	}

	paymentStatus, err := h.paymentService.GetPaymentStatusFromGateway(
		c.Request.Context(),
		domain.Gateway(paymentDomain.Gateway),
		&paymentID,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultResponse{
			Status:  "error",
			Message: fmt.Sprintf("Error trying to fetch payment status: %s", err.Error()),
		})
		return
	}

	var dtoPaymentStatus dto.GetPaymentStatusResponse
	dtoPaymentStatus.FromDomain(paymentDomain, paymentStatus)
	c.JSON(http.StatusOK, dtoPaymentStatus)
}
