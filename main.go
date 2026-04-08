package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"

	"ssh-portfolio/internal/server"
	"ssh-portfolio/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	if os.Getenv("SSH_MODE") == "1" {
		lipgloss.SetColorProfile(termenv.ANSI256)
		if err := server.ListenAndServe("0.0.0.0", 2222); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		return
	}

	p := tea.NewProgram(ui.NewRoot(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
