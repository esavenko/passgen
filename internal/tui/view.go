package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("230")).
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
			Foreground(lipgloss.Color("230"))
)

func (m Model) View() string {
	if m.err != nil {
		return errorStyle.Render(fmt.Sprintf("Error: %v", m.err))
	}

	view := lipgloss.JoinVertical(
		lipgloss.Left,
		titleStyle.Render(`
	This utility is designed for password generation.
Currently, it supports the number of characters, numbers, and special characters.`),
		m.table.View(),
	) + "\n"

	if m.password != "" {
		view += passwordStyle.Render("Generated password: " + m.password)
		view += helpStyle.Render("\nPress q to quiet")
	} else {
		view += helpStyle.Render(`
Navigation:
	↑/↓: Select option    
	←/→: Adjust length    
	Space: Toggle option    
	Enter: Generate password    
	q: Quit`)
	}

	return lipgloss.NewStyle().MarginLeft(1).Render(view)
}
