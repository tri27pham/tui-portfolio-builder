# Tasks

## Phase 1 — local TUI
- [x] go mod init + install deps
- [x] root model with view switching
- [x] about view (landing/hero with ASCII art portrait left, bio text right)
- [x] creations view (scrollable list, enter to expand)
- [x] fun facts view
- [x] contact view
- [x] keybindings + footer on all views
- [x] ASCII art portrait — generate from docs/assets/photo.jpeg using ascii-image-converter -W 60 -H 50 -b

## Phase 2 — portfolio builder
- [x] centralise all content into internal/ui/portfolio.go (Config struct)
- [x] update about, creations, funfacts, contact views to read from Config
- [x] move asciiPortrait and nameArt into Config
- [x] document how to customise (which fields to change, how to regenerate ASCII art)

## Phase 3 — SSH server
- [ ] Wish server in internal/server/
- [ ] no-auth middleware
- [ ] test: ssh localhost -p 2222

## Phase 4 — deploy
- [ ] Dockerfile (multi-stage)
- [ ] fly.toml (TCP, port 22)
- [ ] fly launch + DNS

## Phase 5 — polish
- [ ] colour scheme
- [ ] welcome splash on connect
- [ ] animated transitions between views
