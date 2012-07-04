# go-ee, an EventEmitter for Go

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

