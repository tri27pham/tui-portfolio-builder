# Tasks

## Phase 1 — local TUI
- [x] go mod init + install deps
- [x] root model with view switching (stub — full wiring in step 6)
- [ ] about view (landing/hero with ASCII art portrait left, bio text right)
- [ ] creations view (scrollable list, enter to expand)
- [ ] fun facts view
- [ ] contact view
- [ ] keybindings + footer on all views
- [ ] ASCII art portrait — generate from docs/assets/photo.jpg using ascii-image-converter -W 20 -b -C

## Phase 2 — SSH server
- [ ] Wish server in internal/server/
- [ ] no-auth middleware
- [ ] test: ssh localhost -p 2222

## Phase 3 — deploy
- [ ] Dockerfile (multi-stage)
- [ ] fly.toml (TCP, port 22)
- [ ] fly launch + DNS

## Phase 4 — polish
- [ ] colour scheme
- [ ] welcome splash on connect
- [ ] animated transitions between views
