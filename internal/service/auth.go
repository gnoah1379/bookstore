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

type LoginResponse struct {
	User  model.User `json:"user"`
	Token string     `json:"token"`
}

type AuthService interface {
	Login(ctx context.Context, user UserLoginRequest) (LoginResponse, error)
}

type authService struct {
	userRepo repository.UserRepo
	jwtUser  repository.JWTRepo
}

func NewAuthService(userRepo repository.UserRepo, jwtUser repository.JWTRepo) AuthService {
	return &authService{
		userRepo: userRepo,
		jwtUser:  jwtUser,
	}
}

func (s *authService) Login(ctx context.Context, req UserLoginRequest) (LoginResponse, error) {
	user, err := s.userRepo.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return LoginResponse{}, err
	}
	if !CheckPasswordHash(req.Password, user.Password) {
		return LoginResponse{}, errors.New("invalid password")
	}
	var role = string(user.Role)
	token, err := s.jwtUser.GenerateJWT(user.ID, user.Username, role)
	if err != nil {
		return LoginResponse{}, errors.New("failed to generate JWT")
	}
	return LoginResponse{
		User:  user,
		Token: token,
	}, nil
}
