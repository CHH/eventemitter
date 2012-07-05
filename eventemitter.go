package eventemitter

type Event struct {
	Name   string
	Argv   []interface{}
	Result interface{}
}

type EventListener func(event *Event)

type EventEmitter struct {
	Events map[string][]EventListener
}

func NewEventEmitter() *EventEmitter {
	e := new(EventEmitter)
	e.Init()

	return e
}

// Allocates the EventEmitters memory. Has to be called when
// embedding an EventEmitter in another Type.
func (self *EventEmitter) Init() {
	self.Events = make(map[string][]EventListener)
}

// Alias to AddListener.
func (self *EventEmitter) On(event string, listener EventListener) {
	self.AddListener(event, listener)
}

// AddListener adds an event listener on the given event name.
func (self *EventEmitter) AddListener(event string, listener EventListener) {
	// Check if the event exists, otherwise initialize the list
	// of handlers for this event.
	if _, exists := self.Events[event]; !exists {
		self.Events[event] = []EventListener{listener}
	} else {
		self.Events[event] = append(self.Events[event], listener)
	}
}

// Removes all listeners from the given event.
func (self *EventEmitter) RemoveListeners(event string) {
	delete(self.Events, event)
}

// Emits the given event. Puts all arguments following the event name
// into the Event's `Argv` member. Returns a channel if listeners were
// called, nil otherwise.
func (self *EventEmitter) Emit(event string, argv ...interface{}) chan *Event {
	listeners, exists := self.Events[event]

	if !exists {
		return nil
	}

	c := make(chan *Event)

	for _, listener := range listeners {
		go func(l EventListener) {
			e := &Event{Name: event, Argv: argv}
			l(e)
			c <- e
		}(listener)
	}

	return c
}
