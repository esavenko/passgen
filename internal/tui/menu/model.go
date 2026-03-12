package menu

import tea "github.com/charmbracelet/bubbletea"

type Model struct {
	Choice int
}

func (m *Model) Init() tea.Cmd { return nil }

func NewMenuModel() *Model {
	return &Model{
		Choice: 0,
	}
}
