package menu

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/esavenko/passgen/internal/messages"
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
			if m.Choice > 2 {
				m.Choice = 2
			}

		case "up", "k":
			m.Choice--
			if m.Choice < 0 {
				m.Choice = 0
			}

		case "enter":
			switch m.Choice {
			case 0:
				return m, func() tea.Msg { return messages.SwitchToGenerationMsg{} }
			case 1:
				return m, func() tea.Msg { return messages.SwitchToSettingsMsg{} }
			case 2:
				return m, tea.Quit
			}

		case "q":
			return m, tea.Quit
		}
	}

	return m, nil
}
