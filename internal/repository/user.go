package repository

import (
	"bookstore/internal/model"
	"context"
	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUserByUsername(ctx context.Context, username string) (model.User, error)
	GetUserById(ctx context.Context, id string) (model.User, error)
	GetAllUsers(ctx context.Context) ([]model.User, error)
	UpdateUserById(ctx context.Context, user *model.User) error
	DeleteUserById(ctx context.Context, id string) (model.User, error)
}

var _ UserRepo = (*userRepo)(nil)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) CreateUser(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *userRepo) GetUserByUsername(ctx context.Context, username string) (model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("username = ?", username).First(&user).Error
	return user, err
}

func (r *userRepo) GetUserById(ctx context.Context, id string) (model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	return user, err
}

func (r *userRepo) GetAllUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User
	err := r.db.WithContext(ctx).Find(&users).Error
	return users, err
}

func (r *userRepo) UpdateUserById(ctx context.Context, user *model.User) error {
	result := r.db.WithContext(ctx).Where("id = ?", user.ID).First(&model.User{})
	if result.Error != nil {
		return result.Error
	}
	return r.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", user.ID).Updates(user).Error
}

func (r *userRepo) DeleteUserById(ctx context.Context, id string) (model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&user).Error
	return user, err
}
