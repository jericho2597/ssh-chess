package update

import (
	"ssh-server/config"
	"ssh-server/pkg/app"

	tea "github.com/charmbracelet/bubbletea"
)

func doResize(tui *app.TUI, msg tea.WindowSizeMsg) tea.Cmd {
	tui.Height = msg.Height
	tui.Width = msg.Width
	tui.SufficientSize = tui.Width >= config.MinWidth && tui.Height >= config.MinHeight
	return nil
}
