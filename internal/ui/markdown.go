package ui

import (
	"regexp"

	"github.com/charmbracelet/lipgloss"
)

// renderInline processes a small subset of inline markdown within a string:
//
//	**text** → bold white
//	*text*   → italic
//
// Apply this to any string before passing it to a lipgloss style.
// The base color of the surrounding style is preserved for unmarked text.
func renderInline(s string) string {
	s = boldRe.ReplaceAllStringFunc(s, func(m string) string {
		return inlineBoldStyle.Render(boldRe.FindStringSubmatch(m)[1])
	})
	s = italicRe.ReplaceAllStringFunc(s, func(m string) string {
		return inlineItalicStyle.Render(italicRe.FindStringSubmatch(m)[1])
	})
	return s
}

var (
	boldRe   = regexp.MustCompile(`\*\*(.+?)\*\*`)
	italicRe = regexp.MustCompile(`\*(.+?)\*`)

	inlineBoldStyle   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("255"))
	inlineItalicStyle = lipgloss.NewStyle().Italic(true)
)
