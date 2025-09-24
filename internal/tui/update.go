package tui

import (
	"passgen/internal/app"
	"strconv"

	"github.com/atotto/clipboard"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/evertras/bubble-table/table"
)

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	m.table, cmd = m.table.Update(msg)
	cmds = append(cmds, cmd)
	m.updateFocus(m.table.GetHighlightedRowIndex())

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			cmds = append(cmds, tea.Quit)

		case "enter":
			if m.focused == "generate" {
				cfg := app.GeneratorConfig{
					Length:            m.length,
					UseDigits:         m.useDigits,
					UseSpecialSymbols: m.useSpecialSymbols,
				}

				password, err := app.GeneratePassword(cfg)

				if err != nil {
					m.err = err
				} else {
					m.password = password
				}
			}

		case "c", "C":
			m.copyPassword()

		case "left":
			m.changeLength(-1)

		case "right":
			m.changeLength(1)

		case " ":
			m.toggleCheckboxValue()
		}
	}

	return m, tea.Batch(cmds...)
}

func (m *Model) updateFocus(index int) string {
	switch index {
	case 0:
		m.focused = "length"
	case 1:
		m.focused = "useDigits"
	case 2:
		m.focused = "useSpecialSymbols"
	case 3:
		m.focused = "generate"
	}

	return m.focused
}

func (m *Model) copyPassword() {
	err := clipboard.WriteAll(m.password)
	if err != nil {
		return
	}
}

func (m *Model) changeLength(count int) {
	if m.focused == "length" {
		m.length += count

		if m.length < 1 {
			m.length = 1
		}

		m.updateTable()
	}
}

func (m *Model) toggleCheckboxValue() {
	switch m.focused {
	case "useDigits":
		m.useDigits = !m.useDigits
		m.updateTable()
	case "useSpecialSymbols":
		m.useSpecialSymbols = !m.useSpecialSymbols
		m.updateTable()
	}
}

func (m *Model) getCheckboxValue(checked bool) string {
	if !checked {
		return "✗"
	}

	return "✓"
}

func (m *Model) updateTable() {
	updatedRows := []table.Row{
		makeRow("Length", strconv.Itoa(m.length)),
		makeRow("Use Digits", m.getCheckboxValue(m.useDigits)),
		makeRow("Use Symbols", m.getCheckboxValue(m.useSpecialSymbols)),
		makeRow("Generate", ""),
	}

	m.table = m.table.WithRows(updatedRows)
}
