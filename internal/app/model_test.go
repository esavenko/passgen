package app

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/esavenko/passgen/internal/constants"
	"github.com/esavenko/passgen/internal/messages"
)

func TestGetAppModel(t *testing.T) {
	m := GetAppModel()

	if m.currentScreen != constants.Menu {
		t.Errorf("currentScreen = %d, want Menu(%d)", m.currentScreen, constants.Menu)
	}
	if m.cfg == nil {
		t.Fatal("cfg is nil")
	}
	if m.menuModel == nil {
		t.Fatal("menuModel is nil")
	}
	if m.passwordGenerationModel == nil {
		t.Fatal("passwordGenerationModel is nil")
	}
	if m.settingsModel == nil {
		t.Fatal("settingsModel is nil")
	}
}

func TestCtrlCQuits(t *testing.T) {
	m := GetAppModel()
	_, cmd := m.Update(tea.KeyMsg{Type: tea.KeyCtrlC})

	if cmd == nil {
		t.Fatal("ctrl+c should return a command")
	}
	msg := cmd()
	if _, ok := msg.(tea.QuitMsg); !ok {
		t.Errorf("got %T, want tea.QuitMsg", msg)
	}
}

func TestSwitchToGeneration(t *testing.T) {
	m := GetAppModel()
	updated, _ := m.Update(messages.SwitchToGenerationMsg{})
	m = updated.(*Model)

	if m.currentScreen != constants.PasswordGeneration {
		t.Errorf("currentScreen = %d, want PasswordGeneration(%d)", m.currentScreen, constants.PasswordGeneration)
	}
}

func TestSwitchToSettings(t *testing.T) {
	m := GetAppModel()
	updated, _ := m.Update(messages.SwitchToSettingsMsg{})
	m = updated.(*Model)

	if m.currentScreen != constants.Settings {
		t.Errorf("currentScreen = %d, want Settings(%d)", m.currentScreen, constants.Settings)
	}
}

func TestSwitchToMenu(t *testing.T) {
	m := GetAppModel()

	// go to generation first
	updated, _ := m.Update(messages.SwitchToGenerationMsg{})
	m = updated.(*Model)

	// back to menu
	updated, _ = m.Update(messages.SwitchToMenuMsg{})
	m = updated.(*Model)

	if m.currentScreen != constants.Menu {
		t.Errorf("currentScreen = %d, want Menu(%d)", m.currentScreen, constants.Menu)
	}
}

func TestRoutingToMenu(t *testing.T) {
	m := GetAppModel()
	// send "j" — menu should handle it
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("j")})
	m = updated.(*Model)

	if m.menuModel.Choice != 1 {
		t.Errorf("menu Choice = %d, want 1", m.menuModel.Choice)
	}
}

func TestRoutingToGeneration(t *testing.T) {
	m := GetAppModel()
	updated, _ := m.Update(messages.SwitchToGenerationMsg{})
	m = updated.(*Model)

	// send enter — generation should handle it
	updated, _ = m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	m = updated.(*Model)

	if m.passwordGenerationModel == nil {
		t.Fatal("passwordGenerationModel is nil")
	}
}

func TestRoutingToSettings(t *testing.T) {
	m := GetAppModel()
	updated, _ := m.Update(messages.SwitchToSettingsMsg{})
	m = updated.(*Model)

	// send "j" — settings should handle it (cursor moves)
	updated, _ = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("j")})
	m = updated.(*Model)

	if m.settingsModel == nil {
		t.Fatal("settingsModel is nil")
	}
}

func TestViewPerScreen(t *testing.T) {
	m := GetAppModel()

	// menu view
	if v := m.View(); v == "" {
		t.Error("menu View() is empty")
	}

	// generation view
	m.Update(messages.SwitchToGenerationMsg{})
	m.currentScreen = constants.PasswordGeneration
	if v := m.View(); v == "" {
		t.Error("generation View() is empty")
	}

	// settings view
	m.currentScreen = constants.Settings
	if v := m.View(); v == "" {
		t.Error("settings View() is empty")
	}
}
