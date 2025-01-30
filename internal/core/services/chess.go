package services

import (
	"errors"

	"github.com/zjrb/OpeningTrainer/internal/core/ports"
)

type ChessService struct {
	ChessEngine *ports.ChessEngine
}

func getIndex(white bool, move int) (int, int) {
	if white {
		return ((move - 1) * 3) + 1, move
	} else {
		return ((move - 1) * 3) + 2, move + 1
	}
}

func NewChessService(engine *ports.ChessEngine) *ChessService {
	return &ChessService{
		ChessEngine: engine,
	}
}

func (svc *ChessService) StartGame(opening []string, white bool) string {
	if white {
		return opening[1]
	}
	return ""
}

func (svc *ChessService) PushMove(opening []string, white bool, move int) (string, error) {
	idx, _ := getIndex(white, move)
	if idx > len(opening) {
		return "", errors.New("invalid move")
	}
	return opening[idx], nil
}

func (svc *ChessService) ProcessMove(opening []string, white bool, moveNum int, move string) (string, error) {
	idx, moveNum := getIndex(white, moveNum)
	if opening[idx] == move {
		response, err := svc.PushMove(opening, !white, moveNum)
		if err != nil {
			return "Completed opening!", nil
		}
		return response, nil
	} else {
		return "Incorrect Move, Try Again", errors.New("incorrect move played")
	}
}
