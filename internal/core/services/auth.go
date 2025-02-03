package services

import (
	"fmt"

	"github.com/zjrb/OpeningTrainer/internal/core/domain"
	"github.com/zjrb/OpeningTrainer/internal/core/ports"
)

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

func (s *AuthService) GetOAuthPageURL() (string, string) {
	state := s.OAuthProvider.GenerateStateOauthCookie()
	return s.OAuthProvider.GetAuthURL(state), state
}

func (s *AuthService) Authenticate(code string) (string, error) {
	data, err := s.OAuthProvider.Authenticate(code)
	if err != nil {
		return "", err
	}
	_, err = s.UserRepo.GetUserByEmail(data.Email)
	if err != nil {
		if error := s.UserRepo.CreateUser(data); error != nil {
			fmt.Println("Error creating user: ", error)
			return "", error
		}
		return "", nil
	}
	token, err := s.JWTProvider.GenerateToken(data.Email, data.OAuthProvider)
	return token, nil
}

func (s *AuthService) ValidateToken(token string) (*domain.User, error) {
	email, err := s.JWTProvider.ValidateToken(token)
	if err != nil {
		fmt.Println("Error validating token: ", err)
		return nil, err
	}
	user, err := s.UserRepo.GetUserByEmail(email)
	if err != nil {
		fmt.Println("Error getting user : ", err)
		return nil, err
	}
	return user, nil
}
