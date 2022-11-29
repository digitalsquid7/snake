package layer

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"snake/pkg/state"
)

type AbstractLayer struct {
	batch *pixel.Batch
	state.Subscriber
}

func (l *AbstractLayer) Render(window *pixelgl.Window) {
	l.batch.Draw(window)
}
