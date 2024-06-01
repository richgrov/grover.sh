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

func RenderProjectList(selectedIndex int, viewportWidth int, builder *strings.Builder) {
	leftPadding := calcPaddingToCenter(PROJECT_LIST_WIDTH, viewportWidth)
	builder.WriteString("\n" + AnsiRed + centerText("Up: \u2191, k   Down: \u2193, j", viewportWidth))

	for i, project := range PROJECTS {
		color := AnsiReset
		if i == selectedIndex {
			color = AnsiBlue
		}

		if i == selectedIndex {
			builder.WriteString("\n\n" + leftPadding + color + "\u25CF " + project.name)
		} else {
			builder.WriteString("\n\n" + leftPadding + color + "  " + project.name)
		}

		builder.WriteString("\n" + leftPadding + "  " + color + project.description)
		renderProjectTags(&project, leftPadding+color, builder)
	}
}

func renderProjectTags(project *Project, prefix string, builder *strings.Builder) {
	builder.WriteString("\n" + "  " + prefix)
	lineLength := 0
	for tagIdx, tag := range project.tags {
		if lineLength+len(tag) > PROJECT_LIST_WIDTH {
			builder.WriteString("\n  " + prefix)
			lineLength = 0
		}

		builder.WriteString(tag)
		lineLength += len(tag)

		if tagIdx != len(project.tags)-1 {
			builder.WriteString(" \u2022 ")
			lineLength += 3
		}
	}
}
