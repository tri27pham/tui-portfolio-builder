globs: ["internal/ui/**/*.go"]

# Bubbletea / TUI conventions
- Each view is a separate file in internal/ui/
- Each view implements tea.Model (Init, Update, View)
- Root model holds currentView as a string constant, not an int
- Window size stored in root model, passed down — never re-read per view
- Keybindings defined centrally in internal/ui/keys.go using bubbles/key
- Footer showing available keybindings visible on every view
- Navigation: 1-4 to switch sections (about/creations/fun-facts/contact), q to quit, ? for help
- No blocking operations in Update() — use tea.Cmd for anything async
