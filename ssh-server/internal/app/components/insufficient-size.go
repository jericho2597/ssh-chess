package components

import (
	"fmt"
	"ssh-server/config"
	"ssh-server/internal/app/model"
	"ssh-server/internal/app/styles"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type InsufficientSize struct{}

func NewInsufficientSize() *InsufficientSize {
	return &InsufficientSize{}
}

func (insufficientSize *InsufficientSize) View(state model.TuiState) string {

	boxStyle := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(styles.Accent).
		Padding(1, 2).
		Align(lipgloss.Center)

	content := fmt.Sprintf(
		"%s\n\n%s\n%s",
		lipgloss.NewStyle().Foreground(styles.Accent).Bold(true).Render("Insufficient Terminal Size"),
		"Your terminal window is too small to display the game properly.",
		fmt.Sprintf("Please resize to at least %dx%d.", config.MinWidth, config.MinHeight),
	)

	boxedContent := boxStyle.Render(content)

	centeredContent := lipgloss.Place(
		state.Width,
		state.Height,
		lipgloss.Center,
		lipgloss.Center,
		boxedContent,
	)

	return centeredContent
}

func (insufficientSize *InsufficientSize) Update(msg tea.Msg, tuiState model.TuiState) tea.Cmd {
	return nil
}
