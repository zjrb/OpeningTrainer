package ports

import "github.com/zjrb/OpeningTrainer/internal/core/domain"

type OAuthProvider interface {
	GetAuthURL() string
	Authenticate(code string) (*domain.User, error)
}

type UserRepository interface {
	GetUserByEmail(email string) (*domain.User, error)
	CreateUser(user *domain.User) error
}

type JWTProvider interface {
	GenerateToken(user *domain.User) (string, error)
	ValidateToken(token string) (*domain.User, error)
}
