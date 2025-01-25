package app

import (
	"backend/config"
	v1 "backend/internal/controller/http/v1"
	"backend/internal/usecase"
	"backend/internal/usecase/repo"
	"backend/internal/usecase/webapi"
	"backend/pkg/google"
	"backend/pkg/httpserver"
	"backend/pkg/logger"
	"backend/pkg/postgres"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal("failed to connect to postgres", err)
	}
	defer pg.Close()

	googleAuth := google.New(google.ClientID(cfg.Auth.GoogleClientID), google.ClientSecret(cfg.Auth.GoogleClientSecret), google.RedirectURL(cfg.Auth.RedirectURL))
	authUsecase := usecase.New(repo.New(pg), webapi.New(googleAuth))
	handler := http.NewServeMux()
	v1.NewRouter(handler, l, *authUsecase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))
	l.Debug("starting http server")
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}
	err = httpServer.Shutdown()
	if err != nil {
		l.Fatal("failed to shutdown http server", err)
	}

}
