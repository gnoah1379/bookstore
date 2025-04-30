package handlers

import (
	"bookstore/internal/config"
	"bookstore/internal/service"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	server *http.Server
	user   *UserHandler
	auth   *AuthHandler
	book   *BookHandler
}

func NewServer(cfg config.Server, user *UserHandler, auth *AuthHandler, handler *BookHandler) *Server {
	router := gin.Default()
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: router,
	}
	srv := &Server{
		user:   user,
		auth:   auth,
		book:   handler,
		router: router,
		server: server,
	}
	srv.Register()
	return srv
}

func (srv *Server) Start() error {
	return srv.server.ListenAndServe()
}

func (srv *Server) Shutdown(ctx context.Context) error {
	return srv.server.Shutdown(ctx)
}

func (srv *Server) Register() {
	srv.router.POST("/api/v1/user/register", srv.user.Register)
	srv.router.POST("/api/v1/auth/login", srv.auth.Login)

	protected := srv.router.Group("/api/v1/service")
	protected.Use(service.AuthMiddleware())
	{
		//book service
		protected.POST("/book", srv.book.CreateBook, service.ProtectedHandler)
		protected.GET("/book", srv.book.ListAllBooks, service.ProtectedHandler)
		protected.GET("/book/:id", srv.book.SearchBooks, service.ProtectedHandler)
		protected.PUT("/book/:id", srv.book.UpdateBook, service.ProtectedHandler)
		protected.DELETE("/book/:id", srv.book.DeleteBook, service.ProtectedHandler)
	}
}
