package ui

import (
	"github.com/charmbracelet/lipgloss"

	"ssh-portfolio/portfolio"
)

// Re-export types and data from the portfolio package so the rest of internal/ui
// can reference them without a package prefix.
type Config = portfolio.Config
type Creation = portfolio.Creation
type Link = portfolio.Link

var Portfolio = &portfolio.Portfolio

// accent and dim are derived from Portfolio so styles can reference them at package level.
var (
	accent = lipgloss.Color(Portfolio.AccentColor)
	dim    = lipgloss.Color(Portfolio.DimColor)
)
