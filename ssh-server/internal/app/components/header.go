package components

import (
	"ssh-server/config"
	"ssh-server/internal/app/model"
	s "ssh-server/internal/app/styles"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

type Header struct {
}

func NewHeader() *Header {
	return &Header{}
}

func (h *Header) Update(msg tea.Msg, tuiState model.TuiState) tea.Cmd {
	return nil
}

func (h *Header) View(state model.TuiState) string {

	return lipgloss.Place(
		state.Width,
		config.HeaderHeight,
		lipgloss.Center,
		lipgloss.Bottom,
		table.New().Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(s.Border)).
		Row(getHeaderTabStrings(state.ActiveScreen)...).Render(),
	)
}

func getHeaderTabStrings(activeScreen model.Screen) []string {

	logo := lipgloss.NewStyle().
		Bold(true).
		Foreground(s.Accent)

	highlighted := lipgloss.NewStyle().
		Bold(true).
		Foreground(s.Primary)

	muted := lipgloss.NewStyle().
		Foreground(s.Secondary)

	gameString := muted.Render(" game ")
	aboutString := muted.Render(" about ")

	if activeScreen == model.GameScreen {
		gameString = highlighted.Render(" game ")
	} else {
		aboutString = highlighted.Render(" about ")
	}

	return []string{
		s.PadX2(logo.Padding(0, 2).Render("♜ ssh chess")),
		s.PadX2(highlighted.Render("♞") + gameString),
		s.PadX2(highlighted.Render("♝") + aboutString),
	}
}
