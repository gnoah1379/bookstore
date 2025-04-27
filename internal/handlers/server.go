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
	srv.Login()
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
}

func (srv *Server) Login() {
	srv.router.POST("/api/v1/user/login", srv.auth.Login)
}
