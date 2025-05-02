package handlers

import (
	"bookstore/internal/model"
	"bookstore/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BookHandler struct {
	bookSvc service.BookService
}

func NewBookHandler(bookSrv service.BookService) *BookHandler {
	return &BookHandler{bookSvc: bookSrv}
}

func (b *BookHandler) ListAllBooks(c *gin.Context) {
	books, err := b.bookSvc.ListAllBook(c.Request.Context())
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, err.Error())
	}
	ResponseSuccess(c, books)
}

func (b *BookHandler) SearchBooks(c *gin.Context) {
	var request = c.Param("id")
	books, err := b.bookSvc.SearchBookById(c.Request.Context(), request)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseSuccess(c, books)
}

func (b *BookHandler) SearchBookByName(c *gin.Context) {
	var request = c.Param("bookname")
	books, err := b.bookSvc.SearchBookByName(c.Request.Context(), request)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseSuccess(c, books)
}

func (b *BookHandler) CreateBook(c *gin.Context) {
	var request model.Book
	err := c.ShouldBindJSON(&request)
	if err != nil {
		ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	book, err := b.bookSvc.CreateBook(c.Request.Context(), request)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseSuccess(c, book)
}

func (b *BookHandler) DeleteBook(c *gin.Context) {
	var request = c.Param("id")
	book, err := b.bookSvc.DeleteBookById(c.Request.Context(), request)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseSuccess(c, book)
}

func (b *BookHandler) UpdateBook(c *gin.Context) {
	var request model.Book
	var id = c.Param("id")
	err := c.ShouldBindJSON(&request)
	if err != nil {
		ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	book, err := b.bookSvc.UpdateBook(c.Request.Context(), id, request)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseSuccess(c, book)
}
