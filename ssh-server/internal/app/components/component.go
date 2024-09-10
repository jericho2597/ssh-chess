package components

import (
	"ssh-server/internal/app/model"

	tea "github.com/charmbracelet/bubbletea"
)

type Component interface {
	View(state model.TuiState) string
	Update(msg tea.Msg, tuiState model.TuiState) tea.Cmd
}
