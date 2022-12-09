package layer

import (
	"fmt"
	"github.com/faiface/pixel"
	"snake/pkg/sprite"
	"snake/pkg/state"
)

type Name string

const (
	BACKGROUND Name = "BACKGROUND"
	SNAKE           = "SNAKE"
	RABBIT          = "RABBIT"
	GAME_OVER       = "GAME OVER"
)

type Factory struct {
	layers        map[Name]*Layer
	gameState     *state.Game
	spriteFactory *sprite.Factory
	spriteSheet   pixel.Picture
}

func (f *Factory) CreateLayer(layerName Name) Layer {
	batch := pixel.NewBatch(&pixel.TrianglesData{}, f.spriteSheet)
	switch layerName {
	case BACKGROUND:
		return NewBackgroundLayer(f.gameState, f.spriteFactory, batch)
	case SNAKE:
		return NewSnakeLayer(f.gameState, f.spriteFactory, batch)
	case RABBIT:
		return NewRabbitLayer(f.gameState, f.spriteFactory, batch)
	case GAME_OVER:
		return NewGameOverLayer(f.gameState, batch)
	default:
		panic(fmt.Sprint("layer name provided does not exist: ", layerName))
	}
}

func NewFactory(gameState *state.Game, spriteFactory *sprite.Factory, spriteSheet pixel.Picture) *Factory {
	return &Factory{
		gameState:     gameState,
		spriteFactory: spriteFactory,
		spriteSheet:   spriteSheet,
	}
}
