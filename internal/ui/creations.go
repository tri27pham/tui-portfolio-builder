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

type creation struct {
	name  string
	stack string
	desc  string
}

var creationItems = []creation{
	{
		name:  "Materi",
		stack: "TypeScript, React, TipTap/ProseMirror",
		desc:  "AI-native document editor. Live A4 pagination engine, AI editing features.",
	},
	{
		name:  "Notchpets",
		stack: "Electron, TypeScript, Swift, Supabase",
		desc:  "macOS notch companion app. Pixel art pets that live in your MacBook notch.",
	},
	{
		name:  "SSH Portfolio",
		stack: "Go, Charmbracelet (Wish, Bubbletea, Lip Gloss)",
		desc:  "This site. Terminal-based portfolio, ssh tri.sh.",
	},
}

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
			if m.cursor < len(creationItems)-1 {
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

	for i, item := range creationItems {
		cursor := "  "
		if i == m.cursor {
			cursor = creationCursorStyle.Render("▸ ")
		}

		var title string
		if i == m.cursor {
			title = creationSelectedStyle.Render(item.name)
		} else {
			title = creationTitleStyle.Render(item.name)
		}

		row := fmt.Sprintf("%s%s", cursor, title)
		rows = append(rows, row)

		if i == m.cursor && m.expanded {
			rows = append(rows,
				"    "+creationStackStyle.Render(item.stack),
				"    "+creationDescStyle.Render(item.desc),
				"",
			)
		}
	}

	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}
