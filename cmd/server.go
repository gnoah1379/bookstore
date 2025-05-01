package cmd

import (
	"bookstore/internal/config"
	"bookstore/internal/database"
	"bookstore/internal/handlers"
	"bookstore/internal/repository"
	"bookstore/internal/service"
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/cobra"
)

var configPath string

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the application server",
	Long:  `Run the application server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.LoadConfig(configPath)
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}
		db, err := database.Open(cfg.DB)
		if err != nil {
			return err
		}
		userRepo := repository.NewUserRepo(db)
		userSvc := service.NewUserService(userRepo)
		userHandler := handlers.NewUserHandler(userSvc)
		authRepo := repository.NewUserRepo(db)
		jwtSvc := repository.NewJWTRepo(cfg.Key.JwtSecret)
		authSvc := service.NewAuthService(authRepo, jwtSvc)
		authHandler := handlers.NewAuthHandler(authSvc)
		bookRepo := repository.NewBookRepo(db)
		bookSvc := service.NewBookService(bookRepo)
		bookHandler := handlers.NewBookHandler(bookSvc)
		orderRepo := repository.NewOrderRepo(db)
		orderSvc := service.NewOrderService(orderRepo)
		orderHandler := handlers.NewOrderHandler(orderSvc)
		srv := handlers.NewServer(cfg, userHandler, authHandler, bookHandler, orderHandler)
		go func() {
			sigChan := make(chan os.Signal, 1)
			signal.Notify(sigChan, os.Interrupt, os.Kill)
			<-sigChan
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			err := srv.Shutdown(ctx)
			if err != nil {
				slog.Error("failed to shutdown server", "error", err)
			}
		}()
		slog.Info("server started")
		return srv.Start()
	},
}

func init() {
	serverCmd.Flags().StringVarP(&configPath, "config", "c", "configs/config.json", "Path to config file")
}
