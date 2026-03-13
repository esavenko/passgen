package menu

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/esavenko/passgen/internal/messages"
)

func keyMsg(key string) tea.KeyMsg {
	return tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(key)}
}

func specialKeyMsg(keyType tea.KeyType) tea.KeyMsg {
	return tea.KeyMsg{Type: keyType}
}

func TestNewMenuModel(t *testing.T) {
	m := NewMenuModel()
	if m.Choice != 0 {
		t.Errorf("Choice = %d, want 0", m.Choice)
	}
}

func TestCursorNavigation(t *testing.T) {
	m := NewMenuModel()

	// down
	updated, _ := m.Update(keyMsg("j"))
	m = updated.(*Model)
	if m.Choice != 1 {
		t.Errorf("after down: Choice = %d, want 1", m.Choice)
	}

	updated, _ = m.Update(keyMsg("j"))
	m = updated.(*Model)
	if m.Choice != 2 {
		t.Errorf("after down x2: Choice = %d, want 2", m.Choice)
	}

	// clamp at 2
	updated, _ = m.Update(keyMsg("j"))
	m = updated.(*Model)
	if m.Choice != 2 {
		t.Errorf("after down x3: Choice = %d, want 2", m.Choice)
	}

	// up
	updated, _ = m.Update(keyMsg("k"))
	m = updated.(*Model)
	if m.Choice != 1 {
		t.Errorf("after up: Choice = %d, want 1", m.Choice)
	}

	updated, _ = m.Update(keyMsg("k"))
	m = updated.(*Model)
	if m.Choice != 0 {
		t.Errorf("after up x2: Choice = %d, want 0", m.Choice)
	}

	// clamp at 0
	updated, _ = m.Update(keyMsg("k"))
	m = updated.(*Model)
	if m.Choice != 0 {
		t.Errorf("after up x3: Choice = %d, want 0", m.Choice)
	}
}

func TestEnterGenerate(t *testing.T) {
	m := NewMenuModel()
	// Choice=0 -> Generate
	_, cmd := m.Update(specialKeyMsg(tea.KeyEnter))
	if cmd == nil {
		t.Fatal("expected command")
	}
	msg := cmd()
	if _, ok := msg.(messages.SwitchToGenerationMsg); !ok {
		t.Errorf("got %T, want SwitchToGenerationMsg", msg)
	}
}

func TestEnterSettings(t *testing.T) {
	m := NewMenuModel()
	m.Update(keyMsg("j")) // Choice=1
	updated, _ := m.Update(keyMsg("j"))
	m = updated.(*Model)
	// reset to 1
	m.Choice = 1

	_, cmd := m.Update(specialKeyMsg(tea.KeyEnter))
	if cmd == nil {
		t.Fatal("expected command")
	}
	msg := cmd()
	if _, ok := msg.(messages.SwitchToSettingsMsg); !ok {
		t.Errorf("got %T, want SwitchToSettingsMsg", msg)
	}
}

func TestEnterQuit(t *testing.T) {
	m := NewMenuModel()
	m.Choice = 2

	_, cmd := m.Update(specialKeyMsg(tea.KeyEnter))
	if cmd == nil {
		t.Fatal("expected command")
	}
	msg := cmd()
	if _, ok := msg.(tea.QuitMsg); !ok {
		t.Errorf("got %T, want tea.QuitMsg", msg)
	}
}

func TestQKey(t *testing.T) {
	m := NewMenuModel()
	_, cmd := m.Update(keyMsg("q"))
	if cmd == nil {
		t.Fatal("expected command")
	}
	msg := cmd()
	if _, ok := msg.(tea.QuitMsg); !ok {
		t.Errorf("got %T, want tea.QuitMsg", msg)
	}
}

func TestViewNotEmpty(t *testing.T) {
	m := NewMenuModel()
	if m.View() == "" {
		t.Error("View() returned empty string")
	}
}
