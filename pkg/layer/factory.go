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
)

type Factory struct {
	gameState     *state.Game
	spriteFactory *sprite.Factory
	spriteSheet   pixel.Picture
	notifier      *state.Notifier
}

func (f *Factory) CreateLayer(layerName Name) Layer {
	batch := pixel.NewBatch(&pixel.TrianglesData{}, f.spriteSheet)
	switch layerName {
	case BACKGROUND:
		background := NewBackgroundLayer(f.gameState, f.spriteFactory, batch)
		f.notifier.Subscribe(background, state.EVENT_NEW_GAME)
		return background
	case SNAKE:
		snake := NewSnakeLayer(f.gameState, f.spriteFactory, batch)
		f.notifier.Subscribe(snake, state.EVENT_NEW_GAME, state.EVENT_SNAKE_MOVED)
		return snake
	case RABBIT:
		rabbit := NewRabbitLayer(f.gameState, f.spriteFactory, batch)
		f.notifier.Subscribe(rabbit, state.EVENT_NEW_GAME, state.EVENT_RABBIT_EATEN)
		return rabbit
	default:
		panic(fmt.Sprint("layer name provided does not exist: ", layerName))
	}
}

func NewFactory(gameState *state.Game, spriteFactory *sprite.Factory, spriteSheet pixel.Picture, notifier *state.Notifier) *Factory {
	return &Factory{
		gameState:     gameState,
		spriteFactory: spriteFactory,
		spriteSheet:   spriteSheet,
		notifier:      notifier,
	}
}
