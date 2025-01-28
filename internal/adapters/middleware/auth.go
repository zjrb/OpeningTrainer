package middleware

import "net/http"

// AuthMiddleware is a middleware that checks if the user is authenticated
type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler,
	}
}
func (m *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.Handler.ServeHTTP(w, r)
}
