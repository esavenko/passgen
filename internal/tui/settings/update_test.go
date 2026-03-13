package settings

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

func TestNewSettingsModel(t *testing.T) {
	cfg := &config.Settings{Length: 20, UseDigits: true, UseSpecialSymbols: false}
	m := NewSettingsModel(cfg)

	if m.cursor != 0 {
		t.Errorf("cursor = %d, want 0", m.cursor)
	}
	if m.textInput.Value() != "20" {
		t.Errorf("textInput value = %q, want %q", m.textInput.Value(), "20")
	}
}

func TestCursorNavigation(t *testing.T) {
	cfg := config.NewDefaultSettings()
	m := NewSettingsModel(cfg)

	// down moves cursor
	updated, _ := m.Update(keyMsg("j"))
	m = updated.(*Model)
	if m.cursor != 1 {
		t.Errorf("after down: cursor = %d, want 1", m.cursor)
	}

	updated, _ = m.Update(keyMsg("j"))
	m = updated.(*Model)
	if m.cursor != 2 {
		t.Errorf("after down x2: cursor = %d, want 2", m.cursor)
	}

	// should not go below 2
	updated, _ = m.Update(keyMsg("j"))
	m = updated.(*Model)
	if m.cursor != 2 {
		t.Errorf("after down x3: cursor = %d, want 2", m.cursor)
	}

	// up moves cursor back
	updated, _ = m.Update(keyMsg("k"))
	m = updated.(*Model)
	if m.cursor != 1 {
		t.Errorf("after up: cursor = %d, want 1", m.cursor)
	}

	updated, _ = m.Update(keyMsg("k"))
	m = updated.(*Model)
	if m.cursor != 0 {
		t.Errorf("after up x2: cursor = %d, want 0", m.cursor)
	}

	// should not go above 0
	updated, _ = m.Update(keyMsg("k"))
	m = updated.(*Model)
	if m.cursor != 0 {
		t.Errorf("after up x3: cursor = %d, want 0", m.cursor)
	}
}

func TestToggleDigits(t *testing.T) {
	cfg := config.NewDefaultSettings()
	m := NewSettingsModel(cfg)

	// move to digits field (cursor=1)
	updated, _ := m.Update(keyMsg("j"))
	m = updated.(*Model)

	original := cfg.UseDigits
	updated, _ = m.Update(specialKeyMsg(tea.KeyEnter))
	m = updated.(*Model)

	if cfg.UseDigits == original {
		t.Error("UseDigits was not toggled on enter")
	}

	// toggle back with space
	updated, _ = m.Update(keyMsg(" "))
	m = updated.(*Model)

	if cfg.UseDigits != original {
		t.Error("UseDigits was not toggled back on space")
	}
}

func TestToggleSpecialSymbols(t *testing.T) {
	cfg := config.NewDefaultSettings()
	m := NewSettingsModel(cfg)

	// move to special symbols field (cursor=2)
	updated, _ := m.Update(keyMsg("j"))
	m = updated.(*Model)
	updated, _ = m.Update(keyMsg("j"))
	m = updated.(*Model)

	original := cfg.UseSpecialSymbols
	updated, _ = m.Update(specialKeyMsg(tea.KeyEnter))
	m = updated.(*Model)

	if cfg.UseSpecialSymbols == original {
		t.Error("UseSpecialSymbols was not toggled")
	}
}

func TestEscReturnsSwitchToMenuMsg(t *testing.T) {
	cfg := config.NewDefaultSettings()
	m := NewSettingsModel(cfg)

	_, cmd := m.Update(specialKeyMsg(tea.KeyEscape))
	if cmd == nil {
		t.Fatal("esc should return a command")
	}

	msg := cmd()
	if _, ok := msg.(messages.SwitchToMenuMsg); !ok {
		t.Errorf("esc command returned %T, want SwitchToMenuMsg", msg)
	}
}

func TestValidateLength(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		wantError  string
		wantLength int
	}{
		{"valid", "50", "", 50},
		{"min boundary", "1", "", 1},
		{"max boundary", "999", "", 999},
		{"zero", "0", "Must be between 1 and 999", 15},
		{"too large truncated by charlimit", "1000", "", 100},
		{"negative", "-1", "Must be between 1 and 999", 15},
		{"not a number", "abc", "Must be a number", 15},
		{"empty", "", "", 15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := config.NewDefaultSettings()
			m := NewSettingsModel(cfg)

			m.textInput.SetValue(tt.input)
			m.validateLength()

			if m.inputError != tt.wantError {
				t.Errorf("inputError = %q, want %q", m.inputError, tt.wantError)
			}
			if cfg.Length != tt.wantLength {
				t.Errorf("cfg.Length = %d, want %d", cfg.Length, tt.wantLength)
			}
		})
	}
}

func TestViewNotEmpty(t *testing.T) {
	cfg := config.NewDefaultSettings()
	m := NewSettingsModel(cfg)

	v := m.View()
	if v == "" {
		t.Error("View() returned empty string")
	}
}
