package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/esavenko/passgen/internal/config"
	"github.com/esavenko/passgen/internal/constants"
	"github.com/esavenko/passgen/internal/messages"
	"github.com/esavenko/passgen/internal/tui/generation"
	"github.com/esavenko/passgen/internal/tui/menu"
	"github.com/esavenko/passgen/internal/tui/settings"
)

type Model struct {
	currentScreen           constants.Screen
	cfg                     *config.Settings
	menuModel               *menu.Model
	passwordGenerationModel *generation.Model
	settingsModel           *settings.Model
}

func GetAppModel() *Model {
	cfg := config.NewDefaultSettings()

	return &Model{
		currentScreen:           constants.Menu,
		cfg:                     cfg,
		menuModel:               menu.NewMenuModel(),
		passwordGenerationModel: generation.NewPasswordModel(cfg),
		settingsModel:           settings.NewSettingsModel(cfg),
	}
}

func (m *Model) Init() tea.Cmd { return nil }

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

	case messages.SwitchToMenuMsg:
		m.currentScreen = constants.Menu
		return m, nil

	case messages.SwitchToGenerationMsg:
		m.currentScreen = constants.PasswordGeneration
		return m, nil

	case messages.SwitchToSettingsMsg:
		m.currentScreen = constants.Settings
		return m, m.settingsModel.Init()
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

	case constants.Settings:
		newModel, cmd := m.settingsModel.Update(msg)
		m.settingsModel = newModel.(*settings.Model)
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

	case constants.Settings:
		return m.settingsModel.View()

	default:
		return ""
	}
}
