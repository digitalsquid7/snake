package state

type Coordinates struct {
	X int
	Y int
}

func (c *Coordinates) RelativePosition(other Coordinates) Direction {
	if other.X > c.X {
		return DIRECTION_RIGHT
	} else if other.X < c.X {
		return DIRECTION_LEFT
	} else if other.Y > c.Y {
		return DIRECTION_UP
	}
	return DIRECTION_DOWN
}
