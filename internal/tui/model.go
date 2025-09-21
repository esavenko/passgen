package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/evertras/bubble-table/table"
)

type model struct {
	/** Configuration*/
	length            int
	useDigits         bool
	useSpecialSymbols bool

	/** UI */
	table     table.Model
	tableData []table.RowData
	focused   string

	/** Password */
	password string
	err      error
}

const (
	columnKeyOption = "option"
	columnKeyValue  = "value"
)

func initialModel() model {
	columns := []table.Column{
		table.NewColumn(columnKeyOption, "Option", 20),
		table.NewColumn(columnKeyValue, "Value", 15),
	}

	rows := []table.Row{
		table.NewRow(table.RowData{
			columnKeyOption: "Length",
			columnKeyValue:  "12",
		}),
		table.NewRow(table.RowData{
			columnKeyOption: "Use Digits",
			columnKeyValue:  "✓",
		}),
		table.NewRow(table.RowData{
			columnKeyOption: "Use Symbols",
			columnKeyValue:  "✓",
		}),
		table.NewRow(table.RowData{
			columnKeyOption: "Generate",
			columnKeyValue:  "",
		}),
	}

	initialTable := table.New(columns).
		WithRows(rows).
		Focused(true)

	return model{
		length:            12,
		useDigits:         true,
		useSpecialSymbols: true,
		table:             initialTable,
		focused:           "length",
	}
}

func NewProgram() *tea.Program {
	return tea.NewProgram(initialModel())
}
