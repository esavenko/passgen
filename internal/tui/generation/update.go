package generation

import (
	"github.com/atotto/clipboard"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/esavenko/passgen/internal/generator"
	"github.com/esavenko/passgen/internal/messages"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			pwd, err := generator.GeneratePassword(generator.GeneratorConfig{
				Length:            m.cfg.Length,
				UseDigits:         m.cfg.UseDigits,
				UseSpecialSymbols: m.cfg.UseSpecialSymbols,
			})
			if err == nil {
				m.password = pwd
				m.generated = true
				m.copied = false
			}
			return m, nil

		case "c":
			if m.generated && m.password != "" {
				if err := clipboard.WriteAll(m.password); err == nil {
					m.copied = true
				}
			}
			return m, nil

		case "esc":
			m.password = ""
			m.generated = false
			return m, func() tea.Msg { return messages.SwitchToMenuMsg{} }
		}
	}

	return m, nil
}
