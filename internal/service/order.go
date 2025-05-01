package service

import (
	"bookstore/internal/model"
	"bookstore/internal/repository"
	"context"
	"strconv"
	"time"
)

type OrderService interface {
	// Order detail service
	CreateOrderDetail(ctx context.Context, orderDetail model.OrderDetail) (model.OrderDetail, error)
	ListAllOrderDetailByOrderId(ctx context.Context, orderId string) ([]model.OrderDetail, error)
	UpdateOrderDetailById(ctx context.Context, id string, req model.OrderDetail) (model.OrderDetail, error)
	DeleteOrderDetailById(ctx context.Context, id string) (model.OrderDetail, error)

	// Order service
	CreateOrder(ctx context.Context, order model.Order) (model.Order, error)
	ListAllOrder(ctx context.Context) ([]model.Order, error)
	SearchOrderById(ctx context.Context, id string) (model.Order, error)
	UpdateOrderById(ctx context.Context, id string, req model.Order) (model.Order, error)
	DeleteOrderById(ctx context.Context, id string) (model.Order, error)
}

type orderService struct {
	orderRepo repository.OrderRepo
}

func NewOrderService(orderRepo repository.OrderRepo) OrderService {
	return &orderService{
		orderRepo: orderRepo,
	}
}

// Order Detail
func (s *orderService) CreateOrderDetail(ctx context.Context, orderDetail model.OrderDetail) (model.OrderDetail, error) {
	return orderDetail, s.orderRepo.CreateOrderDetail(ctx, &orderDetail)
}

func (s *orderService) ListAllOrderDetailByOrderId(ctx context.Context, orderId string) ([]model.OrderDetail, error) {
	return s.orderRepo.GetAllOrderDetailByOrderId(ctx, orderId)
}

func (s *orderService) UpdateOrderDetailById(ctx context.Context, id string, req model.OrderDetail) (model.OrderDetail, error) {
	var idInt, _ = strconv.Atoi(id)
	var orderDetail = model.OrderDetail{
		ID:        idInt,
		Quantity:  req.Quantity,
		Total:     req.Total,
		UpdatedAt: time.Now(),
	}
	err := s.orderRepo.UpdateOrderDetailById(ctx, id, orderDetail)
	if err != nil {
		return model.OrderDetail{}, err
	}
	return orderDetail, err
}

func (s *orderService) DeleteOrderDetailById(ctx context.Context, id string) (model.OrderDetail, error) {
	return model.OrderDetail{}, s.orderRepo.DeleteOrderDetailById(ctx, id)
}

// Order
func (s *orderService) CreateOrder(ctx context.Context, order model.Order) (model.Order, error) {
	return order, s.orderRepo.CreateOrder(ctx, &order)
}

func (s *orderService) ListAllOrder(ctx context.Context) ([]model.Order, error) {
	return s.orderRepo.GetAllOrder(ctx)
}

func (s *orderService) SearchOrderById(ctx context.Context, id string) (model.Order, error) {
	return s.orderRepo.GetOrderById(ctx, id)
}

func (s *orderService) UpdateOrderById(ctx context.Context, id string, req model.Order) (model.Order, error) {
	var idInt, _ = strconv.Atoi(id)
	var order = model.Order{
		ID:        idInt,
		Status:    req.Status,
		UpdatedAt: time.Now(),
	}
	err := s.orderRepo.UpdateOrderById(ctx, order)
	if err != nil {
		return model.Order{}, err
	}
	return order, err
}

func (s *orderService) DeleteOrderById(ctx context.Context, id string) (model.Order, error) {
	return model.Order{}, s.orderRepo.DeleteOrderById(ctx, id)
}
