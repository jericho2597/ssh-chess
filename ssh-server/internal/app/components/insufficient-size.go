package components

import (
	"ssh-server/internal/app/model"

	tea "github.com/charmbracelet/bubbletea"
)

type InsufficientSize struct {
}

func NewInsufficientSize() *InsufficientSize {
	return &InsufficientSize{}
}

func (insufficientSize *InsufficientSize) View(state model.TuiState) string {
	return "Insufficient Size..."
}

func (insufficientSize *InsufficientSize) Update(msg tea.Msg) tea.Cmd {
	return nil
}
