package layer

import (
	"github.com/faiface/pixel"
	"snake/pkg/sprite"
	"snake/pkg/state"
)

type RabbitLayer struct {
	gameState     *state.Game
	spriteFactory *sprite.Factory
	AbstractLayer
}

func (s *RabbitLayer) Update(notification *state.Notification) {
	if notification.EventOccurred(state.EVENT_NEW_GAME) {
		rabbit := s.spriteFactory.CreateSprite(sprite.RABBIT)
		position := s.gameState.Rabbit.Position()
		vec := pixel.V(24+float64(position.X)*48, 24+float64(position.Y)*48)
		rabbit.Draw(s.batch, pixel.IM.Scaled(pixel.ZV, 3).Moved(vec))
	} else if notification.EventOccurred(state.EVENT_RABBIT_EATEN) {
		s.batch.Clear()
		rabbit := s.spriteFactory.CreateSprite(sprite.RABBIT)
		position := s.gameState.Rabbit.Position()
		vec := pixel.V(24+float64(position.X)*48, 24+float64(position.Y)*48)
		rabbit.Draw(s.batch, pixel.IM.Scaled(pixel.ZV, 3).Moved(vec))
	}
}

func NewRabbitLayer(gameState *state.Game, spriteFactory *sprite.Factory, batch *pixel.Batch) *RabbitLayer {
	rabbitLayer := &RabbitLayer{
		gameState:     gameState,
		spriteFactory: spriteFactory,
	}
	rabbitLayer.batch = batch
	return rabbitLayer
}
