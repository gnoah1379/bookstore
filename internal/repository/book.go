package repository

import (
	"bookstore/internal/model"
	"context"
	"gorm.io/gorm"
)

type BookRepo interface {
	CreateBook(ctx context.Context, book *model.Book) error
	GetAllBook(ctx context.Context) ([]model.Book, error)
	GetBookById(ctx context.Context, id string) (model.Book, error)
	DeleteById(ctx context.Context, id string) (model.Book, error)
	UpdateBook(ctx context.Context, book *model.Book) error
}

var _ BookRepo = (*bookRepo)(nil)

type bookRepo struct {
	db *gorm.DB
}

func NewBookRepo(db *gorm.DB) BookRepo {
	return &bookRepo{
		db: db,
	}
}

func (r *bookRepo) GetAllBook(ctx context.Context) ([]model.Book, error) {
	var book []model.Book
	err := r.db.WithContext(ctx).Find(&book).Error
	return book, err
}

func (r *bookRepo) GetBookById(ctx context.Context, id string) (model.Book, error) {
	var book model.Book
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&book).Error
	return book, err
}

func (r *bookRepo) CreateBook(ctx context.Context, book *model.Book) error {
	return r.db.WithContext(ctx).Create(book).Error
}

func (r *bookRepo) DeleteById(ctx context.Context, id string) (model.Book, error) {
	var book model.Book
	err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&book).Error
	return book, err
}

func (r *bookRepo) UpdateBook(ctx context.Context, book *model.Book) error {
	r.db.First(&book)
	return r.db.WithContext(ctx).Save(&book).Error
}
