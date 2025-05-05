package handlers

import (
	"bookstore/internal/model"
	"bookstore/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PaymentHandler struct {
	PaymentSvc service.PaymentService
}

func NewPaymentHandler(paymentSvc service.PaymentService) *PaymentHandler {
	return &PaymentHandler{PaymentSvc: paymentSvc}
}

func (i *PaymentHandler) ListAllPayments(c *gin.Context) {
	payments, err := i.PaymentSvc.ListAllPayment(c.Request.Context())
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, err.Error())
	}
	ResponseSuccess(c, payments)
}

func (i *PaymentHandler) SearchPayment(c *gin.Context) {
	var req = c.Param("order id")
	payment, err := i.PaymentSvc.GetPaymentByOrderId(c.Request.Context(), req)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseSuccess(c, payment)
}

func (i *PaymentHandler) CreatePayment(c *gin.Context) {
	var req model.Payment
	if err := c.ShouldBindJSON(&req); err != nil {
		ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	payment, err := i.PaymentSvc.CreatePayment(c.Request.Context(), req)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseSuccess(c, payment)
}

func (i *PaymentHandler) ConfirmPayment(c *gin.Context) {
	var orderId = c.Param("order id")
	var req model.Payment
	if err := c.ShouldBindJSON(&req); err != nil {
		ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	payment, err := i.PaymentSvc.CheckPayment(c.Request.Context(), orderId, req)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseSuccess(c, payment)
}
