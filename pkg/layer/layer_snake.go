package layer

import (
	"github.com/faiface/pixel"
	"snake/pkg/sprite"
	"snake/pkg/state"
)

type SnakeLayer struct {
	gameState     *state.Game
	spriteFactory *sprite.Factory
	snakeSprites  []*pixel.Sprite
	headMap       map[state.Direction]sprite.Name
	tailMap       map[state.Direction]sprite.Name
	middleMap     map[state.Direction]map[state.Direction]sprite.Name
	AbstractLayer
}

func (s *SnakeLayer) Update(notification *state.Notification) {
	if notification.EventOccurred(state.EVENT_NEW_GAME) {
		snakeSprites := []*pixel.Sprite{
			s.spriteFactory.CreateSprite(sprite.SNAKE_TAIL_RIGHT),
			s.spriteFactory.CreateSprite(sprite.SNAKE_HEAD_RIGHT),
		}
		s.updateBatch(snakeSprites)

	} else if notification.EventOccurred(state.EVENT_SNAKE_MOVED) {
		snakeSprites := s.createSnakeSprites()
		s.batch.Clear()
		s.updateBatch(snakeSprites)
	}
}

func (s *SnakeLayer) createSnakeSprites() []*pixel.Sprite {
	snakeSprites := make([]*pixel.Sprite, 0)
	piece := s.gameState.Snake.Tail.PrevSnakePiece
	snakeSprites = append(s.snakeSprites, s.findTailSprite())

	for piece.PrevSnakePiece != nil {
		snakeSprites = append(snakeSprites, s.findMiddleSprite(piece))
		piece = piece.PrevSnakePiece
	}

	snakeSprites = append(snakeSprites, s.findHeadSprite())
	return snakeSprites
}

func (s *SnakeLayer) updateBatch(snakeSprites []*pixel.Sprite) {
	snakePiece := s.gameState.Snake.Tail

	for _, snakeSprite := range snakeSprites {
		vec := pixel.V(24+float64(snakePiece.Position.X)*48, 24+float64(snakePiece.Position.Y)*48)
		snakeSprite.Draw(s.batch, pixel.IM.Scaled(pixel.ZV, 3).Moved(vec))
		snakePiece = snakePiece.PrevSnakePiece
	}
}

func (s *SnakeLayer) findMiddleSprite(snakePiece *state.SnakePiece) *pixel.Sprite {
	prev := snakePiece.Position.RelativePosition(snakePiece.PrevSnakePiece.Position)
	next := snakePiece.Position.RelativePosition(snakePiece.NextSnakePiece.Position)
	return s.spriteFactory.CreateSprite(s.middleMap[prev][next])
}

func (s *SnakeLayer) findTailSprite() *pixel.Sprite {
	current := s.gameState.Snake.Tail.Position
	prev := s.gameState.Snake.Tail.PrevSnakePiece.Position
	direction := current.RelativePosition(prev)
	return s.spriteFactory.CreateSprite(s.tailMap[direction])
}

func (s *SnakeLayer) findHeadSprite() *pixel.Sprite {
	current := s.gameState.Snake.Head.Position
	next := s.gameState.Snake.Head.NextSnakePiece.Position
	direction := next.RelativePosition(current)
	return s.spriteFactory.CreateSprite(s.headMap[direction])
}

func NewSnakeLayer(gameState *state.Game, spriteFactory *sprite.Factory, batch *pixel.Batch) *SnakeLayer {
	headMap := map[state.Direction]sprite.Name{
		state.DIRECTION_UP:    sprite.SNAKE_HEAD_UP,
		state.DIRECTION_DOWN:  sprite.SNAKE_HEAD_DOWN,
		state.DIRECTION_LEFT:  sprite.SNAKE_HEAD_LEFT,
		state.DIRECTION_RIGHT: sprite.SNAKE_HEAD_RIGHT,
	}

	tailMap := map[state.Direction]sprite.Name{
		state.DIRECTION_UP:    sprite.SNAKE_TAIL_UP,
		state.DIRECTION_DOWN:  sprite.SNAKE_TAIL_DOWN,
		state.DIRECTION_LEFT:  sprite.SNAKE_TAIL_LEFT,
		state.DIRECTION_RIGHT: sprite.SNAKE_TAIL_RIGHT,
	}

	middleMap := map[state.Direction]map[state.Direction]sprite.Name{
		state.DIRECTION_LEFT: {
			state.DIRECTION_UP:    sprite.SNAKE_BODY_TOP_LEFT,
			state.DIRECTION_RIGHT: sprite.SNAKE_BODY_HORIZONTAL,
			state.DIRECTION_DOWN:  sprite.SNAKE_BODY_BOTTOM_LEFT,
		},
		state.DIRECTION_RIGHT: {
			state.DIRECTION_UP:   sprite.SNAKE_BODY_TOP_RIGHT,
			state.DIRECTION_LEFT: sprite.SNAKE_BODY_HORIZONTAL,
			state.DIRECTION_DOWN: sprite.SNAKE_BODY_BOTTOM_RIGHT,
		},
		state.DIRECTION_UP: {
			state.DIRECTION_RIGHT: sprite.SNAKE_BODY_TOP_RIGHT,
			state.DIRECTION_LEFT:  sprite.SNAKE_BODY_TOP_LEFT,
			state.DIRECTION_DOWN:  sprite.SNAKE_BODY_VERTICAL,
		},
		state.DIRECTION_DOWN: {
			state.DIRECTION_RIGHT: sprite.SNAKE_BODY_BOTTOM_RIGHT,
			state.DIRECTION_LEFT:  sprite.SNAKE_BODY_BOTTOM_LEFT,
			state.DIRECTION_UP:    sprite.SNAKE_BODY_VERTICAL,
		},
	}

	snakeLayer := &SnakeLayer{
		gameState:     gameState,
		spriteFactory: spriteFactory,
		snakeSprites:  make([]*pixel.Sprite, 0),
		headMap:       headMap,
		tailMap:       tailMap,
		middleMap:     middleMap,
	}
	snakeLayer.batch = batch
	return snakeLayer
}
