package command

import "snake/pkg/state"

type MoveSnake struct {
	state *state.Game
}

func (c MoveSnake) Execute() {
	c.state.Snake.UpdateHeadPosition()
	//c.state.Background.TakePosition(c.state.Snake.Head.Position)
	if c.state.Snake.Head.Position == c.state.Rabbit.Position() {
		c.state.Rabbit.SetPosition(c.state.Background.FindRandomFreePosition())
		c.state.Snake.GrowSnake()
	} else {
		//c.state.Background.FreePosition(c.state.Snake.Tail.Position)
		c.state.Snake.UpdateTailPosition()
	}
}

func NewMoveSnake(state *state.Game) *MoveSnake {
	return &MoveSnake{state: state}
}
