package ui

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/idaunis/waterjug/gauge"
)

type textColor int32

const (
	None  textColor = 0
	Reset textColor = 1
	Green textColor = 2

	spacing = 2
)

var (
	drawingLineState = []map[gauge.JugState]string{
		{gauge.Empty: "┏   ┓", gauge.Half: "┏   ┓", gauge.Full: "┏▃▃▃┓"},
		{gauge.Empty: "┃   ┃", gauge.Half: "┃   ┃", gauge.Full: "┃███┃"},
		{gauge.Empty: "┃   ┃", gauge.Half: "┃███┃", gauge.Full: "┃███┃"},
		{gauge.Empty: "┗━━━┛", gauge.Half: "┗━━━┛", gauge.Full: "┗━━━┛"},
	}

	colors = map[textColor]string{
		Green: "\033[32m",
		Reset: "\033[0m",
	}
)

func readInt(reader *bufio.Reader, msg string) (int, error) {
	fmt.Printf(msg + ": ")
	str, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	val, err := strconv.Atoi(strings.Trim(str, "\n"))
	if err != nil || val <= 0 {
		return val, errors.New("the value must be greater than zero")
	}
	return val, err
}

func inputInt(reader *bufio.Reader, msg string) int {
	for {
		if x, err := readInt(reader, msg); err != nil {
			fmt.Println(err)
		} else {
			return x
		}
	}
}

func calcMaxWidth(jugs []*gauge.Jug) int {
	max := utf8.RuneCountInString(drawingLineState[0][gauge.Empty])
	for _, j := range jugs {
		display := fmt.Sprintf("%d/%d", j.Max, j.Max)
		if len(display) > max {
			max = len(display)
		}
	}
	return max
}

func drawJugLine(j *gauge.Jug, line int, width int, color textColor) string {
	var display string
	if line == len(drawingLineState) {
		display = fmt.Sprintf("%d/%d", j.Size, j.Max)
	} else if line >= 0 && line < len(drawingLineState) {
		display = drawingLineState[line][j.GetState()]
	}

	count := utf8.RuneCountInString(display)
	padding := strings.Repeat(" ", (width-count)/2)

	if color != None {
		return padding + colors[color] + display + colors[Reset] + padding
	}
	return padding + display + padding
}

func renderJugs(jugs []*gauge.Jug, target int) {
	var out string
	width := calcMaxWidth(jugs) + spacing
	for line := 0; line <= len(drawingLineState); line++ {
		for _, jug := range jugs {
			color := None
			if jug.Size == target {
				color = Green
			}
			out += drawJugLine(jug, line, width, color)
		}
		out += "\n"
	}
	out += "\n"
	fmt.Print(out)
}

func Render(sim *gauge.Simulation) {
	for range sim.Stream {
		renderJugs([]*gauge.Jug{sim.X, sim.Y}, sim.Target)
		sim.Ack <- true
	}
}

func Input() *gauge.Simulation {
	reader := bufio.NewReader(os.Stdin)
	x := inputInt(reader, "Enter size of jug X")
	y := inputInt(reader, "Enter size of jug Y")
	z := inputInt(reader, "Enter target measure Z")
	return gauge.NewSimulation(x, y, z)
}
