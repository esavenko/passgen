package generation

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/esavenko/passgen/internal/config"
	"github.com/esavenko/passgen/internal/messages"
)

func keyMsg(key string) tea.KeyMsg {
	return tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(key)}
}

func specialKeyMsg(keyType tea.KeyType) tea.KeyMsg {
	return tea.KeyMsg{Type: keyType}
}

func TestNewPasswordModel(t *testing.T) {
	cfg := config.NewDefaultSettings()
	m := NewPasswordModel(cfg)

	if m.generated {
		t.Error("generated should be false initially")
	}
	if m.password != "" {
		t.Error("password should be empty initially")
	}
	if m.copied {
		t.Error("copied should be false initially")
	}
}

func TestEnterGeneratesPassword(t *testing.T) {
	cfg := config.NewDefaultSettings()
	m := NewPasswordModel(cfg)

	updated, _ := m.Update(specialKeyMsg(tea.KeyEnter))
	m = updated.(*Model)

	if !m.generated {
		t.Error("generated should be true after enter")
	}
	if m.password == "" {
		t.Error("password should not be empty after enter")
	}
	if len(m.password) != cfg.Length {
		t.Errorf("password length = %d, want %d", len(m.password), cfg.Length)
	}
}

func TestEnterResetscopied(t *testing.T) {
	cfg := config.NewDefaultSettings()
	m := NewPasswordModel(cfg)

	// generate first password
	updated, _ := m.Update(specialKeyMsg(tea.KeyEnter))
	m = updated.(*Model)
	m.copied = true

	// generate again — copied should reset
	updated, _ = m.Update(specialKeyMsg(tea.KeyEnter))
	m = updated.(*Model)

	if m.copied {
		t.Error("copied should be reset after new generation")
	}
}

func TestEscReturnsToMenu(t *testing.T) {
	cfg := config.NewDefaultSettings()
	m := NewPasswordModel(cfg)

	// generate a password first
	updated, _ := m.Update(specialKeyMsg(tea.KeyEnter))
	m = updated.(*Model)

	updated, cmd := m.Update(specialKeyMsg(tea.KeyEscape))
	m = updated.(*Model)

	if m.password != "" {
		t.Error("password should be cleared on esc")
	}
	if m.generated {
		t.Error("generated should be false on esc")
	}
	if cmd == nil {
		t.Fatal("expected command from esc")
	}
	msg := cmd()
	if _, ok := msg.(messages.SwitchToMenuMsg); !ok {
		t.Errorf("got %T, want SwitchToMenuMsg", msg)
	}
}

func TestCopyWithoutGeneration(t *testing.T) {
	cfg := config.NewDefaultSettings()
	m := NewPasswordModel(cfg)

	updated, _ := m.Update(keyMsg("c"))
	m = updated.(*Model)

	if m.copied {
		t.Error("copied should remain false when no password generated")
	}
}

func TestViewNotEmpty(t *testing.T) {
	cfg := config.NewDefaultSettings()
	m := NewPasswordModel(cfg)

	if m.View() == "" {
		t.Error("View() returned empty string")
	}
}

func TestViewAfterGeneration(t *testing.T) {
	cfg := config.NewDefaultSettings()
	m := NewPasswordModel(cfg)

	updated, _ := m.Update(specialKeyMsg(tea.KeyEnter))
	m = updated.(*Model)

	v := m.View()
	if v == "" {
		t.Error("View() returned empty string after generation")
	}
}
