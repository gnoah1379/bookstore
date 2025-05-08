package service

import (
	"bookstore/internal/model"
	"bookstore/internal/repository"
	"context"
	"strconv"
	"time"
)

type WhishlistService interface {
	CreateWhishItem(ctx context.Context, whishlist model.Whishlist) (model.Whishlist, error)
	GetWhishlistByUserId(ctx context.Context, userId string) ([]model.Whishlist, error)
	DeleteWhishItemByWishItemId(ctx context.Context, id string, whishlist model.Whishlist) (model.Whishlist, error)
}

type whishlistService struct {
	whishlistRepo repository.WhishlistRepo
}

func NewWhishlistService(whishlistRepo repository.WhishlistRepo) WhishlistService {
	return &whishlistService{
		whishlistRepo: whishlistRepo,
	}
}

func (s *whishlistService) CreateWhishItem(ctx context.Context, whishlistReq model.Whishlist) (model.Whishlist, error) {
	var whishlist = model.Whishlist{
		UserID:    whishlistReq.UserID,
		BookID:    whishlistReq.BookID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := s.whishlistRepo.CreateWhishItem(ctx, &whishlist)
	if err != nil {
		return model.Whishlist{}, err
	}
	return whishlist, err
}

func (s *whishlistService) GetWhishlistByUserId(ctx context.Context, userId string) ([]model.Whishlist, error) {
	whishlist, err := s.whishlistRepo.GetWhishlistByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	return whishlist, nil
}

func (s *whishlistService) DeleteWhishItemByWishItemId(ctx context.Context, id string, whishlistReq model.Whishlist) (model.Whishlist, error) {
	var idInt, _ = strconv.Atoi(id)
	var whishlist = model.Whishlist{
		ID:     idInt,
		UserID: whishlistReq.UserID,
	}
	err := s.whishlistRepo.DeleteWhishItemByWishItemId(ctx, &whishlist)
	if err != nil {
		return model.Whishlist{}, err
	}
	return whishlist, err
}
