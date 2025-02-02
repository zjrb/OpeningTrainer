package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zjrb/OpeningTrainer/internal/core/domain"
)

type GameSessionRepo struct {
	pool *pgxpool.Pool
}

func NewGameSessionRepo(pool *pgxpool.Pool) GameSessionRepo {
	return GameSessionRepo{
		pool: pool,
	}
}

func (g *GameSessionRepo) GetGameSession(id int) (*domain.GameSessionDB, error) {
	var gameSession domain.GameSessionDB
	err := g.pool.QueryRow(context.Background(), `
	SELECT opening_id, user_id, wrong_moves, accuracy
	FROM game_session
	WHERE id = $1
	`, id).Scan(&gameSession.OpeningID, &gameSession.UserID, &gameSession.WrongMoves, &gameSession.Accuracy)
	if err != nil {
		return nil, err
	}
	return &gameSession, nil
}
func (g *GameSessionRepo) CreateGameSession(gameSesh *domain.GameSessionDB) error {
	_, err := g.pool.Exec(context.Background(), `
	INSERT INTO game_session (opening_id, user_id)
	VALUES($1, $2)`, gameSesh.OpeningID, gameSesh.UserID)
	return err
}

func (g *GameSessionRepo) UpdateGameSession(gameSesh *domain.GameSessionDB) error {
	commandTag, err := g.pool.Exec(context.Background(), `
	UPDATE game_session 
	SET wrong_moves = $1, accuracy = $2 
	WHERE id = $3`, gameSesh.WrongMoves, gameSesh.Accuracy, gameSesh.ID)
	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("no user found with ID %d", gameSesh.ID)
	}
	return err
}
