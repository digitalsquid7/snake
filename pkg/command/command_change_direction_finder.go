package command

import (
	"github.com/faiface/pixel/pixelgl"
	"snake/pkg/state"
)

type ChangeDirectionFinder struct {
	window       *pixelgl.Window
	directionMap map[pixelgl.Button]state.Direction
	keysToCheck  []pixelgl.Button
	gameState    *state.Game
}

func (f *ChangeDirectionFinder) Find() Command {
	for _, keyToCheck := range f.keysToCheck {
		if f.window.Pressed(keyToCheck) {
			return NewChangeDirection(f.gameState, f.directionMap[keyToCheck])
		}
	}
	return nil
}

func NewChangeDirectionFinder(window *pixelgl.Window, gameState *state.Game) *ChangeDirectionFinder {
	directionMap := map[pixelgl.Button]state.Direction{
		pixelgl.KeyLeft:  state.DIRECTION_LEFT,
		pixelgl.KeyRight: state.DIRECTION_RIGHT,
		pixelgl.KeyUp:    state.DIRECTION_UP,
		pixelgl.KeyDown:  state.DIRECTION_DOWN,
	}

	keysToCheck := []pixelgl.Button{pixelgl.KeyLeft, pixelgl.KeyRight, pixelgl.KeyUp, pixelgl.KeyDown}

	return &ChangeDirectionFinder{
		window:       window,
		gameState:    gameState,
		directionMap: directionMap,
		keysToCheck:  keysToCheck,
	}
}
