package generation

import tea "github.com/charmbracelet/bubbletea"

type Model struct {
	length            int
	useDigits         bool
	useSpecialSymbols bool
}

func (m *Model) Init() tea.Cmd { return nil }

func NewPasswordModel() *Model {
	return &Model{
		length:            15,
		useDigits:         true,
		useSpecialSymbols: true,
	}
}
