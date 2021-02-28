package gauge

import "testing"

func TestTransfer(t *testing.T) {
	j1, j2 := NewJug(3), NewJug(9)

	j2.Fill()

	j2.Transfer(j1)

	if got, exp := j1.Size, 3; got != exp {
		t.Errorf("Expected %v but got %v", exp, got)
	}
	if got, exp := j2.Size, 6; got != exp {
		t.Errorf("Expected %v but got %v", exp, got)
	}

	j1.Transfer(j2)

	if got, exp := j1.Size, 0; got != exp {
		t.Errorf("Expected %v but got %v", exp, got)
	}
	if got, exp := j2.Size, 9; got != exp {
		t.Errorf("Expected %v but got %v", exp, got)
	}
}
