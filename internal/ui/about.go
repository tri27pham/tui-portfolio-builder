package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	aboutStarStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("255"))

	aboutBioStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("252"))

	aboutBioBoldStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("255")).
				Bold(true)

	aboutBioMutedStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("243"))
)

type About struct{}

func NewAbout() About { return About{} }

func (m About) Init() tea.Cmd { return nil }

func (m About) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return m, nil }

func (m About) View() string {
	rows := []string{aboutStarStyle.Render("*"), ""}

	for _, line := range Portfolio.RoleLines {
		rows = append(rows, aboutBioStyle.Render(line))
	}
	rows = append(rows, "")

	for _, line := range Portfolio.CurrentWork {
		rows = append(rows, aboutBioBoldStyle.Render(line))
	}
	rows = append(rows, "")

	for _, line := range Portfolio.Background {
		rows = append(rows, aboutBioMutedStyle.Render(line))
	}

	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}
