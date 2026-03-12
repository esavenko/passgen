package settings

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	mainStyle      = lipgloss.NewStyle().MarginLeft(2)
	secondaryStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("230"))
	submainStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	checkboxStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	dotStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("236")).Render(" • ")
	errorStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("196"))
)

func (m *Model) View() string {
	var s string

	s += secondaryStyle.Render("Settings") + "\n\n"

	// Length field
	lengthLine := fmt.Sprintf("Length: %s", m.textInput.View())
	if m.cursor == 0 {
		s += checkboxStyle.Render(lengthLine)
	} else {
		s += lengthLine
	}
	s += "\n"

	if m.inputError != "" {
		s += errorStyle.Render(m.inputError) + "\n"
	}

	// Digits toggle
	s += toggleCheckbox("Use digits", m.cfg.UseDigits, m.cursor == 1) + "\n"

	// Special symbols toggle
	s += toggleCheckbox("Use special symbols", m.cfg.UseSpecialSymbols, m.cursor == 2) + "\n"

	s += "\n"
	s += submainStyle.Render("up/down: navigation") + dotStyle +
		submainStyle.Render("enter/space: toggle") + dotStyle +
		submainStyle.Render("esc: back")

	return mainStyle.Render("\n" + s + "\n\n")
}

func toggleCheckbox(label string, checked bool, active bool) string {
	prefix := "[ ] "
	if checked {
		prefix = "[x] "
	}

	if active {
		return checkboxStyle.Render(prefix + label)
	}

	return fmt.Sprintf("%s%s", prefix, label)
}
