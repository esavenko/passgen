package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/evertras/bubble-table/table"
)

type Model struct {
	/** Configuration*/
	length            int
	useDigits         bool
	useSpecialSymbols bool

	/** UI */
	table   table.Model
	focused string

	/** Password */
	password string
	err      error
}

const (
	columnKeyOption = "option"
	columnKeyValue  = "value"
)

func makeRow(option, value string) table.Row {
	return table.NewRow(table.RowData{
		columnKeyOption: option,
		columnKeyValue:  value,
	})
}

func NewModel() Model {
	return Model{
		length:            15,
		useDigits:         true,
		useSpecialSymbols: true,
		focused:           "length",

		table: table.New([]table.Column{
			table.NewColumn(columnKeyOption, "Option", 25),
			table.NewColumn(columnKeyValue, "Value", 25),
		}).WithRows([]table.Row{
			makeRow("Length", "15"),
			makeRow("Use Digits", "✓"),
			makeRow("Use Symbols", "✓"),
			makeRow("Generate", ""),
		}).
			Focused(true),
	}
}

func NewProgram() *tea.Program {
	return tea.NewProgram(NewModel())
}
