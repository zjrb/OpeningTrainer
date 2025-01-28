package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zjrb/OpeningTrainer/internal/core/domain"
)

type UserRepositoryPostgres struct {
	Pool *pgxpool.Pool
}

func NewUserRepositoryPostgres(pool *pgxpool.Pool) *UserRepositoryPostgres {
	return &UserRepositoryPostgres{
		Pool: pool,
	}
}

func (r *UserRepositoryPostgres) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.Pool.QueryRow(context.Background(), `
		SELECT id, email, name
		FROM users
		WHERE email = $1
	`, email).Scan(&user.ID, &user.Email, &user.Name)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryPostgres) CreateUser(user *domain.OAuthResponse) error {
	_, err := r.Pool.Exec(context.Background(), `
		INSERT INTO users (name, email, oauth_provider, oauth_provider_id, profile_picture) 
		VALUES ($1, $2, $3, $4, $5)
	`, user.Name, user.Email, user.OAuthProvider, user.OAuthID, user.ProfilePicture)
	if err != nil {
		return err
	}
	return nil
}
