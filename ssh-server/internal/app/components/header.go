package components

import (
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

func (h *Header) Update(msg tea.Msg) tea.Cmd {
	return nil
}

func (h *Header) View(state model.TuiState) string {

	return table.New().Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(s.Border)).
		Row(getHeaderTabString()...).Render()
}

func getHeaderTabString() []string {

	raised := lipgloss.NewStyle().
		Bold(true).
		Foreground(s.Primary)

	muted := lipgloss.NewStyle().
		Foreground(s.Secondary)

	return []string{
		s.PadX2(raised.Padding(0, 2).Render("♞ ssh chess")),
		s.PadX2(raised.Render("g") + muted.Render(" game")),
		s.PadX2(raised.Render("a") + muted.Render(" about")),
	}
}
