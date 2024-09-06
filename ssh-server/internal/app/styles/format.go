package styles

import "github.com/charmbracelet/lipgloss"

var (
	padx2 = lipgloss.NewStyle().Padding(0, 2)
)

func PadX2(s string) string {
	return padx2.Render(s)
}
