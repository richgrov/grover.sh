package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const PROJECT_LIST_WIDTH int = 80

type Project struct {
	name        string
	description string
}

var PROJECTS = []Project{
	{"dCubed", "Scan & solve a Rubik's Cube with 2 Photos"},
	{"villa", "Clean-room implementation of the Minecraft Beta 1.7.3 client"},
}

var selectedStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#0078BA"))

var navStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#5C5C5C"))

func HandleProjectListInput(msg tea.Msg, selectedIndex int) int {
	key, ok := msg.(tea.KeyMsg)
	if !ok {
		return selectedIndex
	}

	switch key.String() {
	case "up", "k":
		return max(selectedIndex-1, 0)

	case "down", "j":
		return min(selectedIndex+1, len(PROJECTS)-1)
	}

	return selectedIndex
}

func RenderProjectList(selectedIndex int, viewportWidth int) string {
	leftPadding := calcPaddingToCenter(PROJECT_LIST_WIDTH, viewportWidth)
	result := "\n" + navStyle.Render(centerText("Up: \u2191, k   Down: \u2193, j", viewportWidth))

	for i, project := range PROJECTS {
		if i == selectedIndex {
			result += selectedStyle.Render("\n\n" + leftPadding + "\u25CF " + project.name + "\n" + leftPadding + "  " + project.description)
		} else {
			result += "\n\n" + leftPadding + "  " + project.name + "\n" + leftPadding + "  " + project.description
		}
	}

	return result
}

func calcPaddingToCenter(lineWidth int, viewportWidth int) string {
	return strings.Repeat(" ", viewportWidth/2-lineWidth/2)
}

func centerText(text string, viewportWidth int) string {
	textLen := len([]rune(text))
	return calcPaddingToCenter(textLen, viewportWidth) + text
}
