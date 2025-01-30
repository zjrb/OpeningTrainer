package engine

type ChessEngine struct {
}

func (c *ChessEngine) StartGame(opening []string, white bool) string {
	if white {
		return opening[1]
	}
	return ""
}

func (c *ChessEngine) ProcessMove() {}
