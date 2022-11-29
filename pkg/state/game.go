package state

type Game struct {
	Background *Background
	Snake      *Snake
	Rabbit     *Rabbit
}

func NewGame(notifier *Notifier) *Game {
	bg := NewBackground(15, 15)
	return &Game{
		Background: bg,
		Snake:      NewSnake(notifier),
		Rabbit:     NewRabbit(bg.FindRandomFreePosition(), notifier),
	}
}
