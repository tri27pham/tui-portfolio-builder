package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	aboutBioStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("252"))

	aboutBioBoldStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("255"))

	aboutBioMutedStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("243"))
)

type About struct{}

func NewAbout() About { return About{} }

func (m About) Init() tea.Cmd { return nil }

func (m About) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return m, nil }

func (m About) View() string {
	var rows []string

	for _, para := range Portfolio.RoleLines {
		rows = append(rows, aboutBioStyle.Render(renderInline(para)))
	}
	rows = append(rows, "")

	for _, para := range Portfolio.CurrentWork {
		rows = append(rows, aboutBioBoldStyle.Render(renderInline(para)))
	}
	rows = append(rows, "")

	for _, para := range Portfolio.BackgroundParas {
		rows = append(rows, aboutBioMutedStyle.Render(renderInline(para)))
	}

	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}
