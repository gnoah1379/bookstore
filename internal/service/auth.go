package service

import (
	"bookstore/internal/model"
	"bookstore/internal/repository"
	"context"
	"fmt"
)

type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthService interface {
	Login(ctx context.Context, user UserLogin) (model.User, error)
}

type authService struct {
	authRepo repository.AuthRepo
}

func NewAuthService(authRepo repository.AuthRepo) AuthService {
	return &authService{authRepo: authRepo}
}

func (s *authService) Login(ctx context.Context, user UserLogin) (model.User, error) {
	Password, err := HashPassword(user.Password)
	if err != nil {
		return model.User{}, err
	}
	var userLogin = model.User{
		Username: user.Username,
		Password: string(Password),
	}

	var dbUser model.User
	dbUser, err = s.authRepo.Login(ctx, userLogin)
	if err != nil {
		return model.User{}, fmt.Errorf("error login %w", err)
	}
	return dbUser, nil
}
