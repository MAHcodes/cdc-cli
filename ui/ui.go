package ui

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

var (
	white = lipgloss.AdaptiveColor{Light: "236", Dark: "#EBEBD0"}
	green = lipgloss.AdaptiveColor{Light: "236", Dark: "#779455"}

	infoStyle  = lipgloss.NewStyle().
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
)

func PrintInfo(title, text string) {
	title = titleStyle.Render(title + ":")
	text = textStyle.Render(text)

	info := lipgloss.JoinHorizontal(lipgloss.Center, title, text)
	info = infoStyle.Render(info)

	fmt.Println(info)
}

func PrintTag(tag string) {
	tag = tagStyle.Render(tag)
	fmt.Print(tag)
}
