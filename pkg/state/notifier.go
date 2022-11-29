package state

import "fmt"

type Event string

const (
	EVENT_NEW_GAME     Event = "New Game"
	EVENT_SNAKE_MOVED        = "Snake Moved"
	EVENT_SNAKE_DIED         = "Snake Died"
	EVENT_RABBIT_EATEN       = "Rabbit Eaten"
)

type Subscriber interface {
	Update(notification *Notification)
}

type Notification struct {
	eventsOccurred map[Event]bool
}

func (n *Notification) EventOccurred(event Event) bool {
	if occurred, ok := n.eventsOccurred[event]; ok {
		return occurred
	}
	panic(fmt.Sprint("subscriber is not subscribed to the event provided: ", event))
}

type Notifier struct {
	subscribers    map[Subscriber]map[Event]struct{}
	eventsHappened map[Event]struct{}
}

func (n *Notifier) Subscribe(subscriber Subscriber, events ...Event) {
	eventSet := map[Event]struct{}{}
	for _, event := range events {
		eventSet[event] = struct{}{}
	}
	n.subscribers[subscriber] = eventSet
}

func (n *Notifier) SetEvent(event Event) {
	n.eventsHappened[event] = struct{}{}
}

func (n *Notifier) Notify() {
	for subscriber, eventsSubscribed := range n.subscribers {
		result := map[Event]bool{}
		for eventSubscribed, _ := range eventsSubscribed {
			if _, ok := n.eventsHappened[eventSubscribed]; ok {
				result[eventSubscribed] = true
			} else {
				result[eventSubscribed] = false
			}
		}
		subscriber.Update(&Notification{eventsOccurred: result})
	}
	n.eventsHappened = make(map[Event]struct{}, 0)
}

func NewNotifier() *Notifier {
	return &Notifier{
		map[Subscriber]map[Event]struct{}{},
		map[Event]struct{}{EVENT_NEW_GAME: {}},
	}
}
