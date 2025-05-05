package model

import "time"

type Order struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	UserID    string    `json:"user_id"`
	Status    string    `json:"status"`
	Total     int       `json:"total"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OrderDetail struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	OrderID   int       `json:"order_id"`
	BookID    int       `json:"book_id"`
	Quantity  int       `json:"quantity"`
	Price     int       `json:"price"`
	Total     int       `json:"total"`
	CreateAt  time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
