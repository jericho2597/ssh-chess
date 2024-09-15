package components

import (
	"math/rand"
	"ssh-server/internal/app/commands"
	"ssh-server/internal/app/messages"
	"ssh-server/internal/app/model"
	"ssh-server/internal/app/styles"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Loading struct {
	text       string
	typedText  string
	blink      bool
	blinkCount int
	typeIndex  int
}

func NewLoading() *Loading {
	return &Loading{
		text:       "ssh chess",
		typedText:  "",
		blink:      true,
		blinkCount: 0,
		typeIndex:  0,
	}
}

func (loading *Loading) View(state model.TuiState) string {
	cursorStyle := lipgloss.NewStyle().
		Background(styles.Accent).
		SetString(" ")

	cursor := " "
	if loading.blink && loading.typeIndex >= len(loading.text) {
		cursor = cursorStyle.Render()
	}

	return lipgloss.Place(
		state.Width,
		state.Height,
		lipgloss.Center,
		lipgloss.Center,
		loading.typedText+cursor)
}

func (loading *Loading) Update(msg tea.Msg, tuiState model.TuiState) tea.Cmd {

	_, ok := msg.(messages.TickMsg)
	if !ok {
		// only update loading state on a tick message.
		return nil
	}

	if loading.typeIndex < len(loading.text) {
		loading.typedText = loading.text[:loading.typeIndex+1]
		loading.typeIndex++
		return commands.DoTick(rand.Intn(201) + 100)
	}

	loading.blink = !loading.blink
	loading.blinkCount++

	if loading.blinkCount > 5 {
		return func() tea.Msg {
			return messages.FinishedLoadingMsg{}
		}
	}

	return commands.DoTick(700)
}
