package service

import (
	"bookstore/internal/model"
	"bookstore/internal/repository"
	"context"
	"errors"
)

type UserLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthService interface {
	Login(ctx context.Context, user UserLoginRequest) (model.User, error)
}

type authService struct {
	userRepo repository.UserRepo
}

func NewAuthService(userRepo repository.UserRepo) AuthService {
	return &authService{userRepo: userRepo}
}

func (s *authService) Login(ctx context.Context, req UserLoginRequest) (model.User, error) {
	user, err := s.userRepo.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return model.User{}, err
	}
	if !CheckPasswordHash(req.Password, user.Password) {
		return model.User{}, errors.New("invalid password")
	}
	return user, nil
}
