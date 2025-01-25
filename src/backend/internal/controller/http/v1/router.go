package v1

import (
	"backend/internal/usecase"
	"backend/pkg/logger"
	"net/http"
)

func NewRouter(h *http.ServeMux, l logger.Interface, a usecase.AuthUseCase) {
	newAuthRoutes("/v1/auth", h, a, l)
	l.Debug("v1 routes registered")
}
