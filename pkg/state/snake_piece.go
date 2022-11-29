package state

type SnakePiece struct {
	NextSnakePiece *SnakePiece
	PrevSnakePiece *SnakePiece
	Position       Coordinates
}
