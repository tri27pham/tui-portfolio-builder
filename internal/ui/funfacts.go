package ui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
)

var (
	factBulletStyle = lipgloss.NewStyle().Foreground(accent)
	factTextStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("252"))
)

type FunFacts struct {
	width int
}

func NewFunFacts() FunFacts { return FunFacts{} }

func (m FunFacts) Init() tea.Cmd { return nil }

func (m FunFacts) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return m, nil }

func (m FunFacts) View() string {
	title := "Fun Facts"
	rows := []string{
		contactTitleStyle.Render(title),
		contactRuleStyle.Render(strings.Repeat("─", len(title)+16)),
		"",
	}
	bullet := "◆ "
	indent := "  "
	wrapWidth := m.width - len(indent)
	if wrapWidth < 20 {
		wrapWidth = 40
	}
	for _, f := range Portfolio.Facts {
		wrapped := wordwrap.String(f, wrapWidth)
		lines := strings.Split(wrapped, "\n")
		first := factBulletStyle.Render(bullet) + factTextStyle.Render(lines[0])
		if len(lines) > 1 {
			for _, line := range lines[1:] {
				first += "\n" + indent + factTextStyle.Render(line)
			}
		}
		rows = append(rows, first)
	}
	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}
