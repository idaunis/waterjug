package gauge

import (
	"errors"
)

type Simulation struct {
	X      *Jug
	Y      *Jug
	Target int
	Stream chan *Simulation
	Ack    chan bool

	displayEnabled bool
	stateCount     int
}

func NewSimulation(x, y, target int) *Simulation {
	return &Simulation{
		X:      NewJug(x),
		Y:      NewJug(y),
		Target: target,
		Stream: make(chan *Simulation),
		Ack:    make(chan bool),
	}
}

func (s *Simulation) swappedCopy() *Simulation {
	return &Simulation{
		X:      s.Y,
		Y:      s.X,
		Target: s.Target,
		Stream: s.Stream,
		Ack:    s.Ack,
	}
}

func (s *Simulation) hasSolution() bool {
	// there is no jug big enough to hold the target measure
	if s.Target > s.X.Max && s.Target > s.Y.Max {
		return false
	}
	// target must be divisible by the gcd of X and Y to be solvable
	if s.Target%gcd(s.X.Max, s.Y.Max) != 0 {
		return false
	}
	return true
}

func (s *Simulation) updateState() {
	s.stateCount++
	if !s.displayEnabled {
		return
	}
	// send the current state to the display stream
	s.Stream <- s
	// wait to be acknowledged
	<-s.Ack
}

func (s *Simulation) close() {
	if !s.displayEnabled {
		return
	}
	close(s.Stream)
}

func (s *Simulation) run() int {
	defer s.close()

	s.Y.Empty()
	s.X.Fill()
	s.updateState()

	for s.X.Size != s.Target && s.Y.Size != s.Target {
		if s.X.IsEmpty() {
			s.X.Fill()
			s.updateState()
		}

		if s.Y.IsFilled() {
			s.Y.Empty()
			s.updateState()
		}

		s.X.Transfer(s.Y)
		s.updateState()
	}

	return s.stateCount
}

func (s *Simulation) Simulate() error {
	if !s.hasSolution() {
		return errors.New("No Solution")
	}

	// calculate most efficient way (less state changes) by running silent first
	swap := s.swappedCopy()
	if swap.run() < s.run() {
		s = swap
	}

	// display each state change
	s.displayEnabled = true

	// display first them empty jugs
	s.X.Empty()
	s.Y.Empty()
	s.updateState()

	// run simulation displaying the state changes
	s.run()

	return nil
}
