package handlers

import (
	"bookstore/internal/config"
	"bookstore/internal/repository"
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
	order  *OrderHandler
}

func NewServer(cfg config.Config, user *UserHandler, auth *AuthHandler, handler *BookHandler, order *OrderHandler) *Server {
	router := gin.Default()
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: router,
	}
	srv := &Server{
		user:   user,
		auth:   auth,
		book:   handler,
		router: router,
		server: server,
		order:  order,
	}
	srv.Register(cfg)
	return srv
}

func (srv *Server) Start() error {
	return srv.server.ListenAndServe()
}

func (srv *Server) Shutdown(ctx context.Context) error {
	return srv.server.Shutdown(ctx)
}

func (srv *Server) Register(cfg config.Config) {
	srv.router.POST("/api/v1/user/register", srv.user.Register)
	srv.router.POST("/api/v1/auth/login", srv.auth.Login)
	srv.router.GET("/api/v1/book", srv.book.ListAllBooks)

	protected := srv.router.Group("/api/v1/service")
	protected.Use(service.AuthMiddleware(repository.NewJWTRepo(cfg.Key.JwtSecret)))
	{
		//book service
		protected.POST("/book", srv.book.CreateBook, service.ProtectedHandler)
		protected.GET("/book/:id", srv.book.SearchBooks, service.ProtectedHandler)
		protected.PUT("/book/:id", srv.book.UpdateBook, service.ProtectedHandler)
		protected.DELETE("/book/:id", srv.book.DeleteBook, service.ProtectedHandler)

		//user service
		protected.GET("/user", srv.user.ListUsers, service.ProtectedHandler)
		protected.GET("/user/:id", srv.user.SearchUser, service.ProtectedHandler)
		protected.PUT("/user/:id", srv.user.UpdateUser, service.ProtectedHandler)
		protected.DELETE("/user/:id", srv.user.DeleteUser, service.ProtectedHandler)

		//order service
		protected.POST("/order", srv.order.CreateOrder, service.ProtectedHandler)
		protected.GET("/order", srv.order.ListAllOrder, service.ProtectedHandler)
		protected.GET("/order/:id", srv.order.SearchOrder, service.ProtectedHandler)
		protected.PUT("/order/:id", srv.order.UpdateOrder, service.ProtectedHandler)
		protected.DELETE("/order/:id", srv.order.DeleteOrder, service.ProtectedHandler)
	}
}
