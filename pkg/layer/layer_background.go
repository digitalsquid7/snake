package layer

import (
	"github.com/faiface/pixel"
	"snake/pkg/sprite"
	"snake/pkg/state"
)

type BackgroundLayer struct {
	gameState     *state.Game
	spriteFactory *sprite.Factory
	AbstractLayer
}

func (b *BackgroundLayer) Update(notification *state.Notification) {
	if notification.EventOccurred(state.EVENT_NEW_GAME) {
		grass := b.spriteFactory.CreateSprite(sprite.GRASS)
		for width := 0; width < b.gameState.Background.Width()*48; width += 48 {
			for height := 0; height < b.gameState.Background.Width()*48; height += 48 {
				grass.Draw(b.batch, pixel.IM.Scaled(pixel.ZV, 3).Moved(pixel.V(24+float64(width), 24+float64(height))))
			}
		}
	}
}

func NewBackgroundLayer(gameState *state.Game, factory *sprite.Factory, batch *pixel.Batch) *BackgroundLayer {
	backgroundLayer := &BackgroundLayer{
		gameState:     gameState,
		spriteFactory: factory,
	}
	backgroundLayer.batch = batch
	return backgroundLayer
}
