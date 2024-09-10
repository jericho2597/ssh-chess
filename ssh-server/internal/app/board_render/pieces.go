package board_render

import "github.com/charmbracelet/lipgloss"

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

func GetPieceRenderingFunc(pieceColor bool, squareColor bool, size int, pieceType string) func(lipgloss.Color, lipgloss.Color, lipgloss.Color, lipgloss.Color) []string {
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

func GetEmptySquareRenderingFunc(squareColor bool, size int) func(lipgloss.Color, lipgloss.Color, lipgloss.Color, lipgloss.Color) []string {

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
	ColourSquareSizeMap = [2][2][5]map[string](func(lipgloss.Color, lipgloss.Color, lipgloss.Color, lipgloss.Color) []string){
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

func EmptyDarkSquare8(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("████████\x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("████████\x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("████████\x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("████████\x1b[0m"),
	}
}

func EmptyDarkSquare10(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("██████████\x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("██████████\x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("██████████\x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("██████████\x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("██████████\x1b[0m"),
	}
}

func EmptyDarkSquare12(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("████████████\x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("████████████\x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("████████████\x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("████████████\x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("████████████\x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("████████████\x1b[0m"),
	}
}

func EmptyLightSquare8(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("████████\x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("████████\x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("████████\x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("████████\x1b[0m"),
	}
}

func EmptyLightSquare10(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("██████████\x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("██████████\x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("██████████\x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("██████████\x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("██████████\x1b[0m"),
	}
}

func EmptyLightSquare12(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("████████████\x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("████████████\x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("████████████\x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("████████████\x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("████████████\x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("████████████\x1b[0m"),
	}
}

func BlackPawnLightSquare8(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("   ▂▂   \x1b[0m"),
		style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("  ▝") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("  ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▘  \x1b[0m"),
		style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("  ▗▄▄▖  \x1b[0m"),
		style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("  ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▄▄▄") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▘ ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" \x1b[0m"),
	}
}

func BlackPawnLightSquare10(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("   ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▁▁") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("  ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("   ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▎") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("  ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▊  ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("   ▝") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▄▄") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▘  ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("   ▟▇▇▙") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("   ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▅▅▅▅") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("  ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" \x1b[0m"),
	}
}

func BlackPawnLightSquare12(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("    ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render("    ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("   ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("   ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▟") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("  ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▙") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("  ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("   ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▁") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("  ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▁") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("  ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("   ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▂▂▂▂") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("  ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("   ▗") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render(" ⎽⎽ ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▖  ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("   ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▆▆▆▆▆▆") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("  ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" \x1b[0m"),
	}
}

func BlackBishopLightSquare8(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("   ┏▌  ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("  ▗") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render(" ━") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▖ ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Background(BlackPieceColor).Background(LightSquareColor).Render("▙▁▂▟") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("━") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▃▃▃▃") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("━") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" \x1b[0m"),
	}
}

func BlackBishopLightSquare10(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("   ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▅▅") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("  ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▌\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("  ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▄") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render(" ▂") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▌\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("  ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▊") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("  ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▆") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render(" ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("") + style.Background(LightSquareColor).Background(BlackPieceColor).Render(" ") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render(" ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▌\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▁") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▘▄▄▝") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▁") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ▌\x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▝") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▄▄▄▄▄▄") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("╴") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▊\x1b[0m"),
	}
}

func BlackBishopLightSquare12(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("     ▅▄") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render("  ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("  ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render("│\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("   ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▗") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("  ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("  ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render("│\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("   ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▌") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("  ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▝") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▅▖") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render("│\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("   ▝") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render(" ") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("   ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▘  ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render("│\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▁▅▄") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("━━") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▄▅▁") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" │\x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▅▅▅▅▅▅▅▅") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▌\x1b[0m"),
	}
}
func BlackKnightLightSquare8(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" ▄▂▃▃▂") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▌") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("    ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▎") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("   ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▇") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▃▃▃▃") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▘ ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" \x1b[0m"),
	}
}

func BlackKnightLightSquare10(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▗▂▁▂▂▂ ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▊") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▆") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("     ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render(" ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▊") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("    ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▂▄▆") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▌") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("    ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▝") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▙▄▄▄▄▄") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" \x1b[0m"),
	}
}

func BlackKnightLightSquare12(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("  ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▂ ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▁▁▁") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render("  ▊\x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("╺▄") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render(" ") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render(" ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render(" ") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("   ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▖") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▎") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("      ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render(" ▟") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▊") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("     ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("╺▆") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▊") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("     ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("   ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▅▅▅▅▅▅") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▘  ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" \x1b[0m"),
	}
}
func BlackRookLightSquare8(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▎") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▗▄▗▖▄▖ \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▎") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▌▁▁▁▁") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▌ \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▎") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▎") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▇▇▇") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▎ ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▄▄▄▄") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" \x1b[0m"),
	}
}

func BlackRookLightSquare10(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("│") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" ▃▖▃▃▗▃  \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("│") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▊") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("▎") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▝  ▘") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("▊") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▎ \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("│") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▅━━━━▅") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("  \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("│") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▊") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("    ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▎  \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▎ ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▝") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▅▅▅▅") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▘ ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" \x1b[0m"),
	}
}

func BlackRookLightSquare12(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("│") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" ▂▂ ▂▂ ▂▂") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("│") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▎") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render(" ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("┗") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("▎▊") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("┛") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▊  \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("│") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▎▁▁▁▁▁▁") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▊") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("│") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▌") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▆▆▆▆▌") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("  \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("│") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▎") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("    ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render(" ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("  \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▎") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▅▅▅▅▅▅") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" \x1b[0m"),
	}
}
func BlackKingLightSquare8(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▎") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("  ▄▄") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("  \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▃▄") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▜▛") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▄▄") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▊") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("   ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render(" ") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("  ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▎\x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▎") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▙▃▃▃▃▟") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" \x1b[0m"),
	}
}

func BlackKingLightSquare10(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("│") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("   ▄▄") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("   \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("│ ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▁ ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▄▄") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" ▁") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▘") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("▇ ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▙▟") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("  ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▝") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▖") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("      ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render(" ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▎") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▇▄▄▄▄▄▄▇") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" \x1b[0m"),
	}
}

func BlackKingLightSquare12(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("│") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("    ▃▃     \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("│") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("  ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▊▁▁") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▎") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render("  ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("  \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▎") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▗▅▆▅▖▗▅▆▅▖") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▎") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▎") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("   ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("  ") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("   ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▊") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▎") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▌") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("        ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▟") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▎ ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▅▅▅▅▅▅▅▅") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" \x1b[0m"),
	}
}
func BlackQueenLightSquare8(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("      ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Background(BlackPieceColor).Background(LightSquareColor).Render("▝") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▃") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("  ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▃") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▘") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▎") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("    ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▊ \x1b[0m"),
		style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" ▝") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▃▃▃▃") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▘") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" \x1b[0m"),
	}
}

func BlackQueenLightSquare10(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("  ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("      ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ▊\x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▎") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▖ ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▘▝") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" ▗▊") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▌") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▇") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render(" ▇") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▇") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▌") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▊") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("      ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▎") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("  ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▄▄▄▄▄▄") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" \x1b[0m"),
	}
}

func BlackQueenLightSquare12(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("           ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▗▁") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render("  ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▘▝") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render("  ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▁▖") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▊ ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▄") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▘") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("  ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▝") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▄") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render(" ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▎") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▊") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("        ") + style.Foreground(BlackPieceColor).Background(LightSquareColor).Render("▎") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render(" ") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("      ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render(" ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ▊\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(BlackPieceColor).Render("▅▅▅▅▅▅▅▅") + style.Foreground(LightSquareColor).Background(LightSquareColor).Render("  \x1b[0m"),
	}
}

func BlackPawnDarkSquare8(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("   ▂▂   \x1b[0m"),
		style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("  ▝") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("  ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▘  \x1b[0m"),
		style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("  ▗▄▄▖  \x1b[0m"),
		style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("  ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▄▄▄") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▘ ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
	}
}

func BlackPawnDarkSquare10(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("   ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▁▁") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("  ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("   ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▎") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("  ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▊  ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("   ▝") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▄▄") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▘  ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("   ▟▇▇▙") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("   ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▅▅▅▅") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("  ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
	}
}

func BlackPawnDarkSquare12(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("    ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("    ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("   ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("   ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▟") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("  ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▙") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("  ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("   ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▁") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("  ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▁") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("  ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("   ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▂▂▂▂") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("  ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("   ▗") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render(" ⎽⎽ ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▖  ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("   ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▆▆▆▆▆▆") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("  ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
	}
}
func BlackBishopDarkSquare8(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("   ┏▌  ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("  ▗") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render(" ━") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▖ ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▙▁▂▟") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("━") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▃▃▃▃") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("━") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
	}
}

func BlackBishopDarkSquare10(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("   ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▅▅") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("  ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▌\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("  ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▄") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render(" ▂") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▌\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("  ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▊") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("  ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▆") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render(" ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▌\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▁") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▘▄▄▝") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▁") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ▌\x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▝") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▄▄▄▄▄▄") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("╴") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▊\x1b[0m"),
	}
}

func BlackBishopDarkSquare12(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("     ▅▄") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("  ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("  ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("│\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("   ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▗") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("  ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("  ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("│\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("   ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▌") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("  ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▝") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▅▖") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("│\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("   ▝") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render(" ") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("   ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▘  ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("│\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▁▅▄") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("━━") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▄▅▁") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" │\x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▅▅▅▅▅▅▅▅") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▌\x1b[0m"),
	}
}
func BlackKnightDarkSquare8(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" ▄▂▃▃▂") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▌") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("    ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▎") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("   ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▇") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▃▃▃▃") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▘ ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
	}
}

func BlackKnightDarkSquare10(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▗▂▁▂▂▂ ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▊") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▆") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("     ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▊") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("    ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▂▄▆") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▌") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("    ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▝") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▙▄▄▄▄▄") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
	}
}

func BlackKnightDarkSquare12(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("  ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▂ ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▁▁▁") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("  ▊\x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("╺▄") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render(" ") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render(" ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render(" ") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("   ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▖") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▎") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("      ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render(" ▟") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▊") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("     ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("╺▆") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▊") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("     ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("   ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▅▅▅▅▅▅") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▘  ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
	}
}
func BlackRookDarkSquare8(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▎") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▗▄▗▖▄▖ \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▎") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▌▁▁▁▁") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▌ \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▎") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▎") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▇▇▇") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▎ ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▄▄▄▄") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" \x1b[0m"),
	}
}

func BlackRookDarkSquare10(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("│") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" ▃▖▃▃▗▃  \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("│") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▊") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("▎") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▝  ▘") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("▊") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▎ \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("│") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▅━━━━▅") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("  \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("│") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▊") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("    ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▎  \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▎ ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▝") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▅▅▅▅") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▘ ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
	}
}

func BlackRookDarkSquare12(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("│") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" ▂▂ ▂▂ ▂▂") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("│") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▎") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render(" ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("┗") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("▎▊") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("┛") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▊  \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("│") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▎▁▁▁▁▁▁") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▊") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("│") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▌") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▆▆▆▆▌") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("  \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("│") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▎") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("    ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("  \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▎") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▅▅▅▅▅▅") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
	}
}
func BlackKingDarkSquare8(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▎") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("  ▄▄") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("  \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▃▄") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▜▛") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▄▄") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▊") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("   ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render(" ") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("  ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▎\x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▎") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▙▃▃▃▃▟") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
	}
}

func BlackKingDarkSquare10(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("│") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("   ▄▄") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("   \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("│ ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▁ ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▄▄") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" ▁") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▘") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("▇ ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▙▟") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("  ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▝") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▖") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("      ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▎") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▇▄▄▄▄▄▄▇") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" \x1b[0m"),
	}
}

func BlackKingDarkSquare12(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("│") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("    ▃▃     \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("│") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("  ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▊▁▁") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▎") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("  ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("  \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▎") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▗▅▆▅▖▗▅▆▅▖") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▎") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▎") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("   ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("  ") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("   ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▊") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▎") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▌") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("        ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▟") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▎ ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▅▅▅▅▅▅▅▅") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" \x1b[0m"),
	}
}
func BlackQueenDarkSquare8(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("      ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▝") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▃") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("  ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▃") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▘") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▎") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("    ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▊ \x1b[0m"),
		style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" ▝") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▃▃▃▃") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▘") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
	}
}

func BlackQueenDarkSquare10(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("  ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("      ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ▊\x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▎") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▖ ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▘▝") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" ▗▊") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▌") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▇") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render(" ▇") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▇") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▌") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▊") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("      ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▎") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("  ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▄▄▄▄▄▄") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
	}
}

func BlackQueenDarkSquare12(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("           ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▗▁") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("  ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▘▝") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("  ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▁▖") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▊ ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▄") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▘") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("  ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▝") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▄") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render(" ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▎") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▊") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("        ") + style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render("▎") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render(" ") + style.Foreground(BlackPieceColor).Background(BlackPieceColor).Render("      ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ▊\x1b[0m"),
		style.Foreground(BlackPieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(BlackPieceColor).Render("▅▅▅▅▅▅▅▅") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("  \x1b[0m"),
	}
}

func WhitePawnDarkSquare8(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("   ▂▂   \x1b[0m"),
		style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("  ▝") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("  ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▘  \x1b[0m"),
		style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("  ▗▄▄▖  \x1b[0m"),
		style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("  ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▄▄▄") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▘ ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
	}
}

func WhitePawnDarkSquare10(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("   ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▁▁") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("  ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("   ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▎") + style.Foreground(WhitePieceColor).Background(WhitePieceColor).Render("  ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▊  ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("   ▝") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▄▄") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▘  ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("   ▟▇▇▙") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("   ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▅▅▅▅") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("  ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
	}
}

func WhitePawnDarkSquare12(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("    ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("    ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("   ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("   ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▟") + style.Foreground(WhitePieceColor).Background(WhitePieceColor).Render("  ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▙") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("  ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("   ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▁") + style.Foreground(WhitePieceColor).Background(WhitePieceColor).Render("  ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▁") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("  ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("   ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▂▂▂▂") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("  ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("   ▗") + style.Foreground(WhitePieceColor).Background(WhitePieceColor).Render(" ⎽⎽ ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▖  ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("   ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▆▆▆▆▆▆") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("  ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
	}
}
func WhiteBishopDarkSquare8(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("   ┏▌  ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("  ▗") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render(" ━") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▖ ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▙▁▂▟") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("━") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▃▃▃▃") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("━") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
	}
}

func WhiteBishopDarkSquare10(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("   ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▅▅") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("  ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▌\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("  ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▄") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render(" ▂") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▌\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("  ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▊") + style.Foreground(WhitePieceColor).Background(WhitePieceColor).Render("  ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▆") + style.Foreground(WhitePieceColor).Background(WhitePieceColor).Render(" ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▌\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▁") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▘▄▄▝") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▁") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ▌\x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▝") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▄▄▄▄▄▄") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("╴") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▊\x1b[0m"),
	}
}

func WhiteBishopDarkSquare12(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("     ▅▄") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("  ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("  ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("│\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("   ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▗") + style.Foreground(WhitePieceColor).Background(WhitePieceColor).Render("  ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("  ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("│\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("   ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▌") + style.Foreground(WhitePieceColor).Background(WhitePieceColor).Render("  ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▝") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▅▖") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("│\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("   ▝") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render(" ") + style.Foreground(WhitePieceColor).Background(WhitePieceColor).Render("   ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▘  ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("│\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▁▅▄") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("━━") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▄▅▁") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" │\x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▅▅▅▅▅▅▅▅") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▌\x1b[0m"),
	}
}
func WhiteKnightDarkSquare8(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" ▄▂▃▃▂") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▌") + style.Foreground(WhitePieceColor).Background(WhitePieceColor).Render("    ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▎") + style.Foreground(WhitePieceColor).Background(WhitePieceColor).Render("   ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▇") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▃▃▃▃") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▘ ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
	}
}

func WhiteKnightDarkSquare10(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▗▂▁▂▂▂ ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▊") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▆") + style.Foreground(WhitePieceColor).Background(WhitePieceColor).Render("     ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▊") + style.Foreground(WhitePieceColor).Background(WhitePieceColor).Render("    ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▂▄▆") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▌") + style.Foreground(WhitePieceColor).Background(WhitePieceColor).Render("    ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▝") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▙▄▄▄▄▄") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
	}
}

func WhiteKnightDarkSquare12(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("  ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▂ ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▁▁▁") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("  ▊\x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("╺▄") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render(" ") + style.Foreground(WhitePieceColor).Background(WhitePieceColor).Render(" ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render(" ") + style.Foreground(WhitePieceColor).Background(WhitePieceColor).Render("   ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▖") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▎") + style.Foreground(WhitePieceColor).Background(WhitePieceColor).Render("      ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render(" ▟") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▊") + style.Foreground(WhitePieceColor).Background(WhitePieceColor).Render("     ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("╺▆") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▊") + style.Foreground(WhitePieceColor).Background(WhitePieceColor).Render("     ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("   ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▅▅▅▅▅▅") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▘  ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
	}
}
func WhiteRookDarkSquare8(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▎") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▗▄▗▖▄▖ \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▎") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▌▁▁▁▁") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▌ \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▎") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▎") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▇▇▇") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▎ ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▄▄▄▄") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" \x1b[0m"),
	}
}

func WhiteRookDarkSquare10(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("│") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" ▃▖▃▃▗▃  \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("│") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▊") + style.Foreground(WhitePieceColor).Background(WhitePieceColor).Render("▎") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▝  ▘") + style.Foreground(WhitePieceColor).Background(WhitePieceColor).Render("▊") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▎ \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("│") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▅━━━━▅") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("  \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("│") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▊") + style.Foreground(WhitePieceColor).Background(WhitePieceColor).Render("    ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▎  \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▎ ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▝") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▅▅▅▅") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▘ ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
	}
}

func WhiteRookDarkSquare12(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("│") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" ▂▂ ▂▂ ▂▂") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("│") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▎") + style.Foreground(WhitePieceColor).Background(WhitePieceColor).Render(" ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("┗") + style.Foreground(WhitePieceColor).Background(WhitePieceColor).Render("▎▊") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("┛") + style.Foreground(WhitePieceColor).Background(WhitePieceColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▊  \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("│") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▎▁▁▁▁▁▁") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▊") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("│") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▌") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▆▆▆▆▌") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("  \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("│") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▎") + style.Foreground(WhitePieceColor).Background(WhitePieceColor).Render("    ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("  \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▎") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▅▅▅▅▅▅") + style.Background(DarkSquareColor).Background(WhitePieceColor).Render(" ") + style.Foreground(WhitePieceColor).Background(WhitePieceColor).Render(" ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render(" \x1b[0m"),
	}
}
func WhiteKingDarkSquare8(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▎") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("  ▄▄") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("  \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▃▄") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▜▛") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▄▄") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▊") + style.Foreground(WhitePieceColor).Background(WhitePieceColor).Render("   ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render(" ") + style.Foreground(WhitePieceColor).Background(WhitePieceColor).Render("  ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▎\x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▎") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▙▃▃▃▃▟") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
	}
}

func WhiteKingDarkSquare10(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("│") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("   ▄▄") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("   \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("│ ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▁ ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▄▄") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" ▁") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▘") + style.Foreground(WhitePieceColor).Background(WhitePieceColor).Render("▇ ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▙▟") + style.Foreground(WhitePieceColor).Background(WhitePieceColor).Render("  ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▝") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▖") + style.Foreground(WhitePieceColor).Background(WhitePieceColor).Render("      ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▎") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▇▄▄▄▄▄▄▇") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" \x1b[0m"),
	}
}

func WhiteKingDarkSquare12(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("│") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("    ▃▃     \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("│") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("  ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▊▁▁") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▎") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("  ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("  \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▎") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▗▅▆▅▖▗▅▆▅▖") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▎") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▎") + style.Foreground(WhitePieceColor).Background(WhitePieceColor).Render("   ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("  ") + style.Foreground(WhitePieceColor).Background(WhitePieceColor).Render("   ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▊") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▎") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▌") + style.Foreground(WhitePieceColor).Background(WhitePieceColor).Render("        ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▟") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▎ ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▅▅▅▅▅▅▅▅") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" \x1b[0m"),
	}
}
func WhiteQueenDarkSquare8(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("      ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▝") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▃") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("  ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▃") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▘") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▎") + style.Foreground(WhitePieceColor).Background(WhitePieceColor).Render("    ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▊ \x1b[0m"),
		style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" ▝") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▃▃▃▃") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▘") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
	}
}

func WhiteQueenDarkSquare10(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("  ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("      ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ▊\x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▎") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▖ ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▘▝") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" ▗▊") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▌") + style.Foreground(WhitePieceColor).Background(WhitePieceColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▇") + style.Foreground(WhitePieceColor).Background(WhitePieceColor).Render(" ▇") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▇") + style.Foreground(WhitePieceColor).Background(WhitePieceColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▌") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▊") + style.Foreground(WhitePieceColor).Background(WhitePieceColor).Render("      ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▎") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("  ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▄▄▄▄▄▄") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" \x1b[0m"),
	}
}

func WhiteQueenDarkSquare12(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("           ") + style.Foreground(DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(DarkSquareColor).Background(DarkSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render("▗▁") + style.Foreground(DarkSquareColor).Render("  ") + style.Background(WhitePieceColor).Render("▘▝") + style.Background(DarkSquareColor).Render("  ") + style.Foreground(WhitePieceColor).Render("▁▖") + style.Foreground(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▊ ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▄") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▘") + style.Foreground(WhitePieceColor).Render("  ") + style.Foreground(DarkSquareColor).Render("▝") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▄") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render(" ") + style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render("▎") + style.Foreground(DarkSquareColor).Render(" \x1b[0m"),
		style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Background(WhitePieceColor).Render("▊") + style.Foreground(WhitePieceColor).Render("        ") + style.Background(DarkSquareColor).Render("▎") + style.Foreground(DarkSquareColor).Render("▊\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Render(" ") + style.Background(WhitePieceColor).Render(" ") + style.Foreground(WhitePieceColor).Render("      ") + style.Foreground(DarkSquareColor).Render(" ") + style.Background(DarkSquareColor).Render(" ▊\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(DarkSquareColor).Render(" ") + style.Foreground(DarkSquareColor).Render(" ") + style.Background(WhitePieceColor).Render("▅▅▅▅▅▅▅▅") + style.Background(DarkSquareColor).Render("  \x1b[0m"),
	}
}

func WhitePawnLightSquare8(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("   ▂▂   \x1b[0m"),
		style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("  ▝") + style.Foreground(LightSquareColor).Background(WhitePieceColor).Render("  ") + style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("▘  \x1b[0m"),
		style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("  ▗▄▄▖  \x1b[0m"),
		style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("  ") + style.Foreground(LightSquareColor).Background(WhitePieceColor).Render("▄▄▄") + style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("▘ ") + style.Foreground(LightSquareColor).Render(" \x1b[0m"),
	}
}

func WhitePawnLightSquare10(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("   ") + style.Foreground(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render("▁▁") + style.Foreground(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render("  ") + style.Foreground(LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("   ") + style.Foreground(LightSquareColor).Background(WhitePieceColor).Render("▎") + style.Foreground(WhitePieceColor).Render("  ") + style.Background(LightSquareColor).Render("▊  ") + style.Foreground(LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("   ▝") + style.Foreground(LightSquareColor).Background(WhitePieceColor).Render("▄▄") + style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("▘  ") + style.Foreground(LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("   ▟▇▇▙") + style.Foreground(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render(" ") + style.Foreground(LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("   ") + style.Foreground(LightSquareColor).Background(WhitePieceColor).Render("▅▅▅▅") + style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("  ") + style.Foreground(LightSquareColor).Render(" \x1b[0m"),
	}
}

func WhitePawnLightSquare12(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("    ") + style.Foreground(LightSquareColor).Render("    ") + style.Foreground(WhitePieceColor).Render("   ") + style.Foreground(LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("   ") + style.Foreground(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render("▟") + style.Background(WhitePieceColor).Render("  ") + style.Background(LightSquareColor).Render("▙") + style.Foreground(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render("  ") + style.Foreground(LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("   ") + style.Foreground(LightSquareColor).Render(" ") + style.Background(WhitePieceColor).Render("▁") + style.Foreground(WhitePieceColor).Render("  ") + style.Foreground(LightSquareColor).Render("▁") + style.Background(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render("  ") + style.Foreground(LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("   ") + style.Foreground(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render("▂▂▂▂") + style.Foreground(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render("  ") + style.Foreground(LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("   ▗") + style.Background(WhitePieceColor).Render(" ⎽⎽ ") + style.Background(LightSquareColor).Render("▖  ") + style.Foreground(LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("   ") + style.Foreground(LightSquareColor).Background(WhitePieceColor).Render("▆▆▆▆▆▆") + style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("  ") + style.Foreground(LightSquareColor).Render(" \x1b[0m"),
	}
}
func WhiteBishopLightSquare8(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("   ┏▌  ") + style.Foreground(LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("  ▗") + style.Foreground(LightSquareColor).Background(WhitePieceColor).Render(" ━") + style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("▖ ") + style.Foreground(LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Render(" ") + style.Background(WhitePieceColor).Render("▙▁▂▟") + style.Foreground(WhitePieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render("━") + style.Foreground(LightSquareColor).Background(WhitePieceColor).Render("▃▃▃▃") + style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("━") + style.Foreground(LightSquareColor).Render(" \x1b[0m"),
	}
}

func WhiteBishopLightSquare10(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("   ") + style.Foreground(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render("▅▅") + style.Foreground(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render("  ") + style.Foreground(LightSquareColor).Render("▌\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("  ") + style.Foreground(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render("▄") + style.Foreground(LightSquareColor).Background(WhitePieceColor).Render(" ▂") + style.Foreground(WhitePieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render(" ") + style.Foreground(LightSquareColor).Render("▌\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("  ") + style.Foreground(LightSquareColor).Background(WhitePieceColor).Render("▊") + style.Foreground(WhitePieceColor).Render("  ") + style.Background(LightSquareColor).Render("▆") + style.Background(WhitePieceColor).Render(" ") + style.Foreground(LightSquareColor).Render("") + style.Background(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render(" ") + style.Foreground(LightSquareColor).Render("▌\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render("▁") + style.Foreground(LightSquareColor).Background(WhitePieceColor).Render("▘▄▄▝") + style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("▁") + style.Foreground(LightSquareColor).Render(" ▌\x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render("▝") + style.Foreground(LightSquareColor).Background(WhitePieceColor).Render("▄▄▄▄▄▄") + style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("╴") + style.Foreground(LightSquareColor).Render("▊\x1b[0m"),
	}
}

func WhiteBishopLightSquare12(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("     ▅▄") + style.Foreground(LightSquareColor).Render("  ") + style.Foreground(WhitePieceColor).Render("  ") + style.Foreground(LightSquareColor).Render("│\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("   ") + style.Foreground(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render("▗") + style.Background(WhitePieceColor).Render("  ") + style.Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render("  ") + style.Foreground(LightSquareColor).Render("│\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("   ") + style.Foreground(LightSquareColor).Background(WhitePieceColor).Render("▌") + style.Foreground(WhitePieceColor).Render("  ") + style.Foreground(LightSquareColor).Render("▝") + style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("▅▖") + style.Foreground(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render(" ") + style.Foreground(LightSquareColor).Render("│\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("   ▝") + style.Foreground(LightSquareColor).Background(WhitePieceColor).Render(" ") + style.Foreground(WhitePieceColor).Render("   ") + style.Background(LightSquareColor).Render("▘  ") + style.Foreground(LightSquareColor).Render("│\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render("▁▅▄") + style.Foreground(LightSquareColor).Background(WhitePieceColor).Render("━━") + style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("▄▅▁") + style.Foreground(LightSquareColor).Render(" │\x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render(" ") + style.Foreground(LightSquareColor).Background(WhitePieceColor).Render("▅▅▅▅▅▅▅▅") + style.Foreground(WhitePieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Render("▌\x1b[0m"),
	}
}
func WhiteKnightLightSquare8(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render(" ▄▂▃▃▂") + style.Foreground(LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Background(WhitePieceColor).Render("▌") + style.Foreground(WhitePieceColor).Render("    ") + style.Foreground(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Background(WhitePieceColor).Render("▎") + style.Foreground(WhitePieceColor).Render("   ") + style.Foreground(LightSquareColor).Render("▇") + style.Background(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render(" \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Background(WhitePieceColor).Render("▃▃▃▃") + style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("▘ ") + style.Foreground(LightSquareColor).Render(" \x1b[0m"),
	}
}

func WhiteKnightLightSquare10(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(WhitePieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render("▗▂▁▂▂▂ ") + style.Foreground(LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Background(WhitePieceColor).Render("▊") + style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("▆") + style.Background(WhitePieceColor).Render("     ") + style.Foreground(LightSquareColor).Render(" ") + style.Background(LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(WhitePieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(WhitePieceColor).Render("▊") + style.Foreground(WhitePieceColor).Render("    ") + style.Foreground(LightSquareColor).Render("▂▄▆") + style.Background(LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Background(WhitePieceColor).Render("▌") + style.Foreground(WhitePieceColor).Render("    ") + style.Foreground(LightSquareColor).Render("▝") + style.Background(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render(" ") + style.Foreground(LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Background(WhitePieceColor).Render("▙▄▄▄▄▄") + style.Background(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render(" ") + style.Foreground(LightSquareColor).Render(" \x1b[0m"),
	}
}

func WhiteKnightLightSquare12(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("  ") + style.Foreground(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render("▂ ") + style.Foreground(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render("▁▁▁") + style.Foreground(LightSquareColor).Render("  ▊\x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render("╺▄") + style.Foreground(LightSquareColor).Background(WhitePieceColor).Render(" ") + style.Foreground(WhitePieceColor).Render(" ") + style.Foreground(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render("   ") + style.Foreground(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("▖") + style.Foreground(LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(WhitePieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Render(" ") + style.Background(WhitePieceColor).Render("▎") + style.Foreground(WhitePieceColor).Render("      ") + style.Foreground(LightSquareColor).Render(" ▟") + style.Background(LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(WhitePieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(WhitePieceColor).Render("▊") + style.Foreground(WhitePieceColor).Render("     ") + style.Foreground(LightSquareColor).Render("╺▆") + style.Background(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render(" ") + style.Foreground(LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(WhitePieceColor).Render("▊") + style.Foreground(WhitePieceColor).Render("     ") + style.Foreground(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("   ") + style.Foreground(LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render(" ") + style.Foreground(LightSquareColor).Background(WhitePieceColor).Render("▅▅▅▅▅▅") + style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("▘  ") + style.Foreground(LightSquareColor).Render(" \x1b[0m"),
	}
}
func WhiteRookLightSquare8(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▎") + style.Foreground(WhitePieceColor).Render("▗▄▗▖▄▖ \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▎") + style.Background(WhitePieceColor).Render("▌▁▁▁▁") + style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("▌ \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▎") + style.Foreground(WhitePieceColor).Render(" ") + style.Foreground(LightSquareColor).Background(WhitePieceColor).Render("▎") + style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("▇▇▇") + style.Foreground(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render(" \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▎ ") + style.Background(WhitePieceColor).Render("▄▄▄▄") + style.Background(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render(" \x1b[0m"),
	}
}

func WhiteRookLightSquare10(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("│") + style.Foreground(WhitePieceColor).Render(" ▃▖▃▃▗▃  \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("│") + style.Background(WhitePieceColor).Render("▊") + style.Foreground(WhitePieceColor).Render("▎") + style.Foreground(LightSquareColor).Render("▝  ▘") + style.Foreground(WhitePieceColor).Render("▊") + style.Background(LightSquareColor).Render("▎ \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("│") + style.Foreground(WhitePieceColor).Render(" ") + style.Foreground(LightSquareColor).Background(WhitePieceColor).Render("▅━━━━▅") + style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("  \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("│") + style.Foreground(WhitePieceColor).Render(" ") + style.Foreground(LightSquareColor).Background(WhitePieceColor).Render("▊") + style.Foreground(WhitePieceColor).Render("    ") + style.Background(LightSquareColor).Render("▎  \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▎ ") + style.Foreground(WhitePieceColor).Render("▝") + style.Foreground(LightSquareColor).Background(WhitePieceColor).Render("▅▅▅▅") + style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("▘ ") + style.Foreground(LightSquareColor).Render(" \x1b[0m"),
	}
}

func WhiteRookLightSquare12(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("│") + style.Foreground(WhitePieceColor).Render(" ▂▂ ▂▂ ▂▂") + style.Foreground(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render(" \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("│") + style.Foreground(WhitePieceColor).Render(" ") + style.Foreground(LightSquareColor).Background(WhitePieceColor).Render("▎") + style.Foreground(WhitePieceColor).Render(" ") + style.Foreground(LightSquareColor).Render("┗") + style.Foreground(WhitePieceColor).Render("▎▊") + style.Foreground(LightSquareColor).Render("┛") + style.Foreground(WhitePieceColor).Render(" ") + style.Background(LightSquareColor).Render("▊  \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("│") + style.Foreground(WhitePieceColor).Render(" ") + style.Foreground(LightSquareColor).Background(WhitePieceColor).Render("▎▁▁▁▁▁▁") + style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("▊") + style.Foreground(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render(" \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("│") + style.Foreground(WhitePieceColor).Render(" ") + style.Foreground(LightSquareColor).Render(" ") + style.Background(WhitePieceColor).Render("▌") + style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("▆▆▆▆▌") + style.Foreground(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render("  \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("│") + style.Foreground(WhitePieceColor).Render(" ") + style.Foreground(LightSquareColor).Render(" ") + style.Background(WhitePieceColor).Render("▎") + style.Foreground(WhitePieceColor).Render("    ") + style.Foreground(LightSquareColor).Render(" ") + style.Background(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render("  \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▎") + style.Foreground(WhitePieceColor).Render(" ") + style.Foreground(LightSquareColor).Render(" ") + style.Background(WhitePieceColor).Render("▅▅▅▅▅▅") + style.Background(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render(" ") + style.Foreground(LightSquareColor).Render(" \x1b[0m"),
	}
}
func WhiteKingLightSquare8(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▎") + style.Foreground(WhitePieceColor).Render("  ▄▄") + style.Foreground(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render("  \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render("▃▄") + style.Foreground(LightSquareColor).Background(WhitePieceColor).Render("▜▛") + style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("▄▄") + style.Foreground(LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(LightSquareColor).Background(WhitePieceColor).Render("▊") + style.Foreground(WhitePieceColor).Render("   ") + style.Foreground(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render("  ") + style.Background(LightSquareColor).Render("▎\x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▎") + style.Background(WhitePieceColor).Render("▙▃▃▃▃▟") + style.Background(LightSquareColor).Render(" \x1b[0m"),
	}
}

func WhiteKingLightSquare10(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("│") + style.Foreground(WhitePieceColor).Render("   ▄▄") + style.Foreground(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render("   \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("│ ") + style.Foreground(WhitePieceColor).Render("▁ ") + style.Foreground(LightSquareColor).Background(WhitePieceColor).Render("▄▄") + style.Foreground(WhitePieceColor).Background(LightSquareColor).Render(" ▁") + style.Foreground(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render(" \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Background(WhitePieceColor).Render("▘") + style.Foreground(WhitePieceColor).Render("▇ ") + style.Background(LightSquareColor).Render("▙▟") + style.Background(WhitePieceColor).Render("  ") + style.Foreground(LightSquareColor).Render("▝") + style.Background(LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Background(WhitePieceColor).Render("▖") + style.Foreground(WhitePieceColor).Render("      ") + style.Foreground(LightSquareColor).Render(" ") + style.Background(LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▎") + style.Background(WhitePieceColor).Render("▇▄▄▄▄▄▄▇") + style.Foreground(WhitePieceColor).Background(LightSquareColor).Render(" \x1b[0m"),
	}
}

func WhiteKingLightSquare12(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("│") + style.Foreground(WhitePieceColor).Render("    ▃▃     \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("│") + style.Foreground(WhitePieceColor).Render("  ") + style.Foreground(LightSquareColor).Render(" ") + style.Background(WhitePieceColor).Render("▊▁▁") + style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("▎") + style.Foreground(LightSquareColor).Render("  ") + style.Foreground(WhitePieceColor).Render("  \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▎") + style.Foreground(WhitePieceColor).Render("▗▅▆▅▖▗▅▆▅▖") + style.Foreground(LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▎") + style.Background(WhitePieceColor).Render("▎") + style.Foreground(WhitePieceColor).Render("   ") + style.Foreground(LightSquareColor).Render("  ") + style.Foreground(WhitePieceColor).Render("   ") + style.Background(LightSquareColor).Render("▊") + style.Foreground(LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▎") + style.Background(WhitePieceColor).Render("▌") + style.Foreground(WhitePieceColor).Render("        ") + style.Foreground(LightSquareColor).Render("▟") + style.Background(LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("▎ ") + style.Background(WhitePieceColor).Render("▅▅▅▅▅▅▅▅") + style.Background(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render(" \x1b[0m"),
	}
}
func WhiteQueenLightSquare8(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render("      ") + style.Foreground(LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Background(WhitePieceColor).Render("▝") + style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("▃") + style.Foreground(LightSquareColor).Background(WhitePieceColor).Render("  ") + style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("▃") + style.Foreground(LightSquareColor).Background(WhitePieceColor).Render("▘") + style.Background(LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(WhitePieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(WhitePieceColor).Render("▎") + style.Foreground(WhitePieceColor).Render("    ") + style.Background(LightSquareColor).Render("▊ \x1b[0m"),
		style.Foreground(WhitePieceColor).Background(LightSquareColor).Render(" ▝") + style.Foreground(LightSquareColor).Background(WhitePieceColor).Render("▃▃▃▃") + style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("▘") + style.Foreground(LightSquareColor).Render(" \x1b[0m"),
	}
}

func WhiteQueenLightSquare10(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {
	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render("  ") + style.Foreground(WhitePieceColor).Render("      ") + style.Foreground(LightSquareColor).Render(" ▊\x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Background(WhitePieceColor).Render("▎") + style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("▖ ") + style.Foreground(LightSquareColor).Background(WhitePieceColor).Render("▘▝") + style.Foreground(WhitePieceColor).Background(LightSquareColor).Render(" ▗▊") + style.Foreground(LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(WhitePieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(WhitePieceColor).Render("▌") + style.Foreground(WhitePieceColor).Render(" ") + style.Background(LightSquareColor).Render("▇") + style.Background(WhitePieceColor).Render(" ▇") + style.Background(LightSquareColor).Render("▇") + style.Background(WhitePieceColor).Render(" ") + style.Background(LightSquareColor).Render("▌") + style.Foreground(LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(WhitePieceColor).Render("▊") + style.Foreground(WhitePieceColor).Render("      ") + style.Background(LightSquareColor).Render("▎") + style.Foreground(LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("  ") + style.Foreground(LightSquareColor).Background(WhitePieceColor).Render("▄▄▄▄▄▄") + style.Foreground(WhitePieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Render(" \x1b[0m"),
	}
}

func WhiteQueenLightSquare12(WhitePieceColor lipgloss.Color, BlackPieceColor lipgloss.Color, LightSquareColor lipgloss.Color, DarkSquareColor lipgloss.Color) []string {

	style := lipgloss.NewStyle()

	return []string{
		style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("           ") + style.Foreground(LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(LightSquareColor).Background(LightSquareColor).Render(" ") + style.Foreground(WhitePieceColor).Render("▗▁") + style.Foreground(LightSquareColor).Render("  ") + style.Background(WhitePieceColor).Render("▘▝") + style.Background(LightSquareColor).Render("  ") + style.Foreground(WhitePieceColor).Render("▁▖") + style.Foreground(LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(WhitePieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(WhitePieceColor).Render("▊ ") + style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("▄") + style.Foreground(LightSquareColor).Background(WhitePieceColor).Render("▘") + style.Foreground(WhitePieceColor).Render("  ") + style.Foreground(LightSquareColor).Render("▝") + style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("▄") + style.Foreground(LightSquareColor).Background(WhitePieceColor).Render(" ") + style.Foreground(WhitePieceColor).Background(LightSquareColor).Render("▎") + style.Foreground(LightSquareColor).Render(" \x1b[0m"),
		style.Foreground(WhitePieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Background(WhitePieceColor).Render("▊") + style.Foreground(WhitePieceColor).Render("        ") + style.Background(LightSquareColor).Render("▎") + style.Foreground(LightSquareColor).Render("▊\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Render(" ") + style.Background(WhitePieceColor).Render(" ") + style.Foreground(WhitePieceColor).Render("      ") + style.Foreground(LightSquareColor).Render(" ") + style.Background(LightSquareColor).Render(" ▊\x1b[0m"),
		style.Foreground(WhitePieceColor).Background(LightSquareColor).Render(" ") + style.Foreground(LightSquareColor).Render(" ") + style.Background(WhitePieceColor).Render("▅▅▅▅▅▅▅▅") + style.Background(LightSquareColor).Render("  \x1b[0m"),
	}
}
