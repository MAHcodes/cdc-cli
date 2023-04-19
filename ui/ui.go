package ui

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

type Style struct {
	info,
	title,
	text lipgloss.Style
}

var style = Style{
	info: lipgloss.NewStyle().
		Margin(0, 1),

	title: lipgloss.NewStyle().
		Foreground(lipgloss.Color("#779455")).
		Bold(true).
		Padding(0, 1).
    Width(15).
		MarginRight(1),

	text: lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ebebd0")),
}

func PrintInfo(title, text string) {
	title = style.title.Render(title + ":")
	text = style.text.Render(text)

	info := lipgloss.JoinHorizontal(lipgloss.Center, title, text)
	info = style.info.Render(info)

	fmt.Println(info)
}
