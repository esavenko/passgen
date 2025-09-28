package menu

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/esavenko/passgen/common/constants"
)

var (
	mainStyle      = lipgloss.NewStyle().MarginLeft(2)
	secondaryStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("230"))
	checkboxStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	errorStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("1"))
)

func (m *Model) View() string {
	var s string

	s += secondaryStyle.Render(constants.Logo + "\n")
	s += secondaryStyle.Render("\nThis utility is designed for password generation.\nCurrently, it supports the number of characters, numbers, and special \ncharacters.\n")
	s += secondaryStyle.Render("\nChoose the option:\n")
	s += secondaryStyle.Render("\n" + m.choicesView())

	return mainStyle.Render("\n" + s + "\n\n")
}

func (m *Model) choicesView() string {
	c := m.Choice

	choices := fmt.Sprintf(
		"%s\n%s\n",
		toggleCheckbox("Generate password", c == 0),
		toggleCheckbox("Quit", c == 1),
	)

	return fmt.Sprintf(choices)
}

func toggleCheckbox(label string, checked bool) string {
	if checked {
		return checkboxStyle.Render("[x] " + label)
	}

	return fmt.Sprintf("[ ] %s", label)
}
