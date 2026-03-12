package settings

import (
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/esavenko/passgen/internal/messages"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return m, func() tea.Msg { return messages.SwitchToMenuMsg{} }

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
				m.updateFocus()
			}
			return m, nil

		case "down", "j":
			if m.cursor < 2 {
				m.cursor++
				m.updateFocus()
			}
			return m, nil

		case "enter", " ":
			if m.cursor == 1 {
				m.cfg.UseDigits = !m.cfg.UseDigits
				return m, nil
			}
			if m.cursor == 2 {
				m.cfg.UseSpecialSymbols = !m.cfg.UseSpecialSymbols
				return m, nil
			}
		}
	}

	if m.cursor == 0 {
		var cmd tea.Cmd
		m.textInput, cmd = m.textInput.Update(msg)
		m.validateLength()
		return m, cmd
	}

	return m, nil
}

func (m *Model) updateFocus() {
	if m.cursor == 0 {
		m.textInput.Focus()
	} else {
		m.textInput.Blur()
	}
}

func (m *Model) validateLength() {
	val := m.textInput.Value()
	if val == "" {
		m.inputError = ""
		return
	}

	n, err := strconv.Atoi(val)
	if err != nil {
		m.inputError = "Must be a number"
		return
	}

	if n < 1 || n > 999 {
		m.inputError = "Must be between 1 and 999"
		return
	}

	m.inputError = ""
	m.cfg.Length = n
}
