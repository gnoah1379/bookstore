package handlers

import (
	"bookstore/internal/model"
	"bookstore/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderHandler struct {
	orderSvc service.OrderService
}

func NewOrderHandler(orderSvc service.OrderService) *OrderHandler {
	return &OrderHandler{
		orderSvc: orderSvc,
	}
}

// Order Detail
func (h *OrderHandler) CreateOrderDetail(c *gin.Context) {
	var req model.OrderDetail
	if err := c.ShouldBindJSON(&req); err != nil {
		ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	orderDetail, err := h.orderSvc.CreateOrderDetail(c.Request.Context(), req)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseSuccess(c, orderDetail)
}

func (h *OrderHandler) UpdateOrderDetail(c *gin.Context) {
	var id = c.Param("id")
	var req model.OrderDetail
	if err := c.ShouldBindJSON(&req); err != nil {
		ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	orderDetail, err := h.orderSvc.UpdateOrderDetailById(c.Request.Context(), id, req)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseSuccess(c, orderDetail)
}

func (h *OrderHandler) DeleteOrderDetail(c *gin.Context) {
	var id = c.Param("id")
	orderDetail, err := h.orderSvc.DeleteOrderDetailById(c.Request.Context(), id)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseSuccess(c, orderDetail)
}

func (h *OrderHandler) ListAllOrderDetail(c *gin.Context) {
	orderDetails, err := h.orderSvc.ListAllOrderDetailByOrderId(c.Request.Context(), c.Param("id"))
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseSuccess(c, orderDetails)
}

// order
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req model.Order
	if err := c.ShouldBindJSON(&req); err != nil {
		ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	order, err := h.orderSvc.CreateOrder(c.Request.Context(), req)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseSuccess(c, order)
}

func (h *OrderHandler) SearchOrder(c *gin.Context) {
	var id = c.Param("id")
	order, err := h.orderSvc.SearchOrderById(c.Request.Context(), id)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseSuccess(c, order)
}

func (h *OrderHandler) ListAllOrder(c *gin.Context) {
	orders, err := h.orderSvc.ListAllOrder(c.Request.Context())
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseSuccess(c, orders)
}

func (h *OrderHandler) ListAllOrderByUserId(c *gin.Context) {
	orders, err := h.orderSvc.ListAllOrderByUserId(c.Request.Context(), c.Param("user id"))
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseSuccess(c, orders)
}

func (h *OrderHandler) UpdateOrder(c *gin.Context) {
	var id = c.Param("id")
	var req model.Order
	if err := c.ShouldBindJSON(&req); err != nil {
		ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	order, err := h.orderSvc.UpdateOrderById(c.Request.Context(), id, req)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseSuccess(c, order)
}

func (h *OrderHandler) DeleteOrder(c *gin.Context) {
	var id = c.Param("id")
	order, err := h.orderSvc.DeleteOrderById(c.Request.Context(), id)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseSuccess(c, order)
}
