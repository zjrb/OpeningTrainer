package services

import (
	"github.com/zjrb/OpeningTrainer/internal/core/domain"
	"github.com/zjrb/OpeningTrainer/internal/core/ports"
	"github.com/zjrb/OpeningTrainer/internal/logger"
)

type ChessService struct {
	ChessEngine ports.ChessEngine
	Cache       ports.GameCache
	l           *logger.Logger
}

func NewChessService(engine ports.ChessEngine, cache ports.GameCache, l *logger.Logger) *ChessService {
	return &ChessService{
		ChessEngine: engine,
		Cache:       cache,
		l:           l,
	}
}

func (svc *ChessService) HandleMessage(gameSession *domain.GameSession) *domain.GameSession {
	result := svc.ChessEngine.ProcessMove(gameSession)
	if result == domain.GAMECOMPLETE {
		gameSession.Message = domain.GAMECOMPLETE
	} else if result == domain.WRONGMOVE {
		gameSession.Message = domain.WRONGMOVE
		gameSession.MoveNum -= 1
		gameSession.LastMove = ""
	} else {
		gameSession.MoveNum += 1
		gameSession.LastMove = result
	}
	return gameSession
}
