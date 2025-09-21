package main

import (
	"os"
	"passgen/internal/tui"
)

func main() {
	if err := tui.NewProgram().Start(); err != nil {
		println("Error running program", err.Error())
		os.Exit(1)
	}
}
