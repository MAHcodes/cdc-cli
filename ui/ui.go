package ui

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

func PrintInfo(title, text string) {
	fmt.Println(SprintInfo(title, text))
}

func SprintInfo(title, text string) (info string) {
	title = titleStyle.Render(title + ":")
	text = textStyle.Render(text)

	info = lipgloss.JoinHorizontal(lipgloss.Center, title, text)
	info = infoStyle.Render(info)

	return info
}

func PrintTag(tag string) {
	fmt.Print(SprintTag(tag))
}

func SprintTag(tag string) string {
	tag = tagStyle.Render(tag)
	return tag
}
