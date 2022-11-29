package sprite

import (
	"fmt"
	"github.com/faiface/pixel"
)

type Name string

const (
	GRASS                   Name = "Grass"
	RABBIT                       = "Rabbit"
	SNAKE_BODY_HORIZONTAL        = "Snake Body Horizontal"
	SNAKE_BODY_VERTICAL          = "Snake Body Vertical"
	SNAKE_BODY_TOP_LEFT          = "Snake Body Top Left"
	SNAKE_BODY_TOP_RIGHT         = "Snake Body Top Right"
	SNAKE_BODY_BOTTOM_LEFT       = "Snake Body Bottom Left"
	SNAKE_BODY_BOTTOM_RIGHT      = "Snake Body Bottom Right"
	SNAKE_HEAD_UP                = "Snake Head Up"
	SNAKE_HEAD_DOWN              = "Snake Head Down"
	SNAKE_HEAD_LEFT              = "Snake Head Left"
	SNAKE_HEAD_RIGHT             = "Snake Head Right"
	SNAKE_TAIL_UP                = "Snake Tail Up"
	SNAKE_TAIL_DOWN              = "Snake Tail Down"
	SNAKE_TAIL_LEFT              = "Snake Tail Left"
	SNAKE_TAIL_RIGHT             = "Snake Tail Right"
)

type Factory struct {
	SpriteSheet  pixel.Picture
	SpriteWidth  float64
	SpriteHeight float64
	spriteMap    map[Name]*pixel.Sprite
}

func (f *Factory) CreateSprite(spriteName Name) *pixel.Sprite {

	sprite, ok := f.spriteMap[spriteName]

	if ok {
		return sprite
	}

	panic(fmt.Sprintf("provided sprite name does not exist: %s", spriteName))
}

func (f *Factory) createRect(offsetX float64, offsetY float64) pixel.Rect {
	bounds := f.SpriteSheet.Bounds()
	minX := bounds.Min.X + (offsetX * f.SpriteWidth)
	minY := bounds.Min.Y + ((offsetY) * f.SpriteWidth)
	maxX := bounds.Min.X + ((offsetX + 1) * f.SpriteWidth)
	maxY := bounds.Min.Y + ((offsetY + 1) * f.SpriteWidth)
	return pixel.R(minX, minY, maxX, maxY)
}

func NewFactory(spriteSheet pixel.Picture, spriteWidth float64, spriteHeight float64) *Factory {
	coordinatesMap := map[Name][2]int{
		GRASS:                   {3, 0},
		RABBIT:                  {2, 0},
		SNAKE_BODY_HORIZONTAL:   {1, 0},
		SNAKE_BODY_VERTICAL:     {0, 0},
		SNAKE_BODY_TOP_LEFT:     {3, 1},
		SNAKE_BODY_TOP_RIGHT:    {0, 1},
		SNAKE_BODY_BOTTOM_LEFT:  {2, 1},
		SNAKE_BODY_BOTTOM_RIGHT: {1, 1},
		SNAKE_HEAD_UP:           {0, 3},
		SNAKE_HEAD_DOWN:         {2, 3},
		SNAKE_HEAD_LEFT:         {3, 3},
		SNAKE_HEAD_RIGHT:        {1, 3},
		SNAKE_TAIL_UP:           {0, 2},
		SNAKE_TAIL_DOWN:         {2, 2},
		SNAKE_TAIL_LEFT:         {3, 2},
		SNAKE_TAIL_RIGHT:        {1, 2},
	}

	spriteMap := make(map[Name]*pixel.Sprite, 0)

	for name, coor := range coordinatesMap {
		bounds := spriteSheet.Bounds()
		minX := bounds.Min.X + (float64(coor[0]) * spriteWidth)
		minY := bounds.Min.Y + ((float64(coor[1])) * spriteWidth)
		maxX := bounds.Min.X + ((float64(coor[0]) + 1) * spriteWidth)
		maxY := bounds.Min.Y + ((float64(coor[1]) + 1) * spriteWidth)
		pixel.NewSprite(spriteSheet, pixel.R(minX, minY, maxX, maxY))
		spriteMap[name] = pixel.NewSprite(spriteSheet, pixel.R(minX, minY, maxX, maxY))
	}

	return &Factory{
		SpriteSheet:  spriteSheet,
		SpriteWidth:  spriteWidth,
		SpriteHeight: spriteHeight,
		spriteMap:    spriteMap,
	}
}
