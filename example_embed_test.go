package eventemitter

type SampleServer struct {
	EventEmitter
}

func NewServer() *Server {
	s := new(Server)

	// Initialize Maps
	s.EventEmitter.Init()
	return s
}

func ExampleEventEmitter_Init() {
}
