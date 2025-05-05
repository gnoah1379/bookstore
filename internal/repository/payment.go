package repository

import (
	"bookstore/internal/model"
	"context"
	"errors"
	"gorm.io/gorm"
)

type PaymentRepo interface {
	CreatePayment(ctx context.Context, payment *model.Payment) error
	CheckPayment(ctx context.Context, payment *model.Payment) error
	GetAllPayment(ctx context.Context) ([]model.Payment, error)
	GetPaymentByOrderId(ctx context.Context, id string) (model.Payment, error)
}

type paymentRepo struct {
	db *gorm.DB
}

func NewPaymentRepo(db *gorm.DB) PaymentRepo {
	return &paymentRepo{
		db: db,
	}
}

func (r *paymentRepo) CreatePayment(ctx context.Context, payment *model.Payment) error {
	if payment.Status != "pending" {
		return errors.New("trạng thái thanh toán không hợp lệ")
	}
	return r.db.WithContext(ctx).Create(payment).Error
}

func (r *paymentRepo) CheckPayment(ctx context.Context, payment *model.Payment) error {
	if payment.Status != "completed" {
		return errors.New("trạng thái thanh toán không hợp lệ")
	}
	tx := r.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}
	paymentUpdate := tx.Model(&model.Payment{}).Where("order_id = ?", payment.OrderID).Update("status", "completed")
	if paymentUpdate.Error != nil {
		return paymentUpdate.Error
	}
	if paymentUpdate.RowsAffected == 0 {
		return errors.New("không tìm thấy đơn hàng hoặc đã được cập nhật")
	}
	orderUpdate := tx.Model(&model.Order{}).Where("id = ?", payment.OrderID).Update("status", "completed")
	if orderUpdate.Error != nil {
		return orderUpdate.Error
	}
	if orderUpdate.RowsAffected == 0 {
		return errors.New("không tìm thấy đơn hàng hoặc đã được cập nhật")
	}
	return tx.Commit().Error
}

func (r *paymentRepo) GetAllPayment(ctx context.Context) ([]model.Payment, error) {
	var payments []model.Payment
	err := r.db.WithContext(ctx).Find(&payments).Error
	return payments, err
}

func (r *paymentRepo) GetPaymentByOrderId(ctx context.Context, id string) (model.Payment, error) {
	var payment model.Payment
	err := r.db.WithContext(ctx).Where("order_id = ?", id).First(&payment).Error
	return payment, err
}
