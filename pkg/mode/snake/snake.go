package snake

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	_ "image/png"
	"path/filepath"
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

	spriteSheet, err := picture.Load(filepath.Join("assets", "images", "snake.png"))
	if err != nil {
		panic(err)
	}

	stateNotifier := state.NewNotifier()
	gameState := state.NewGame(stateNotifier)
	spriteFactory := sprite.NewFactory(spriteSheet, 16.0, 16.0)
	layerFactory := layer.NewFactory(gameState, spriteFactory, spriteSheet)

	timer := time.Tick(time.Millisecond * 200)
	commandFinderFactory := command.NewFinderFactory(win, gameState, timer)

	commandFinders := []command.Finder{
		commandFinderFactory.CreateFinder(command.FINDER_MOVE_SNAKE),
		commandFinderFactory.CreateFinder(command.FINDER_CHANGE_DIRECTION),
	}

	windowRenderer := window.NewRenderer(layerFactory, stateNotifier)
	stateNotifier.Subscribe(windowRenderer, state.EVENT_NEW_GAME, state.EVENT_SNAKE_DIED)

	var (
		frames = 0
		second = time.Tick(time.Second)
	)

	for !win.Closed() {
		commands := command.Find(commandFinders)
		command.Execute(commands)
		stateNotifier.Notify()
		windowRenderer.Render(win)

		frames++
		select {
		case <-second:
			win.SetTitle(fmt.Sprintf("%s | FPS: %d", cfg.Title, frames))
			frames = 0
		default:
		}
	}
}
