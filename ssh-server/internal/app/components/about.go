package components

import (
	"ssh-server/internal/app/model"
	"ssh-server/internal/app/styles"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
)

type About struct {
}

const (
	maxWidth = 70
)

func NewAbout() *About {
	return &About{}
}

func (about *About) View(state model.TuiState) string {

	width := min(maxWidth, state.Width)

	title := "- ssh chess -"
	description := wordwrap.String(`
this isn't your typical chess AI that would crush you within 7 moves.
it's actually just a digital version of me on the chessboard.

here's what makes it unique:

- playstyle - this AI has been trained exclusively on my Chess.com game history. it plays just like i do, with my playstyle and openings.

- level - The AI reflects my current skill level. it's not trying to be magnus, just a digital clone of me. expect blunders.

- terminal-based - it's all SSH-based. perfect for playing a game in between deployments from your cli

thanks for checking it out!
- Jericho
`, width)

	return lipgloss.JoinVertical(
		lipgloss.Center,
		"\n",
		lipgloss.Place(width, 1, lipgloss.Center, lipgloss.Center, styles.Bold(title)),
		lipgloss.NewStyle().MaxWidth(width).Render(description),
	)
}

func (about *About) Update(msg tea.Msg) tea.Cmd {
	return nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
