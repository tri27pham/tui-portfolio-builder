# Architecture

## Overview
main.go → Wish SSH server → on connect, spawns tea.Program for that session

## Portfolio builder
All user-configurable content lives in a single file: `internal/ui/portfolio.go`.
To customise the portfolio, users only need to edit that file — no view code changes required.

## Views
- about      — hero/intro, role, background, current work
- creations  — scrollable list of projects, enter to expand
- fun-facts  — personal facts, interests, hobbies
- contact    — links and handles

## Key files
- internal/ui/portfolio.go   — all user content (name, bio, creations, facts, contact, ASCII art)
- internal/server/server.go  — Wish setup, middleware, host key
- internal/ui/root.go        — root model, view switching, layout
- internal/ui/about.go
- internal/ui/creations.go
- internal/ui/funfacts.go
- internal/ui/contact.go
- internal/ui/keys.go        — all keybindings

## Data flow
SSH connect → new tea.Program instance per session
WindowSizeMsg → root model → passed to active view on render
KeyMsg → root model → handled globally (nav) or forwarded to active view (scroll/expand)

## Layout
Left column (fixed): ASCII portrait + always visible
Right column (changes): name art → active view content → nav tabs → footer
Footer is pinned to bottom of left column height via lipgloss.Height measurement
