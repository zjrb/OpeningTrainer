package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/zjrb/OpeningTrainer/internal/core/services"
	"github.com/zjrb/OpeningTrainer/internal/logger"
)

type AuthHandler struct {
	AuthService *services.AuthService
	Logger      *logger.Logger
}

func NewAuthHandler(authService *services.AuthService, logger *logger.Logger) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
		Logger:      logger,
	}
}

func (h *AuthHandler) GetOAuthPageURL() http.Handler {
	url, state := h.AuthService.GetOAuthPageURL()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.Logger.Info("Redirecting to OAuth page")
		var expiration = time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	})
}

func (h *AuthHandler) GetOAuthCallback() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		oauthState, _ := r.Cookie("oauthstate")
		if r.FormValue("state") != oauthState.Value {
			h.Logger.Error("Invalid oauth state")
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
		token, err := h.AuthService.Authenticate((r.FormValue("code")))
		if err != nil {
			log.Println(err.Error())
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
		var expiration = time.Now().Add(24 * time.Hour)
		cookie := http.Cookie{Name: "token", Value: token, Expires: expiration}
		http.SetCookie(w, &cookie)
		w.Write([]byte("Login success"))
	})

}
