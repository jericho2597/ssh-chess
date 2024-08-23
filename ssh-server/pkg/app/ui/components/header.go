package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const header = `
 @@@@@@    @@@@@@   @@@  @@@      @@@@@@@  @@@  @@@  @@@@@@@@   @@@@@@    @@@@@@   
@@@@@@@   @@@@@@@   @@@  @@@     @@@@@@@@  @@@  @@@  @@@@@@@@  @@@@@@@   @@@@@@@   
!@@       !@@       @@!  @@@     !@@       @@!  @@@  @@!       !@@       !@@       
!@!       !@!       !@!  @!@     !@!       !@!  @!@  !@!       !@!       !@!       
!!@@!!    !!@@!!    @!@!@!@!     !@!       @!@!@!@!  @!!!:!    !!@@!!    !!@@!!    
 !!@!!!    !!@!!!   !!!@!!!!     !!!       !!!@!!!!  !!!!!:     !!@!!!    !!@!!!   
     !:!       !:!  !!:  !!!     :!!       !!:  !!!  !!:            !:!       !:!  
    !:!       !:!   :!:  !:!     :!:       :!:  !:!  :!:           !:!       !:!   
:::: ::   :::: ::   ::   :::      ::: :::  ::   :::   :: ::::  :::: ::   :::: ::   
:: : :    :: : :     :   : :      :: :: :   :   : :  : :: ::   :: : :    :: : :   
`

var gradientColors = []lipgloss.Color{
	"#000080", // Dark Blue
	"#0000FF",
	"#1E90FF",
	"#00BFFF",
	"#87CEFA",
	"#B0E0E6",
	"#F0F8FF",
	"#FFE4B5", // Light Orange
	"#FFA500",
	"#FF8C00", // Dark Orange
}

// Render returns the string representation of the header banner
func RenderHeaderBanner(width int) string {

	lines := strings.Split(strings.TrimSpace(header), "\n")
	styledLines := make([]string, len(lines))

	for i, line := range lines {
		color := gradientColors[i%len(gradientColors)]
		style := lipgloss.NewStyle().
			Bold(true).
			Foreground(color).
			PaddingLeft(2).
			PaddingRight(2).
			Width(width)

		styledLines[i] = style.Render(line)
	}

	return lipgloss.JoinVertical(lipgloss.Left, styledLines...)
}
