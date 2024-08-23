package view

import (
	"ssh-server/pkg/app"
	"ssh-server/pkg/app/ui/components"

	"github.com/charmbracelet/lipgloss"
)

func RenderLoadingScreen(tui *app.TUI) string {

	headerStr := components.RenderHeaderBanner(tui.Width)

	return lipgloss.JoinVertical(lipgloss.Center, headerStr, "Loading...")
}
