package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/esavenko/passgen/common/constants"
	"github.com/esavenko/passgen/common/messages"
	"github.com/esavenko/passgen/internal/tui/generation"
	"github.com/esavenko/passgen/internal/tui/menu"
)

type Model struct {
	currentScreen           constants.Screen
	menuModel               *menu.Model
	passwordGenerationModel *generation.Model
}

func GetAppModel() *Model {
	return &Model{
		currentScreen:           constants.Menu,
		menuModel:               menu.NewMenuModel(),
		passwordGenerationModel: generation.NewPasswordModel(),
	}
}

func (m *Model) Init() tea.Cmd { return nil }

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl + c" || msg.String() == "q" {
			return m, tea.Quit
		}

	case messages.SwitchToMenuMsg:
		m.currentScreen = constants.Menu
		return m, nil

	case messages.SwitchToGenerationMsg:
		m.currentScreen = constants.PasswordGeneration
		return m, nil
	}

	/** Current state */
	switch m.currentScreen {
	case constants.Menu:
		newModel, cmd := m.menuModel.Update(msg)
		m.menuModel = newModel.(*menu.Model)
		return m, cmd

	case constants.PasswordGeneration:
		newModel, cmd := m.passwordGenerationModel.Update(msg)
		m.passwordGenerationModel = newModel.(*generation.Model)
		return m, cmd
	}

	return m, nil
}

func (m *Model) View() string {
	switch m.currentScreen {
	case constants.Menu:
		return m.menuModel.View()

	case constants.PasswordGeneration:
		return m.passwordGenerationModel.View()

	default:
		return ""
	}
}
