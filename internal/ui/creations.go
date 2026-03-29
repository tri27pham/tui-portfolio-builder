package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	creationTitleStyle    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("255"))
	creationSelectedStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("39"))
	creationStackStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("243"))
	creationDescStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("252"))
	creationCursorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("39"))
)

type Creations struct {
	cursor   int
	expanded bool
}

func NewCreations() Creations { return Creations{} }

func (m Creations) Init() tea.Cmd { return nil }

func (m Creations) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, Keys.Up):
			if m.cursor > 0 {
				m.cursor--
				m.expanded = false
			}
		case key.Matches(msg, Keys.Down):
			if m.cursor < len(Portfolio.Creations)-1 {
				m.cursor++
				m.expanded = false
			}
		case key.Matches(msg, Keys.Expand):
			m.expanded = !m.expanded
		}
	}
	return m, nil
}

func (m Creations) View() string {
	var rows []string
	rows = append(rows, creationStackStyle.Render("select one to learn more ↓"), "")

	for i, item := range Portfolio.Creations {
		cursor := "  "
		if i == m.cursor {
			cursor = creationCursorStyle.Render("▸ ")
		}

		var title string
		if i == m.cursor {
			title = creationSelectedStyle.Render(item.Name)
		} else {
			title = creationTitleStyle.Render(item.Name)
		}

		row := fmt.Sprintf("%s%s", cursor, title)
		rows = append(rows, row)

		if i == m.cursor && m.expanded {
			rows = append(rows,
				"    "+creationStackStyle.Render(item.Stack),
				"    "+creationDescStyle.Render(item.Desc),
				"",
			)
		}
	}

	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}
