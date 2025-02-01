package services

import (
	"github.com/zjrb/OpeningTrainer/internal/core/ports"
	"github.com/zjrb/OpeningTrainer/internal/logger"
)

type ChessService struct {
	ChessEngine ports.ChessEngine
	Cache       ports.GameCache
	l           *logger.Logger
}

func NewChessService(engine *ports.ChessEngine, cache *ports.GameCache, l *logger.Logger) *ChessService {
	return &ChessService{
		ChessEngine: *engine,
		Cache:       *cache,
		l:           l,
	}
}

func (svc *ChessService) handleMessage(message string) string {
	opening, err := svc.Cache.GetOpening(message)
	if err != nil {
		svc.l.Error("Error Getting Opening", err)
	}

}

func (svc *ChessService) StartGame(opening []string, white bool) string {
	return svc.ChessEngine.StartGame(opening, white)
}

func (svc *ChessService) PushMove(opening []string, white bool, move int) (string, error) {
	return svc.ChessEngine.PushMove(opening, white, move)
}

func (svc *ChessService) ProcessMove(opening []string, white bool, moveNum int, move string) (string, error) {
	return svc.ChessEngine.ProcessMove(opening, white, moveNum, move)
}
