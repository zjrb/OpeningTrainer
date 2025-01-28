package handler

import (
	"net/http"

	"github.com/zjrb/OpeningTrainer/internal/adapters/middleware"
)

func AddRoutes(mux *http.ServeMux, authHandler *AuthHandler, authMiddleware *middleware.AuthMiddleware) {
	mux.Handle("/v1/auth/google/login", authHandler.GetOAuthPageURL())
	mux.Handle("/v1/auth/google/callback", authHandler.GetOAuthCallback())
	mux.Handle("/v1/auth/google/protected", authMiddleware.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Protected"))
	})))

}
