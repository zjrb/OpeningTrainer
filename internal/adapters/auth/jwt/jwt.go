package jwt

import "github.com/zjrb/OpeningTrainer/internal/core/domain"

type JWT struct {
	secret string
}

func NewJWT(secret string) *JWT {
	return &JWT{
		secret: secret,
	}
}

func (j *JWT) GenerateToken(user *domain.User) (string, error) {
	return "", nil
}

func (j *JWT) ValidateToken(token string) (*domain.User, error) {
	return &domain.User{}, nil
}
