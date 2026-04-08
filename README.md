# TUI Portfolio Builder

A terminal-based portfolio site. Visitors connect via SSH and get an interactive TUI — no browser needed.

Built with Go and the [Charmbracelet](https://charm.sh) stack (Bubbletea, Lip Gloss, Wish).

## Quick start

```bash
# 1. Clone the repo
git clone https://github.com/yourusername/tui-portfolio-builder.git
cd tui-portfolio-builder

# 2. Edit portfolio/portfolio.go with your own content (see below)

# 3. Run it
go run .
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
| `go run .` | Preview the TUI locally in your terminal |
| `SSH_MODE=1 go run .` | Start as an SSH server on port 2222 |
| `ssh localhost -p 2222 -o StrictHostKeyChecking=no` | Connect to the local SSH server |
| `go build -o portfolio .` | Build a single binary |
| `fly deploy` | Deploy to Fly.io |
| `go vet ./...` | Lint |

## Deploy

> **Note:** Fly.io requires a payment method (credit/debit card) to deploy, even on the free tier. This app is lightweight enough that it should stay well within the free allowance.

```bash
# 1. Install the Fly CLI
brew install flyctl

# 2. Sign up or log in
fly auth signup   # or: fly auth login

# 3. Launch the app (first time only — creates the app on Fly)
fly launch

# 4. Deploy
fly deploy

# 5. Connect to your live site
ssh your-app-name.fly.dev
```

### Custom domain

To use a custom domain (e.g. `ssh yourdomain.sh`):

```bash
# Add your domain to Fly
fly certs add yourdomain.sh

# Then point your domain's DNS A/AAAA records to your Fly app's IP
fly ips list
```
