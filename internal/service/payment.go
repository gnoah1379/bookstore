package service

import (
	"bookstore/internal/model"
	"bookstore/internal/repository"
	"context"
	"strconv"
	"time"
)

type PaymentService interface {
	CreatePayment(ctx context.Context, payment model.Payment) (model.Payment, error)
	CheckPayment(ctx context.Context, orderId string, payment model.Payment) (model.Payment, error)
	ListAllPayment(ctx context.Context) ([]model.Payment, error)
	GetPaymentByOrderId(ctx context.Context, id string) (model.Payment, error)
}

type paymentService struct {
	paymentRepo repository.PaymentRepo
}

func NewPaymentService(paymentRepo repository.PaymentRepo) PaymentService {
	return &paymentService{
		paymentRepo: paymentRepo,
	}
}

func (s *paymentService) CreatePayment(ctx context.Context, paymentReq model.Payment) (model.Payment, error) {
	var payment = model.Payment{
		ID:        paymentReq.ID,
		OrderID:   paymentReq.OrderID,
		Payer:     paymentReq.Payer,
		Amount:    paymentReq.Amount,
		Method:    paymentReq.Method,
		Status:    "pending",
		CreatedAt: time.Now(),
	}

	err := s.paymentRepo.CreatePayment(ctx, &payment)
	if err != nil {
		return model.Payment{}, err
	}
	return payment, nil
}

func (s *paymentService) CheckPayment(ctx context.Context, orderId string, paymentReq model.Payment) (model.Payment, error) {
	var cnvInt, _ = strconv.Atoi(orderId)
	var payment = model.Payment{
		ID:        cnvInt,
		OrderID:   paymentReq.OrderID,
		Payer:     paymentReq.Payer,
		Amount:    paymentReq.Amount,
		Method:    paymentReq.Method,
		Status:    "completed",
		UpdatedAt: time.Now(),
	}

	err := s.paymentRepo.CheckPayment(ctx, &payment)
	if err != nil {
		return model.Payment{}, err
	}
	return payment, nil
}

func (s *paymentService) ListAllPayment(ctx context.Context) ([]model.Payment, error) {
	payments, err := s.paymentRepo.GetAllPayment(ctx)
	if err != nil {
		return []model.Payment{}, err
	}
	return payments, nil
}

func (s *paymentService) GetPaymentByOrderId(ctx context.Context, id string) (model.Payment, error) {
	payment, err := s.paymentRepo.GetPaymentByOrderId(ctx, id)
	if err != nil {
		return model.Payment{}, err
	}
	return payment, nil
}
