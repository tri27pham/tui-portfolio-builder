package main

import (
	"fmt"
	"os"

	"ssh-portfolio/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(ui.NewRoot(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
