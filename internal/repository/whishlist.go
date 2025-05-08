package repository

import (
	"bookstore/internal/model"
	"context"
	"gorm.io/gorm"
)

type WhishlistRepo interface {
	CreateWhishItem(ctx context.Context, whishlist *model.Whishlist) error
	GetWhishlistByUserId(ctx context.Context, userId string) ([]model.Whishlist, error)
	DeleteWhishItemByWishItemId(ctx context.Context, whishlist *model.Whishlist) error
}

var _ WhishlistRepo = (*whishlistRepo)(nil)

type whishlistRepo struct {
	db *gorm.DB
}

func NewWhishlistRepo(db *gorm.DB) WhishlistRepo {
	return &whishlistRepo{
		db: db,
	}
}

func (r *whishlistRepo) CreateWhishItem(ctx context.Context, whishlist *model.Whishlist) error {
	return r.db.WithContext(ctx).Create(whishlist).Error
}

func (r *whishlistRepo) GetWhishlistByUserId(ctx context.Context, userId string) ([]model.Whishlist, error) {
	var whishlists []model.Whishlist
	result := r.db.WithContext(ctx).Where("user_id = ?", userId).Find(&whishlists)
	if result.Error != nil {
		return nil, result.Error
	}

	return whishlists, nil
}

func (r *whishlistRepo) DeleteWhishItemByWishItemId(ctx context.Context, whishlist *model.Whishlist) error {
	result := r.db.WithContext(ctx).Where("id = ? AND user_id = ?", whishlist.ID, whishlist.UserID).First(&model.Whishlist{})
	if result.Error != nil {
		return result.Error
	}
	return r.db.WithContext(ctx).Where("id = ?", whishlist.ID).Delete(&whishlist).Error
}
