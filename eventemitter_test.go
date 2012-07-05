package eventemitter

import (
	"fmt"
	"testing"
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

func ExampleEmitReturnsEventOnChan() {
	emitter := NewEventEmitter()

	emitter.On("foo", func(event *Event) {
	})

	e := <-emitter.Emit("foo")

	fmt.Println(e.Name)
	// Output:
	// foo
}

func BenchmarkEmit(b *testing.B) {
	b.StopTimer()
	emitter := NewEventEmitter()
	nListeners := 100

	for i := 0; i < nListeners; i++ {
		emitter.On("hello", func(event *Event) {
			event.Result = "Hello World " + event.Argv[0].(string)
		})
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		<-emitter.Emit("hello", "John")
	}
}
