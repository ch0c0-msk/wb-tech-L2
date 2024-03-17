package pattern

import "fmt"

// State interface
type State interface {
	Handle(msg string)
}

// Object in the state
type Server struct {
	state State
}

func (s *Server) SetState(state State) {
	s.state = state
}

func (s *Server) Request(msg string) {
	s.state.Handle(msg)
}

// Concrete states
type NormalState struct{}

func (n *NormalState) Handle(msg string) {
	fmt.Printf("Your message \"%s\" has been processed\n", msg)
}

type WrongState struct{}

func (w *WrongState) Handle(msg string) {
	fmt.Println("Something`s wrong")
}
