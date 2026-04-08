package ui

import (
	"math/rand"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	tickInterval  = 80 * time.Millisecond
	maxStars    = 2
	spawnChance = 0.15 // probability of spawning a new star each tick
)

var (
	starHeadStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("255")).Bold(true)
	starTrail1Style = lipgloss.NewStyle().Foreground(lipgloss.Color("250"))
	starTrail2Style = lipgloss.NewStyle().Foreground(lipgloss.Color("243"))
	starTrail3Style = lipgloss.NewStyle().Foreground(lipgloss.Color("238"))
)

type starSegment struct{ x, y, age int }

type shootingStar struct {
	segments []starSegment
	dx, dy   int
}

type Rain struct {
	stars  []shootingStar
	width  int
	height int
}

type rainTickMsg struct{}

func newRain(nameArt string) Rain {
	lines := nameArtLines(nameArt)
	w := 0
	for _, l := range lines {
		if len([]rune(l)) > w {
			w = len([]rune(l))
		}
	}
	return Rain{width: w, height: len(lines)}
}

func (r Rain) newStar() shootingStar {
	canvasH := r.height + 4
	canvasW := r.width
	if canvasW < 1 {
		canvasW = 1
	}
	if canvasH < 1 {
		canvasH = 1
	}

	dx := []int{-1, 1}[rand.Intn(2)]
	dy := []int{-1, -1, 1, 1}[rand.Intn(4)]

	var startX, startY int
	if dx > 0 {
		startX = 0
	} else {
		startX = canvasW - 1
	}
	if dy > 0 {
		startY = 0
	} else {
		startY = canvasH - 1
	}

	return shootingStar{
		segments: []starSegment{{x: startX, y: startY, age: 0}},
		dx:       dx,
		dy:       dy,
	}
}

func (r Rain) tickCmd() tea.Cmd {
	return tea.Tick(tickInterval, func(time.Time) tea.Msg { return rainTickMsg{} })
}

func (r Rain) update(msg tea.Msg) (Rain, tea.Cmd) {
	if _, ok := msg.(rainTickMsg); !ok {
		return r, nil
	}

	canvasH := r.height + 4

	// advance each star
	next := r.stars[:0]
	for _, s := range r.stars {
		var alive []starSegment
		for _, seg := range s.segments {
			seg.age++
			if seg.age <= 3 {
				alive = append(alive, seg)
			}
		}

		head := s.segments[0]
		head.x += s.dx
		head.y += s.dy
		head.age = 0

		if head.x >= 0 && head.x < r.width && head.y >= 0 && head.y < canvasH {
			s.segments = append([]starSegment{head}, alive...)
			next = append(next, s)
		} else if len(alive) > 0 {
			s.segments = alive
			next = append(next, s)
		}
		// star fully exited — drop it
	}
	r.stars = next

	// maybe spawn a new star
	if len(r.stars) < maxStars && rand.Float32() < spawnChance {
		r.stars = append(r.stars, r.newStar())
	}

	return r, r.tickCmd()
}

func (r Rain) render(nameArt string, canvasH int) string {
	artLines := nameArtLines(nameArt)
	artOffsetY := 2

	type pos struct{ x, y int }
	// keep lowest age (brightest) when stars overlap
	segMap := make(map[pos]int)
	for _, s := range r.stars {
		for _, seg := range s.segments {
			p := pos{seg.x, seg.y}
			if existing, ok := segMap[p]; !ok || seg.age < existing {
				segMap[p] = seg.age
			}
		}
	}

	out := make([]string, canvasH)
	for y := 0; y < canvasH; y++ {
		artY := y - artOffsetY
		var artRunes []rune
		if artY >= 0 && artY < len(artLines) {
			artRunes = []rune(artLines[artY])
		}

		rowW := len(artRunes)
		for _, s := range r.stars {
			for _, seg := range s.segments {
				if seg.y == y && seg.x+1 > rowW {
					rowW = seg.x + 1
				}
			}
		}

		var sb strings.Builder
		for x := 0; x < rowW; x++ {
			ch := ' '
			if x < len(artRunes) {
				ch = artRunes[x]
			}
			if age, hit := segMap[pos{x, y}]; hit && ch == ' ' {
				switch age {
				case 0:
					sb.WriteString(starHeadStyle.Render("✦"))
				case 1:
					sb.WriteString(starTrail1Style.Render("✦"))
				case 2:
					sb.WriteString(starTrail2Style.Render("✦"))
				default:
					sb.WriteString(starTrail3Style.Render("✦"))
				}
			} else {
				sb.WriteString(nameArtStyle.Render(string(ch)))
			}
		}
		out[y] = sb.String()
	}
	return strings.Join(out, "\n")
}

func nameArtLines(nameArt string) []string {
	s := strings.TrimPrefix(nameArt, "\n")
	s = strings.TrimRight(s, "\n")
	return strings.Split(s, "\n")
}
