package state

type Background struct {
	width         int
	height        int
	freePositions map[Coordinates]struct{}
}

func (f *Background) Width() int {
	return f.width
}

func (f *Background) Height() int {
	return f.height
}

func (f *Background) FindRandomFreePosition() Coordinates {
	for k := range f.freePositions {
		return k
	}
	panic("no free positions")
}

func (f *Background) TakePosition(position Coordinates) {
	f.freePositions[position] = struct{}{}
}

func (f *Background) FreePosition(position Coordinates) {
	delete(f.freePositions, position)
}

func NewBackground(width int, height int) *Background {
	freePositions := map[Coordinates]struct{}{}

	for w := 0; w < width; w++ {
		for h := 0; h < height; h++ {
			freePositions[Coordinates{X: w, Y: h}] = struct{}{}
		}
	}

	delete(freePositions, Coordinates{X: 1, Y: 1})
	delete(freePositions, Coordinates{X: 2, Y: 1})

	return &Background{
		width:         width,
		height:        height,
		freePositions: freePositions,
	}
}
