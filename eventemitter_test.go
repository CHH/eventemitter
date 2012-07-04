
package eventemitter

import(
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

    e := <- s.Emit("recv")

    if res := e.Result.(string); res != "bar" {
        t.Errorf("Expected %s, got %s", "bar", res)
    }
}

func TestEmitReturnsChan(t *testing.T) {
	emitter := NewEventEmitter()

	emitter.On("foo", func(event *Event) {
	})

    e := <- emitter.Emit("foo")

    if (e.Name != "foo") {
        t.Errorf("Expected event name %s, got %s", "foo", e.Name)
    }
}

