# SSH Portfolio

## What this is
Terminal portfolio site. Users connect via `ssh tri.sh` and get an
interactive TUI. No web frontend.

## Stack
- Go 1.22+
- Wish — SSH server (github.com/charmbracelet/wish)
- Bubbletea — TUI framework (github.com/charmbracelet/bubbletea)
- Lip Gloss — styling (github.com/charmbracelet/lipgloss)
- Fly.io — hosting

## Commands
- Run locally: `go run .`
- Test SSH locally: `ssh localhost -p 2222 -o StrictHostKeyChecking=no`
- Build: `go build -o portfolio .`
- Deploy: `fly deploy`
- Lint: `go vet ./...`

## Structure
- `portfolio/` — user content (the only file someone needs to edit)
- `internal/ui/` — Bubbletea models and views
- `internal/server/` — Wish SSH server setup
- `main.go` — entrypoint

## Views
- about      — hero/intro, role, background, current work
- creations  — projects and things built
- fun-facts  — personal facts, interests, hobbies
- contact    — links and handles

## Hard rules
- Bubbletea UI and Wish SSH wiring must stay completely separate
- Never hardcode terminal size — always use tea.WindowSizeMsg
- SSH server runs with no-auth (anonymous connections only)
- All Lip Gloss styles defined at package level, never inside View()
- Single binary output, no runtime dependencies
