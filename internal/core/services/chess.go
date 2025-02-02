package services

import (
	"fmt"

	"github.com/zjrb/OpeningTrainer/internal/core/domain"
	"github.com/zjrb/OpeningTrainer/internal/core/ports"
	"github.com/zjrb/OpeningTrainer/internal/logger"
)

type ChessService struct {
	chessEngine ports.ChessEngine
	cache       ports.GameCache
	l           *logger.Logger
	gameRepo    ports.GameSessionRepo
	openingRepo ports.OpeningRepository
}

func NewChessService(engine ports.ChessEngine, cache ports.GameCache, l *logger.Logger, gameRepo ports.GameSessionRepo, openingRepo ports.OpeningRepository) *ChessService {
	return &ChessService{
		chessEngine: engine,
		cache:       cache,
		l:           l,
		gameRepo:    gameRepo,
		openingRepo: openingRepo,
	}
}

func (svc *ChessService) HandleMessage(gameSession *domain.GameSession) *domain.GameSession {
	result := svc.chessEngine.ProcessMove(gameSession)
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
func (svc *ChessService) CreateGameSession(userId int, openingId int, white bool) (*domain.GameSession, error) {
	var gameSesh domain.GameSession
	var gameDbSesh domain.GameSessionDB
	opening, err := svc.openingRepo.GetSingleOpeningById(openingId)
	if err != nil {
		return nil, fmt.Errorf("error getting opening %v", err)
	}
	gameSesh.Opening = opening.MoveArray
	gameSesh.White = white
	gameDbSesh.OpeningID = openingId
	gameDbSesh.UserID = userId
	err = svc.gameRepo.CreateGameSession(&gameDbSesh)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	return &gameSesh, nil

}
