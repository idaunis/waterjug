package gauge

import (
	"testing"
)

type testStates struct {
	X int
	Y int
}

func TestSimulation(t *testing.T) {
	s := NewSimulation(4, 5, 3)
	go s.Simulate()

	steps := []testStates{{0, 0}, {4, 0}, {0, 4}, {4, 4}, {3, 5}}
	for _, expected := range steps {
		<-s.Stream
		if got, exp := s.X.Size, expected.X; got != exp {
			t.Errorf("Expected %v but got %v", exp, got)
		}
		if got, exp := s.Y.Size, expected.Y; got != exp {
			t.Errorf("Expected %v but got %v", exp, got)
		}
		s.Ack <- true
	}

	s = NewSimulation(3, 5, 4)
	go s.Simulate()

	i := 0
	steps = []testStates{{0, 0}, {0, 5}, {3, 2}, {0, 2}, {2, 0}, {2, 5}, {3, 4}}
	for range s.Stream {
		expected := steps[i]
		if got, exp := s.X.Size, expected.X; got != exp {
			t.Errorf("Expected %v but got %v", exp, got)
		}
		if got, exp := s.Y.Size, expected.Y; got != exp {
			t.Errorf("Expected %v but got %v", exp, got)
		}
		i++
		s.Ack <- true
	}
}

func TestSteps(t *testing.T) {
	if got, exp := NewSimulation(4, 5, 3).run(), 4; got != exp {
		t.Errorf("Expected %v but got %v", exp, got)
	}
	if got, exp := NewSimulation(5, 4, 3).run(), 10; got != exp {
		t.Errorf("Expected %v but got %v", exp, got)
	}
}

func TestHasSolution(t *testing.T) {
	s := NewSimulation(3, 6, 4)
	if got, exp := s.hasSolution(), false; got != exp {
		t.Errorf("Expected %v but got %v", exp, got)
	}
}
