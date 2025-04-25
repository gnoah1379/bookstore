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
	Password  string    `json:"password"`
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

type UserInfo struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      UserRole  `json:"role"`
	Gender    string    `json:"gender"`
	Address   string    `json:"address"`
	Birthday  time.Time `json:"birthday"`
	Phone     string    `json:"phone"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (r *UserInfo) FromUser(user User) {
	r.ID = user.ID
	r.Username = user.Username
	r.Email = user.Email
	r.Role = user.Role
	r.Gender = user.Gender
	r.Address = user.Address
	r.Birthday = user.Birthday
	r.Phone = user.Phone
	r.Avatar = user.Avatar
	r.CreatedAt = user.CreatedAt
	r.UpdatedAt = user.UpdatedAt
}
