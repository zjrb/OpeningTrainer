package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/zjrb/OpeningTrainer/internal/adapters/auth/jwt"
	"github.com/zjrb/OpeningTrainer/internal/adapters/auth/oauth"
	"github.com/zjrb/OpeningTrainer/internal/adapters/storage/postgres"
	"github.com/zjrb/OpeningTrainer/internal/config"
	"github.com/zjrb/OpeningTrainer/internal/core/services"
	"github.com/zjrb/OpeningTrainer/internal/logger"
	"github.com/zjrb/OpeningTrainer/pkg/db"
	"github.com/zjrb/OpeningTrainer/pkg/httpserver"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func main() {
	logger := logger.New("debug")
	logger.Info("Logger Initiatied")
	cfg, err := config.NewConfig()
	if err != nil {
		logger.Fatal("Error reading config", err)
	}
	pg, err := db.New(cfg.PG.URL, db.MaxPoolSize(cfg.PG.PoolMax), db.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		logger.Fatal("Error connecting to database", err)
	}
	defer pg.Close()

	handler := http.NewServeMux()
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))
	logger.Debug("Starting http server")
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	// Building AuthService
	jwtProvider := jwt.NewJWT(cfg.JWT.Secret)
	userRepo := postgres.NewUserRepositoryPostgres(pg.Pool)
	authconfig := oauth2.Config{
		ClientID:     cfg.Auth.GoogleClientID,
		ClientSecret: cfg.Auth.GoogleClientSecret,
		RedirectURL:  cfg.Auth.RedirectURL,
		Endpoint:     google.Endpoint,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	}
	oauthProvider := oauth.NewOauth2google(&authconfig)
	authService := services.NewAuthService(oauthProvider, userRepo, jwtProvider)
	if authService != nil {
		logger.Info("AuthService initiated")
	}
	select {
	case s := <-interrupt:
		logger.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		logger.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}
	err = httpServer.Shutdown()
	if err != nil {
		logger.Fatal("failed to shutdown http server", err)
	}

}
