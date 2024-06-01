package main

import (
	"net"
	"strings"
	"unicode/utf8"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	"github.com/charmbracelet/wish/bubbletea"
)

var HEADER = []string{
	"╱╱╱╱╱╱╱╱╱╭╮╱╱╱╱╱╱╱╱╱╱     ╭╮                          ",
	"╱╱╱╱╱╱╱╱╱┃┃╱╱╱╱╱╱╱╱╱╱     ┃┃               ╱╱╱╱╱╱╱╱   ",
	"╭━━┳━┳━━━┫╰━━━┳━━━━┳━━┳━━━╯┃ ╭━━━━┳━━┳━━━━┳━┳━┳━━━┳━━╮",
	"┃ ╭╋━┫ ╭━┫ ╭╮ ┃ ╭╮ ┃ ╭┫ ╭╮ ┃ ┃ ╭╮ ┃ ╭┫ ╭╮ ┃ ┃ ┃ ┃━┫ ╭╯",
	"┃ ┃┃ ┃ ╰━┫ ┃┃ ┃ ╭╮ ┃ ┃┃ ╰╯ ┃╱┃ ╰╯ ┃ ┃┃ ╰╯ ┣╮┃╭┫ ┃━┫ ┃ ",
	"╰━╯╰━┻━━━┻━╯╰━┻━╯╰━┻━╯╰━━━━╯╱╰━━╮ ┣━╯╰━━━━╯╰━╯╰━━━┻━╯ ",
	"           ╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╭━━╯ ┃╱╱╱╱╱╱╱            ",
	"           ╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╰━━━━╯╱╱╱╱╱╱╱            ",
	"",
}

var headerWidth int
var headerHeight int = len(HEADER)

func init() {
	for _, line := range HEADER {
		headerWidth = max(utf8.RuneCountInString(line), headerWidth)
	}
}

func main() {
	server, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort("localhost", "23234")),
		wish.WithHostKeyPath(".ssh/id_ed25519"),
		wish.WithMiddleware(
			bubbletea.Middleware(teaHandler),
			activeterm.Middleware(),
		),
	)

	if err != nil {
		panic(err)
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

	print("Stopping")
}

func teaHandler(session ssh.Session) (tea.Model, []tea.ProgramOption) {
	pty, _, _ := session.Pty()

	return model{
		pty.Window.Width,
		pty.Window.Height,
		calcPaddingToCenter(headerWidth, pty.Window.Width),
		0,
	}, []tea.ProgramOption{tea.WithAltScreen()}
}

type model struct {
	width                int
	height               int
	headerPadding        string
	selectedProjectIndex int
}

func (mod model) Init() tea.Cmd {
	return nil
}

func (mod model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		mod.width = msg.Width
		mod.height = msg.Height
		mod.headerPadding = calcPaddingToCenter(headerWidth, mod.width)
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return mod, tea.Quit
		}
	}

	var cmd tea.Cmd
	mod.selectedProjectIndex = HandleProjectListInput(msg, mod.selectedProjectIndex)
	return mod, cmd
}

func (mod model) View() string {
	var builder strings.Builder

	builder.WriteString(mod.headerPadding)
	builder.WriteString(strings.Join(HEADER, "\n"+mod.headerPadding))
	RenderProjectList(mod.selectedProjectIndex, mod.width, &builder)
	return builder.String()
}
