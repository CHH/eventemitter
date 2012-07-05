package eventemitter

import(
	"reflect"
)

type Event struct {
	Name   string
	Argv   []interface{}
	Result interface{}
}

type EventEmitter struct {
	Events map[string][]reflect.Value
}

func NewEventEmitter() *EventEmitter {
	e := new(EventEmitter)
	e.Init()

	return e
}

// Allocates the EventEmitters memory. Has to be called when
// embedding an EventEmitter in another Type.
func (self *EventEmitter) Init() {
	self.Events = make(map[string][]reflect.Value)
}

// Alias to AddListener.
func (self *EventEmitter) On(event string, listener interface{}) {
	self.AddListener(event, listener)
}

// AddListener adds an event listener on the given event name.
func (self *EventEmitter) AddListener(event string, listener interface{}) {
	// Check if the event exists, otherwise initialize the list
	// of handlers for this event.
	if _, exists := self.Events[event]; !exists {
		self.Events[event] = []reflect.Value{}
	}
	
	if l, ok := listener.(reflect.Value); ok {
		self.Events[event] = append(self.Events[event], l)
	} else {
		l := reflect.ValueOf(listener)
		self.Events[event] = append(self.Events[event], l)
	}
}

// Removes all listeners from the given event.
func (self *EventEmitter) RemoveListeners(event string) {
	delete(self.Events, event)
}

// Emits the given event. Puts all arguments following the event name
// into the Event's `Argv` member. Returns a channel if listeners were
// called, nil otherwise.
func (self *EventEmitter) Emit(event string, argv ...interface{}) <- chan []interface{} {
	listeners, exists := self.Events[event]

	if !exists {
		return nil
	}

	var callArgv []reflect.Value
	c := make(chan []interface{})

	for _, a := range argv {
		callArgv = append(callArgv, reflect.ValueOf(a))
	}

	for _, listener := range listeners {
		go func() {
			retVals := listener.Call(callArgv)

			response := []interface{}{}

			for _, r := range retVals {
				response = append(response, r.Interface())
			}

			c <- response
		}()
	}

	return c
}

