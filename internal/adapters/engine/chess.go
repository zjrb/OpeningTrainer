package engine

import (
	"github.com/zjrb/OpeningTrainer/internal/core/domain"
)

type ChessEngine struct {
}

func NewChessEngine() *ChessEngine {
	return &ChessEngine{}
}

func (c *ChessEngine) ProcessMove(gs *domain.GameSession) string {
	if gs.Opening[gs.MoveNum] != gs.LastMove {
		return domain.WRONGMOVE
	}
	if len(gs.Opening) == gs.MoveNum+1 {
		return domain.GAMECOMPLETE
	}
	return gs.Opening[gs.MoveNum+1]
}
