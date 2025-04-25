package handlers

import (
	"bookstore/internal/config"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	router *gin.Engine
	server *http.Server
	user   *UserHandler
}

func NewServer(cfg config.Server, user *UserHandler) *Server {
	router := gin.Default()
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: router,
	}
	srv := &Server{
		user:   user,
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
}
