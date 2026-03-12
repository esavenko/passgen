package generation

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/esavenko/passgen/internal/messages"
)

type (
	errMsg error
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return m, func() tea.Msg { return messages.SwitchToMenuMsg{} }
		}

	case errMsg:
		m.inputError = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}
