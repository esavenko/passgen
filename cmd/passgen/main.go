package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/esavenko/passgen/internal/app"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	initialModel := app.GetAppModel()

	p := tea.NewProgram(initialModel)
	if _, err := p.Run(); err != nil {
		fmt.Println("Could not start program", err)
	}
}
