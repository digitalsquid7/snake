package snake

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	_ "image/png"
	"snake/pkg/command"
	"snake/pkg/layer"
	"snake/pkg/picture"
	"snake/pkg/sprite"
	"snake/pkg/state"
	"snake/pkg/window"
	"time"
)

func Run() {
	cfg := pixelgl.WindowConfig{
		Title:  "SNAKE",
		Bounds: pixel.R(0, 0, 720, 720),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	spriteSheet, err := picture.Load("assets\\images\\snake.png")
	if err != nil {
		panic(err)
	}

	stateNotifier := state.NewNotifier()
	gameState := state.NewGame(stateNotifier)
	spriteFactory := sprite.NewFactory(spriteSheet, 16.0, 16.0)
	layerFactory := layer.NewFactory(gameState, spriteFactory, spriteSheet, stateNotifier)

	layers := []layer.Layer{
		layerFactory.CreateLayer(layer.BACKGROUND),
		layerFactory.CreateLayer(layer.RABBIT),
		layerFactory.CreateLayer(layer.SNAKE),
	}

	timer := time.Tick(time.Millisecond * 200)
	commandFinderFactory := command.NewFinderFactory(win, gameState, timer)

	commandFinders := []command.Finder{
		commandFinderFactory.CreateFinder(command.FINDER_MOVE_SNAKE),
		commandFinderFactory.CreateFinder(command.FINDER_CHANGE_DIRECTION),
	}

	windowRenderer := window.NewRenderer(win)

	var (
		frames = 0
		second = time.Tick(time.Second)
	)

	for !win.Closed() {
		commands := command.Find(commandFinders)
		command.Execute(commands)
		stateNotifier.Notify()
		windowRenderer.Render(layers)

		frames++
		select {
		case <-second:
			win.SetTitle(fmt.Sprintf("%s | FPS: %d", cfg.Title, frames))
			frames = 0
		default:
		}
	}
}
