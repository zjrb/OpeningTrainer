package handler

import "net/http"

func AddRoutes(mux *http.ServeMux, authHandler *AuthHandler) {
	mux.Handle("/v1/auth/google/login", authHandler.GetOAuthPageURL())
	mux.Handle("/v1/auth/google/callback", authHandler.GetOAuthCallback())
}
