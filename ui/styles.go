package ui

import "github.com/charmbracelet/lipgloss"

var (
	white = lipgloss.AdaptiveColor{Light: "#312E2B", Dark: "#EBEBD0"}
	gray  = lipgloss.Color("#989795")
	green = lipgloss.AdaptiveColor{Light: "#769656", Dark: "#4E7838"}

	infoStyle = lipgloss.NewStyle().
			Margin(0, 1)

	titleStyle = lipgloss.NewStyle().
			Foreground(green).
			Bold(true).
			MarginRight(1)

	textStyle = lipgloss.NewStyle().
			Foreground(white)

	tagStyle = lipgloss.NewStyle().
			Foreground(white).
			Background(green).
			Bold(true).
			Padding(0, 1).
			Margin(0, 1)

	BoxStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderLeft(true).
			BorderLeftForeground(green).
			Margin(0, 1).
			Padding(1, 0, 0, 1)
)
