package engine

import "errors"

type ChessEngine struct {
}

func getIndex(white bool, move int) (int, int) {
	if white {
		return ((move - 1) * 3) + 1, move
	} else {
		return ((move - 1) * 3) + 2, move + 1
	}
}
func (c *ChessEngine) StartGame(opening []string, white bool) string {
	if white {
		return opening[1]
	}
	return ""
}
func (c *ChessEngine) PushMove(opening []string, white bool, move int) (string, error) {
	idx, _ := getIndex(white, move)
	if idx > len(opening) {
		return "", errors.New("invalid move")
	}
	return opening[idx], nil
}

func (c *ChessEngine) ProcessMove(opening []string, white bool, moveNum int, move string) (string, error) {
	idx, moveNum := getIndex(white, moveNum)
	if opening[idx] == move {
		response, err := c.PushMove(opening, !white, moveNum)
		if err != nil {
			return "Completed opening!", nil
		}
		return response, nil
	} else {
		return "Incorrect Move, Try Again", errors.New("incorrect move played")
	}
}
