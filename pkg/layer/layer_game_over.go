package layer

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
	"snake/pkg/state"
)

type GameOverLayer struct {
	gameState *state.Game
	batch     *pixel.Batch
}

func (l *GameOverLayer) Update(notification *state.Notification) {

}

func (l *GameOverLayer) Render(window *pixelgl.Window) {
	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt := text.New(pixel.V(100, 500), basicAtlas)
	fmt.Fprintln(basicTxt, "Snake Died :(")
	basicTxt.Draw(window, pixel.IM)
}

func NewGameOverLayer(gameState *state.Game, batch *pixel.Batch) *GameOverLayer {
	gameOverLayer := &GameOverLayer{
		gameState: gameState,
	}
	gameOverLayer.batch = batch
	return gameOverLayer
}
