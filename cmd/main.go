package main

import (
	"fmt"

	"github.com/idaunis/waterjug/ui"
)

func main() {
	simulation := ui.Input()
	go ui.Render(simulation)
	if err := simulation.Simulate(); err != nil {
		fmt.Println(err)
	}
}
