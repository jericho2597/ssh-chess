package model

import (
	"github.com/notnil/chess"
)

type Game struct {
	State *chess.Game
}

func NewGame() Game {

	return Game{
		State: chess.NewGame(),
	}
}
