package app

import (
	"log"
	config "ssh-server/config"
	"ssh-server/internal/app/board_render"
	"ssh-server/internal/app/commands"
	"ssh-server/internal/app/components"
	"ssh-server/internal/app/messages"
	"ssh-server/internal/app/model"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TUI struct {
	Loading          components.Component
	Header           components.Component
	Footer           components.Component
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
		Footer:           components.NewFooter(),
		Game:             components.NewGame(),
		About:            components.NewAbout(),
		InsufficientSize: components.NewInsufficientSize(),

		State: model.TuiState{
			Width:          width,
			Height:         height,
			SufficientSize: initialSizeSufficient,
			ActiveScreen:   model.LoadingScreen,

			ContentWidth:  min(config.ContentMaxWidth, width),
			ContentHeight: min(config.ContentMaxHeight, height-config.HeaderHeight-config.FooterHeight),

			SquareSize: board_render.CalculateSquareSize(min(config.ContentMaxWidth, width), min(config.ContentMaxHeight, height-config.HeaderHeight-config.FooterHeight)),
		},
	}
}

func (tui *TUI) View() string {

	state := tui.State

	if state.ActiveScreen == model.LoadingScreen {
		return tui.Loading.View(state)
	}

	if !state.SufficientSize {
		return tui.InsufficientSize.View(state)
	}

	var content string
	switch state.ActiveScreen {
	case model.AboutScreen:
		content = tui.About.View(state)
	case model.GameScreen:
		content = tui.Game.View(state)
	default:
		content = "Unknown screen"
	}

	centeredContent := lipgloss.Place(
		state.ContentWidth,
		state.ContentHeight,
		lipgloss.Center,
		lipgloss.Center,
		content,
	)

	header := lipgloss.Place(
		tui.State.Width,
		config.HeaderHeight,
		lipgloss.Center,
		lipgloss.Center,
		tui.Header.View(state),
	)

	footer := lipgloss.Place(
		tui.State.Width,
		config.FooterHeight,
		lipgloss.Center,
		lipgloss.Center,
		tui.Footer.View(state),
	)

	joinedContent := lipgloss.JoinVertical(
		lipgloss.Center,
		header,
		centeredContent,
		footer,
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
		tui.About.Update(msg, tui.State),
		tui.Game.Update(msg, tui.State),
		tui.Loading.Update(msg, tui.State),
		tui.Header.Update(msg, tui.State),
		tui.Footer.Update(msg, tui.State),
		tui.InsufficientSize.Update(msg, tui.State),
	)

	return tui, cmd
}

func (tui *TUI) doResize(msg tea.WindowSizeMsg) {
	tui.State.Height = msg.Height
	tui.State.Width = msg.Width
	tui.State.SufficientSize = tui.State.Width >= config.MinWidth && tui.State.Height >= config.MinHeight

	tui.State.ContentWidth = min(config.ContentMaxWidth, msg.Width)
	tui.State.ContentHeight = min(config.ContentMaxHeight, msg.Height-config.HeaderHeight-config.FooterHeight)

	tui.State.SquareSize = board_render.CalculateSquareSize(tui.State.ContentWidth, tui.State.ContentHeight)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
