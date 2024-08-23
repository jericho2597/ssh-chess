package app

import (
	config "ssh-server/config"
	model "ssh-server/pkg/app/model"

	tea "github.com/charmbracelet/bubbletea"
)

type Screen int

const (
	LoadingScreen Screen = iota
	AboutScreen
	GameScreen
)

type ViewerIfc interface {
	Render(tui *TUI) string
}

type UpdaterIfc interface {
	Update(tui *TUI, msg tea.Msg) (tea.Model, tea.Cmd)
}

type TUI struct {
	viewer  ViewerIfc
	updater UpdaterIfc

	Game model.Game

	Height         int
	Width          int
	SufficientSize bool
	ActiveScreen   Screen
	Input          string
	ErrorMsg       string
}

func NewTUI(v ViewerIfc, u UpdaterIfc, width int, height int) *TUI {

	initialSizeSufficient := (height > config.MinHeight && width > config.MinWidth)

	return &TUI{
		viewer:  v,
		updater: u,
		Game:    model.NewGame(),

		Width:          width,
		Height:         height,
		SufficientSize: initialSizeSufficient,
		ActiveScreen:   LoadingScreen,
	}
}

func (tui *TUI) View() string {
	return tui.viewer.Render(tui)
}

func (tui *TUI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return tui.updater.Update(tui, msg)
}

func (tui *TUI) Init() tea.Cmd {
	return nil
}
