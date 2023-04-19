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
		Margin(0, 2),

	title: lipgloss.NewStyle().
		Foreground(lipgloss.Color("#7D56F4")).
		Bold(true).
		Padding(0, 1).
		MarginRight(2),

	text: lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")),
}

func PrintInfo(title, text string) {
	title = style.title.Render(title + ":")
	text = style.text.Render(text)

	info := lipgloss.JoinHorizontal(lipgloss.Center, title, text)
	info = style.info.Render(info)

	fmt.Println(info)
}
