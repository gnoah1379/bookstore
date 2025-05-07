package repository

import (
	"bookstore/internal/model"
	"context"
	"errors"
	"gorm.io/gorm"
)

type ReviewRepo interface {
	//review
	CreateReview(ctx context.Context, review *model.Review) error
	GetAllReview(ctx context.Context) ([]model.Review, error)
	GetReviewByBookId(ctx context.Context, id string) ([]model.Review, error)
	UpdateReviewByReviewId(ctx context.Context, review *model.Review) error
	DeleteReviewByReviewId(ctx context.Context, review *model.Review) error

	//reply review
	CreateReply(ctx context.Context, reply *model.ReplyReview) error
	GetAllReply(ctx context.Context) ([]model.ReplyReview, error)
	GetReplyByReviewId(ctx context.Context, id string) ([]model.ReplyReview, error)
	UpdateReplyByReplyId(ctx context.Context, reply *model.ReplyReview) error
	DeleteReplyByReplyId(ctx context.Context, reply *model.ReplyReview) error
}

var _ ReviewRepo = (*reviewRepo)(nil)

type reviewRepo struct {
	db *gorm.DB
}

func NewReviewRepo(db *gorm.DB) ReviewRepo {
	return &reviewRepo{
		db: db,
	}
}

// review
func (r *reviewRepo) CreateReview(ctx context.Context, review *model.Review) error {
	return r.db.WithContext(ctx).Create(review).Error
}

func (r *reviewRepo) GetAllReview(ctx context.Context) ([]model.Review, error) {
	var review []model.Review
	err := r.db.WithContext(ctx).Find(review).Error
	return review, err
}

func (r *reviewRepo) GetReviewByBookId(ctx context.Context, id string) ([]model.Review, error) {
	var review []model.Review
	err := r.db.WithContext(ctx).Where("book_id = ?", id).Find(review).Error
	return review, err
}

func (r *reviewRepo) UpdateReviewByReviewId(ctx context.Context, review *model.Review) error {
	var role string
	checkRole := r.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", review.UserID).Select("role").Scan(&role)
	if checkRole.Error != nil {
		return errors.New("không tìm thấy người dùng")
	}
	if role == "admin" {
		return r.db.WithContext(ctx).Model(&model.Review{}).Where("id = ?", review.ID).Updates(review).Error
	}

	result := r.db.WithContext(ctx).Where("id = ? AND user_id = ? ", review.ID, review.UserID).First(&model.Review{})
	if result.Error != nil {
		return result.Error
	}

	return r.db.WithContext(ctx).Model(&model.Review{}).Where("id = ?", review.ID).Updates(review).Error
}

func (r *reviewRepo) DeleteReviewByReviewId(ctx context.Context, review *model.Review) error {
	var role string
	checkRole := r.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", review.UserID).Select("role").Scan(&role)
	if checkRole.Error != nil {
		return errors.New("không tìm thấy người dùng")
	}
	if role == "admin" {
		return r.db.WithContext(ctx).Where("id = ?", review.ID).Delete(review).Error
	}

	result := r.db.WithContext(ctx).Where("id = ? AND user_id = ? ", review.ID, review.UserID).First(&model.Review{})
	if result.Error != nil {
		return result.Error
	}
	return r.db.WithContext(ctx).Where("id = ?", review.ID).Delete(review).Error
}

// reply reviewRepo
func (r *reviewRepo) CreateReply(ctx context.Context, reply *model.ReplyReview) error {
	tx := r.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	reviewUpdate := tx.Model(&model.Review{}).Where("id = ?", reply.ReviewID).Update("count_reply", gorm.Expr("count_reply + 1"))
	if reviewUpdate.Error != nil {
		return reviewUpdate.Error
	}
	if reviewUpdate.RowsAffected == 0 {
		return errors.New("không tìm thấy đánh giá hoặc đã được cập nhật")
	}

	createReply := tx.Create(reply)
	if createReply.Error != nil {
		return createReply.Error
	}

	return tx.Commit().Error
}

func (r *reviewRepo) GetAllReply(ctx context.Context) ([]model.ReplyReview, error) {
	var reply []model.ReplyReview
	err := r.db.WithContext(ctx).Find(reply).Error
	return reply, err
}

func (r *reviewRepo) GetReplyByReviewId(ctx context.Context, id string) ([]model.ReplyReview, error) {
	var reply []model.ReplyReview
	err := r.db.WithContext(ctx).Where("review_id = ?", id).Find(reply).Error
	return reply, err
}

func (r *reviewRepo) UpdateReplyByReplyId(ctx context.Context, reply *model.ReplyReview) error {
	var role string
	checkRole := r.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", reply.UserID).Select("role").Scan(&role)
	if checkRole.Error != nil {
		return errors.New("không tìm thấy người dùng")
	}
	if role == "admin" {
		return r.db.WithContext(ctx).Model(&model.ReplyReview{}).Where("id = ?", reply.ID).Updates(reply).Error
	}

	result := r.db.WithContext(ctx).Where("id = ? AND user_id = ? ", reply.ID, reply.UserID).First(&model.Review{})
	if result.Error != nil {
		return result.Error
	}
	return r.db.WithContext(ctx).Model(&model.ReplyReview{}).Where("id = ?", reply.ID).Updates(reply).Error
}

func (r *reviewRepo) DeleteReplyByReplyId(ctx context.Context, reply *model.ReplyReview) error {
	var role string
	checkRole := r.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", reply.UserID).Select("role").Scan(&role)
	if checkRole.Error != nil {
		return errors.New("không tìm thấy người dùng")
	}
	if role == "admin" {
		return r.db.WithContext(ctx).Where("id = ?", reply.ID).Delete(reply).Error
	}

	result := r.db.WithContext(ctx).Where("id = ? AND user_id = ? ", reply.ID, reply.UserID).First(&model.Review{})
	if result.Error != nil {
		return result.Error
	}
	return r.db.WithContext(ctx).Where("id = ?", reply.ID).Delete(reply).Error
}
