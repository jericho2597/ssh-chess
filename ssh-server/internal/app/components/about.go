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

	width := min(maxWidth, state.ContentWidth)

	title := "- ssh chess -"
	description := wordwrap.String(`
- playstyle - this bot has been trained exclusively on my Chess.com game history. it plays just like i do, with my playstyle and openings.

- level - the bot reflects my current skill level. it's not trying to be magnus, just a digital clone of me. expect blunders.

- terminal-based - it's all SSH-based. perfect for playing a game in between deployments from your cli

`, width)
	signoff := "thanks for checking it out!\n- Jericho"

	return lipgloss.JoinVertical(
		lipgloss.Center,
		"\n",
		lipgloss.Place(width, 1, lipgloss.Center, lipgloss.Center, styles.Bold(title)),
		description,
		lipgloss.Place(width, 1, lipgloss.Center, lipgloss.Center, signoff),
	)
}

func (about *About) Update(msg tea.Msg, tuiState model.TuiState) tea.Cmd {
	return nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
