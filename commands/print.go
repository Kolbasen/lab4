package commands

import (
	"fmt"

	"github.com/Kolbasen/lab4/engine"
)

// PrintCommand - represends command for printing input
type PrintCommand struct {
	Arg string
}

// Execute - prints input into console.
func (p *PrintCommand) Execute(loop engine.Handler) {
	fmt.Println(p.Arg)
}
