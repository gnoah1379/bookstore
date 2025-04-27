package handlers

import (
	"bookstore/internal/config"
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
}

func NewServer(cfg config.Server, user *UserHandler, auth *AuthHandler) *Server {
	router := gin.Default()
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: router,
	}
	srv := &Server{
		user:   user,
		auth:   auth,
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
}
