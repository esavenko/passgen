package generation

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	// Input
	textInput  textinput.Model
	inputError error

	// Password
	password          string
	length            int
	useDigits         bool
	useSpecialSymbols bool
}

func (m *Model) Init() tea.Cmd {
	return textinput.Blink
}

func NewPasswordModel() *Model {
	ti := textinput.New()
	ti.Placeholder = "Enter the numbers of digits (ex: 15)"
	ti.Focus()
	ti.CharLimit = 150
	ti.Width = 40
	ti.Prompt = "~> "
	
	return &Model{
		textInput:         ti,
		inputError:        nil,
		length:            15,
		useDigits:         true,
		useSpecialSymbols: true,
	}
}
