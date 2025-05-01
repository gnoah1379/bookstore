package service

import (
	"bookstore/internal/model"
	"bookstore/internal/repository"
	"context"
	"fmt"
	"strconv"
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
	Register(ctx context.Context, user UserRegistration) (model.User, error)
	ListAllUser(ctx context.Context) ([]model.User, error)
	SearchUserById(ctx context.Context, id string) (model.User, error)
	UpdateUserById(ctx context.Context, id string, req model.User) (model.User, error)
	DeleteUserById(ctx context.Context, id string) (model.User, error)
}

type userService struct {
	userRepo repository.UserRepo
}

func NewUserService(userRepo repository.UserRepo) UserService {
	return &userService{userRepo: userRepo}
}

func (u *userService) Register(ctx context.Context, userReg UserRegistration) (model.User, error) {
	password, err := HashPassword(userReg.Password)
	if err != nil {
		return model.User{}, fmt.Errorf("error hashing password %w", err)
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
		return model.User{}, fmt.Errorf("error creating user %w", err)
	}
	return user, err
}

func (u *userService) ListAllUser(ctx context.Context) ([]model.User, error) {
	user, err := u.userRepo.GetAllUsers(ctx)
	if err != nil {
		return []model.User{}, err
	}
	return user, err
}

func (u *userService) SearchUserById(ctx context.Context, id string) (model.User, error) {
	user, err := u.userRepo.GetUserById(ctx, id)
	if err != nil {
		return model.User{}, err
	}
	return user, err
}

func (u *userService) UpdateUserById(ctx context.Context, id string, req model.User) (model.User, error) {
	var idInt, _ = strconv.Atoi(id)
	var user = model.User{
		ID:        idInt,
		Username:  req.Username,
		Password:  req.Password,
		Email:     req.Email,
		Birthday:  req.Birthday,
		Phone:     req.Phone,
		Avatar:    req.Avatar,
		Address:   req.Address,
		Gender:    req.Gender,
		UpdatedAt: time.Now(),
	}
	err := u.userRepo.UpdateUserById(ctx, &user)
	if err != nil {
		return model.User{}, err
	}
	return user, err
}

func (u *userService) DeleteUserById(ctx context.Context, id string) (model.User, error) {
	user, err := u.userRepo.DeleteUserById(ctx, id)
	if err != nil {
		return model.User{}, err
	}
	return user, err
}
