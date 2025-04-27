package repository

import (
	"bookstore/internal/model"
	"context"

	"gorm.io/gorm"
)

type AuthRepo interface {
	Login(ctx context.Context, user model.User) (model.User, error)
}

var _ AuthRepo = (*authRepo)(nil)

type authRepo struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) AuthRepo {
	return &authRepo{db: db}
}

func (r *authRepo) Login(ctx context.Context, user model.User) (model.User, error) {
	var dbUser model.User
	if err := r.db.WithContext(ctx).Where("username = ?", user.Username).Where("password = ?", user.Password).First(&dbUser).Error; err != nil {
		return model.User{}, err
	}
	return dbUser, nil
}
