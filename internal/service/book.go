package service

import (
	"bookstore/internal/model"
	"bookstore/internal/repository"
	"context"
	"fmt"
	"time"
)

type BookService interface {
	CreateBook(ctx context.Context, book model.Book) (model.Book, error)
	ListAllBook(ctx context.Context) ([]model.Book, error)
	SearchBookById(ctx context.Context, req string) (model.Book, error)
	DeleteBookById(ctx context.Context, req string) (model.Book, error)
	UpdateBook(ctx context.Context, id string, bookUpdate model.Book) (model.Book, error)
}

type bookService struct {
	bookRepo repository.BookRepo
}

func NewBookService(bookRepo repository.BookRepo) BookService {
	return &bookService{bookRepo: bookRepo}
}

func (s *bookService) ListAllBook(ctx context.Context) ([]model.Book, error) {
	books, err := s.bookRepo.GetAllBook(ctx)
	if err != nil {
		return []model.Book{}, err
	}
	return books, nil
}

func (s *bookService) SearchBookById(ctx context.Context, req string) (model.Book, error) {
	book, err := s.bookRepo.GetBookById(ctx, req)
	if err != nil {
		return model.Book{}, err
	}
	return book, nil
}

func (s *bookService) CreateBook(ctx context.Context, bookAdd model.Book) (model.Book, error) {
	var book = model.Book{
		Bookname:    bookAdd.Bookname,
		Description: bookAdd.Description,
		Category:    bookAdd.Category,
		Author:      bookAdd.Author,
		Stock:       bookAdd.Stock,
		Price:       bookAdd.Price,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err := s.bookRepo.CreateBook(ctx, &book)
	if err != nil {
		return model.Book{}, fmt.Errorf("error creating book %w", err)
	}
	return book, err
}

func (s *bookService) DeleteBookById(ctx context.Context, req string) (model.Book, error) {
	book, err := s.bookRepo.DeleteById(ctx, req)
	if err != nil {
		return model.Book{}, err
	}
	return book, nil
}

func (s *bookService) UpdateBook(ctx context.Context, id string, bookUpdate model.Book) (model.Book, error) {
	var book = model.Book{
		ID:          id,
		Bookname:    bookUpdate.Bookname,
		Description: bookUpdate.Description,
		Category:    bookUpdate.Category,
		Author:      bookUpdate.Author,
		Stock:       bookUpdate.Stock,
		Price:       bookUpdate.Price,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err := s.bookRepo.UpdateBook(ctx, book)
	if err != nil {
		return model.Book{}, err
	}
	return book, nil
}
