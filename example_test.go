package eventemitter

import (
	"fmt"
)

func ExampleEventEmitter() {
	// Construct a new EventEmitter instance
	emitter := NewEventEmitter()

	emitter.On("hello", func(event *Event) {
		fmt.Println("Hello World")
	})

	emitter.On("hello", func(event *Event) {
		fmt.Println("Hello Hello World")
	})

	// Wait until all handlers have finished
	<-emitter.Emit("hello")
	// Output:
	// Hello World
	// Hello Hello World
}

func ExampleEventEmitter_Emit() {
	emitter := NewEventEmitter()

	emitter.On("hello", func(event *Event) {
		fmt.Printf("Hello World %s\n", event.Argv[0].(string))
	})

	<-emitter.Emit("hello", "John")
	// Output:
	// Hello World John
}
