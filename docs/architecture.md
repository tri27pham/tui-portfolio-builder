# Architecture

## Overview
main.go → Wish SSH server → on connect, spawns tea.Program for that session

## Views
- about      — hero/intro, role, background, current work
- creations  — scrollable list of projects, enter to expand
- fun-facts  — personal facts, interests, hobbies
- contact    — links and handles

## Key files
- internal/server/server.go  — Wish setup, middleware, host key
- internal/ui/root.go        — root model, view switching
- internal/ui/about.go
- internal/ui/creations.go
- internal/ui/funfacts.go
- internal/ui/contact.go
- internal/ui/keys.go        — all keybindings

## Data flow
SSH connect → new tea.Program instance per session
WindowSizeMsg → root model → passed to active view on render
KeyMsg → root model → handled globally or forwarded to active view
