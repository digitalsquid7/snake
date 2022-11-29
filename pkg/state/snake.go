package state

type Snake struct {
	Direction   Direction
	Alive       bool
	Coordinates map[Coordinates]*SnakePiece
	Head        *SnakePiece
	Tail        *SnakePiece
	notifier    *Notifier
	length      int
}

func (s *Snake) GrowSnake() {
	s.length += 1
}

func (s *Snake) Length() int {
	return s.length
}

func (s *Snake) PositionOccupied(coordinates Coordinates) bool {
	_, ok := s.Coordinates[coordinates]
	return ok
}

func (s *Snake) UpdateHeadPosition() {
	newHead := &SnakePiece{
		NextSnakePiece: s.Head,
		PrevSnakePiece: nil,
		Position:       s.findNewHeadPosition(s.Direction),
	}
	s.Head.PrevSnakePiece = newHead
	s.Head = newHead
	if s.PositionOccupied(newHead.Position) {
		s.Alive = false
		s.notifier.SetEvent(EVENT_SNAKE_DIED)
	}
	s.Coordinates[newHead.Position] = newHead
	s.notifier.SetEvent(EVENT_SNAKE_MOVED)
}

func (s *Snake) UpdateTailPosition() {
	newTail := s.Tail.PrevSnakePiece
	delete(s.Coordinates, s.Tail.Position)
	newTail.NextSnakePiece = nil
	s.Tail = newTail
}

func (s *Snake) findNewHeadPosition(direction Direction) Coordinates {
	switch direction {
	case DIRECTION_UP:
		return Coordinates{s.Head.Position.X, s.Head.Position.Y + 1}
	case DIRECTION_DOWN:
		return Coordinates{s.Head.Position.X, s.Head.Position.Y - 1}
	case DIRECTION_RIGHT:
		return Coordinates{s.Head.Position.X + 1, s.Head.Position.Y}
	default:
		return Coordinates{s.Head.Position.X - 1, s.Head.Position.Y}
	}
}

func NewSnake(notifier *Notifier) *Snake {
	head := &SnakePiece{Position: Coordinates{2, 1}}
	tail := &SnakePiece{PrevSnakePiece: head, Position: Coordinates{1, 1}}
	head.NextSnakePiece = tail
	coordinatesMap := map[Coordinates]*SnakePiece{
		head.Position: head,
		tail.Position: tail,
	}

	return &Snake{
		Direction:   DIRECTION_RIGHT,
		Alive:       true,
		Coordinates: coordinatesMap,
		Head:        head,
		Tail:        tail,
		notifier:    notifier,
		length:      2,
	}
}
