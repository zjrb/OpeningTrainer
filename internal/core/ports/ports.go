package ports

import (
	"github.com/zjrb/OpeningTrainer/internal/core/domain"
)

type OAuthProvider interface {
	GetAuthURL(string) string
	Authenticate(code string) (*domain.OAuthResponse, error)
	GenerateStateOauthCookie() string
}

type UserRepository interface {
	GetUserByEmail(email string) (*domain.User, error)
	CreateUser(user *domain.OAuthResponse) error
}

type JWTProvider interface {
	GenerateToken(email string, oauthprovider string) (string, error)
	ValidateToken(token string) (string, error)
}
