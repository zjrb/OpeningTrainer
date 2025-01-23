package v1

import (
	"backend/internal/usecase"
	"backend/pkg/logger"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type authRoutes struct {
	a usecase.AuthUseCase
	l logger.Interface
}

func newAuthRoutes(handler *mux.Router, a usecase.AuthUseCase, l logger.Interface) {
	r := &authRoutes{
		a: a,
		l: l,
	}
	handler.HandleFunc("/auth/google/url", r.getGoogleAuthUrl).Methods("GET")
	handler.HandleFunc("/auth/google", r.authGoogle).Methods("GET")
}

func (r *authRoutes) getGoogleAuthUrl(w http.ResponseWriter, req *http.Request) {
	url := r.a.GetGoogleAuthUrl()
	http.Redirect(w, req, url, http.StatusTemporaryRedirect)
}
func (r *authRoutes) authGoogle(w http.ResponseWriter, req *http.Request) {
	code := req.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Code is missing", http.StatusBadRequest)
		return
	}
	response, err := r.a.AuthGoogle(code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "User Info: %s", response)
}
