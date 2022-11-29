package window

import (
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"snake/pkg/layer"
	"snake/pkg/state"
)

type Renderer struct {
	window       *pixelgl.Window
	layerFactory layer.Factory
}

func (r *Renderer) Update(notification *state.Notification) {
	// Adjust layers to be rendered
}

func (r *Renderer) Render(layers []layer.Layer) {
	r.window.Clear(colornames.Black)

	for _, l := range layers {
		l.Render(r.window)
	}

	r.window.Update()
}

func NewRenderer(window *pixelgl.Window) *Renderer {
	return &Renderer{window: window}
}
