package model

import "time"

type Payment struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	OrderID   int       `json:"order_id"`
	Payer     string    `json:"payer"`
	Amount    int       `json:"amount"`
	Method    string    `json:"method"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
