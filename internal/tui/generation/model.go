package generation

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/esavenko/passgen/internal/config"
)

type Model struct {
	cfg       *config.Settings
	password  string
	generated bool
	copied    bool
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func NewPasswordModel(cfg *config.Settings) *Model {
	return &Model{
		cfg: cfg,
	}
}
