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
		authRepo := repository.NewAuthRepo(db)
		authSvc := service.NewAuthService(authRepo)
		authHandler := handlers.NewAuthHandler(authSvc)
		srv := handlers.NewServer(cfg.Server, userHandler, authHandler)
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
