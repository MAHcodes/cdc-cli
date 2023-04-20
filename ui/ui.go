package ui

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
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

func PrintStats(stats string)
