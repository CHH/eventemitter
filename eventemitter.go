
package eventemitter

import(
    "fmt"
)

type Event struct {
	Name	string
	Argv	[]interface{}
    Result  interface{}
}

type EventListener func(event *Event)

type EventError struct {
	EventName string
	Message string
}

func (self EventError) Error() string {
	return fmt.Sprintf("%s (Event: '%s')", self.Message, self.EventName)
}

type EventEmitter struct {
	Events map[string][]EventListener
}

func NewEventEmitter() *EventEmitter {
    e := new(EventEmitter)
    e.Init()

    return e
}

func (self *EventEmitter) Init() {
    self.Events = make(map[string][]EventListener)
}

func (self *EventEmitter) RemoveListeners(event string) {
	delete(self.Events, event)
}

func (self *EventEmitter) On(event string, listener EventListener) {
	// Check if the event exists, otherwise initialize the list
	// of handlers for this event.

	if _, exists := self.Events[event]; !exists {
		self.Events[event] = []EventListener{listener}
	} else {
        self.Events[event] = append(self.Events[event], listener)
    }
}

func (self *EventEmitter) Emit(event string, argv ...interface{}) (chan *Event) {
	listeners, exists := self.Events[event]

	if !exists {
		return nil
	}

	c := make(chan *Event)

	for _, listener := range listeners {
		go func() {
            e := &Event{Name: event, Argv: argv}
			listener(e)
			c <- e
		}()
	}

	return c
}

