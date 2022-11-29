package state

type Rabbit struct {
	position Coordinates
	notifier *Notifier
}

func (r *Rabbit) Position() Coordinates {
	return r.position
}

func (r *Rabbit) SetPosition(position Coordinates) {
	r.notifier.SetEvent(EVENT_RABBIT_EATEN)
	r.position = position
}

func NewRabbit(position Coordinates, notifier *Notifier) *Rabbit {
	return &Rabbit{position: position, notifier: notifier}
}
