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

func (r *UserRepositoryPostgres) CreateUser(user *domain.User) error {
	_, err := r.Pool.Exec(context.Background(), `
		INSERT INTO users (email, name)
		VALUES ($1, $2)
	`, user.Email, user.Name)
	if err != nil {
		return err
	}
	return nil
}
