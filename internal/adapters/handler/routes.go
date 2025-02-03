package handler

import (
	"net/http"

	"github.com/zjrb/OpeningTrainer/internal/adapters/middleware"
)

func AddRoutes(mux *http.ServeMux, authHandler *AuthHandler, authMiddleware *middleware.AuthMiddleware,
	openingHandler *OpeningHandler) {
	mux.Handle("/v1/auth/google/login", authHandler.GetOAuthPageURL())
	mux.Handle("/v1/auth/google/callback", authHandler.GetOAuthCallback())
	mux.Handle("GET /v1/openings/{name}", authMiddleware.AuthMiddleware(openingHandler.GetOpening()))

}
