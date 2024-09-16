package components

import (
	"ssh-server/config"
	"ssh-server/internal/app/model"
	s "ssh-server/internal/app/styles"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Footer struct {
}

func NewFooter() *Footer {
	return &Footer{}
}

func (h *Footer) Update(msg tea.Msg, tuiState model.TuiState) tea.Cmd {
	return nil
}

func (f *Footer) View(state model.TuiState) string {
	separator := lipgloss.NewStyle().
		Foreground(s.Secondary).
		Render(strings.Repeat("─", state.ContentWidth))

	highlighted := lipgloss.NewStyle().
		Bold(true).
		Foreground(s.Primary)

	muted := lipgloss.NewStyle().
		Foreground(s.Secondary)

	infoText := highlighted.Render("tab") +
		muted.Render(" - switch tabs   ") +
		highlighted.Render("n") +
		muted.Render(" - new game   ") +
		highlighted.Render("ctrl+c") +
		muted.Render(" - exit")

	return lipgloss.Place(
		state.Width,
		config.FooterHeight,
		lipgloss.Center,
		lipgloss.Top,
		separator + "\n" + infoText,
	)
}
