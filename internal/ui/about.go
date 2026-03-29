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
	return lipgloss.JoinVertical(lipgloss.Left,
		aboutStarStyle.Render("*"),
		"",
		aboutBioStyle.Render("is a Founding Engineer @ Materi AI,"),
		aboutBioStyle.Render("building AI-native document tooling,"),
		aboutBioStyle.Render("previously at Fitch Ratings."),
		"",
		aboutBioBoldStyle.Render("Currently building an AI document editor —"),
		aboutBioBoldStyle.Render("think Cursor, but for docs."),
		"",
		aboutBioMutedStyle.Render("MSci Computer Science, King's College London (First Class)."),
		aboutBioMutedStyle.Render("Based in London."),
	)
}
