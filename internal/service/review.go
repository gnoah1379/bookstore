package service

import (
	"bookstore/internal/model"
	"bookstore/internal/repository"
	"context"
	"fmt"
	"strconv"
	"time"
)

type ReviewService interface {
	//review
	CreateReview(ctx context.Context, reviewReq model.Review) (model.Review, error)
	ListAllReview(ctx context.Context) ([]model.Review, error)
	ListReviewByBookId(ctx context.Context, id string) ([]model.Review, error)
	UpdateReviewByReviewId(ctx context.Context, id string, reviewUpdate model.Review) (model.Review, error)
	DeleteReviewByReviewId(ctx context.Context, id string, reviewDelete model.Review) (model.Review, error)

	//reply review
	CreateReply(ctx context.Context, replyReq model.ReplyReview) (model.ReplyReview, error)
	ListAllReply(ctx context.Context) ([]model.ReplyReview, error)
	ListReplyByReviewId(ctx context.Context, id string) ([]model.ReplyReview, error)
	UpdateReplyByReplyId(ctx context.Context, id string, replyUpdate model.ReplyReview) (model.ReplyReview, error)
	DeleteReplyByReplyId(ctx context.Context, id string, replyDelete model.ReplyReview) (model.ReplyReview, error)
}

type reviewService struct {
	reviewRepo repository.ReviewRepo
}

func NewReviewService(reviewRepo repository.ReviewRepo) ReviewService {
	return &reviewService{
		reviewRepo: reviewRepo,
	}
}

// review
func (r *reviewService) CreateReview(ctx context.Context, reviewReq model.Review) (model.Review, error) {
	var review = model.Review{
		UserID:    reviewReq.ID,
		BookID:    reviewReq.BookID,
		Rating:    reviewReq.Rating,
		Comment:   reviewReq.Comment,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := r.reviewRepo.CreateReview(ctx, &review)
	if err != nil {
		return model.Review{}, fmt.Errorf("error creating review %w", err)
	}
	return review, err
}

func (r *reviewService) ListAllReview(ctx context.Context) ([]model.Review, error) {
	reviews, err := r.reviewRepo.GetAllReview(ctx)
	if err != nil {
		return []model.Review{}, err
	}
	return reviews, nil
}

func (r *reviewService) ListReviewByBookId(ctx context.Context, id string) ([]model.Review, error) {
	reviews, err := r.reviewRepo.GetReviewByBookId(ctx, id)
	if err != nil {
		return []model.Review{}, err
	}
	return reviews, err
}

func (r *reviewService) UpdateReviewByReviewId(ctx context.Context, id string, reviewUpdate model.Review) (model.Review, error) {
	var idInt, _ = strconv.Atoi(id)
	var review = model.Review{
		ID:        idInt,
		UserID:    reviewUpdate.UserID,
		Rating:    reviewUpdate.Rating,
		Comment:   reviewUpdate.Comment,
		UpdatedAt: time.Now(),
	}
	err := r.reviewRepo.UpdateReviewByReviewId(ctx, &review)
	if err != nil {
		return model.Review{}, fmt.Errorf("error update review %w", err)
	}
	return review, err
}

func (r *reviewService) DeleteReviewByReviewId(ctx context.Context, id string, reviewDelete model.Review) (model.Review, error) {
	var idInt, _ = strconv.Atoi(id)
	var review = model.Review{
		ID:     idInt,
		UserID: reviewDelete.UserID,
	}
	err := r.reviewRepo.DeleteReviewByReviewId(ctx, &review)
	if err != nil {
		return model.Review{}, err
	}
	return review, err
}

// reply review
func (r *reviewService) CreateReply(ctx context.Context, replyReq model.ReplyReview) (model.ReplyReview, error) {
	var reply = model.ReplyReview{
		ReviewID:  replyReq.ReviewID,
		UserID:    replyReq.UserID,
		Comment:   replyReq.Comment,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := r.reviewRepo.CreateReply(ctx, &reply)
	if err != nil {
		return model.ReplyReview{}, fmt.Errorf("error creating reply %w", err)
	}
	return reply, err
}

func (r *reviewService) ListAllReply(ctx context.Context) ([]model.ReplyReview, error) {
	reply, err := r.reviewRepo.GetAllReply(ctx)
	if err != nil {
		return []model.ReplyReview{}, err
	}
	return reply, nil
}

func (r *reviewService) ListReplyByReviewId(ctx context.Context, id string) ([]model.ReplyReview, error) {
	reply, err := r.reviewRepo.GetReplyByReviewId(ctx, id)
	if err != nil {
		return []model.ReplyReview{}, err
	}
	return reply, err
}

func (r *reviewService) UpdateReplyByReplyId(ctx context.Context, id string, replyUpdate model.ReplyReview) (model.ReplyReview, error) {
	var idInt, _ = strconv.Atoi(id)
	var reply = model.ReplyReview{
		ID:        idInt,
		Comment:   replyUpdate.Comment,
		UpdatedAt: time.Now(),
	}
	err := r.reviewRepo.UpdateReplyByReplyId(ctx, &reply)
	if err != nil {
		return model.ReplyReview{}, fmt.Errorf("error update review %w", err)
	}
	return reply, err
}

func (r *reviewService) DeleteReplyByReplyId(ctx context.Context, id string, replyDelete model.ReplyReview) (model.ReplyReview, error) {
	var idInt, _ = strconv.Atoi(id)
	var reply = model.ReplyReview{
		ID:     idInt,
		UserID: replyDelete.UserID,
	}
	err := r.reviewRepo.DeleteReplyByReplyId(ctx, &reply)
	if err != nil {
		return model.ReplyReview{}, err
	}
	return reply, err
}
