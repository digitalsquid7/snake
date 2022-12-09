package window

import (
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"snake/pkg/layer"
	"snake/pkg/state"
)

type Renderer struct {
	layerMap       map[layer.Name]layer.Layer
	layersToRender []layer.Layer
	notifier       *state.Notifier
}

func (r *Renderer) Update(notification *state.Notification) {
	if notification.EventOccurred(state.EVENT_NEW_GAME) {
		r.layersToRender = []layer.Layer{
			r.layerMap[layer.BACKGROUND],
			r.layerMap[layer.RABBIT],
			r.layerMap[layer.SNAKE],
		}
		r.notifier.Subscribe(r.layerMap[layer.BACKGROUND], state.EVENT_NEW_GAME)
		r.notifier.Subscribe(r.layerMap[layer.SNAKE], state.EVENT_NEW_GAME, state.EVENT_SNAKE_MOVED)
		r.notifier.Subscribe(r.layerMap[layer.RABBIT], state.EVENT_NEW_GAME, state.EVENT_RABBIT_EATEN)
		//r.notifier.Subscribe(r.layerMap[layer.GAME_OVER], state.EVENT_SNAKE_DIED)

	} else if notification.EventOccurred(state.EVENT_SNAKE_DIED) {
		r.layersToRender = append(r.layersToRender, r.layerMap[layer.GAME_OVER])
		r.notifier.Unsubscribe(r.layerMap[layer.BACKGROUND], r.layerMap[layer.SNAKE], r.layerMap[layer.RABBIT])
	}
}

func (r *Renderer) Render(window *pixelgl.Window) {
	window.Clear(colornames.Black)

	for _, l := range r.layersToRender {
		l.Render(window)
	}

	window.Update()
}

func NewRenderer(layerFactory *layer.Factory, notifier *state.Notifier) *Renderer {
	layerMap := map[layer.Name]layer.Layer{
		layer.BACKGROUND: layerFactory.CreateLayer(layer.BACKGROUND),
		layer.RABBIT:     layerFactory.CreateLayer(layer.RABBIT),
		layer.SNAKE:      layerFactory.CreateLayer(layer.SNAKE),
		layer.GAME_OVER:  layerFactory.CreateLayer(layer.GAME_OVER),
	}

	return &Renderer{layerMap: layerMap, notifier: notifier}
}
