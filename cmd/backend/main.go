package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/redis/go-redis/v9"
	"github.com/zjrb/OpeningTrainer/internal/adapters/auth/jwt"
	"github.com/zjrb/OpeningTrainer/internal/adapters/auth/oauth"
	"github.com/zjrb/OpeningTrainer/internal/adapters/engine"
	"github.com/zjrb/OpeningTrainer/internal/adapters/handler"
	"github.com/zjrb/OpeningTrainer/internal/adapters/middleware"
	"github.com/zjrb/OpeningTrainer/internal/adapters/storage/postgres"
	rediscli "github.com/zjrb/OpeningTrainer/internal/adapters/storage/redis"
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

	h := http.NewServeMux()
	httpServer := httpserver.New(h, httpserver.Port(cfg.HTTP.Port))
	logger.Debug("Starting http server")
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.REDIS.Address,
		Password: cfg.REDIS.Password,
		DB:       0,
	})
	cache := rediscli.NewRedisRepo(rdb)
	// Building AuthService
	jwtProvider := jwt.NewJWT(cfg.JWT.Secret)
	userRepo := postgres.NewUserRepositoryPostgres(pg.Pool)
	authconfig := oauth2.Config{
		ClientID:     cfg.Auth.GoogleClientID,
		ClientSecret: cfg.Auth.GoogleClientSecret,
		RedirectURL:  cfg.Auth.RedirectURL,
		Endpoint:     google.Endpoint,
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile"},
	}
	openingRepo := postgres.NewOpeningRepoPostgres(pg.Pool)
	oauthProvider := oauth.NewOauth2google(&authconfig)

	openingService := services.NewOpeningService(openingRepo)
	authService := services.NewAuthService(oauthProvider, userRepo, jwtProvider)
	authHandler := handler.NewAuthHandler(authService, logger)
	openingHandler := handler.NewOpeningHandler(openingService)
	authMiddleware := middleware.NewAuthMiddleware(authService)
	chessEngine := engine.NewChessEngine()
	gameSessionRepo := postgres.NewGameSessionRepo(pg.Pool)
	chessService := services.NewChessService(chessEngine, cache, logger, &gameSessionRepo, openingRepo)
	websocketHandler := handler.NewWebSocketHandler(chessService, logger)
	h.HandleFunc("/ws", websocketHandler.HandleConnections())
	handler.AddRoutes(h, authHandler, authMiddleware, openingHandler)

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
