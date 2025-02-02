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

type OpeningRepository interface {
	GetOpeningByName(name string) ([]domain.Opening, error)
	GetSingleOpeningById(id int) (*domain.Opening, error)
}

type ChessEngine interface {
	ProcessMove(gs *domain.GameSession) string
}

type GameCache interface {
	GetOpening(key string) (*domain.GameSession, error)
	SetOpening(key string, opening *domain.GameSession) error
}

type GameSessionRepo interface {
	GetGameSession(id int) (*domain.GameSessionDB, error)
	CreateGameSession(gameSesh *domain.GameSessionDB) error
	UpdateGameSession(gameSesh *domain.GameSessionDB) error
}
