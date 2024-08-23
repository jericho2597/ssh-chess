package update

import (
	"fmt"
	"ssh-server/pkg/app"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/notnil/chess"
)

type Update struct{}

func (u Update) Update(tui *app.TUI, msg tea.Msg) (tea.Model, tea.Cmd) {
	// Implement update logic
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return tui, tea.Quit
		case "enter":
			// Process move
			move, err := chess.UCINotation{}.Decode(tui.Game.State.Position(), tui.Input)
			if err != nil {
				tui.ErrorMsg = fmt.Sprintf("Invalid move: %v", err)
			} else if err := tui.Game.State.Move(move); err != nil {
				tui.ErrorMsg = fmt.Sprintf("Illegal move: %v", err)
			} else {
				tui.Input = ""
				tui.ErrorMsg = ""
			}
		case "backspace":
			if len(tui.Input) > 0 {
				tui.Input = tui.Input[:len(tui.Input)-1]
			}
		default:
			if len(msg.String()) == 1 {
				tui.Input += msg.String()
			}
		}
	case tea.WindowSizeMsg:
		return tui, doResize(tui, msg)
	}
	return tui, nil
}
