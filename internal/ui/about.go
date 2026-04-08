package ui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
)

var (
	aboutBioStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("252"))

	aboutBioBoldStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("255"))

	aboutBioMutedStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("240"))
)

type About struct {
	width int
}

func NewAbout() About { return About{} }

func (m About) Init() tea.Cmd { return nil }

func (m About) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return m, nil }

func (m About) View() string {
	var rows []string

	for _, para := range Portfolio.RoleLines {
		rows = append(rows, styledWrap(para, m.width, aboutBioStyle))
	}
	rows = append(rows, "")

	for _, para := range Portfolio.CurrentWork {
		rows = append(rows, styledWrap(para, m.width, aboutBioBoldStyle))
	}
	rows = append(rows, "")

	for _, para := range Portfolio.BackgroundParas {
		rows = append(rows, styledWrap(para, m.width, aboutBioMutedStyle))
	}

	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}

// styledWrap word-wraps the paragraph first, then applies the style to each
// line so that ANSI colors survive the wrap.
func styledWrap(para string, width int, style lipgloss.Style) string {
	if width < 20 {
		width = 40
	}
	wrapped := wordwrap.String(para, width)
	lines := strings.Split(wrapped, "\n")
	for i, line := range lines {
		lines[i] = style.Render(renderInline(line))
	}
	return strings.Join(lines, "\n")
}
