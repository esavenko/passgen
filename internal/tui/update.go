package tui

import (
	"passgen/internal/app"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/evertras/bubble-table/table"
)

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

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

				return m, nil
			}

		case " ":
			m.toogleBool()

		case "left":
			m.changeLength(-1)

		case "right":
			m.changeLength(1)
		}
	}
	m.table, cmd = m.table.Update(msg)
	m.updateFocus()

	return m, cmd
}

func (m *model) updateFocus() {
	switch m.table.GetHighlightedRowIndex() {
	case 0:
		m.focused = "length"
	case 1:
		m.focused = "useDigits"
	case 2:
		m.focused = "useSymbols"
	case 3:
		m.focused = "generate"
	}
}

func (m *model) toogleBool() {
	switch m.focused {
	case "useDigits":
		m.useDigits = !m.useDigits
		m.updateTableValue()

	case "useSymbols":
		m.useSpecialSymbols = !m.useSpecialSymbols
		m.updateTableValue()
	}
}

func (m *model) changeLength(change int) {
	if m.focused == "length" {
		m.length += change
		if m.length < 1 {
			m.length = 1
		}
		m.updateTableValue()
	}
}

func (m *model) updateTableValue() {
	newRows := []table.Row{
		table.NewRow(table.RowData{
			columnKeyOption: "Length",
			columnKeyValue:  strconv.Itoa(m.length),
		}),
		table.NewRow(table.RowData{
			columnKeyOption: "Use Digits",
			columnKeyValue:  m.getCheckboxValue(m.useDigits),
		}),
		table.NewRow(table.RowData{
			columnKeyOption: "Use Symbols",
			columnKeyValue:  m.getCheckboxValue(m.useSpecialSymbols),
		}),
		table.NewRow(table.RowData{
			columnKeyOption: "Generate",
			columnKeyValue:  "",
		}),
	}

	m.table = m.table.WithRows(newRows)
}

func (m *model) getCheckboxValue(checked bool) string {
	if checked {
		return "✓"
	}

	return "✗"
}