package middleware

import (
	"context"
	"net/http"

	"github.com/zjrb/OpeningTrainer/internal/core/domain"
	"github.com/zjrb/OpeningTrainer/internal/core/services"
)

type AuthMiddleware struct {
	authService *services.AuthService
}

func NewAuthMiddleware(authService *services.AuthService) *AuthMiddleware {
	return &AuthMiddleware{
		authService: authService,
	}
}

func (m *AuthMiddleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		usr, err := m.authService.ValidateToken(token.Value)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), domain.EmailContextKey, usr.Email)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
