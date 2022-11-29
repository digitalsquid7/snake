package command

import (
	"github.com/faiface/pixel/pixelgl"
	"snake/pkg/state"
	"time"
)

type MoveSnakeFinder struct {
	window    *pixelgl.Window
	gameState *state.Game
	timer     <-chan time.Time
}

func (f *MoveSnakeFinder) Find() Command {
	select {
	case <-f.timer:
		return NewMoveSnake(f.gameState)
	default:
		return nil
	}
}

func NewMoveSnakeFinder(window *pixelgl.Window, gameState *state.Game, timer <-chan time.Time) *MoveSnakeFinder {
	return &MoveSnakeFinder{
		window:    window,
		gameState: gameState,
		timer:     timer,
	}
}
