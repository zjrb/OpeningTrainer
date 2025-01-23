package app

import (
	"backend/config"
	"backend/internal/usecase"
	"backend/internal/usecase/repo"
	"backend/internal/usecase/webapi"
	"backend/pkg/httpserver"
	"backend/pkg/logger"
	"backend/pkg/postgres"

	"github.com/gorilla/mux"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal("failed to connect to postgres", err)
	}
	defer pg.Close()

	authUsecase := usecase.New(repo.New(pg), webapi.New(cfg.Auth.GoogleClientID, cfg.Auth.GoogleClientSecret, cfg.Auth.RedirectURL))
	handler := mux.NewRouter()
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))
	err = httpServer.Shutdown()

}
