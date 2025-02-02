package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zjrb/OpeningTrainer/internal/core/domain"
)

type OpeningRepoPostgres struct {
	pool *pgxpool.Pool
}

func NewOpeningRepoPostgres(p *pgxpool.Pool) *OpeningRepoPostgres {
	return &OpeningRepoPostgres{
		pool: p,
	}
}

func (r *OpeningRepoPostgres) GetOpeningByName(name string) ([]domain.Opening, error) {
	var openings []domain.Opening
	rows, err := r.pool.Query(context.Background(), `
		SELECT opening_name, pgn, eco 
		FROM openings 
		WHERE LOWER(opening_name) LIKE '%' || $1 || '%'
	`, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var opening domain.Opening
		if err := rows.Scan(&opening.OpeningName, &opening.PGN, &opening.ECO); err != nil {
			return nil, err
		}
		openings = append(openings, opening)
	}

	return openings, nil
}

func (r *OpeningRepoPostgres) GetSingleOpeningById(id int) (*domain.Opening, error) {
	var opening domain.Opening
	err := r.pool.QueryRow(context.Background(), `
	SELECT opening_name, eco, pgn, uci, fen, move_arry
	FROM openings
	WHERE id = $1`, id, &opening.OpeningName, &opening.ECO, &opening.PGN, &opening.UCI, &opening.MoveArray)
	if err != nil {
		return nil, fmt.Errorf("problem with the sql query")
	}
	return &opening, nil
}
