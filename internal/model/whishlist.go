package model

import "time"

type Whishlist struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	UserID    int       `json:"user_id"`
	BookID    int       `json:"book_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (w *Whishlist) tableName() string {
	return "whishlists"
}
