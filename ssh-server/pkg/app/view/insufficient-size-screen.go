package view

import (
	"fmt"
	"ssh-server/config"
	"ssh-server/pkg/app"

	"github.com/charmbracelet/lipgloss"
)

func RenderInsufficientSizeScreen(tui *app.TUI) string {

	message := fmt.Sprintf("Terminal too small\nPlease resize to at least %dx%d", config.MinWidth, config.MinHeight)
	style := lipgloss.NewStyle().
		Foreground(green).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(green).
		Padding(1).
		Bold(true)

	return lipgloss.Place(tui.Width, tui.Height,
		lipgloss.Center, lipgloss.Center,
		style.Render(message),
	)
}
