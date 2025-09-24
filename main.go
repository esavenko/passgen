package main

import (
	"os"

	"github.com/esavenko/passgen/internal/tui"
)

func main() {
	if err := tui.NewProgram().Start(); err != nil {
		println("Error running program", err.Error())
		os.Exit(1)
	}
}
