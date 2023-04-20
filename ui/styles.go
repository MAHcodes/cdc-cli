package ui

import "github.com/charmbracelet/lipgloss"

var (
	white = lipgloss.AdaptiveColor{Light: "#312E2B", Dark: "#EBEBD0"}
	green = lipgloss.AdaptiveColor{Light: "#769656", Dark: "#4E7838"}

	infoStyle = lipgloss.NewStyle().
			Margin(0, 1)

	titleStyle = lipgloss.NewStyle().
			Foreground(green).
			Bold(true).
			Width(15).
			MarginRight(1)

	textStyle = lipgloss.NewStyle().
			Foreground(white)

	tagStyle = lipgloss.NewStyle().
			Foreground(white).
			Background(green).
			Bold(true).
			Padding(0, 1).
			MarginLeft(1)

	boxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(green).
			Padding(1)
)
