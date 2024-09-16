package board_render

import (
	color "ssh-server/internal/app/styles"

	"github.com/charmbracelet/lipgloss"
)

func CalculateSquareSize(boardWidth int, boardHeight int) int {
	maxSquareSizeByWidth := 8

	if boardWidth >= 96 {
		maxSquareSizeByWidth = 12
	} else if boardWidth >= 80 {
		maxSquareSizeByWidth = 10
	}

	maxSquareSizeByHeight := 8

	if boardHeight >= 48 {
		maxSquareSizeByHeight = 12
	} else if boardHeight >= 40 {
		maxSquareSizeByHeight = 10
	}

	return min(maxSquareSizeByHeight, maxSquareSizeByWidth)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func GetPieceRenderingFunc(pieceColor bool, squareColor bool, size int, pieceType string) func() []string {
	colorIndex := map[bool]int{true: 0, false: 1}[pieceColor]   // true for white, false for black
	squareIndex := map[bool]int{true: 0, false: 1}[squareColor] // true for light square, false for dark square
	sizeIndex := 0
	if size == 10 {
		sizeIndex = 1
	} else if size == 12 {
		sizeIndex = 2
	}

	return ColourSquareSizeMap[colorIndex][squareIndex][sizeIndex][pieceType]
}

func GetEmptySquareRenderingFunc(squareColor bool, size int) func() []string {

	if squareColor {
		switch size {
		case 8:
			return EmptyLightSquare8
		case 10:
			return EmptyLightSquare10
		case 12:
			return EmptyLightSquare12
		}
	}
	switch size {
	case 8:
		return EmptyDarkSquare8
	case 10:
		return EmptyDarkSquare10
	case 12:
		return EmptyDarkSquare12
	}

	return EmptyDarkSquare12
}

var (
	ColourSquareSizeMap = [2][2][5]map[string](func() []string){
		{ // White pieces
			{ // Light squares
				{ // Size 8
					"p": WhitePawnLightSquare8,
					"n": WhiteKnightLightSquare8,
					"b": WhiteBishopLightSquare8,
					"r": WhiteRookLightSquare8,
					"q": WhiteQueenLightSquare8,
					"k": WhiteKingLightSquare8,
				},
				{ // Size 10
					"p": WhitePawnLightSquare10,
					"n": WhiteKnightLightSquare10,
					"b": WhiteBishopLightSquare10,
					"r": WhiteRookLightSquare10,
					"q": WhiteQueenLightSquare10,
					"k": WhiteKingLightSquare10,
				},
				{ // Size 12
					"p": WhitePawnLightSquare12,
					"n": WhiteKnightLightSquare12,
					"b": WhiteBishopLightSquare12,
					"r": WhiteRookLightSquare12,
					"q": WhiteQueenLightSquare12,
					"k": WhiteKingLightSquare12,
				},
			},
			{ // Dark squares
				{ // Size 8
					"p": WhitePawnDarkSquare8,
					"n": WhiteKnightDarkSquare8,
					"b": WhiteBishopDarkSquare8,
					"r": WhiteRookDarkSquare8,
					"q": WhiteQueenDarkSquare8,
					"k": WhiteKingDarkSquare8,
				},
				{ // Size 10
					"p": WhitePawnDarkSquare10,
					"n": WhiteKnightDarkSquare10,
					"b": WhiteBishopDarkSquare10,
					"r": WhiteRookDarkSquare10,
					"q": WhiteQueenDarkSquare10,
					"k": WhiteKingDarkSquare10,
				},
				{ // Size 12
					"p": WhitePawnDarkSquare12,
					"n": WhiteKnightDarkSquare12,
					"b": WhiteBishopDarkSquare12,
					"r": WhiteRookDarkSquare12,
					"q": WhiteQueenDarkSquare12,
					"k": WhiteKingDarkSquare12,
				},
			},
		},
		{ // Black pieces
			{ // Light squares
				{ // Size 8
					"p": BlackPawnLightSquare8,
					"n": BlackKnightLightSquare8,
					"b": BlackBishopLightSquare8,
					"r": BlackRookLightSquare8,
					"q": BlackQueenLightSquare8,
					"k": BlackKingLightSquare8,
				},
				{ // Size 10
					"p": BlackPawnLightSquare10,
					"n": BlackKnightLightSquare10,
					"b": BlackBishopLightSquare10,
					"r": BlackRookLightSquare10,
					"q": BlackQueenLightSquare10,
					"k": BlackKingLightSquare10,
				},
				{ // Size 12
					"p": BlackPawnLightSquare12,
					"n": BlackKnightLightSquare12,
					"b": BlackBishopLightSquare12,
					"r": BlackRookLightSquare12,
					"q": BlackQueenLightSquare12,
					"k": BlackKingLightSquare12,
				},
			},
			{ // Dark squares
				{ // Size 8
					"p": BlackPawnDarkSquare8,
					"n": BlackKnightDarkSquare8,
					"b": BlackBishopDarkSquare8,
					"r": BlackRookDarkSquare8,
					"q": BlackQueenDarkSquare8,
					"k": BlackKingDarkSquare8,
				},
				{ // Size 10
					"p": BlackPawnDarkSquare10,
					"n": BlackKnightDarkSquare10,
					"b": BlackBishopDarkSquare10,
					"r": BlackRookDarkSquare10,
					"q": BlackQueenDarkSquare10,
					"k": BlackKingDarkSquare10,
				},
				{ // Size 12
					"p": BlackPawnDarkSquare12,
					"n": BlackKnightDarkSquare12,
					"b": BlackBishopDarkSquare12,
					"r": BlackRookDarkSquare12,
					"q": BlackQueenDarkSquare12,
					"k": BlackKingDarkSquare12,
				},
			},
		},
	}
)

func EmptyDarkSquare8() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("████████\x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("████████\x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("████████\x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("████████\x1b[0m"),
	}
}

func EmptyDarkSquare10() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("██████████\x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("██████████\x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("██████████\x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("██████████\x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("██████████\x1b[0m"),
	}
}

func EmptyDarkSquare12() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("████████████\x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("████████████\x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("████████████\x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("████████████\x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("████████████\x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("████████████\x1b[0m"),
	}
}

func EmptyLightSquare8() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("████████\x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("████████\x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("████████\x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("████████\x1b[0m"),
	}
}

func EmptyLightSquare10() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("██████████\x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("██████████\x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("██████████\x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("██████████\x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("██████████\x1b[0m"),
	}
}

func EmptyLightSquare12() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("████████████\x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("████████████\x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("████████████\x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("████████████\x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("████████████\x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("████████████\x1b[0m"),
	}
}

func BlackPawnLightSquare8() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("   ▂▂   \x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("  ▝") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("  ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▘  \x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("  ▗▄▄▖  \x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("  ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▄▄▄") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▘ ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
	}
}

func BlackPawnLightSquare10() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("   ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▁▁") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("  ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("   ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▎") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("  ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▊  ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("   ▝") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▄▄") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▘  ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("   ▟▇▇▙") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("   ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▅▅▅▅") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("  ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
	}
}

func BlackPawnLightSquare12() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("    ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("    ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("   ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("   ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▟") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("  ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▙") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("  ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("   ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▁") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("  ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▁") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("  ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("   ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▂▂▂▂") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("  ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("   ▗") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render(" ⎽⎽ ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▖  ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("   ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▆▆▆▆▆▆") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("  ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
	}
}

func BlackBishopLightSquare8() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("   ┏▌  ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("  ▗") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render(" ━") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▖ ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▙▁▂▟") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("━") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▃▃▃▃") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("━") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
	}
}

func BlackBishopLightSquare10() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("   ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▅▅") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("  ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▌\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("  ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▄") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render(" ▂") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▌\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("  ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▊") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("  ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▆") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▌\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▁") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▘▄▄▝") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▁") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ▌\x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▝") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▄▄▄▄▄▄") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("╴") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▊\x1b[0m"),
	}
}

func BlackBishopLightSquare12() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("     ▅▄") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("  ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("  ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("│\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("   ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▗") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("  ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("  ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("│\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("   ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▌") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("  ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▝") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▅▖") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("│\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("   ▝") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("   ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▘  ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("│\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▁▅▄") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("━━") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▄▅▁") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" │\x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▅▅▅▅▅▅▅▅") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▌\x1b[0m"),
	}
}
func BlackKnightLightSquare8() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" ▄▂▃▃▂") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▌") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("    ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▎") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("   ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▇") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▃▃▃▃") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▘ ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
	}
}

func BlackKnightLightSquare10() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▗▂▁▂▂▂ ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▊") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▆") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("     ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▊") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("    ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▂▄▆") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▌") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("    ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▝") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▙▄▄▄▄▄") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
	}
}

func BlackKnightLightSquare12() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("  ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▂ ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▁▁▁") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("  ▊\x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("╺▄") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("   ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▖") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▎") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("      ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render(" ▟") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▊") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("     ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("╺▆") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▊") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("     ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("   ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▅▅▅▅▅▅") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▘  ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
	}
}
func BlackRookLightSquare8() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▎") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▗▄▗▖▄▖ \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▎") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▌▁▁▁▁") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▌ \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▎") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▎") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▇▇▇") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▎ ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▄▄▄▄") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
	}
}

func BlackRookLightSquare10() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("│") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" ▃▖▃▃▗▃  \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("│") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▊") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("▎") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▝  ▘") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("▊") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▎ \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("│") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▅━━━━▅") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("  \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("│") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▊") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("    ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▎  \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▎ ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▝") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▅▅▅▅") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▘ ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
	}
}

func BlackRookLightSquare12() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("│") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" ▂▂ ▂▂ ▂▂") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("│") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▎") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("┗") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("▎▊") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("┛") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▊  \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("│") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▎▁▁▁▁▁▁") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▊") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("│") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▌") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▆▆▆▆▌") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("  \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("│") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▎") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("    ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("  \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▎") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▅▅▅▅▅▅") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
	}
}

func BlackKingLightSquare8() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▎") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("  ▄▄") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("  \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▃▄") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▜▛") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▄▄") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▊") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("   ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("  ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▎\x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▎") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▙▃▃▃▃▟") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
	}
}

func BlackKingLightSquare10() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("│") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("   ▄▄") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("   \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("│ ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▁ ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▄▄") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" ▁") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▘") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("▇ ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▙▟") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("  ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▝") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▖") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("      ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▎") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▇▄▄▄▄▄▄▇") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
	}
}

func BlackKingLightSquare12() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("│") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("    ▃▃     \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("│") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("  ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▊▁▁") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▎") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("  ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("  \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▎") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▗▅▆▅▖▗▅▆▅▖") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▎") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▎") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("   ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("  ") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("   ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▊") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▎") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▌") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("        ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▟") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▎ ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▅▅▅▅▅▅▅▅") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
	}
}

func BlackQueenLightSquare8() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("      ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▝") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▃") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("  ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▃") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▘") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▎") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("    ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▊ \x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" ▝") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▃▃▃▃") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▘") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
	}
}

func BlackQueenLightSquare10() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("  ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("      ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ▊\x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▎") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▖ ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▘▝") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" ▗▊") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▌") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▇") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render(" ▇") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▇") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▌") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▊") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("      ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▎") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("  ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▄▄▄▄▄▄") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
	}
}

func BlackQueenLightSquare12() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("           ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▗▁") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("  ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▘▝") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("  ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▁▖") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▊ ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▄") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▘") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("  ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▝") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▄") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▎") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▊") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("        ") + style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render("▎") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("      ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ▊\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.BlackPieceColor).Render("▅▅▅▅▅▅▅▅") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("  \x1b[0m"),
	}
}

func BlackPawnDarkSquare8() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("   ▂▂   \x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("  ▝") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("  ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▘  \x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("  ▗▄▄▖  \x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("  ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▄▄▄") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▘ ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
	}
}

func BlackPawnDarkSquare10() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("   ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▁▁") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("  ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("   ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▎") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("  ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▊  ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("   ▝") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▄▄") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▘  ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("   ▟▇▇▙") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("   ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▅▅▅▅") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("  ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
	}
}

func BlackPawnDarkSquare12() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("    ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("    ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("   ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("   ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▟") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("  ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▙") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("  ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("   ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▁") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("  ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▁") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("  ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("   ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▂▂▂▂") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("  ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("   ▗") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render(" ⎽⎽ ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▖  ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("   ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▆▆▆▆▆▆") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("  ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
	}
}

func BlackBishopDarkSquare8() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("   ┏▌  ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("  ▗") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render(" ━") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▖ ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▙▁▂▟") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("━") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▃▃▃▃") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("━") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
	}
}

func BlackBishopDarkSquare10() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("   ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▅▅") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("  ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▌\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("  ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▄") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render(" ▂") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▌\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("  ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▊") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("  ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▆") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▌\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▁") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▘▄▄▝") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▁") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ▌\x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▝") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▄▄▄▄▄▄") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("╴") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▊\x1b[0m"),
	}
}

func BlackBishopDarkSquare12() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("     ▅▄") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("  ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("  ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("│\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("   ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▗") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("  ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("  ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("│\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("   ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▌") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("  ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▝") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▅▖") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("│\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("   ▝") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("   ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▘  ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("│\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▁▅▄") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("━━") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▄▅▁") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" │\x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▅▅▅▅▅▅▅▅") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▌\x1b[0m"),
	}
}
func BlackKnightDarkSquare8() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" ▄▂▃▃▂") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▌") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("    ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▎") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("   ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▇") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▃▃▃▃") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▘ ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
	}
}

func BlackKnightDarkSquare10() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▗▂▁▂▂▂ ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▊") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▆") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("     ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▊") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("    ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▂▄▆") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▌") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("    ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▝") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▙▄▄▄▄▄") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
	}
}

func BlackKnightDarkSquare12() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("  ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▂ ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▁▁▁") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("  ▊\x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("╺▄") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("   ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▖") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▎") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("      ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render(" ▟") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▊") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("     ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("╺▆") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▊") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("     ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("   ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▅▅▅▅▅▅") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▘  ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
	}
}
func BlackRookDarkSquare8() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▎") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▗▄▗▖▄▖ \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▎") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▌▁▁▁▁") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▌ \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▎") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▎") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▇▇▇") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▎ ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▄▄▄▄") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
	}
}

func BlackRookDarkSquare10() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("│") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" ▃▖▃▃▗▃  \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("│") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▊") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("▎") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▝  ▘") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("▊") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▎ \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("│") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▅━━━━▅") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("  \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("│") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▊") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("    ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▎  \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▎ ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▝") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▅▅▅▅") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▘ ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
	}
}

func BlackRookDarkSquare12() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("│") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" ▂▂ ▂▂ ▂▂") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("│") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▎") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("┗") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("▎▊") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("┛") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▊  \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("│") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▎▁▁▁▁▁▁") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▊") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("│") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▌") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▆▆▆▆▌") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("  \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("│") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▎") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("    ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("  \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▎") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▅▅▅▅▅▅") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
	}
}
func BlackKingDarkSquare8() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▎") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("  ▄▄") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("  \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▃▄") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▜▛") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▄▄") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▊") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("   ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("  ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▎\x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▎") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▙▃▃▃▃▟") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
	}
}

func BlackKingDarkSquare10() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("│") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("   ▄▄") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("   \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("│ ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▁ ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▄▄") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" ▁") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▘") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("▇ ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▙▟") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("  ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▝") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▖") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("      ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▎") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▇▄▄▄▄▄▄▇") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
	}
}

func BlackKingDarkSquare12() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("│") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("    ▃▃     \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("│") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("  ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▊▁▁") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▎") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("  ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("  \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▎") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▗▅▆▅▖▗▅▆▅▖") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▎") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▎") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("   ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("  ") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("   ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▊") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▎") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▌") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("        ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▟") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▎ ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▅▅▅▅▅▅▅▅") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
	}
}
func BlackQueenDarkSquare8() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("      ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▝") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▃") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("  ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▃") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▘") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▎") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("    ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▊ \x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" ▝") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▃▃▃▃") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▘") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
	}
}

func BlackQueenDarkSquare10() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("  ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("      ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ▊\x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▎") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▖ ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▘▝") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" ▗▊") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▌") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▇") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render(" ▇") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▇") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▌") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▊") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("      ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▎") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("  ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▄▄▄▄▄▄") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
	}
}

func BlackQueenDarkSquare12() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("           ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▗▁") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("  ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▘▝") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("  ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▁▖") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▊ ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▄") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▘") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("  ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▝") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▄") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▎") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▊") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("        ") + style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render("▎") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render(" ") + style.Foreground(color.BlackPieceColor).Background(color.BlackPieceColor).Render("      ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ▊\x1b[0m"),
		style.Foreground(color.BlackPieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.BlackPieceColor).Render("▅▅▅▅▅▅▅▅") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("  \x1b[0m"),
	}
}

func WhitePawnDarkSquare8() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("   ▂▂   \x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("  ▝") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("  ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▘  \x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("  ▗▄▄▖  \x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("  ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▄▄▄") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▘ ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
	}
}

func WhitePawnDarkSquare10() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("   ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▁▁") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("  ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("   ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▎") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("  ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▊  ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("   ▝") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▄▄") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▘  ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("   ▟▇▇▙") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("   ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▅▅▅▅") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("  ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
	}
}

func WhitePawnDarkSquare12() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("    ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("    ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("   ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("   ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▟") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("  ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▙") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("  ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("   ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▁") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("  ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▁") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("  ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("   ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▂▂▂▂") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("  ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("   ▗") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render(" ⎽⎽ ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▖  ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("   ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▆▆▆▆▆▆") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("  ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
	}
}
func WhiteBishopDarkSquare8() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("   ┏▌  ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("  ▗") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render(" ━") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▖ ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▙▁▂▟") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("━") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▃▃▃▃") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("━") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
	}
}

func WhiteBishopDarkSquare10() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("   ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▅▅") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("  ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▌\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("  ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▄") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render(" ▂") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▌\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("  ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▊") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("  ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▆") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▌\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▁") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▘▄▄▝") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▁") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ▌\x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▝") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▄▄▄▄▄▄") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("╴") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▊\x1b[0m"),
	}
}

func WhiteBishopDarkSquare12() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("     ▅▄") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("  ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("  ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("│\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("   ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▗") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("  ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("  ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("│\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("   ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▌") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("  ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▝") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▅▖") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("│\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("   ▝") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("   ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▘  ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("│\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▁▅▄") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("━━") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▄▅▁") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" │\x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▅▅▅▅▅▅▅▅") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▌\x1b[0m"),
	}
}
func WhiteKnightDarkSquare8() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" ▄▂▃▃▂") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▌") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("    ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▎") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("   ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▇") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▃▃▃▃") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▘ ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
	}
}

func WhiteKnightDarkSquare10() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▗▂▁▂▂▂ ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▊") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▆") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("     ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▊") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("    ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▂▄▆") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▌") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("    ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▝") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▙▄▄▄▄▄") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
	}
}

func WhiteKnightDarkSquare12() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("  ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▂ ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▁▁▁") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("  ▊\x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("╺▄") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("   ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▖") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▎") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("      ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render(" ▟") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▊") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("     ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("╺▆") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▊") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("     ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("   ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▅▅▅▅▅▅") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▘  ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
	}
}
func WhiteRookDarkSquare8() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▎") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▗▄▗▖▄▖ \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▎") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▌▁▁▁▁") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▌ \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▎") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▎") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▇▇▇") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▎ ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▄▄▄▄") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
	}
}

func WhiteRookDarkSquare10() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("│") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" ▃▖▃▃▗▃  \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("│") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▊") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("▎") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▝  ▘") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("▊") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▎ \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("│") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▅━━━━▅") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("  \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("│") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▊") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("    ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▎  \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▎ ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▝") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▅▅▅▅") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▘ ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
	}
}

func WhiteRookDarkSquare12() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("│") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" ▂▂ ▂▂ ▂▂") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("│") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▎") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("┗") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("▎▊") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("┛") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▊  \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("│") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▎▁▁▁▁▁▁") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▊") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("│") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▌") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▆▆▆▆▌") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("  \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("│") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▎") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("    ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("  \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▎") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▅▅▅▅▅▅") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
	}
}
func WhiteKingDarkSquare8() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▎") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("  ▄▄") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("  \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▃▄") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▜▛") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▄▄") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▊") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("   ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("  ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▎\x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▎") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▙▃▃▃▃▟") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
	}
}

func WhiteKingDarkSquare10() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("│") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("   ▄▄") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("   \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("│ ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▁ ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▄▄") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" ▁") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▘") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("▇ ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▙▟") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("  ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▝") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▖") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("      ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▎") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▇▄▄▄▄▄▄▇") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
	}
}

func WhiteKingDarkSquare12() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("│") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("    ▃▃     \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("│") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("  ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▊▁▁") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▎") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("  ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("  \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▎") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▗▅▆▅▖▗▅▆▅▖") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▎") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▎") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("   ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("  ") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("   ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▊") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▎") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▌") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("        ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▟") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▎ ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▅▅▅▅▅▅▅▅") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
	}
}
func WhiteQueenDarkSquare8() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("      ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▝") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▃") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("  ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▃") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▘") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▎") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("    ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▊ \x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" ▝") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▃▃▃▃") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▘") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
	}
}

func WhiteQueenDarkSquare10() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("  ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("      ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ▊\x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▎") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▖ ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▘▝") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" ▗▊") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▌") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▇") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render(" ▇") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▇") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▌") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▊") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("      ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▎") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("  ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▄▄▄▄▄▄") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
	}
}

func WhiteQueenDarkSquare12() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("           ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▗▁") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("  ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▘▝") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("  ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▁▖") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▊ ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▄") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▘") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("  ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▝") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▄") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▎") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▊") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("        ") + style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render("▎") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("      ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ▊\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render(" ") + style.Foreground(color.DarkSquareColor).Background(color.WhitePieceColor).Render("▅▅▅▅▅▅▅▅") + style.Foreground(color.DarkSquareColor).Background(color.DarkSquareColor).Render("  \x1b[0m"),
	}
}

func WhitePawnLightSquare8() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("   ▂▂   \x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("  ▝") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("  ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▘  \x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("  ▗▄▄▖  \x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("  ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▄▄▄") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▘ ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
	}
}

func WhitePawnLightSquare10() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("   ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▁▁") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("  ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("   ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▎") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("  ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▊  ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("   ▝") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▄▄") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▘  ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("   ▟▇▇▙") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("   ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▅▅▅▅") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("  ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
	}
}

func WhitePawnLightSquare12() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("    ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("    ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("   ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("   ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▟") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("  ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▙") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("  ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("   ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▁") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("  ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▁") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("  ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("   ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▂▂▂▂") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("  ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("   ▗") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render(" ⎽⎽ ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▖  ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("   ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▆▆▆▆▆▆") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("  ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
	}
}
func WhiteBishopLightSquare8() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("   ┏▌  ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("  ▗") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render(" ━") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▖ ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▙▁▂▟") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("━") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▃▃▃▃") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("━") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
	}
}

func WhiteBishopLightSquare10() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("   ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▅▅") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("  ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▌\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("  ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▄") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render(" ▂") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▌\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("  ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▊") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("  ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▆") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▌\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▁") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▘▄▄▝") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▁") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ▌\x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▝") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▄▄▄▄▄▄") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("╴") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▊\x1b[0m"),
	}
}

func WhiteBishopLightSquare12() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("     ▅▄") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("  ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("  ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("│\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("   ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▗") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("  ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("  ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("│\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("   ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▌") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("  ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▝") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▅▖") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("│\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("   ▝") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("   ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▘  ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("│\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▁▅▄") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("━━") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▄▅▁") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" │\x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▅▅▅▅▅▅▅▅") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▌\x1b[0m"),
	}
}
func WhiteKnightLightSquare8() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" ▄▂▃▃▂") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▌") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("    ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▎") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("   ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▇") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▃▃▃▃") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▘ ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
	}
}

func WhiteKnightLightSquare10() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▗▂▁▂▂▂ ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▊") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▆") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("     ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▊") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("    ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▂▄▆") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▌") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("    ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▝") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▙▄▄▄▄▄") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
	}
}

func WhiteKnightLightSquare12() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("  ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▂ ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▁▁▁") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("  ▊\x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("╺▄") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("   ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▖") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▎") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("      ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render(" ▟") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▊") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("     ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("╺▆") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▊") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("     ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("   ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▅▅▅▅▅▅") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▘  ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
	}
}
func WhiteRookLightSquare8() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▎") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▗▄▗▖▄▖ \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▎") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▌▁▁▁▁") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▌ \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▎") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▎") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▇▇▇") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▎ ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▄▄▄▄") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
	}
}

func WhiteRookLightSquare10() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("│") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" ▃▖▃▃▗▃  \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("│") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▊") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("▎") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▝  ▘") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("▊") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▎ \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("│") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▅━━━━▅") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("  \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("│") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▊") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("    ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▎  \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▎ ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▝") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▅▅▅▅") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▘ ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
	}
}

func WhiteRookLightSquare12() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("│") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" ▂▂ ▂▂ ▂▂") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("│") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▎") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("┗") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("▎▊") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("┛") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▊  \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("│") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▎▁▁▁▁▁▁") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▊") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("│") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▌") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▆▆▆▆▌") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("  \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("│") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▎") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("    ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("  \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▎") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▅▅▅▅▅▅") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
	}
}

func WhiteKingLightSquare8() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▎") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("  ▄▄") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("  \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▃▄") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▜▛") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▄▄") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▊") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("   ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("  ") + style.Background(color.LightSquareColor).Background(color.WhitePieceColor).Render("▎\x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▎") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▙▃▃▃▃▟") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
	}
}

func WhiteKingLightSquare10() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("│") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("   ▄▄") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("   \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("│ ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▁ ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▄▄") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" ▁") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▘") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("▇ ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▙▟") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("  ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▝") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▖") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("      ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▎") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▇▄▄▄▄▄▄▇") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
	}
}

func WhiteKingLightSquare12() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("│") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("    ▃▃     \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("│") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("  ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▊▁▁") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▎") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("  ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("  \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▎") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▗▅▆▅▖▗▅▆▅▖") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▎") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▎") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("   ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("  ") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("   ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▊") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▎") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▌") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("        ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▟") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▎ ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▅▅▅▅▅▅▅▅") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
	}
}

func WhiteQueenLightSquare8() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("      ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▝") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▃") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("  ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▃") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▘") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▎") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("    ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▊ \x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" ▝") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▃▃▃▃") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▘") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
	}
}

func WhiteQueenLightSquare10() []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("  ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("      ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ▊\x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▎") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▖ ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▘▝") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" ▗▊") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▌") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▇") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render(" ▇") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▇") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▌") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▊") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("      ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▎") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("  ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▄▄▄▄▄▄") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
	}
}

func WhiteQueenLightSquare12() []string {

	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("           ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▗▁") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("  ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▘▝") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("  ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▁▖") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▊ ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▄") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▘") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("  ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▝") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▄") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▎") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▊") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("        ") + style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render("▎") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render(" ") + style.Foreground(color.WhitePieceColor).Background(color.WhitePieceColor).Render("      ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ▊\x1b[0m"),
		style.Foreground(color.WhitePieceColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render(" ") + style.Foreground(color.LightSquareColor).Background(color.WhitePieceColor).Render("▅▅▅▅▅▅▅▅") + style.Foreground(color.LightSquareColor).Background(color.LightSquareColor).Render("  \x1b[0m"),
	}
}
