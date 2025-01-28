package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	secret string
}

func NewJWT(secret string) *JWT {
	return &JWT{
		secret: secret,
	}
}

func (j *JWT) GenerateToken(email string, oauthprovider string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = jwt.MapClaims{
		"email":         email,
		"oath_provider": oauthprovider,
		"exp":           time.Now().Add(time.Hour * 24).Unix(),
		"iss":           "OpeningTrainer",
	}
	tokenString, err := token.SignedString([]byte(j.secret))
	if err != nil {
		fmt.Println("Error signing token: ", err)
		return "", err
	}

	return tokenString, err
}

func (j *JWT) ValidateToken(token string) (string, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok || !t.Valid {
		return "", err
	}
	email, ok := claims["email"].(string)
	if !ok {
		return "", err
	}
	return email, nil
}
