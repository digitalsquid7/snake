package command

import (
	"fmt"
	"github.com/faiface/pixel/pixelgl"
	"snake/pkg/state"
	"time"
)

type FinderName string

const (
	FINDER_MOVE_SNAKE       FinderName = "Move Snake"
	FINDER_CHANGE_DIRECTION            = "Change Direction"
)

type FinderFactory struct {
	window    *pixelgl.Window
	gameState *state.Game
	timer     <-chan time.Time
}

func (f *FinderFactory) CreateFinder(name FinderName) Finder {
	switch name {
	case FINDER_MOVE_SNAKE:
		return NewMoveSnakeFinder(f.window, f.gameState, f.timer)
	case FINDER_CHANGE_DIRECTION:
		return NewChangeDirectionFinder(f.window, f.gameState)
	default:
		panic(fmt.Sprint("factory name provided does not exist: ", name))
	}
}

func NewFinderFactory(window *pixelgl.Window, gameState *state.Game, timer <-chan time.Time) *FinderFactory {
	return &FinderFactory{
		window:    window,
		gameState: gameState,
		timer:     timer,
	}
}
