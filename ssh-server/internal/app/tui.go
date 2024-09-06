package app

import (
	config "ssh-server/config"
	"ssh-server/internal/app/commands"
	"ssh-server/internal/app/components"
	"ssh-server/internal/app/messages"
	"ssh-server/internal/app/model"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	contentMaxWidth  = 90
	contentMaxHeight = 50
)

type TUI struct {
	Loading          components.Component
	Header           components.Component
	Game             components.Component
	About            components.Component
	InsufficientSize components.Component

	State model.TuiState
}

func (tui *TUI) Init() tea.Cmd {
	return commands.DoTick()
}

func NewTUI(width int, height int) *TUI {

	initialSizeSufficient := (height > config.MinHeight && width > config.MinWidth)

	return &TUI{
		Loading:          components.NewLoading(),
		Header:           components.NewHeader(),
		Game:             components.NewGame(),
		About:            components.NewAbout(),
		InsufficientSize: components.NewInsufficientSize(),

		State: model.TuiState{
			Width:          width,
			Height:         height,
			SufficientSize: initialSizeSufficient,
			ActiveScreen:   model.LoadingScreen,
		},
	}
}

func (tui *TUI) View() string {

	state := tui.State

	if !state.SufficientSize {
		return tui.InsufficientSize.View(state)
	}

	var content string
	switch state.ActiveScreen {
	case model.LoadingScreen:
		return tui.Loading.View(state)
	case model.AboutScreen:
		content = tui.About.View(state)
	case model.GameScreen:
		content = tui.Game.View(state)
	default:
		content = "Unknown screen"
	}

	contentWidth := min(contentMaxWidth, state.Width)
	centeredContent := lipgloss.Place(
		contentWidth,
		contentMaxHeight,
		lipgloss.Center,
		lipgloss.Top,
		content,
	)

	header := lipgloss.Place(
		tui.State.Width,
		1, // header height
		lipgloss.Center,
		lipgloss.Center,
		tui.Header.View(state),
	)

	joinedContent := lipgloss.JoinVertical(
		lipgloss.Center,
		header,
		centeredContent,
	)

	return lipgloss.Place(
		tui.State.Width,
		tui.State.Height,
		lipgloss.Center,
		lipgloss.Center,
		joinedContent,
	)
}

func (tui *TUI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case messages.FinishedLoadingMsg:
		tui.State.ActiveScreen = model.GameScreen
		return tui, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return tui, tea.Quit
		case "tab":
			tui.State.ActiveScreen = model.NextScreen(tui.State.ActiveScreen)
			return tui, nil
		}
	case tea.WindowSizeMsg:
		tui.doResize(msg)
		return tui, nil
	}

	cmd := tea.Batch(
		tui.About.Update(msg),
		tui.Game.Update(msg),
		tui.Loading.Update(msg),
		tui.Header.Update(msg),
		tui.InsufficientSize.Update(msg),
	)

	return tui, cmd
}

func (tui *TUI) doResize(msg tea.WindowSizeMsg) {
	tui.State.Height = msg.Height
	tui.State.Width = msg.Width
	tui.State.SufficientSize = tui.State.Width >= config.MinWidth && tui.State.Height >= config.MinHeight
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
