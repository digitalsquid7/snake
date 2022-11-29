package main

import (
	"github.com/faiface/pixel/pixelgl"
	"snake/pkg/mode/snake"
)

func main() {
	pixelgl.Run(snake.Run)
}
