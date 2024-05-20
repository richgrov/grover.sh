package main

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var PROJECTS = []list.Item{
	project{"dCubed", "Scan & solve a Rubik's Cube with 2 Photos"},
	project{"villa", "Clean-room implementation of the Minecraft Beta 1.7.3 client"},
}

var projectStyle = lipgloss.NewStyle().
	PaddingLeft(2)

var selectedProjectStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#0078BA"))

type project struct {
	title       string
	description string
}

type itemDelegate struct{}

func (i project) FilterValue() string { return "" }

func (del itemDelegate) Height() int {
	return 2
}

func (del itemDelegate) Spacing() int {
	return 1
}

func (del itemDelegate) Update(tag tea.Msg, model *list.Model) tea.Cmd { return nil }

func (del itemDelegate) Render(w io.Writer, model list.Model, index int, listItem list.Item) {
	it := listItem.(project)

	var style lipgloss.Style
	var text string

	if model.Index() == index {
		style = selectedProjectStyle
		text = "\u25CF " + it.title + "\n  " + it.description
	} else {
		style = projectStyle
		text = it.title + "\n" + it.description
	}

	fmt.Fprint(w, style.Render(text))
}

func newProjectList(width int, height int) list.Model {
	projectList := list.New(PROJECTS, itemDelegate{}, min(width, 80), height)
	projectList.SetShowTitle(false)
	projectList.SetShowStatusBar(false)
	projectList.SetFilteringEnabled(false)

	return projectList
}
