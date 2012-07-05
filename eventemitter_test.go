package eventemitter

import (
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

	s.On("recv", func(msg string) string {
		return msg
	})

	ret := <-s.Emit("recv", "Hello World")

	expected := "Hello World"

	if res := ret[0].(string); res != expected {
		t.Errorf("Expected %s, got %s", expected, res)
	}
}

func TestEmitReturnsChan(t *testing.T) {
	emitter := NewEventEmitter()

	emitter.On("hello", func(name string) string {
		return "Hello World " + name
	})

	ret := <-emitter.Emit("hello", "John")
	expected := "Hello World John"

	if ret[0].(string) != expected {
		t.Errorf("Expected %s, but got %q", expected, ret)
	}
}

func BenchmarkEmit(b *testing.B) {
	b.StopTimer()
	emitter := NewEventEmitter()

	for i := 0; i < 100; i++ {
		emitter.On("hello", func(name string) string {
			return "Hello World " + name
		})
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		<-emitter.Emit("hello", "John")
	}
}
