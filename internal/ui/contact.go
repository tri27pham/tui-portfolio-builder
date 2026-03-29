package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	contactLabelStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("243")).
				Width(12)

	contactValueStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("39"))
)

type Contact struct{}

func NewContact() Contact { return Contact{} }

func (m Contact) Init() tea.Cmd { return nil }

func (m Contact) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return m, nil }

func (m Contact) View() string {
	var rows []string
	for _, l := range Portfolio.Links {
		row := contactLabelStyle.Render(l.Label) + contactValueStyle.Render(l.Value)
		rows = append(rows, row)
	}
	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}
