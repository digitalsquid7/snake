package layer

import (
	"github.com/faiface/pixel/pixelgl"
	"snake/pkg/state"
)

type Layer interface {
	state.Subscriber
	Render(window *pixelgl.Window)
}
