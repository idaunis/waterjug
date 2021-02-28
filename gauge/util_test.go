package gauge

import "testing"

func TestGCD(t *testing.T) {
	if got, exp := gcd(3, 6), 3; got != exp {
		t.Errorf("Expected %v but got %v", exp, got)
	}
	if got, exp := gcd(3, 7), 1; got != exp {
		t.Errorf("Expected %v but got %v", exp, got)
	}
	if got, exp := gcd(0, 7), 7; got != exp {
		t.Errorf("Expected %v but got %v", exp, got)
	}
}
