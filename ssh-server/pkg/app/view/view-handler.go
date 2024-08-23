package view

import (
	"ssh-server/pkg/app"

	"github.com/charmbracelet/lipgloss"
)

var (
	green     = lipgloss.Color("#00FF00")
	darkGreen = lipgloss.Color("#005500")
)

type View struct{}

func (v View) Render(tui *app.TUI) string {

	if !tui.SufficientSize {
		return RenderInsufficientSizeScreen(tui)
	}

	switch tui.ActiveScreen {
	case app.LoadingScreen:
		return RenderLoadingScreen(tui)
	case app.AboutScreen:
		return RenderAboutScreen(tui)
	case app.GameScreen:
		return RenderGameScreen(tui)
	default:
		return "Unknown view"
	}
}
