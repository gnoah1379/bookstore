package model

import "time"

type Book struct {
	ID          string    `json:"id"`
	Bookname    string    `json:"bookname"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Author      string    `json:"author"`
	Stock       int       `json:"stock"`
	Price       int       `json:"price"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
