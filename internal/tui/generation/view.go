package generation

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/esavenko/passgen/internal/generator"
)

var (
	mainStyle      = lipgloss.NewStyle().MarginLeft(2)
	secondaryStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("230"))
	submainStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	checkboxStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	dotStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("236")).Render(" • ")

	panelStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Padding(1, 2).
			Width(30)
)

func (m *Model) View() string {
	var s string

	s += secondaryStyle.Render("Password Generator") + "\n\n"

	// Settings panel (left)
	digitsStr := "no"
	if m.cfg.UseDigits {
		digitsStr = "yes"
	}
	specialStr := "no"
	if m.cfg.UseSpecialSymbols {
		specialStr = "yes"
	}

	settingsContent := fmt.Sprintf(
		"Length: %d\nDigits: %s\nSpecial: %s",
		m.cfg.Length, digitsStr, specialStr,
	)

	settingsPanel := panelStyle.
		BorderForeground(lipgloss.Color("212")).
		Render(settingsContent)

	// Output panel (right)
	cfg := generator.GeneratorConfig{
		Length:            m.cfg.Length,
		UseDigits:         m.cfg.UseDigits,
		UseSpecialSymbols: m.cfg.UseSpecialSymbols,
	}
	entropy := generator.Entropy(cfg)
	strength := generator.Strength(entropy)

	var outputContent string
	if m.generated {
		outputContent = checkboxStyle.Render(m.password) + "\n\n"
		outputContent += fmt.Sprintf("Entropy:  %.1f bits\n", entropy)
		outputContent += fmt.Sprintf("Strength: %s", strength)
		if m.copied {
			outputContent += "\n\n" + secondaryStyle.Render("Copied!")
		}
	} else {
		outputContent = submainStyle.Render("Press enter to generate")
	}

	outputPanel := panelStyle.
		BorderForeground(lipgloss.Color("230")).
		Render(outputContent)

	// Join panels horizontally
	panels := lipgloss.JoinHorizontal(lipgloss.Top, settingsPanel, "  ", outputPanel)

	s += panels + "\n\n"

	s += submainStyle.Render("enter: generate") + dotStyle +
		submainStyle.Render("c: copy") + dotStyle +
		submainStyle.Render("esc: back")

	return mainStyle.Render("\n" + s + "\n\n")
}
