package portfolio

import _ "embed"

// ============================================================================
// CONFIG — this is the only file you need to edit to make the portfolio yours.
//
// Everything you see when someone connects via SSH is defined here:
// your photo, your name, your bio, your projects, your links, and your colors.
// Just change the text below and the site updates — no other code to touch.
//
// HOW TO GENERATE YOUR OWN ASCII ART:
//
//	Portrait (the picture on the left side):
//	  1. Put your photo at portfolio/assets/[yourphotoname]
//	  2. Run:  ascii-image-converter portfolio/assets/[yourphotoname] -W 60 -b
//
// .      eg. ascii-image-converter portfolio/assets/photo.jpg -W 60 -b
//
//  3. Paste the output into portfolio/portrait.txt
//
//     Name art (the big stylised name at the top right):
//
//  1. Run:  figlet -f slant "yourname"
//
//  2. Paste the output into portfolio/name.txt
//
// ============================================================================
type Config struct {

	// ---------- LEFT SIDE ----------

	// Portrait is the ASCII art picture shown on the left side of the screen.
	// Replace it with your own (see instructions above).
	Portrait string

	// ---------- TOP RIGHT — YOUR NAME ----------

	// NameArt is the big stylised text that appears at the top right.
	// Replace it with your own name (see instructions above).
	NameArt string

	// ---------- ABOUT TAB ----------
	// These three fields make up the "About" page. Each is a list of paragraphs.
	// Long text wraps automatically — you don't need to worry about line length.
	// You can use **bold** and *italic* inside the text.

	// RoleLines — your headline / intro. Shown in normal white text.
	// Example: "is a software engineer building cool things"
	RoleLines []string

	// CurrentWork — what you're doing right now. Shown in bright white (stands out).
	// Example: "Currently building an AI editor at Acme Corp"
	CurrentWork []string

	// BackgroundParas — education, background, extra context. Shown in grey.
	// Example: "Studied CS at MIT, previously interned at Google"
	BackgroundParas []string

	// ---------- CREATIONS TAB ----------

	// Creations — your projects. Each one has a Name, Subheading, and Desc.
	// Users can arrow-key through them and press Enter to expand details.
	// See the Creation type further down for the fields.
	Creations []Creation

	// ---------- FUN FACTS TAB ----------

	// Facts — a list of bullet points about you. Each string = one bullet.
	// Example: "Ran a marathon", "Speaks three languages"
	Facts []string

	// ---------- CONTACT TAB ----------

	// Links — your contact/social links. Each one has a short Label and a Value.
	// Example: {Label: "GH", Value: "github.com/you"}
	// See the Link type further down for the fields.
	Links []Link

	// ---------- COLORS ----------
	// These control the accent color used everywhere: nav highlights, name art,
	// bullet points, selected items, and footer keys.
	//
	// Pick a number from the 256-color terminal palette. Some good options:
	//
	//   "39"  — bright blue         "205" — hot pink
	//   "51"  — electric cyan       "220" — amber/gold
	//   "82"  — bright green        "141" — soft purple
	//
	// Browse the full palette:
	//   https://en.wikipedia.org/wiki/ANSI_escape_code#8-bit
	// Or preview colors in your terminal:
	//   for i in {0..255}; do printf "\e[38;5;${i}m %3d \e[0m" $i; done

	// AccentColor — the main highlight color.
	AccentColor string

	// DimColor — a muted version of your accent, used for the active tab when
	// your cursor has moved to a different tab. Pick a darker shade of your accent.
	//
	// Suggested pairings:
	//   AccentColor "39"/"51" → DimColor "68"   (muted blue)
	//   AccentColor "82"     → DimColor "64"   (muted green)
	//   AccentColor "205"    → DimColor "162"  (muted pink)
	//   AccentColor "220"    → DimColor "136"  (muted amber)
	//   AccentColor "141"    → DimColor "97"   (muted purple)
	DimColor string
}

// Creation — a single project in your Creations tab.
//
//	Name       — the project title (shown in the list)
//	Subheading — tech stack or short subtitle (shown when expanded)
//	Desc       — one-line description (shown when expanded)
type Creation struct {
	Name       string
	Subheading string
	Desc       string
}

// Link — a single entry in your Contact tab.
//
//	Label — a short tag like "GH", "LI", or an emoji like "✉"
//	Value — the URL, email, or handle
type Link struct {
	Label string
	Value string
}

// ============================================================================
// YOUR PORTFOLIO — edit everything below to make it yours.
// ============================================================================

// Portrait — your ASCII art photo.
//  1. Run:  ascii-image-converter portfolio/assets/photo.jpg -W 60 -b
//  2. Paste the output into portfolio/portrait.txt
//
//go:embed portrait.txt
var portrait string

// NameArt — your big stylised name.
//  1. Run:  figlet -f slant "yourname"
//  2. Paste the output into portfolio/name.txt
//
//go:embed name.txt
var nameArt string

var Portfolio = Config{
	Portrait: portrait,
	NameArt:  nameArt,

	// Your colors. Change the numbers to pick different colors (see options above).
	AccentColor: "39",
	DimColor:    "68",

	// --- ABOUT TAB ---

	// Your intro headline. Each line in quotes is a paragraph.
	RoleLines: []string{
		"is a ***software engineer*** who loves building things for the terminal",
	},
	// What you're working on right now. Displayed in bright white to stand out.
	CurrentWork: []string{
		"Currently building ***open-source tools*** and exploring the intersection of",
		"***design*** and ***developer experience***.",
	},
	// Your background (education, past experience, etc). Displayed in grey.
	BackgroundParas: []string{
		"Studied computer science, previously worked on web platforms and infrastructure.",
	},

	// --- CREATIONS TAB ---
	// Your projects. Add, remove, or edit entries. Copy the format to add more.
	Creations: []Creation{
		{
			Name:       "Example Project One",
			Subheading: "Go, PostgreSQL, Docker",
			Desc:       "A brief description of your first project goes here.",
		},
		{
			Name:       "Example Project Two",
			Subheading: "TypeScript, React, Node.js",
			Desc:       "A brief description of your second project goes here.",
		},
		{
			Name:       "SSH Portfolio",
			Subheading: "Go, Charmbracelet (Wish, Bubbletea, Lip Gloss)",
			Desc:       "This site! A terminal-based portfolio served over SSH.",
		},
	},

	// --- FUN FACTS TAB ---
	// One string per bullet point. Add or remove as many as you like.
	Facts: []string{
		"Replace these with your own fun facts",
		"Each string becomes a bullet point",
		"Add as many as you like",
		"Hobbies, achievements, interests — anything goes",
	},

	// --- CONTACT TAB ---
	// Your links. Format: {"Short Label", "url or handle"}
	Links: []Link{
		{"GH", "github.com/yourusername"},
		{"LI", "linkedin.com/in/yourusername"},
		{"✉", "you@example.com"},
		{"X", "@yourusername"},
	},
}
