package menu

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/esavenko/passgen/internal/constants"
)

const (
	dotChar = " • "
)

var (
	mainStyle      = lipgloss.NewStyle().MarginLeft(2)
	secondaryStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("230"))
	submainStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	checkboxStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	dotStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("236")).Render(dotChar)
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

	tpl := "%s\n\n"
	tpl += submainStyle.Render("up/down: navigation") + dotStyle +
		submainStyle.Render("enter: choose") + dotStyle +
		submainStyle.Render("q: quit")

	choices := fmt.Sprintf(
		"%s\n%s\n",
		toggleCheckbox("Generate password", c == 0),
		toggleCheckbox("Quit", c == 1),
	)

	return fmt.Sprintf(tpl, choices)
}

func toggleCheckbox(label string, checked bool) string {
	if checked {
		return checkboxStyle.Render("[x] " + label)
	}

	return fmt.Sprintf("[ ] %s", label)
}
