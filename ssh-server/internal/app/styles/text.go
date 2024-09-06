package styles

import (
	"github.com/charmbracelet/lipgloss"
)

func Bold(s string) string {
	return lipgloss.NewStyle().Bold(true).Render(s)
}

func Italic(s string) string {
	return lipgloss.NewStyle().Italic(true).Render(s)
}
