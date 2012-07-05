package eventemitter

import (
	"testing"
	"fmt"
)

// Struct for testing Embedding of EventEmitters
type Server struct {
	EventEmitter
}

func TestEmbedding(t *testing.T) {
	s := new(Server)

	// Don't forget to allocate the memory when
	// used as sub type.
	s.EventEmitter.Init()

	s.On("recv", func(event *Event) {
		event.Result = "bar"
	})

	e := <-s.Emit("recv")

	if res := e.Result.(string); res != "bar" {
		t.Errorf("Expected %s, got %s", "bar", res)
	}
}

func TestEmitReturnsChan(t *testing.T) {
	emitter := NewEventEmitter()

	emitter.On("foo", func(event *Event) {
	})

	e := <-emitter.Emit("foo")

	if e.Name != "foo" {
		t.Errorf("Expected event name %s, got %s", "foo", e.Name)
	}
}

func BenchmarkEmit(b *testing.B) {
	b.StopTimer()
	emitter := NewEventEmitter()

	for i := 0; i < 100; i++ {
		emitter.On("hello", func(event *Event) {
			event.Result = "Hello World " + event.Argv[0].(string)
		})
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		<- emitter.Emit("hello", "John")
	}
}

func ExampleEmit() {
	// Construct a new EventEmitter instance
	emitter := NewEventEmitter()

	emitter.On("hello", func(event *Event) {
		fmt.Println("Hello World")
	})

	emitter.On("hello", func(event *Event) {
		fmt.Println("Hello Hello World")
	})

	// Wait until all handlers have finished
	<- emitter.Emit("hello")
	// Output:
	// Hello World
	// Hello Hello World
}
