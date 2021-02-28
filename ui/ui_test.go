package ui

import (
	"strings"
	"testing"

	"github.com/idaunis/waterjug/gauge"
)

func testLine(j *gauge.Jug, line int) string {
	return drawJugLine(j, line, 5, None)
}

func TestRenderLine(t *testing.T) {
	j := gauge.NewJug(5)
	j.Fill()

	if got, exp := strings.TrimSpace(testLine(j, -1)), ""; got != exp {
		t.Errorf(`Expected "%s" but got "%s"`, exp, got)
	}

	if got, exp := strings.TrimSpace(testLine(j, 5)), ""; got != exp {
		t.Errorf(`Expected "%s" but got "%s"`, exp, got)
	}

	if got, exp := testLine(j, 2), drawingLineState[2][gauge.Full]; got != exp {
		t.Errorf(`Expected "%s" but got "%s"`, exp, got)
	}

	if got, exp := testLine(j, 4), " 5/5 "; got != exp {
		t.Errorf(`Expected "%s" but got "%s"`, exp, got)
	}

	j.Empty()

	if got, exp := testLine(j, 2), drawingLineState[2][gauge.Empty]; got != exp {
		t.Errorf(`Expected "%s" but got "%s"`, exp, got)
	}
}
