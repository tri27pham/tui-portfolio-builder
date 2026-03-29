globs: ["**/*.go"]

# Go conventions
- Use standard Go project layout (cmd/, internal/)
- Handle all errors explicitly — no blank identifier on errors
- Wrap errors with fmt.Errorf("context: %w", err)
- No init() functions
- No global mutable state outside main
- Keep functions under 40 lines
- Run go vet ./... before marking any task done
