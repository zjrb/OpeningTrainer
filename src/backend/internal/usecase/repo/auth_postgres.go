package repo

import "backend/pkg/postgres"

type AuthRepo struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) *AuthRepo {
	return &AuthRepo{
		Postgres: pg,
	}
}
