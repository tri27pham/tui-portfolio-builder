package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/keygen"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	bm "github.com/charmbracelet/wish/bubbletea"
	"github.com/muesli/termenv"

	"ssh-portfolio/internal/ui"
)

func handler(sess ssh.Session) (tea.Model, []tea.ProgramOption) {
	return ui.NewRoot(), []tea.ProgramOption{tea.WithAltScreen()}
}

func ensureHostKey(path string) error {
	if _, err := os.Stat(path); err == nil {
		return nil
	}
	if err := os.MkdirAll(".ssh", 0o700); err != nil {
		return fmt.Errorf("creating .ssh dir: %w", err)
	}
	log.Info("generating SSH host key", "path", path)
	kp, err := keygen.New(path)
	if err != nil {
		return fmt.Errorf("generating key: %w", err)
	}
	return kp.WriteKeys()
}

// ListenAndServe starts the SSH server on the given host and port.
// It blocks until a SIGINT or SIGTERM is received, then shuts down gracefully.
func ListenAndServe(host string, port int) error {
	keyPath := ".ssh/id_ed25519"
	if err := ensureHostKey(keyPath); err != nil {
		return fmt.Errorf("host key: %w", err)
	}

	addr := fmt.Sprintf("%s:%d", host, port)

	srv, err := wish.NewServer(
		wish.WithAddress(addr),
		wish.WithHostKeyPath(".ssh/id_ed25519"),
		wish.WithMiddleware(bm.MiddlewareWithColorProfile(handler, termenv.ANSI256)),
	)
	if err != nil {
		return fmt.Errorf("creating server: %w", err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)

	log.Info("starting SSH server", "addr", addr)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Error("server error", "error", err)
			done <- os.Interrupt
		}
	}()

	<-done
	log.Info("shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return srv.Shutdown(ctx)
}
