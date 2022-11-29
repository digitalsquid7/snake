package command

import "snake/pkg/state"

type ChangeDirection struct {
	state     *state.Game
	direction state.Direction
}

func (c ChangeDirection) Execute() {
	current := c.state.Snake.Head.Position
	next := c.state.Snake.Head.NextSnakePiece.Position
	if current.RelativePosition(next) != c.direction {
		c.state.Snake.Direction = c.direction
	}
}

func NewChangeDirection(state *state.Game, direction state.Direction) *ChangeDirection {
	return &ChangeDirection{
		state:     state,
		direction: direction,
	}
}
