package service

import (
	"bookstore/internal/model"
	"bookstore/internal/repository"
	"context"
	"fmt"
	"time"
)

type UserRegistration struct {
	Username string    `json:"username" binding:"required"`
	Email    string    `json:"email" binding:"required"`
	Password string    `json:"password" binding:"required"`
	Gender   string    `json:"gender" binding:"required,oneof=Male Female "`
	Address  string    `json:"address" binding:"required"`
	Birthday time.Time `json:"birthday" binding:"required"`
	Phone    string    `json:"phone" binding:"required"`
	Avatar   string    `json:"avatar"`
}

type UserService interface {
	Register(ctx context.Context, user UserRegistration) (model.UserInfo, error)
}

type userService struct {
	userRepo repository.UserRepo
}

func NewUserService(userRepo repository.UserRepo) UserService {
	return &userService{userRepo: userRepo}
}

func (u *userService) Register(ctx context.Context, userReg UserRegistration) (model.UserInfo, error) {
	password, err := HashPassword(userReg.Password)
	if err != nil {
		return model.UserInfo{}, fmt.Errorf("error hashing password %w", err)
	}
	var user = model.User{
		Username:  userReg.Username,
		Password:  string(password),
		Email:     userReg.Email,
		Role:      model.RoleUser,
		Gender:    userReg.Gender,
		Address:   userReg.Address,
		Birthday:  userReg.Birthday,
		Phone:     userReg.Phone,
		Avatar:    userReg.Avatar,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = u.userRepo.CreateUser(ctx, &user)
	if err != nil {
		return model.UserInfo{}, fmt.Errorf("error creating user %w", err)
	}
	var result model.UserInfo
	result.FromUser(user)
	return result, err
}
