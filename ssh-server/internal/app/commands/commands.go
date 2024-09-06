package commands

import (
	"ssh-server/internal/app/messages"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// DoTick returns a tea.Cmd that sends a TickMsg after the specified duration.
// If no duration is provided, it defaults to 1000 milliseconds (1 second).
func DoTick(ms ...int) tea.Cmd {
	duration := 1000 * time.Millisecond // Default to 1000ms
	if len(ms) > 0 && ms[0] > 0 {
		duration = time.Duration(ms[0]) * time.Millisecond
	}

	return tea.Tick(duration, func(t time.Time) tea.Msg {
		return messages.TickMsg(t)
	})
}
