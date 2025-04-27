package model

import (
	"time"
)

type UserRole string

const (
	RoleUser  UserRole = "user"
	RoleAdmin UserRole = "admin"
)

type User struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Role      UserRole  `json:"role"`
	Gender    string    `json:"gender"`
	Address   string    `json:"address"`
	Birthday  time.Time `json:"birthday"`
	Phone     string    `json:"phone"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u User) TableName() string {
	return "users"
}
