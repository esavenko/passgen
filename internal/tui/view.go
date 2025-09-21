package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("63")).
			PaddingBottom(1)

	errorStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("1"))

	passwordStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("10")).
			PaddingTop(1).
			PaddingBottom(1)

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240"))
)

func (m model) View() string {
	if m.err != nil {
		return errorStyle.Render(fmt.Sprintf("Error: %v", m.err))
	}

	s := titleStyle.Render("Password Generator\n")
	s += m.table.View()

	if m.password != "" {
		s += passwordStyle.Render(fmt.Sprintf("\nGenerated Password %s", m.password))
		s += helpStyle.Render("\n\nPress q to quiet")
	} else {
		s += helpStyle.Render(`
Navigation:
	↑/↓: Select option
	←/→: Adjust length
	Space: Toggle option
	Enter: Generate password
	q: Quit
`)
	}

	return s
}
