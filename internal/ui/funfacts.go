package ui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	factBulletStyle = lipgloss.NewStyle().Foreground(accent)
	factTextStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("252"))
)

type FunFacts struct{}

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
	for _, f := range Portfolio.Facts {
		row := factBulletStyle.Render("◆ ") + factTextStyle.Render(f)
		rows = append(rows, row)
	}
	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}
