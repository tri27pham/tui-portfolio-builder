package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	factBulletStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("39"))
	factTextStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("252"))
)

var facts = []string{
	"Ran the Barcelona Marathon 2026",
	"Training toward sub-3:30 marathon (target: Fukuoka, Nov 2026)",
	"Currently learning Korean (intermediate, 2A level)",
	"Strong interest in geopolitics and defence tech",
	"Enjoys cooking — recently made chorizo croquettes, beef ragu, mango sticky rice",
	"Building toward starting my own company",
	"Side hobby: robotics (Raspberry Pi + Python)",
}

type FunFacts struct{}

func NewFunFacts() FunFacts { return FunFacts{} }

func (m FunFacts) Init() tea.Cmd { return nil }

func (m FunFacts) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return m, nil }

func (m FunFacts) View() string {
	var rows []string
	for _, f := range facts {
		row := factBulletStyle.Render("◆ ") + factTextStyle.Render(f)
		rows = append(rows, row)
	}
	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}
