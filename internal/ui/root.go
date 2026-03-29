package ui

import tea "github.com/charmbracelet/bubbletea"

type Root struct {
	width  int
	height int
}

func NewRoot() Root {
	return Root{}
}

func (m Root) Init() tea.Cmd {
	return nil
}

func (m Root) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Root) View() string {
	return "loading...\n\npress q to quit"
}
