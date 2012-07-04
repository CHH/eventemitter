# An EventEmitter for Go

## Usage

A new EventEmitter is created by the `NewEventEmitter` function.

    import ee "github.com/CHH/eventemitter"

    func main() {
        emitter := ee.NewEventEmitter()
    }

A listener is of type `func (event *ee.Event)`
Listeners can be bound to event names with the `On` method:

    emitter.On("foo", func(event *ee.Event) {
        fmt.Printf("Received event '%s'.", event.Name)
    })

An event can be triggered by calling the `Emit` method:

    <- emitter.Emit("foo")

When `Emit` is called, each registered listener is called in
its own Goroutine. They all share a common channel, which is
returned by the `Emit` function.

This channel can be used to trigger events synchronously:

    // Waits until all events have finished
    <- emitter.Emit("foo")

Each listener sends an `*Event` out on the channel when he's finished.
This can be used to run some code everytime a listener has returned:

    c := emitter.Emit("foo")

    for event := <- c {
        // Do something
    }

### Embedding

EventEmitters can also be embedded in other types. When embedding you've
to call the `Init` function on the EventEmitter, so the memory is
correctly allocated:

    type Server struct {
        ee.EventEmitter
    }

    func NewServer() *Server {
        s := new(Server)

        // Allocates the EventEmitter's memory.
        s.EventEmitter.Init()
    }

