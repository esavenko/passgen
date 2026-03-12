package settings

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/esavenko/passgen/internal/config"
)

type Model struct {
	cfg        *config.Settings
	cursor     int
	textInput  textinput.Model
	inputError string
}

func NewSettingsModel(cfg *config.Settings) *Model {
	ti := textinput.New()
	ti.SetValue(fmt.Sprintf("%d", cfg.Length))
	ti.Placeholder = "Password length (1-999)"
	ti.CharLimit = 3
	ti.Width = 20
	ti.Prompt = "~> "
	ti.Focus()

	return &Model{
		cfg:       cfg,
		cursor:    0,
		textInput: ti,
	}
}

func (m *Model) Init() tea.Cmd {
	return textinput.Blink
}
