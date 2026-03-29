package ui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	contactTitleStyle = lipgloss.NewStyle().
				Foreground(accent).
				Bold(true)

	contactRuleStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("240"))

	contactLabelStyle = lipgloss.NewStyle().
				Foreground(accent).
				Bold(true).
				Width(6)

	contactValueStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("252"))
)

type Contact struct{}

func NewContact() Contact { return Contact{} }

func (m Contact) Init() tea.Cmd { return nil }

func (m Contact) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return m, nil }

func (m Contact) View() string {
	title := "Contacts"
	rows := []string{
		contactTitleStyle.Render(title),
		contactRuleStyle.Render(strings.Repeat("─", len(title)+16)),
		"",
	}
	for _, l := range Portfolio.Links {
		row := contactLabelStyle.Render(l.Label) + contactValueStyle.Render(l.Value)
		rows = append(rows, row)
	}
	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}
