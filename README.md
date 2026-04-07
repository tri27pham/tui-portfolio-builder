# TUI Portfolio Builder

A terminal-based portfolio site. Visitors connect via SSH and get an interactive TUI — no browser needed.

Built with Go and the [Charmbracelet](https://charm.sh) stack (Bubbletea, Lip Gloss, Wish).

## Quick start

```bash
# 1. Clone the repo
git clone https://github.com/tri27pham/tui-portfolio-builder.git
cd tui-portfolio-builder

# 2. Edit portfolio/portfolio.go with your own content (see below)

# 3. Run locally
go run .

# 4. In another terminal, connect via SSH to see it
ssh localhost -p 2222 -o StrictHostKeyChecking=no
```

## Make it yours

**The only file you need to edit is [`portfolio/portfolio.go`](portfolio/portfolio.go).**

Open it up and replace the content with your own:

| Section        | What to change                                                |
|----------------|---------------------------------------------------------------|
| `Portrait`     | Your ASCII art photo (see generation instructions below)      |
| `NameArt`      | Your stylised name (see generation instructions below)        |
| `RoleLines`    | Your headline intro (e.g. "is a software engineer at ...")    |
| `CurrentWork`  | What you're working on now                                    |
| `BackgroundParas` | Education, past experience, extra context                  |
| `Creations`    | Your projects — each has a Name, Subheading, and Desc         |
| `Facts`        | Fun facts about you — one bullet point per string             |
| `Links`        | Your contact links — each has a Label (e.g. "GH") and Value  |
| `AccentColor`  | Your highlight color (pick from the 256-color palette)        |
| `DimColor`     | A muted version of your accent color                          |

### Generate your ASCII art

**Portrait** (the picture on the left side):

```bash
# Install: https://github.com/TheZoraworski/ascii-image-converter
ascii-image-converter path/to/your-photo.jpg -W 60 -b
```

Copy the output and paste it into the `Portrait` field.

**Name art** (the big stylised name at the top):

```bash
# Install: brew install figlet  (or apt install figlet)
figlet -f slant "yourname"
```

Copy the output and paste it into the `NameArt` field.

### Pick your colors

The `AccentColor` and `DimColor` fields accept 256-color terminal codes as strings. Some good options:

| Color          | AccentColor | DimColor |
|----------------|-------------|----------|
| Bright blue    | `"39"`      | `"68"`   |
| Electric cyan  | `"51"`      | `"68"`   |
| Bright green   | `"82"`      | `"64"`   |
| Hot pink       | `"205"`     | `"162"`  |
| Amber/gold     | `"220"`     | `"136"`  |
| Soft purple    | `"141"`     | `"97"`   |

Preview all 256 colors in your terminal:

```bash
for i in {0..255}; do printf "\e[38;5;${i}m %3d \e[0m" $i; done; echo
```

## Project structure

```
portfolio/portfolio.go   <-- YOUR CONTENT (edit this)
main.go                  <-- entrypoint
internal/ui/             <-- TUI views and layout (you don't need to touch this)
internal/server/         <-- SSH server setup
docs/assets/             <-- place your photo here for ASCII conversion
```

## Commands

| Command | What it does |
|---------|-------------|
| `go run .` | Run the TUI locally |
| `ssh localhost -p 2222 -o StrictHostKeyChecking=no` | Connect to it via SSH |
| `go build -o portfolio .` | Build a single binary |
| `go vet ./...` | Lint |

## Deploy

<!-- TODO: Add deployment instructions (Fly.io setup, fly.toml, custom domain, etc.) -->

Deployment guide coming soon. This project is designed to run on [Fly.io](https://fly.io) as a single binary with no runtime dependencies.
