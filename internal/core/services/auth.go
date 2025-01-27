package services

import "github.com/zjrb/OpeningTrainer/internal/core/ports"

type AuthService struct {
	OAuthProvider ports.OAuthProvider
	UserRepo      ports.UserRepository
	JWTProvider   ports.JWTProvider
}

func NewAuthService(
	oAuthProvider ports.OAuthProvider,
	userRepo ports.UserRepository,
	jwtProvider ports.JWTProvider,
) *AuthService {
	return &AuthService{
		OAuthProvider: oAuthProvider,
		UserRepo:      userRepo,
		JWTProvider:   jwtProvider,
	}
}

func (s *AuthService) GetOAuthPageURL() string {
	return "Fix me"
}

func (s *AuthService) Authenticate(code string) (string, error) {
	return "Fix me", nil
}
