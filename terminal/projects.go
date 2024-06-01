package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

const PROJECT_LIST_WIDTH int = 80

type Project struct {
	name        string
	description string
	tags        []string
}

func project(name string, description string, tags ...string) Project {
	return Project{
		name:        name,
		description: description,
		tags:        tags,
	}
}

var PROJECTS = []Project{
	project(
		"dCubed", "Scan & solve a Rubik's Cube with 2 Photos",
		"Java", "Javalin", "OpenCV", "Python", "PyTorch", "Flask", "TypeScript", "React", "TailwindCSS", "Three.js",
	),
	project(
		"villa", "Clean-room implementation of the Minecraft Beta 1.7.3 client",
		"Rust", "WGPU", "WGSL", "TCP",
	),
	project(
		"grover.sh", "Source code for my website and this SSH server",
		"Hugo", "TailwindCSS", "wish", "SSH",
	),
}

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
	result := "\n" + AnsiRed + centerText("Up: \u2191, k   Down: \u2193, j", viewportWidth)

	for i, project := range PROJECTS {
		color := AnsiReset
		if i == selectedIndex {
			color = AnsiBlue
		}

		if i == selectedIndex {
			result += "\n\n" + leftPadding + color + "\u25CF " + project.name
		} else {
			result += "\n\n" + leftPadding + "  " + color + project.name
		}

		result += "\n" + leftPadding + "  " + color + project.description
		result += "\n" + leftPadding + "  " + color + strings.Join(project.tags, " \u2022 ")
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
