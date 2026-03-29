package ui

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	viewAbout     = "about"
	viewCreations = "creations"
	viewFunFacts  = "funfacts"
	viewContact   = "contact"
)


var navItems = []struct {
	label string
	view  string
}{
	{"About", viewAbout},
	{"Creations", viewCreations},
	{"Fun Facts", viewFunFacts},
	{"Contact", viewContact},
}

var (
	portraitStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			PaddingRight(2)

	nameArtStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("39")).
			Bold(true)

	// cursor/hovered tab — full bright blue
	navActiveStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("39")).
			Bold(true)

	// active view when cursor has moved elsewhere — pale blue
	navCurrentStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("68"))

	navInactiveStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("243"))

	footerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("243"))

	footerKeyStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("39"))
)

type Root struct {
	currentView string
	navCursor   int
	width       int
	height      int
	about       About
	creations   Creations
	funfacts    FunFacts
	contact     Contact
}

func NewRoot() Root {
	return Root{
		currentView: viewAbout,
		navCursor:   0,
		about:       NewAbout(),
		creations:   NewCreations(),
		funfacts:    NewFunFacts(),
		contact:     NewContact(),
	}
}

func (m Root) Init() tea.Cmd {
	return nil
}

func (m Root) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, Keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, Keys.Left):
			if m.navCursor > 0 {
				m.navCursor--
			}
			return m, nil

		case key.Matches(msg, Keys.Right):
			if m.navCursor < len(navItems)-1 {
				m.navCursor++
			}
			return m, nil

		case key.Matches(msg, Keys.Expand):
			// navigate if cursor points to a different view than the current one
			if navItems[m.navCursor].view != m.currentView {
				m.currentView = navItems[m.navCursor].view
				return m, nil
			}
		}
	}

	// forward remaining messages to the active sub-model
	switch m.currentView {
	case viewCreations:
		updated, cmd := m.creations.Update(msg)
		m.creations = updated.(Creations)
		return m, cmd
	}

	return m, nil
}

func (m Root) View() string {
	left := portraitStyle.Render(Portfolio.Portrait)

	footer := m.footer()
	maxWidth := lipgloss.Width(footer)
	contentStyle := lipgloss.NewStyle().MaxWidth(maxWidth)

	rightTop := lipgloss.JoinVertical(lipgloss.Left,
		nameArtStyle.Render(Portfolio.NameArt),
		contentStyle.Render(m.activeView()),
	)

	rightBottom := lipgloss.JoinVertical(lipgloss.Left,
		m.renderNav(),
		"",
		footer,
	)

	// pad between content and footer so footer aligns with bottom of portrait
	leftHeight := lipgloss.Height(left)
	topHeight := lipgloss.Height(rightTop)
	bottomHeight := lipgloss.Height(rightBottom)
	pad := leftHeight - topHeight - bottomHeight
	if pad < 1 {
		pad = 1
	}

	right := lipgloss.JoinVertical(lipgloss.Left,
		rightTop,
		strings.Repeat("\n", pad-1),
		rightBottom,
	)

	return lipgloss.JoinHorizontal(lipgloss.Top, left, right)
}

func (m Root) activeView() string {
	switch m.currentView {
	case viewAbout:
		return m.about.View()
	case viewCreations:
		return m.creations.View()
	case viewFunFacts:
		return m.funfacts.View()
	case viewContact:
		return m.contact.View()
	}
	return ""
}

func (m Root) renderNav() string {
	var items []string
	for i, item := range navItems {
		isCurrent := m.currentView == item.view
		isCursor := i == m.navCursor

		switch {
		case isCursor:
			items = append(items, navActiveStyle.Render("✦ "+item.label))
		case isCurrent:
			items = append(items, navCurrentStyle.Render("  "+item.label))
		default:
			items = append(items, navInactiveStyle.Render("  "+item.label))
		}
	}
	return items[0] + "   " + items[1] + "   " + items[2] + "   " + items[3]
}

func (m Root) footer() string {
	return "[" +
		footerKeyStyle.Render("← →") +
		footerStyle.Render(" navigate · ") +
		footerKeyStyle.Render("↑/↓") +
		footerStyle.Render(" scroll · ") +
		footerKeyStyle.Render("enter") +
		footerStyle.Render(" open/expand · ") +
		footerKeyStyle.Render("q") +
		footerStyle.Render(" quit]")
}
