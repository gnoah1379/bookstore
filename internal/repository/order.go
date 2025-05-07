package repository

import (
	"bookstore/internal/model"
	"context"
	"gorm.io/gorm"
)

type OrderRepo interface {
	// Order Detail CRUD
	CreateOrderDetail(ctx context.Context, orderDetail *model.OrderDetail) error
	GetAllOrderDetailByOrderId(ctx context.Context, orderId string) ([]model.OrderDetail, error)
	UpdateOrderDetailById(ctx context.Context, id string, updateReq model.OrderDetail) error
	DeleteOrderDetailById(ctx context.Context, id string) error

	// Order CRUD
	CreateOrder(ctx context.Context, order *model.Order) error
	GetAllOrder(ctx context.Context) ([]model.Order, error)
	GetAllOrderByUserId(ctx context.Context, userId string) ([]model.Order, error)
	GetOrderById(ctx context.Context, id string) (model.Order, error)
	UpdateOrderById(ctx context.Context, updateReq model.Order) error
	DeleteOrderById(ctx context.Context, id string) error
}

var _ OrderRepo = (*orderRepo)(nil)

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) OrderRepo {
	return &orderRepo{
		db: db,
	}
}

// Order Detail
func (r *orderRepo) CreateOrderDetail(ctx context.Context, orderDetail *model.OrderDetail) error {
	return r.db.WithContext(ctx).Create(orderDetail).Error
}

func (r *orderRepo) GetAllOrderDetailByOrderId(ctx context.Context, orderId string) ([]model.OrderDetail, error) {
	var orderDetails []model.OrderDetail
	err := r.db.WithContext(ctx).Where("order_id = ?", orderId).Find(&orderDetails).Error
	return orderDetails, err
}

func (r *orderRepo) UpdateOrderDetailById(ctx context.Context, id string, updateReq model.OrderDetail) error {
	result := r.db.WithContext(ctx).Where("id = ?", id).First(&model.OrderDetail{})
	if result.Error != nil {
		return result.Error
	}
	return r.db.WithContext(ctx).Model(&model.OrderDetail{}).Where("id = ?", id).Updates(updateReq).Error
}

func (r *orderRepo) DeleteOrderDetailById(ctx context.Context, id string) error {
	var book model.Order
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&book).Error
}

// Order
func (r *orderRepo) CreateOrder(ctx context.Context, order *model.Order) error {
	return r.db.WithContext(ctx).Create(order).Error
}

func (r *orderRepo) GetAllOrder(ctx context.Context) ([]model.Order, error) {
	var orders []model.Order
	err := r.db.WithContext(ctx).Find(&orders).Error
	return orders, err
}

func (r *orderRepo) GetAllOrderByUserId(ctx context.Context, userId string) ([]model.Order, error) {
	var orders []model.Order
	err := r.db.WithContext(ctx).Where("user_id = ?", userId).Find(&orders).Error
	return orders, err
}

func (r *orderRepo) GetOrderById(ctx context.Context, id string) (model.Order, error) {
	var order model.Order
	err := r.db.WithContext(ctx).Where("id = ?", id).Find(&order).Error
	return order, err
}

func (r *orderRepo) UpdateOrderById(ctx context.Context, updateReq model.Order) error {
	result := r.db.WithContext(ctx).Where("id = ?", updateReq.ID).First(&model.Order{})
	if result.Error != nil {
		return result.Error
	}
	return r.db.WithContext(ctx).Model(&model.Order{}).Where("id = ?", updateReq.ID).Updates(updateReq).Error
}

func (r *orderRepo) DeleteOrderById(ctx context.Context, id string) error {
	var order model.Order
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&order).Error
}
