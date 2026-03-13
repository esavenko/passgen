package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/esavenko/passgen/internal/app"
)

var (
	version = "dev"     //nolint:unused // injected via ldflags by goreleaser
	commit  = "none"    //nolint:unused // injected via ldflags by goreleaser
	date    = "unknown" //nolint:unused // injected via ldflags by goreleaser
)

func main() {
	initialModel := app.GetAppModel()

	p := tea.NewProgram(initialModel)
	if _, err := p.Run(); err != nil {
		fmt.Println("Could not start program", err)
	}
}
