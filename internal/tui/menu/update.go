package menu

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/esavenko/passgen/common/messages"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m.updateChoices(msg)
}

func (m *Model) updateChoices(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "down", "j":
			m.Choice++
			if m.Choice > 1 {
				m.Choice = 1
			}

		case "up", "k":
			m.Choice--
			if m.Choice < 0 {
				m.Choice = 0
			}

		case "enter":
			m.Chosen = true
			if m.Choice == 1 {
				return m, tea.Quit
			}
			return m, func() tea.Msg { return messages.SwitchToGenerationMsg{} }
		}
	}

	return m, nil
}
