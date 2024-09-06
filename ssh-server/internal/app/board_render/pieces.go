package board_render

func GetPieceRendering(pieceColor bool, squareColor bool, size int, pieceType string) []string {
	colorIndex := map[bool]int{true: 0, false: 1}[pieceColor]   // true for white, false for black
	squareIndex := map[bool]int{true: 0, false: 1}[squareColor] // true for light square, false for dark square
	sizeIndex := size - 8                                       // sizes start from 8, so we subtract 8 to get the correct index

	return ColourSquareSizeMap[colorIndex][squareIndex][sizeIndex][pieceType]
}

func GetEmptySquareRendering(squareColor bool, size int) []string {

	if squareColor {
		switch size {
		case 8:
			return EmptyLightSquare8
		case 9:
			return EmptyLightSquare9
		case 10:
			return EmptyLightSquare10
		case 11:
			return EmptyLightSquare11
		case 12:
			return EmptyLightSquare12
		}
	}
	switch size {
	case 8:
		return EmptyDarkSquare8
	case 9:
		return EmptyDarkSquare9
	case 10:
		return EmptyDarkSquare10
	case 11:
		return EmptyDarkSquare11
	case 12:
		return EmptyDarkSquare12
	}

	return EmptyDarkSquare12
}

var (
	ColourSquareSizeMap = [2][2][5]map[string][]string{
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
				{ // Size 9
					"p": WhitePawnLightSquare9,
					"n": WhiteKnightLightSquare9,
					"b": WhiteBishopLightSquare9,
					"r": WhiteRookLightSquare9,
					"q": WhiteQueenLightSquare9,
					"k": WhiteKingLightSquare9,
				},
				{ // Size 10
					"p": WhitePawnLightSquare10,
					"n": WhiteKnightLightSquare10,
					"b": WhiteBishopLightSquare10,
					"r": WhiteRookLightSquare10,
					"q": WhiteQueenLightSquare10,
					"k": WhiteKingLightSquare10,
				},
				{ // Size 11
					"p": WhitePawnLightSquare11,
					"n": WhiteKnightLightSquare11,
					"b": WhiteBishopLightSquare11,
					"r": WhiteRookLightSquare11,
					"q": WhiteQueenLightSquare11,
					"k": WhiteKingLightSquare11,
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
				{ // Size 9
					"p": WhitePawnDarkSquare9,
					"n": WhiteKnightDarkSquare9,
					"b": WhiteBishopDarkSquare9,
					"r": WhiteRookDarkSquare9,
					"q": WhiteQueenDarkSquare9,
					"k": WhiteKingDarkSquare9,
				},
				{ // Size 10
					"p": WhitePawnDarkSquare10,
					"n": WhiteKnightDarkSquare10,
					"b": WhiteBishopDarkSquare10,
					"r": WhiteRookDarkSquare10,
					"q": WhiteQueenDarkSquare10,
					"k": WhiteKingDarkSquare10,
				},
				{ // Size 11
					"p": WhitePawnDarkSquare11,
					"n": WhiteKnightDarkSquare11,
					"b": WhiteBishopDarkSquare11,
					"r": WhiteRookDarkSquare11,
					"q": WhiteQueenDarkSquare11,
					"k": WhiteKingDarkSquare11,
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
				{ // Size 9
					"p": BlackPawnLightSquare9,
					"n": BlackKnightLightSquare9,
					"b": BlackBishopLightSquare9,
					"r": BlackRookLightSquare9,
					"q": BlackQueenLightSquare9,
					"k": BlackKingLightSquare9,
				},
				{ // Size 10
					"p": BlackPawnLightSquare10,
					"n": BlackKnightLightSquare10,
					"b": BlackBishopLightSquare10,
					"r": BlackRookLightSquare10,
					"q": BlackQueenLightSquare10,
					"k": BlackKingLightSquare10,
				},
				{ // Size 11
					"p": BlackPawnLightSquare11,
					"n": BlackKnightLightSquare11,
					"b": BlackBishopLightSquare11,
					"r": BlackRookLightSquare11,
					"q": BlackQueenLightSquare11,
					"k": BlackKingLightSquare11,
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
				{ // Size 9
					"p": BlackPawnDarkSquare9,
					"n": BlackKnightDarkSquare9,
					"b": BlackBishopDarkSquare9,
					"r": BlackRookDarkSquare9,
					"q": BlackQueenDarkSquare9,
					"k": BlackKingDarkSquare9,
				},
				{ // Size 10
					"p": BlackPawnDarkSquare10,
					"n": BlackKnightDarkSquare10,
					"b": BlackBishopDarkSquare10,
					"r": BlackRookDarkSquare10,
					"q": BlackQueenDarkSquare10,
					"k": BlackKingDarkSquare10,
				},
				{ // Size 11
					"p": BlackPawnDarkSquare11,
					"n": BlackKnightDarkSquare11,
					"b": BlackBishopDarkSquare11,
					"r": BlackRookDarkSquare11,
					"q": BlackQueenDarkSquare11,
					"k": BlackKingDarkSquare11,
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

const (
	BlackPieceForegroundColourEscape = "\x1b[38;2;0;5;15m"
	BlackPieceBackgroundColourEscape = "\x1b[48;2;0;5;15m"

	WhitePieceForegroundColourEscape = "\x1b[38;2;255;255;255m"
	WhitePieceBackgroundColourEscape = "\x1b[48;2;255;255;255m"

	LightSquareForegroundColourEscape = "\x1b[38;2;189;175;141m"
	LightSquareBackgroundColourEscape = "\x1b[48;2;189;175;141m"

	DarkSquareForegroundColourEscape = "\x1b[38;2;41;30;2m"
	DarkSquareBackgroundColourEscape = "\x1b[48;2;41;30;2m"
)

var (
	EmptyDarkSquare8 = []string{
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "████████\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "████████\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "████████\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "████████\x1b[0m",
	}

	EmptyDarkSquare9 = []string{
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "█████████\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "█████████\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "█████████\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "█████████\x1b[0m",
	}

	EmptyDarkSquare10 = []string{
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "██████████\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "██████████\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "██████████\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "██████████\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "██████████\x1b[0m",
	}

	EmptyDarkSquare11 = []string{
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "███████████\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "███████████\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "███████████\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "███████████\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "███████████\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "███████████\x1b[0m",
	}

	EmptyDarkSquare12 = []string{
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "████████████\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "████████████\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "████████████\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "████████████\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "████████████\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "████████████\x1b[0m",
	}

	EmptyLightSquare8 = []string{
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "████████\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "████████\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "████████\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "████████\x1b[0m",
	}

	EmptyLightSquare9 = []string{
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "█████████\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "█████████\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "█████████\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "█████████\x1b[0m",
	}

	EmptyLightSquare10 = []string{
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "██████████\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "██████████\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "██████████\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "██████████\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "██████████\x1b[0m",
	}

	EmptyLightSquare11 = []string{
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "███████████\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "███████████\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "███████████\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "███████████\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "███████████\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "███████████\x1b[0m",
	}

	EmptyLightSquare12 = []string{
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "████████████\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "████████████\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "████████████\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "████████████\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "████████████\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "████████████\x1b[0m",
	}
)

var (
	BlackPawnLightSquare8 = []string{
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "   ▂▂   \x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  ▝" + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "  " + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▘  \x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  ▗▄▄▖  \x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  " + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▄▄▄" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▘ " + LightSquareForegroundColourEscape + " \x1b[0m",
	}
	BlackPawnLightSquare9 = []string{
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  " + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "▁▃▁" + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  " + LightSquareForegroundColourEscape + " " + BlackPieceBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + " " + LightSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "   ▄▄▄  " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  ▝" + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▄▄▄" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▘ " + LightSquareForegroundColourEscape + " \x1b[0m",
	}
	BlackPawnLightSquare10 = []string{
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "   " + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "▁▁" + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "  " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "   " + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + "  " + LightSquareBackgroundColourEscape + "▊  " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "   ▝" + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▄▄" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▘  " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "   ▟▇▇▙" + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "   " + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▅▅▅▅" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  " + LightSquareForegroundColourEscape + " \x1b[0m",
	}
	BlackPawnLightSquare11 = []string{
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "    " + LightSquareForegroundColourEscape + " ▁ " + BlackPieceForegroundColourEscape + "   " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "   ▗▇" + BlackPieceBackgroundColourEscape + " " + LightSquareBackgroundColourEscape + "▇▖  " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "   ▝" + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + "▁" + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▘  " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "    ▂▂▂   " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  " + LightSquareForegroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▘" + BlackPieceForegroundColourEscape + "⎽⎽⎽" + LightSquareForegroundColourEscape + "▝" + LightSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  " + LightSquareForegroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▆▆▆▆▆" + LightSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + " \x1b[0m",
	}
	BlackPawnLightSquare12 = []string{
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "    " + LightSquareForegroundColourEscape + "    " + BlackPieceForegroundColourEscape + "   " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "   " + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "▟" + BlackPieceBackgroundColourEscape + "  " + LightSquareBackgroundColourEscape + "▙" + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "  " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "   " + LightSquareForegroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▁" + BlackPieceForegroundColourEscape + "  " + LightSquareForegroundColourEscape + "▁" + LightSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + "  " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "   " + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "▂▂▂▂" + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "  " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "   ▗" + BlackPieceBackgroundColourEscape + " ⎽⎽ " + LightSquareBackgroundColourEscape + "▖  " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "   " + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▆▆▆▆▆▆" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  " + LightSquareForegroundColourEscape + " \x1b[0m",
	}
	BlackBishopLightSquare8 = []string{
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "   ┏▌  " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  ▗" + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + " ━" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▖ " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▙▁▂▟" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + "━" + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▃▃▃▃" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "━" + LightSquareForegroundColourEscape + " \x1b[0m",
	}
	BlackBishopLightSquare9 = []string{
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "   " + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▊" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▆   " + LightSquareForegroundColourEscape + "▌\x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  ▗" + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + " ▗┛" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▖ " + LightSquareForegroundColourEscape + "▌\x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  ▝" + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▂▁▂▇" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + "▌\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + "━" + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▃▃▃▃▃" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "━" + LightSquareForegroundColourEscape + "▌\x1b[0m",
	}
	BlackBishopLightSquare10 = []string{
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "   " + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "▅▅" + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "  " + LightSquareForegroundColourEscape + "▌\x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  " + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "▄" + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + " ▂" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + "▌\x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  " + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▊" + BlackPieceForegroundColourEscape + "  " + LightSquareBackgroundColourEscape + "▆" + BlackPieceBackgroundColourEscape + " " + LightSquareForegroundColourEscape + "" + LightSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + "▌\x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "▁" + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▘▄▄▝" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▁" + LightSquareForegroundColourEscape + " ▌\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + "▝" + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▄▄▄▄▄▄" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "╴" + LightSquareForegroundColourEscape + "▊\x1b[0m",
	}
	BlackBishopLightSquare11 = []string{
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "    ▗▅▖   " + LightSquareForegroundColourEscape + "▌\x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  " + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▘" + BlackPieceForegroundColourEscape + " " + LightSquareBackgroundColourEscape + "▌" + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "  " + LightSquareForegroundColourEscape + "▌\x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  " + LightSquareForegroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▘" + BlackPieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + "╺" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▄▊" + LightSquareForegroundColourEscape + "╴" + BlackPieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + "▌\x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  " + LightSquareForegroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▖" + BlackPieceForegroundColourEscape + "   " + LightSquareBackgroundColourEscape + "▌" + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + "▌\x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "▂▅▄" + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "━" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▄▅▂" + LightSquareForegroundColourEscape + " ▌\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + "▝" + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▅▅▅▅▅▅▅" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▘" + LightSquareForegroundColourEscape + "▊\x1b[0m",
	}
	BlackBishopLightSquare12 = []string{
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "     ▅▄" + LightSquareForegroundColourEscape + "  " + BlackPieceForegroundColourEscape + "  " + LightSquareForegroundColourEscape + "│\x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "   " + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "▗" + BlackPieceBackgroundColourEscape + "  " + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "  " + LightSquareForegroundColourEscape + "│\x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "   " + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▌" + BlackPieceForegroundColourEscape + "  " + LightSquareForegroundColourEscape + "▝" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▅▖" + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + "│\x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "   ▝" + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + "   " + LightSquareBackgroundColourEscape + "▘  " + LightSquareForegroundColourEscape + "│\x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "▁▅▄" + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "━━" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▄▅▁" + LightSquareForegroundColourEscape + " │\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▅▅▅▅▅▅▅▅" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + "▌\x1b[0m",
	}
	BlackKnightLightSquare8 = []string{
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + " ▄▂▃▃▂" + LightSquareForegroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▌" + BlackPieceForegroundColourEscape + "    " + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + "   " + LightSquareForegroundColourEscape + "▇" + LightSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▃▃▃▃" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▘ " + LightSquareForegroundColourEscape + " \x1b[0m",
	}
	BlackKnightLightSquare9 = []string{
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  ▗▃▃▃▃▁" + LightSquareForegroundColourEscape + "▊\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▌" + BlackPieceForegroundColourEscape + "     " + LightSquareForegroundColourEscape + "▁" + LightSquareBackgroundColourEscape + " \x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + "   " + LightSquareForegroundColourEscape + " ▇" + LightSquareBackgroundColourEscape + " ▊\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + "┗" + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▃▃▃▃" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  " + LightSquareForegroundColourEscape + " \x1b[0m",
	}
	BlackKnightLightSquare10 = []string{
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "▗▂▁▂▂▂ " + LightSquareForegroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▊" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▆" + BlackPieceBackgroundColourEscape + "     " + LightSquareForegroundColourEscape + " " + LightSquareBackgroundColourEscape + " \x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▊" + BlackPieceForegroundColourEscape + "    " + LightSquareForegroundColourEscape + "▂▄▆" + LightSquareBackgroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▌" + BlackPieceForegroundColourEscape + "    " + LightSquareForegroundColourEscape + "▝" + LightSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▙▄▄▄▄▄" + LightSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + " \x1b[0m",
	}
	BlackKnightLightSquare11 = []string{
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + " ▂▁▂▂▂▁" + LightSquareForegroundColourEscape + " ▊\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▊" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▆" + BlackPieceBackgroundColourEscape + "      " + LightSquareForegroundColourEscape + "▝" + LightSquareBackgroundColourEscape + " \x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▊ " + BlackPieceForegroundColourEscape + "    " + LightSquareForegroundColourEscape + "▂▄▆" + LightSquareBackgroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▌" + BlackPieceForegroundColourEscape + "     " + LightSquareBackgroundColourEscape + "▎  " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + "▝" + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▄▄▄▄▄" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▘" + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + " \x1b[0m",
	}
	BlackKnightLightSquare12 = []string{
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  " + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "▂ " + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "▁▁▁" + LightSquareForegroundColourEscape + "  ▊\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + "╺▄" + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "   " + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▖" + LightSquareForegroundColourEscape + " \x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + "      " + LightSquareForegroundColourEscape + " ▟" + LightSquareBackgroundColourEscape + " \x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▊" + BlackPieceForegroundColourEscape + "     " + LightSquareForegroundColourEscape + "╺▆" + LightSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▊" + BlackPieceForegroundColourEscape + "     " + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "   " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▅▅▅▅▅▅" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▘  " + LightSquareForegroundColourEscape + " \x1b[0m",
	}
	BlackRookLightSquare8 = []string{
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + "▗▄▗▖▄▖ \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎" + BlackPieceBackgroundColourEscape + "▌▁▁▁▁" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▌ \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▇▇▇" + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎ " + BlackPieceBackgroundColourEscape + "▄▄▄▄" + LightSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + " \x1b[0m",
	}
	BlackRookLightSquare9 = []string{
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎" + BlackPieceBackgroundColourEscape + "▊" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▄" + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▊" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▄" + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▊" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▄  \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎" + BlackPieceBackgroundColourEscape + "▊▁▁▁▁▁" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▎ \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▌" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▇▇▇▌  \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎ " + BlackPieceBackgroundColourEscape + "▄▄▄▄▄" + LightSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + " \x1b[0m",
	}
	BlackRookLightSquare10 = []string{
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "│" + BlackPieceForegroundColourEscape + " ▃▖▃▃▗▃  \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "│" + BlackPieceBackgroundColourEscape + "▊" + BlackPieceForegroundColourEscape + "▎" + LightSquareForegroundColourEscape + "▝  ▘" + BlackPieceForegroundColourEscape + "▊" + LightSquareBackgroundColourEscape + "▎ \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "│" + BlackPieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▅━━━━▅" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "│" + BlackPieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▊" + BlackPieceForegroundColourEscape + "    " + LightSquareBackgroundColourEscape + "▎  \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎ " + BlackPieceForegroundColourEscape + "▝" + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▅▅▅▅" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▘ " + LightSquareForegroundColourEscape + " \x1b[0m",
	}
	BlackRookLightSquare11 = []string{
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "│ " + BlackPieceForegroundColourEscape + "▃▃ ▃ ▃▃" + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "│▊" + BlackPieceForegroundColourEscape + "" + BlackPieceBackgroundColourEscape + "▎ " + LightSquareBackgroundColourEscape + "▅" + BlackPieceBackgroundColourEscape + " " + LightSquareBackgroundColourEscape + "▅" + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + "▊" + LightSquareForegroundColourEscape + "" + LightSquareBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "│" + BlackPieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▅━━━━━▅" + LightSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "│" + BlackPieceForegroundColourEscape + "  " + BlackPieceBackgroundColourEscape + "     " + LightSquareBackgroundColourEscape + "   \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + "  " + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▅▅▅▅▅" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  " + LightSquareForegroundColourEscape + " \x1b[0m",
	}
	BlackRookLightSquare12 = []string{
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "│" + BlackPieceForegroundColourEscape + " ▂▂ ▂▂ ▂▂" + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "│" + BlackPieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + "┗" + BlackPieceForegroundColourEscape + "▎▊" + LightSquareForegroundColourEscape + "┛" + BlackPieceForegroundColourEscape + " " + LightSquareBackgroundColourEscape + "▊  \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "│" + BlackPieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▎▁▁▁▁▁▁" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▊" + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "│" + BlackPieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▌" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▆▆▆▆▌" + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "  \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "│" + BlackPieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + "    " + LightSquareForegroundColourEscape + " " + LightSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + "  \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▅▅▅▅▅▅" + LightSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + " \x1b[0m",
	}
	BlackKingLightSquare8 = []string{
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + "  ▄▄" + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "  \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + "▃▄" + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▜▛" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▄▄" + LightSquareForegroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▊" + BlackPieceForegroundColourEscape + "   " + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "  " + LightSquareBackgroundColourEscape + "▎\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎" + BlackPieceBackgroundColourEscape + "▙▃▃▃▃▟" + LightSquareBackgroundColourEscape + " \x1b[0m",
	}
	BlackKingLightSquare9 = []string{
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "│" + BlackPieceForegroundColourEscape + "  ▗▅▖   \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + "▂▄▃" + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▆" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▃▄▂" + LightSquareForegroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▊" + BlackPieceForegroundColourEscape + "   " + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "   " + LightSquareBackgroundColourEscape + "▎\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎" + BlackPieceBackgroundColourEscape + "▙▃▃▃▃▃▟" + LightSquareBackgroundColourEscape + " \x1b[0m",
	}
	BlackKingLightSquare10 = []string{
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "│" + BlackPieceForegroundColourEscape + "   ▄▄" + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "   \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "│ " + BlackPieceForegroundColourEscape + "▁ " + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▄▄" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + " ▁" + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▘" + BlackPieceForegroundColourEscape + "▇ " + LightSquareBackgroundColourEscape + "▙▟" + BlackPieceBackgroundColourEscape + "  " + LightSquareForegroundColourEscape + "▝" + LightSquareBackgroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▖" + BlackPieceForegroundColourEscape + "      " + LightSquareForegroundColourEscape + " " + LightSquareBackgroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎" + BlackPieceBackgroundColourEscape + "▇▄▄▄▄▄▄▇" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + " \x1b[0m",
	}
	BlackKingLightSquare11 = []string{
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "│" + BlackPieceForegroundColourEscape + "   ▗▄▖" + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "   \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "│ " + BlackPieceForegroundColourEscape + "▁▁" + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▙▃▟" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▁▁" + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + "▟" + BlackPieceBackgroundColourEscape + "  " + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▄▇" + BlackPieceBackgroundColourEscape + "  " + LightSquareForegroundColourEscape + "▝" + LightSquareBackgroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎" + BlackPieceBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + "       " + LightSquareForegroundColourEscape + "▗" + LightSquareBackgroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▄▄▄▄▄▄▄" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  \x1b[0m",
	}
	BlackKingLightSquare12 = []string{
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "│" + BlackPieceForegroundColourEscape + "    ▃▃     \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "│" + BlackPieceForegroundColourEscape + "  " + LightSquareForegroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▊▁▁" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▎" + LightSquareForegroundColourEscape + "  " + BlackPieceForegroundColourEscape + "  \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + "▗▅▆▅▖▗▅▆▅▖" + LightSquareForegroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎" + BlackPieceBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + "   " + LightSquareForegroundColourEscape + "  " + BlackPieceForegroundColourEscape + "   " + LightSquareBackgroundColourEscape + "▊" + LightSquareForegroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎" + BlackPieceBackgroundColourEscape + "▌" + BlackPieceForegroundColourEscape + "        " + LightSquareForegroundColourEscape + "▟" + LightSquareBackgroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎ " + BlackPieceBackgroundColourEscape + "▅▅▅▅▅▅▅▅" + LightSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + " \x1b[0m",
	}
	BlackQueenLightSquare8 = []string{
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + "      " + LightSquareForegroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▝" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▃" + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "  " + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▃" + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▘" + LightSquareBackgroundColourEscape + " \x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + "    " + LightSquareBackgroundColourEscape + "▊ \x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + " ▝" + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▃▃▃▃" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▘" + LightSquareForegroundColourEscape + " \x1b[0m",
	}
	BlackQueenLightSquare9 = []string{
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "  " + BlackPieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "▁" + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + "  \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▄▟" + BlackPieceBackgroundColourEscape + " " + LightSquareBackgroundColourEscape + "▙▄▊" + LightSquareForegroundColourEscape + " \x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▌" + BlackPieceForegroundColourEscape + "     " + LightSquareBackgroundColourEscape + "▌" + LightSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + " ▝" + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▃▃▃▃▃" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▘" + LightSquareForegroundColourEscape + " \x1b[0m",
	}
	BlackQueenLightSquare10 = []string{
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "  " + BlackPieceForegroundColourEscape + "      " + LightSquareForegroundColourEscape + " ▊\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▖ " + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▘▝" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + " ▗▊" + LightSquareForegroundColourEscape + " \x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▌" + BlackPieceForegroundColourEscape + " " + LightSquareBackgroundColourEscape + "▇" + BlackPieceBackgroundColourEscape + " ▇" + LightSquareBackgroundColourEscape + "▇" + BlackPieceBackgroundColourEscape + " " + LightSquareBackgroundColourEscape + "▌" + LightSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▊" + BlackPieceForegroundColourEscape + "      " + LightSquareBackgroundColourEscape + "▎" + LightSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  " + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▄▄▄▄▄▄" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + " \x1b[0m",
	}
	BlackQueenLightSquare11 = []string{
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "  " + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "  " + LightSquareForegroundColourEscape + "  \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + "▜▂" + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "▗" + BlackPieceBackgroundColourEscape + " " + LightSquareBackgroundColourEscape + "▖" + LightSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "▂▛" + LightSquareForegroundColourEscape + " \x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▊" + BlackPieceForegroundColourEscape + " " + LightSquareBackgroundColourEscape + "▇" + BlackPieceBackgroundColourEscape + "   " + LightSquareBackgroundColourEscape + "▇" + BlackPieceBackgroundColourEscape + " " + LightSquareBackgroundColourEscape + "▌" + LightSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▊" + BlackPieceForegroundColourEscape + "       " + LightSquareBackgroundColourEscape + "▎" + LightSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▄▄▄▄▄▄▄" + LightSquareBackgroundColourEscape + "  \x1b[0m",
	}
	BlackQueenLightSquare12 = []string{
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "           " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + "▗▁" + LightSquareForegroundColourEscape + "  " + BlackPieceBackgroundColourEscape + "▘▝" + LightSquareBackgroundColourEscape + "  " + BlackPieceForegroundColourEscape + "▁▖" + LightSquareForegroundColourEscape + " \x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▊ " + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▄" + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▘" + BlackPieceForegroundColourEscape + "  " + LightSquareForegroundColourEscape + "▝" + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▄" + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▎" + LightSquareForegroundColourEscape + " \x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▊" + BlackPieceForegroundColourEscape + "        " + LightSquareBackgroundColourEscape + "▎" + LightSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + " " + BlackPieceBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + "      " + LightSquareForegroundColourEscape + " " + LightSquareBackgroundColourEscape + " ▊\x1b[0m",
		BlackPieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▅▅▅▅▅▅▅▅" + LightSquareBackgroundColourEscape + "  \x1b[0m",
	}
)

var (
	BlackPawnDarkSquare8 = []string{
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "   ▂▂   \x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  ▝" + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "  " + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▘  \x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  ▗▄▄▖  \x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  " + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▄▄▄" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▘ " + DarkSquareForegroundColourEscape + " \x1b[0m",
	}
	BlackPawnDarkSquare9 = []string{
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  " + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "▁▃▁" + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  " + DarkSquareForegroundColourEscape + " " + BlackPieceBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + " " + DarkSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "   ▄▄▄  " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  ▝" + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▄▄▄" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▘ " + DarkSquareForegroundColourEscape + " \x1b[0m",
	}
	BlackPawnDarkSquare10 = []string{
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "   " + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "▁▁" + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "  " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "   " + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + "  " + DarkSquareBackgroundColourEscape + "▊  " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "   ▝" + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▄▄" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▘  " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "   ▟▇▇▙" + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "   " + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▅▅▅▅" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  " + DarkSquareForegroundColourEscape + " \x1b[0m",
	}
	BlackPawnDarkSquare11 = []string{
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "    " + DarkSquareForegroundColourEscape + " ▁ " + BlackPieceForegroundColourEscape + "   " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "   ▗▇" + BlackPieceBackgroundColourEscape + " " + DarkSquareBackgroundColourEscape + "▇▖  " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "   ▝" + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + "▁" + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▘  " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "    ▂▂▂   " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  " + DarkSquareForegroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▘" + BlackPieceForegroundColourEscape + "⎽⎽⎽" + DarkSquareForegroundColourEscape + "▝" + DarkSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  " + DarkSquareForegroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▆▆▆▆▆" + DarkSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + " \x1b[0m",
	}
	BlackPawnDarkSquare12 = []string{
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "    " + DarkSquareForegroundColourEscape + "    " + BlackPieceForegroundColourEscape + "   " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "   " + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "▟" + BlackPieceBackgroundColourEscape + "  " + DarkSquareBackgroundColourEscape + "▙" + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "  " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "   " + DarkSquareForegroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▁" + BlackPieceForegroundColourEscape + "  " + DarkSquareForegroundColourEscape + "▁" + DarkSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + "  " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "   " + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "▂▂▂▂" + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "  " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "   ▗" + BlackPieceBackgroundColourEscape + " ⎽⎽ " + DarkSquareBackgroundColourEscape + "▖  " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "   " + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▆▆▆▆▆▆" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  " + DarkSquareForegroundColourEscape + " \x1b[0m",
	}
	BlackBishopDarkSquare8 = []string{
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "   ┏▌  " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  ▗" + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + " ━" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▖ " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▙▁▂▟" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + "━" + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▃▃▃▃" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "━" + DarkSquareForegroundColourEscape + " \x1b[0m",
	}
	BlackBishopDarkSquare9 = []string{
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "   " + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▊" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▆   " + DarkSquareForegroundColourEscape + "▌\x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  ▗" + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + " ▗┛" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▖ " + DarkSquareForegroundColourEscape + "▌\x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  ▝" + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▂▁▂▇" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + "▌\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + "━" + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▃▃▃▃▃" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "━" + DarkSquareForegroundColourEscape + "▌\x1b[0m",
	}
	BlackBishopDarkSquare10 = []string{
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "   " + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "▅▅" + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "  " + DarkSquareForegroundColourEscape + "▌\x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  " + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "▄" + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + " ▂" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + "▌\x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  " + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▊" + BlackPieceForegroundColourEscape + "  " + DarkSquareBackgroundColourEscape + "▆" + BlackPieceBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + "" + DarkSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + "▌\x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "▁" + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▘▄▄▝" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▁" + DarkSquareForegroundColourEscape + " ▌\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + "▝" + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▄▄▄▄▄▄" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "╴" + DarkSquareForegroundColourEscape + "▊\x1b[0m",
	}
	BlackBishopDarkSquare11 = []string{
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "    ▗▅▖   " + DarkSquareForegroundColourEscape + "▌\x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  " + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▘" + BlackPieceForegroundColourEscape + " " + DarkSquareBackgroundColourEscape + "▌" + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "  " + DarkSquareForegroundColourEscape + "▌\x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  " + DarkSquareForegroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▘" + BlackPieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + "╺" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▄▊" + DarkSquareForegroundColourEscape + "╴" + BlackPieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + "▌\x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  " + DarkSquareForegroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▖" + BlackPieceForegroundColourEscape + "   " + DarkSquareBackgroundColourEscape + "▌" + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + "▌\x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "▂▅▄" + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "━" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▄▅▂" + DarkSquareForegroundColourEscape + " ▌\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + "▝" + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▅▅▅▅▅▅▅" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▘" + DarkSquareForegroundColourEscape + "▊\x1b[0m",
	}
	BlackBishopDarkSquare12 = []string{
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "     ▅▄" + DarkSquareForegroundColourEscape + "  " + BlackPieceForegroundColourEscape + "  " + DarkSquareForegroundColourEscape + "│\x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "   " + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "▗" + BlackPieceBackgroundColourEscape + "  " + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "  " + DarkSquareForegroundColourEscape + "│\x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "   " + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▌" + BlackPieceForegroundColourEscape + "  " + DarkSquareForegroundColourEscape + "▝" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▅▖" + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + "│\x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "   ▝" + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + "   " + DarkSquareBackgroundColourEscape + "▘  " + DarkSquareForegroundColourEscape + "│\x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "▁▅▄" + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "━━" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▄▅▁" + DarkSquareForegroundColourEscape + " │\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▅▅▅▅▅▅▅▅" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + "▌\x1b[0m",
	}
	BlackKnightDarkSquare8 = []string{
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + " ▄▂▃▃▂" + DarkSquareForegroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▌" + BlackPieceForegroundColourEscape + "    " + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + "   " + DarkSquareForegroundColourEscape + "▇" + DarkSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▃▃▃▃" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▘ " + DarkSquareForegroundColourEscape + " \x1b[0m",
	}
	BlackKnightDarkSquare9 = []string{
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  ▗▃▃▃▃▁" + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▌" + BlackPieceForegroundColourEscape + "     " + DarkSquareForegroundColourEscape + "▁" + DarkSquareBackgroundColourEscape + " \x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + "   " + DarkSquareForegroundColourEscape + " ▇" + DarkSquareBackgroundColourEscape + " ▊\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + "┗" + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▃▃▃▃" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  " + DarkSquareForegroundColourEscape + " \x1b[0m",
	}
	BlackKnightDarkSquare10 = []string{
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "▗▂▁▂▂▂ " + DarkSquareForegroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▊" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▆" + BlackPieceBackgroundColourEscape + "     " + DarkSquareForegroundColourEscape + " " + DarkSquareBackgroundColourEscape + " \x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▊" + BlackPieceForegroundColourEscape + "    " + DarkSquareForegroundColourEscape + "▂▄▆" + DarkSquareBackgroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▌" + BlackPieceForegroundColourEscape + "    " + DarkSquareForegroundColourEscape + "▝" + DarkSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▙▄▄▄▄▄" + DarkSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + " \x1b[0m",
	}
	BlackKnightDarkSquare11 = []string{
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + " ▂▁▂▂▂▁" + DarkSquareForegroundColourEscape + " ▊\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▊" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▆" + BlackPieceBackgroundColourEscape + "      " + DarkSquareForegroundColourEscape + "▝" + DarkSquareBackgroundColourEscape + " \x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▊ " + BlackPieceForegroundColourEscape + "    " + DarkSquareForegroundColourEscape + "▂▄▆" + DarkSquareBackgroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▌" + BlackPieceForegroundColourEscape + "     " + DarkSquareBackgroundColourEscape + "▎  " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + "▝" + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▄▄▄▄▄" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▘" + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + " \x1b[0m",
	}
	BlackKnightDarkSquare12 = []string{
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  " + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "▂ " + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "▁▁▁" + DarkSquareForegroundColourEscape + "  ▊\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + "╺▄" + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "   " + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▖" + DarkSquareForegroundColourEscape + " \x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + "      " + DarkSquareForegroundColourEscape + " ▟" + DarkSquareBackgroundColourEscape + " \x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▊" + BlackPieceForegroundColourEscape + "     " + DarkSquareForegroundColourEscape + "╺▆" + DarkSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▊" + BlackPieceForegroundColourEscape + "     " + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "   " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▅▅▅▅▅▅" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▘  " + DarkSquareForegroundColourEscape + " \x1b[0m",
	}
	BlackRookDarkSquare8 = []string{
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + "▗▄▗▖▄▖ \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎" + BlackPieceBackgroundColourEscape + "▌▁▁▁▁" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▌ \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▇▇▇" + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎ " + BlackPieceBackgroundColourEscape + "▄▄▄▄" + DarkSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + " \x1b[0m",
	}
	BlackRookDarkSquare9 = []string{
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎" + BlackPieceBackgroundColourEscape + "▊" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▄" + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▊" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▄" + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▊" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▄  \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎" + BlackPieceBackgroundColourEscape + "▊▁▁▁▁▁" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎ \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▌" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▇▇▇▌  \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎ " + BlackPieceBackgroundColourEscape + "▄▄▄▄▄" + DarkSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + " \x1b[0m",
	}
	BlackRookDarkSquare10 = []string{
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "│" + BlackPieceForegroundColourEscape + " ▃▖▃▃▗▃  \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "│" + BlackPieceBackgroundColourEscape + "▊" + BlackPieceForegroundColourEscape + "▎" + DarkSquareForegroundColourEscape + "▝  ▘" + BlackPieceForegroundColourEscape + "▊" + DarkSquareBackgroundColourEscape + "▎ \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "│" + BlackPieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▅━━━━▅" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "│" + BlackPieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▊" + BlackPieceForegroundColourEscape + "    " + DarkSquareBackgroundColourEscape + "▎  \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎ " + BlackPieceForegroundColourEscape + "▝" + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▅▅▅▅" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▘ " + DarkSquareForegroundColourEscape + " \x1b[0m",
	}
	BlackRookDarkSquare11 = []string{
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "│ " + BlackPieceForegroundColourEscape + "▃▃ ▃ ▃▃" + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "│▊" + BlackPieceForegroundColourEscape + "" + BlackPieceBackgroundColourEscape + "▎ " + DarkSquareBackgroundColourEscape + "▅" + BlackPieceBackgroundColourEscape + " " + DarkSquareBackgroundColourEscape + "▅" + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + "▊" + DarkSquareForegroundColourEscape + "" + DarkSquareBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "│" + BlackPieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▅━━━━━▅" + DarkSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "│" + BlackPieceForegroundColourEscape + "  " + BlackPieceBackgroundColourEscape + "     " + DarkSquareBackgroundColourEscape + "   \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + "  " + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▅▅▅▅▅" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  " + DarkSquareForegroundColourEscape + " \x1b[0m",
	}
	BlackRookDarkSquare12 = []string{
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "│" + BlackPieceForegroundColourEscape + " ▂▂ ▂▂ ▂▂" + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "│" + BlackPieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + "┗" + BlackPieceForegroundColourEscape + "▎▊" + DarkSquareForegroundColourEscape + "┛" + BlackPieceForegroundColourEscape + " " + DarkSquareBackgroundColourEscape + "▊  \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "│" + BlackPieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▎▁▁▁▁▁▁" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▊" + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "│" + BlackPieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▌" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▆▆▆▆▌" + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "  \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "│" + BlackPieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + "    " + DarkSquareForegroundColourEscape + " " + DarkSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + "  \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▅▅▅▅▅▅" + DarkSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + " \x1b[0m",
	}
	BlackKingDarkSquare8 = []string{
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + "  ▄▄" + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "  \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + "▃▄" + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▜▛" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▄▄" + DarkSquareForegroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▊" + BlackPieceForegroundColourEscape + "   " + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "  " + DarkSquareBackgroundColourEscape + "▎\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎" + BlackPieceBackgroundColourEscape + "▙▃▃▃▃▟" + DarkSquareBackgroundColourEscape + " \x1b[0m",
	}
	BlackKingDarkSquare9 = []string{
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "│" + BlackPieceForegroundColourEscape + "  ▗▅▖   \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + "▂▄▃" + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▆" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▃▄▂" + DarkSquareForegroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▊" + BlackPieceForegroundColourEscape + "   " + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "   " + DarkSquareBackgroundColourEscape + "▎\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎" + BlackPieceBackgroundColourEscape + "▙▃▃▃▃▃▟" + DarkSquareBackgroundColourEscape + " \x1b[0m",
	}
	BlackKingDarkSquare10 = []string{
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "│" + BlackPieceForegroundColourEscape + "   ▄▄" + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "   \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "│ " + BlackPieceForegroundColourEscape + "▁ " + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▄▄" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " ▁" + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▘" + BlackPieceForegroundColourEscape + "▇ " + DarkSquareBackgroundColourEscape + "▙▟" + BlackPieceBackgroundColourEscape + "  " + DarkSquareForegroundColourEscape + "▝" + DarkSquareBackgroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▖" + BlackPieceForegroundColourEscape + "      " + DarkSquareForegroundColourEscape + " " + DarkSquareBackgroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎" + BlackPieceBackgroundColourEscape + "▇▄▄▄▄▄▄▇" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " \x1b[0m",
	}
	BlackKingDarkSquare11 = []string{
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "│" + BlackPieceForegroundColourEscape + "   ▗▄▖" + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "   \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "│ " + BlackPieceForegroundColourEscape + "▁▁" + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▙▃▟" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▁▁" + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + "▟" + BlackPieceBackgroundColourEscape + "  " + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▄▇" + BlackPieceBackgroundColourEscape + "  " + DarkSquareForegroundColourEscape + "▝" + DarkSquareBackgroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎" + BlackPieceBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + "       " + DarkSquareForegroundColourEscape + "▗" + DarkSquareBackgroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▄▄▄▄▄▄▄" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  \x1b[0m",
	}
	BlackKingDarkSquare12 = []string{
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "│" + BlackPieceForegroundColourEscape + "    ▃▃     \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "│" + BlackPieceForegroundColourEscape + "  " + DarkSquareForegroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▊▁▁" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎" + DarkSquareForegroundColourEscape + "  " + BlackPieceForegroundColourEscape + "  \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + "▗▅▆▅▖▗▅▆▅▖" + DarkSquareForegroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎" + BlackPieceBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + "   " + DarkSquareForegroundColourEscape + "  " + BlackPieceForegroundColourEscape + "   " + DarkSquareBackgroundColourEscape + "▊" + DarkSquareForegroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎" + BlackPieceBackgroundColourEscape + "▌" + BlackPieceForegroundColourEscape + "        " + DarkSquareForegroundColourEscape + "▟" + DarkSquareBackgroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎ " + BlackPieceBackgroundColourEscape + "▅▅▅▅▅▅▅▅" + DarkSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + " \x1b[0m",
	}
	BlackQueenDarkSquare8 = []string{
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + "      " + DarkSquareForegroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▝" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▃" + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "  " + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▃" + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▘" + DarkSquareBackgroundColourEscape + " \x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + "    " + DarkSquareBackgroundColourEscape + "▊ \x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " ▝" + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▃▃▃▃" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▘" + DarkSquareForegroundColourEscape + " \x1b[0m",
	}
	BlackQueenDarkSquare9 = []string{
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "  " + BlackPieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "▁" + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + "  \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▄▟" + BlackPieceBackgroundColourEscape + " " + DarkSquareBackgroundColourEscape + "▙▄▊" + DarkSquareForegroundColourEscape + " \x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▌" + BlackPieceForegroundColourEscape + "     " + DarkSquareBackgroundColourEscape + "▌" + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " ▝" + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▃▃▃▃▃" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▘" + DarkSquareForegroundColourEscape + " \x1b[0m",
	}
	BlackQueenDarkSquare10 = []string{
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "  " + BlackPieceForegroundColourEscape + "      " + DarkSquareForegroundColourEscape + " ▊\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▎" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▖ " + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▘▝" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " ▗▊" + DarkSquareForegroundColourEscape + " \x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▌" + BlackPieceForegroundColourEscape + " " + DarkSquareBackgroundColourEscape + "▇" + BlackPieceBackgroundColourEscape + " ▇" + DarkSquareBackgroundColourEscape + "▇" + BlackPieceBackgroundColourEscape + " " + DarkSquareBackgroundColourEscape + "▌" + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▊" + BlackPieceForegroundColourEscape + "      " + DarkSquareBackgroundColourEscape + "▎" + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  " + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▄▄▄▄▄▄" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + " \x1b[0m",
	}
	BlackQueenDarkSquare11 = []string{
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "  " + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "  " + DarkSquareForegroundColourEscape + "  \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + "▜▂" + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "▗" + BlackPieceBackgroundColourEscape + " " + DarkSquareBackgroundColourEscape + "▖" + DarkSquareForegroundColourEscape + " " + BlackPieceForegroundColourEscape + "▂▛" + DarkSquareForegroundColourEscape + " \x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▊" + BlackPieceForegroundColourEscape + " " + DarkSquareBackgroundColourEscape + "▇" + BlackPieceBackgroundColourEscape + "   " + DarkSquareBackgroundColourEscape + "▇" + BlackPieceBackgroundColourEscape + " " + DarkSquareBackgroundColourEscape + "▌" + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▊" + BlackPieceForegroundColourEscape + "       " + DarkSquareBackgroundColourEscape + "▎" + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▄▄▄▄▄▄▄" + DarkSquareBackgroundColourEscape + "  \x1b[0m",
	}
	BlackQueenDarkSquare12 = []string{
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "           " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + "▗▁" + DarkSquareForegroundColourEscape + "  " + BlackPieceBackgroundColourEscape + "▘▝" + DarkSquareBackgroundColourEscape + "  " + BlackPieceForegroundColourEscape + "▁▖" + DarkSquareForegroundColourEscape + " \x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▊ " + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▄" + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▘" + BlackPieceForegroundColourEscape + "  " + DarkSquareForegroundColourEscape + "▝" + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▄" + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎" + DarkSquareForegroundColourEscape + " \x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + BlackPieceBackgroundColourEscape + "▊" + BlackPieceForegroundColourEscape + "        " + DarkSquareBackgroundColourEscape + "▎" + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + " " + BlackPieceBackgroundColourEscape + " " + BlackPieceForegroundColourEscape + "      " + DarkSquareForegroundColourEscape + " " + DarkSquareBackgroundColourEscape + " ▊\x1b[0m",
		BlackPieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + " " + BlackPieceBackgroundColourEscape + "▅▅▅▅▅▅▅▅" + DarkSquareBackgroundColourEscape + "  \x1b[0m",
	}
)

var (
	WhitePawnDarkSquare8 = []string{
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "   ▂▂   \x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  ▝" + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "  " + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▘  \x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  ▗▄▄▖  \x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  " + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▄▄▄" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▘ " + DarkSquareForegroundColourEscape + " \x1b[0m",
	}
	WhitePawnDarkSquare9 = []string{
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  " + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "▁▃▁" + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  " + DarkSquareForegroundColourEscape + " " + WhitePieceBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + " " + DarkSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "   ▄▄▄  " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  ▝" + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▄▄▄" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▘ " + DarkSquareForegroundColourEscape + " \x1b[0m",
	}
	WhitePawnDarkSquare10 = []string{
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "   " + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "▁▁" + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "  " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "   " + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + "  " + DarkSquareBackgroundColourEscape + "▊  " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "   ▝" + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▄▄" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▘  " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "   ▟▇▇▙" + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "   " + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▅▅▅▅" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  " + DarkSquareForegroundColourEscape + " \x1b[0m",
	}
	WhitePawnDarkSquare11 = []string{
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "    " + DarkSquareForegroundColourEscape + " ▁ " + WhitePieceForegroundColourEscape + "   " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "   ▗▇" + WhitePieceBackgroundColourEscape + " " + DarkSquareBackgroundColourEscape + "▇▖  " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "   ▝" + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + "▁" + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▘  " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "    ▂▂▂   " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  " + DarkSquareForegroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▘" + WhitePieceForegroundColourEscape + "⎽⎽⎽" + DarkSquareForegroundColourEscape + "▝" + DarkSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  " + DarkSquareForegroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▆▆▆▆▆" + DarkSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + " \x1b[0m",
	}
	WhitePawnDarkSquare12 = []string{
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "    " + DarkSquareForegroundColourEscape + "    " + WhitePieceForegroundColourEscape + "   " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "   " + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "▟" + WhitePieceBackgroundColourEscape + "  " + DarkSquareBackgroundColourEscape + "▙" + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "  " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "   " + DarkSquareForegroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▁" + WhitePieceForegroundColourEscape + "  " + DarkSquareForegroundColourEscape + "▁" + DarkSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + "  " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "   " + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "▂▂▂▂" + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "  " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "   ▗" + WhitePieceBackgroundColourEscape + " ⎽⎽ " + DarkSquareBackgroundColourEscape + "▖  " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "   " + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▆▆▆▆▆▆" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  " + DarkSquareForegroundColourEscape + " \x1b[0m",
	}
	WhiteBishopDarkSquare8 = []string{
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "   ┏▌  " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  ▗" + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + " ━" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▖ " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▙▁▂▟" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + "━" + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▃▃▃▃" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "━" + DarkSquareForegroundColourEscape + " \x1b[0m",
	}
	WhiteBishopDarkSquare9 = []string{
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "   " + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▊" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▆   " + DarkSquareForegroundColourEscape + "▌\x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  ▗" + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + " ▗┛" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▖ " + DarkSquareForegroundColourEscape + "▌\x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  ▝" + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▂▁▂▇" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + "▌\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + "━" + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▃▃▃▃▃" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "━" + DarkSquareForegroundColourEscape + "▌\x1b[0m",
	}
	WhiteBishopDarkSquare10 = []string{
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "   " + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "▅▅" + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "  " + DarkSquareForegroundColourEscape + "▌\x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  " + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "▄" + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + " ▂" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + "▌\x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  " + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▊" + WhitePieceForegroundColourEscape + "  " + DarkSquareBackgroundColourEscape + "▆" + WhitePieceBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + "" + DarkSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + "▌\x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "▁" + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▘▄▄▝" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▁" + DarkSquareForegroundColourEscape + " ▌\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + "▝" + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▄▄▄▄▄▄" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "╴" + DarkSquareForegroundColourEscape + "▊\x1b[0m",
	}
	WhiteBishopDarkSquare11 = []string{
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "    ▗▅▖   " + DarkSquareForegroundColourEscape + "▌\x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  " + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▘" + WhitePieceForegroundColourEscape + " " + DarkSquareBackgroundColourEscape + "▌" + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "  " + DarkSquareForegroundColourEscape + "▌\x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  " + DarkSquareForegroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▘" + WhitePieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + "╺" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▄▊" + DarkSquareForegroundColourEscape + "╴" + WhitePieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + "▌\x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  " + DarkSquareForegroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▖" + WhitePieceForegroundColourEscape + "   " + DarkSquareBackgroundColourEscape + "▌" + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + "▌\x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "▂▅▄" + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "━" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▄▅▂" + DarkSquareForegroundColourEscape + " ▌\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + "▝" + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▅▅▅▅▅▅▅" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▘" + DarkSquareForegroundColourEscape + "▊\x1b[0m",
	}
	WhiteBishopDarkSquare12 = []string{
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "     ▅▄" + DarkSquareForegroundColourEscape + "  " + WhitePieceForegroundColourEscape + "  " + DarkSquareForegroundColourEscape + "│\x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "   " + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "▗" + WhitePieceBackgroundColourEscape + "  " + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "  " + DarkSquareForegroundColourEscape + "│\x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "   " + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▌" + WhitePieceForegroundColourEscape + "  " + DarkSquareForegroundColourEscape + "▝" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▅▖" + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + "│\x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "   ▝" + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + "   " + DarkSquareBackgroundColourEscape + "▘  " + DarkSquareForegroundColourEscape + "│\x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "▁▅▄" + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "━━" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▄▅▁" + DarkSquareForegroundColourEscape + " │\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▅▅▅▅▅▅▅▅" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + "▌\x1b[0m",
	}
	WhiteKnightDarkSquare8 = []string{
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + " ▄▂▃▃▂" + DarkSquareForegroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▌" + WhitePieceForegroundColourEscape + "    " + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + "   " + DarkSquareForegroundColourEscape + "▇" + DarkSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▃▃▃▃" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▘ " + DarkSquareForegroundColourEscape + " \x1b[0m",
	}
	WhiteKnightDarkSquare9 = []string{
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  ▗▃▃▃▃▁" + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▌" + WhitePieceForegroundColourEscape + "     " + DarkSquareForegroundColourEscape + "▁" + DarkSquareBackgroundColourEscape + " \x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + "   " + DarkSquareForegroundColourEscape + " ▇" + DarkSquareBackgroundColourEscape + " ▊\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + "┗" + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▃▃▃▃" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  " + DarkSquareForegroundColourEscape + " \x1b[0m",
	}
	WhiteKnightDarkSquare10 = []string{
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "▗▂▁▂▂▂ " + DarkSquareForegroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▊" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▆" + WhitePieceBackgroundColourEscape + "     " + DarkSquareForegroundColourEscape + " " + DarkSquareBackgroundColourEscape + " \x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▊" + WhitePieceForegroundColourEscape + "    " + DarkSquareForegroundColourEscape + "▂▄▆" + DarkSquareBackgroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▌" + WhitePieceForegroundColourEscape + "    " + DarkSquareForegroundColourEscape + "▝" + DarkSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▙▄▄▄▄▄" + DarkSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + " \x1b[0m",
	}
	WhiteKnightDarkSquare11 = []string{
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + " ▂▁▂▂▂▁" + DarkSquareForegroundColourEscape + " ▊\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▊" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▆" + WhitePieceBackgroundColourEscape + "      " + DarkSquareForegroundColourEscape + "▝" + DarkSquareBackgroundColourEscape + " \x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▊ " + WhitePieceForegroundColourEscape + "    " + DarkSquareForegroundColourEscape + "▂▄▆" + DarkSquareBackgroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▌" + WhitePieceForegroundColourEscape + "     " + DarkSquareBackgroundColourEscape + "▎  " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + "▝" + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▄▄▄▄▄" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▘" + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + " \x1b[0m",
	}
	WhiteKnightDarkSquare12 = []string{
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  " + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "▂ " + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "▁▁▁" + DarkSquareForegroundColourEscape + "  ▊\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + "╺▄" + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "   " + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▖" + DarkSquareForegroundColourEscape + " \x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + "      " + DarkSquareForegroundColourEscape + " ▟" + DarkSquareBackgroundColourEscape + " \x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▊" + WhitePieceForegroundColourEscape + "     " + DarkSquareForegroundColourEscape + "╺▆" + DarkSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▊" + WhitePieceForegroundColourEscape + "     " + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "   " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▅▅▅▅▅▅" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▘  " + DarkSquareForegroundColourEscape + " \x1b[0m",
	}
	WhiteRookDarkSquare8 = []string{
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + "▗▄▗▖▄▖ \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎" + WhitePieceBackgroundColourEscape + "▌▁▁▁▁" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▌ \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▇▇▇" + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎ " + WhitePieceBackgroundColourEscape + "▄▄▄▄" + DarkSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + " \x1b[0m",
	}
	WhiteRookDarkSquare9 = []string{
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎" + WhitePieceBackgroundColourEscape + "▊" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▄" + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▊" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▄" + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▊" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▄  \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎" + WhitePieceBackgroundColourEscape + "▊▁▁▁▁▁" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎ \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▌" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▇▇▇▌  \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎ " + WhitePieceBackgroundColourEscape + "▄▄▄▄▄" + DarkSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + " \x1b[0m",
	}
	WhiteRookDarkSquare10 = []string{
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "│" + WhitePieceForegroundColourEscape + " ▃▖▃▃▗▃  \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "│" + WhitePieceBackgroundColourEscape + "▊" + WhitePieceForegroundColourEscape + "▎" + DarkSquareForegroundColourEscape + "▝  ▘" + WhitePieceForegroundColourEscape + "▊" + DarkSquareBackgroundColourEscape + "▎ \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "│" + WhitePieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▅━━━━▅" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "│" + WhitePieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▊" + WhitePieceForegroundColourEscape + "    " + DarkSquareBackgroundColourEscape + "▎  \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎ " + WhitePieceForegroundColourEscape + "▝" + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▅▅▅▅" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▘ " + DarkSquareForegroundColourEscape + " \x1b[0m",
	}
	WhiteRookDarkSquare11 = []string{
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "│ " + WhitePieceForegroundColourEscape + "▃▃ ▃ ▃▃" + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "│▊" + WhitePieceForegroundColourEscape + "" + WhitePieceBackgroundColourEscape + "▎ " + DarkSquareBackgroundColourEscape + "▅" + WhitePieceBackgroundColourEscape + " " + DarkSquareBackgroundColourEscape + "▅" + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + "▊" + DarkSquareForegroundColourEscape + "" + DarkSquareBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "│" + WhitePieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▅━━━━━▅" + DarkSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "│" + WhitePieceForegroundColourEscape + "  " + WhitePieceBackgroundColourEscape + "     " + DarkSquareBackgroundColourEscape + "   \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + "  " + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▅▅▅▅▅" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  " + DarkSquareForegroundColourEscape + " \x1b[0m",
	}
	WhiteRookDarkSquare12 = []string{
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "│" + WhitePieceForegroundColourEscape + " ▂▂ ▂▂ ▂▂" + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "│" + WhitePieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + "┗" + WhitePieceForegroundColourEscape + "▎▊" + DarkSquareForegroundColourEscape + "┛" + WhitePieceForegroundColourEscape + " " + DarkSquareBackgroundColourEscape + "▊  \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "│" + WhitePieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▎▁▁▁▁▁▁" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▊" + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "│" + WhitePieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▌" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▆▆▆▆▌" + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "  \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "│" + WhitePieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + "    " + DarkSquareForegroundColourEscape + " " + DarkSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + "  \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▅▅▅▅▅▅" + DarkSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + " \x1b[0m",
	}
	WhiteKingDarkSquare8 = []string{
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + "  ▄▄" + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "  \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + "▃▄" + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▜▛" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▄▄" + DarkSquareForegroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▊" + WhitePieceForegroundColourEscape + "   " + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "  " + DarkSquareBackgroundColourEscape + "▎\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎" + WhitePieceBackgroundColourEscape + "▙▃▃▃▃▟" + DarkSquareBackgroundColourEscape + " \x1b[0m",
	}
	WhiteKingDarkSquare9 = []string{
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "│" + WhitePieceForegroundColourEscape + "  ▗▅▖   \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + "▂▄▃" + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▆" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▃▄▂" + DarkSquareForegroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▊" + WhitePieceForegroundColourEscape + "   " + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "   " + DarkSquareBackgroundColourEscape + "▎\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎" + WhitePieceBackgroundColourEscape + "▙▃▃▃▃▃▟" + DarkSquareBackgroundColourEscape + " \x1b[0m",
	}
	WhiteKingDarkSquare10 = []string{
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "│" + WhitePieceForegroundColourEscape + "   ▄▄" + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "   \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "│ " + WhitePieceForegroundColourEscape + "▁ " + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▄▄" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " ▁" + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▘" + WhitePieceForegroundColourEscape + "▇ " + DarkSquareBackgroundColourEscape + "▙▟" + WhitePieceBackgroundColourEscape + "  " + DarkSquareForegroundColourEscape + "▝" + DarkSquareBackgroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▖" + WhitePieceForegroundColourEscape + "      " + DarkSquareForegroundColourEscape + " " + DarkSquareBackgroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎" + WhitePieceBackgroundColourEscape + "▇▄▄▄▄▄▄▇" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " \x1b[0m",
	}
	WhiteKingDarkSquare11 = []string{
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "│" + WhitePieceForegroundColourEscape + "   ▗▄▖" + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "   \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "│ " + WhitePieceForegroundColourEscape + "▁▁" + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▙▃▟" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▁▁" + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + "▟" + WhitePieceBackgroundColourEscape + "  " + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▄▇" + WhitePieceBackgroundColourEscape + "  " + DarkSquareForegroundColourEscape + "▝" + DarkSquareBackgroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎" + WhitePieceBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + "       " + DarkSquareForegroundColourEscape + "▗" + DarkSquareBackgroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▄▄▄▄▄▄▄" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  \x1b[0m",
	}
	WhiteKingDarkSquare12 = []string{
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "│" + WhitePieceForegroundColourEscape + "    ▃▃     \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "│" + WhitePieceForegroundColourEscape + "  " + DarkSquareForegroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▊▁▁" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎" + DarkSquareForegroundColourEscape + "  " + WhitePieceForegroundColourEscape + "  \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + "▗▅▆▅▖▗▅▆▅▖" + DarkSquareForegroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎" + WhitePieceBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + "   " + DarkSquareForegroundColourEscape + "  " + WhitePieceForegroundColourEscape + "   " + DarkSquareBackgroundColourEscape + "▊" + DarkSquareForegroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎" + WhitePieceBackgroundColourEscape + "▌" + WhitePieceForegroundColourEscape + "        " + DarkSquareForegroundColourEscape + "▟" + DarkSquareBackgroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎ " + WhitePieceBackgroundColourEscape + "▅▅▅▅▅▅▅▅" + DarkSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + " \x1b[0m",
	}
	WhiteQueenDarkSquare8 = []string{
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + "      " + DarkSquareForegroundColourEscape + " \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▝" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▃" + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "  " + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▃" + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▘" + DarkSquareBackgroundColourEscape + " \x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + "    " + DarkSquareBackgroundColourEscape + "▊ \x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " ▝" + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▃▃▃▃" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▘" + DarkSquareForegroundColourEscape + " \x1b[0m",
	}
	WhiteQueenDarkSquare9 = []string{
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "  " + WhitePieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "▁" + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + "  \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▄▟" + WhitePieceBackgroundColourEscape + " " + DarkSquareBackgroundColourEscape + "▙▄▊" + DarkSquareForegroundColourEscape + " \x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▌" + WhitePieceForegroundColourEscape + "     " + DarkSquareBackgroundColourEscape + "▌" + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " ▝" + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▃▃▃▃▃" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▘" + DarkSquareForegroundColourEscape + " \x1b[0m",
	}
	WhiteQueenDarkSquare10 = []string{
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + "  " + WhitePieceForegroundColourEscape + "      " + DarkSquareForegroundColourEscape + " ▊\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▖ " + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▘▝" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " ▗▊" + DarkSquareForegroundColourEscape + " \x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▌" + WhitePieceForegroundColourEscape + " " + DarkSquareBackgroundColourEscape + "▇" + WhitePieceBackgroundColourEscape + " ▇" + DarkSquareBackgroundColourEscape + "▇" + WhitePieceBackgroundColourEscape + " " + DarkSquareBackgroundColourEscape + "▌" + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▊" + WhitePieceForegroundColourEscape + "      " + DarkSquareBackgroundColourEscape + "▎" + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "  " + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▄▄▄▄▄▄" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + " \x1b[0m",
	}
	WhiteQueenDarkSquare11 = []string{
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "  " + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + " " + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "  " + DarkSquareForegroundColourEscape + "  \x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + "▜▂" + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "▗" + WhitePieceBackgroundColourEscape + " " + DarkSquareBackgroundColourEscape + "▖" + DarkSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "▂▛" + DarkSquareForegroundColourEscape + " \x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▊" + WhitePieceForegroundColourEscape + " " + DarkSquareBackgroundColourEscape + "▇" + WhitePieceBackgroundColourEscape + "   " + DarkSquareBackgroundColourEscape + "▇" + WhitePieceBackgroundColourEscape + " " + DarkSquareBackgroundColourEscape + "▌" + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▊" + WhitePieceForegroundColourEscape + "       " + DarkSquareBackgroundColourEscape + "▎" + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▄▄▄▄▄▄▄" + DarkSquareBackgroundColourEscape + "  \x1b[0m",
	}
	WhiteQueenDarkSquare12 = []string{
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "           " + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		DarkSquareForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + "▗▁" + DarkSquareForegroundColourEscape + "  " + WhitePieceBackgroundColourEscape + "▘▝" + DarkSquareBackgroundColourEscape + "  " + WhitePieceForegroundColourEscape + "▁▖" + DarkSquareForegroundColourEscape + " \x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▊ " + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▄" + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▘" + WhitePieceForegroundColourEscape + "  " + DarkSquareForegroundColourEscape + "▝" + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▄" + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + "▎" + DarkSquareForegroundColourEscape + " \x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▊" + WhitePieceForegroundColourEscape + "        " + DarkSquareBackgroundColourEscape + "▎" + DarkSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + " " + WhitePieceBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + "      " + DarkSquareForegroundColourEscape + " " + DarkSquareBackgroundColourEscape + " ▊\x1b[0m",
		WhitePieceForegroundColourEscape + DarkSquareBackgroundColourEscape + " " + DarkSquareForegroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▅▅▅▅▅▅▅▅" + DarkSquareBackgroundColourEscape + "  \x1b[0m",
	}
)

var (
	WhitePawnLightSquare8 = []string{
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "   ▂▂   \x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  ▝" + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "  " + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▘  \x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  ▗▄▄▖  \x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  " + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▄▄▄" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▘ " + LightSquareForegroundColourEscape + " \x1b[0m",
	}
	WhitePawnLightSquare9 = []string{
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  " + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "▁▃▁" + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  " + LightSquareForegroundColourEscape + " " + WhitePieceBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + " " + LightSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "   ▄▄▄  " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  ▝" + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▄▄▄" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▘ " + LightSquareForegroundColourEscape + " \x1b[0m",
	}
	WhitePawnLightSquare10 = []string{
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "   " + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "▁▁" + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "  " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "   " + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + "  " + LightSquareBackgroundColourEscape + "▊  " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "   ▝" + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▄▄" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▘  " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "   ▟▇▇▙" + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "   " + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▅▅▅▅" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  " + LightSquareForegroundColourEscape + " \x1b[0m",
	}
	WhitePawnLightSquare11 = []string{
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "    " + LightSquareForegroundColourEscape + " ▁ " + WhitePieceForegroundColourEscape + "   " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "   ▗▇" + WhitePieceBackgroundColourEscape + " " + LightSquareBackgroundColourEscape + "▇▖  " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "   ▝" + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + "▁" + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▘  " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "    ▂▂▂   " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  " + LightSquareForegroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▘" + WhitePieceForegroundColourEscape + "⎽⎽⎽" + LightSquareForegroundColourEscape + "▝" + LightSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  " + LightSquareForegroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▆▆▆▆▆" + LightSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + " \x1b[0m",
	}
	WhitePawnLightSquare12 = []string{
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "    " + LightSquareForegroundColourEscape + "    " + WhitePieceForegroundColourEscape + "   " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "   " + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "▟" + WhitePieceBackgroundColourEscape + "  " + LightSquareBackgroundColourEscape + "▙" + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "  " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "   " + LightSquareForegroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▁" + WhitePieceForegroundColourEscape + "  " + LightSquareForegroundColourEscape + "▁" + LightSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + "  " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "   " + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "▂▂▂▂" + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "  " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "   ▗" + WhitePieceBackgroundColourEscape + " ⎽⎽ " + LightSquareBackgroundColourEscape + "▖  " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "   " + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▆▆▆▆▆▆" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  " + LightSquareForegroundColourEscape + " \x1b[0m",
	}
	WhiteBishopLightSquare8 = []string{
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "   ┏▌  " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  ▗" + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + " ━" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▖ " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▙▁▂▟" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + "━" + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▃▃▃▃" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "━" + LightSquareForegroundColourEscape + " \x1b[0m",
	}
	WhiteBishopLightSquare9 = []string{
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "   " + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▊" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▆   " + LightSquareForegroundColourEscape + "▌\x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  ▗" + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + " ▗┛" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▖ " + LightSquareForegroundColourEscape + "▌\x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  ▝" + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▂▁▂▇" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + "▌\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + "━" + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▃▃▃▃▃" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "━" + LightSquareForegroundColourEscape + "▌\x1b[0m",
	}
	WhiteBishopLightSquare10 = []string{
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "   " + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "▅▅" + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "  " + LightSquareForegroundColourEscape + "▌\x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  " + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "▄" + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + " ▂" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + "▌\x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  " + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▊" + WhitePieceForegroundColourEscape + "  " + LightSquareBackgroundColourEscape + "▆" + WhitePieceBackgroundColourEscape + " " + LightSquareForegroundColourEscape + "" + LightSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + "▌\x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "▁" + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▘▄▄▝" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▁" + LightSquareForegroundColourEscape + " ▌\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + "▝" + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▄▄▄▄▄▄" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "╴" + LightSquareForegroundColourEscape + "▊\x1b[0m",
	}
	WhiteBishopLightSquare11 = []string{
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "    ▗▅▖   " + LightSquareForegroundColourEscape + "▌\x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  " + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▘" + WhitePieceForegroundColourEscape + " " + LightSquareBackgroundColourEscape + "▌" + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "  " + LightSquareForegroundColourEscape + "▌\x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  " + LightSquareForegroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▘" + WhitePieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + "╺" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▄▊" + LightSquareForegroundColourEscape + "╴" + WhitePieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + "▌\x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  " + LightSquareForegroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▖" + WhitePieceForegroundColourEscape + "   " + LightSquareBackgroundColourEscape + "▌" + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + "▌\x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "▂▅▄" + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "━" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▄▅▂" + LightSquareForegroundColourEscape + " ▌\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + "▝" + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▅▅▅▅▅▅▅" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▘" + LightSquareForegroundColourEscape + "▊\x1b[0m",
	}
	WhiteBishopLightSquare12 = []string{
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "     ▅▄" + LightSquareForegroundColourEscape + "  " + WhitePieceForegroundColourEscape + "  " + LightSquareForegroundColourEscape + "│\x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "   " + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "▗" + WhitePieceBackgroundColourEscape + "  " + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "  " + LightSquareForegroundColourEscape + "│\x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "   " + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▌" + WhitePieceForegroundColourEscape + "  " + LightSquareForegroundColourEscape + "▝" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▅▖" + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + "│\x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "   ▝" + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + "   " + LightSquareBackgroundColourEscape + "▘  " + LightSquareForegroundColourEscape + "│\x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "▁▅▄" + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "━━" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▄▅▁" + LightSquareForegroundColourEscape + " │\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▅▅▅▅▅▅▅▅" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + "▌\x1b[0m",
	}
	WhiteKnightLightSquare8 = []string{
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + " ▄▂▃▃▂" + LightSquareForegroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▌" + WhitePieceForegroundColourEscape + "    " + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + "   " + LightSquareForegroundColourEscape + "▇" + LightSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▃▃▃▃" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▘ " + LightSquareForegroundColourEscape + " \x1b[0m",
	}
	WhiteKnightLightSquare9 = []string{
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  ▗▃▃▃▃▁" + LightSquareForegroundColourEscape + "▊\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▌" + WhitePieceForegroundColourEscape + "     " + LightSquareForegroundColourEscape + "▁" + LightSquareBackgroundColourEscape + " \x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + "   " + LightSquareForegroundColourEscape + " ▇" + LightSquareBackgroundColourEscape + " ▊\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + "┗" + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▃▃▃▃" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  " + LightSquareForegroundColourEscape + " \x1b[0m",
	}
	WhiteKnightLightSquare10 = []string{
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "▗▂▁▂▂▂ " + LightSquareForegroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▊" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▆" + WhitePieceBackgroundColourEscape + "     " + LightSquareForegroundColourEscape + " " + LightSquareBackgroundColourEscape + " \x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▊" + WhitePieceForegroundColourEscape + "    " + LightSquareForegroundColourEscape + "▂▄▆" + LightSquareBackgroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▌" + WhitePieceForegroundColourEscape + "    " + LightSquareForegroundColourEscape + "▝" + LightSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▙▄▄▄▄▄" + LightSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + " \x1b[0m",
	}
	WhiteKnightLightSquare11 = []string{
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + " ▂▁▂▂▂▁" + LightSquareForegroundColourEscape + " ▊\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▊" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▆" + WhitePieceBackgroundColourEscape + "      " + LightSquareForegroundColourEscape + "▝" + LightSquareBackgroundColourEscape + " \x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▊ " + WhitePieceForegroundColourEscape + "    " + LightSquareForegroundColourEscape + "▂▄▆" + LightSquareBackgroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▌" + WhitePieceForegroundColourEscape + "     " + LightSquareBackgroundColourEscape + "▎  " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + "▝" + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▄▄▄▄▄" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▘" + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + " \x1b[0m",
	}
	WhiteKnightLightSquare12 = []string{
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  " + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "▂ " + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "▁▁▁" + LightSquareForegroundColourEscape + "  ▊\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + "╺▄" + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "   " + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▖" + LightSquareForegroundColourEscape + " \x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + "      " + LightSquareForegroundColourEscape + " ▟" + LightSquareBackgroundColourEscape + " \x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▊" + WhitePieceForegroundColourEscape + "     " + LightSquareForegroundColourEscape + "╺▆" + LightSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▊" + WhitePieceForegroundColourEscape + "     " + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "   " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▅▅▅▅▅▅" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▘  " + LightSquareForegroundColourEscape + " \x1b[0m",
	}
	WhiteRookLightSquare8 = []string{
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + "▗▄▗▖▄▖ \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎" + WhitePieceBackgroundColourEscape + "▌▁▁▁▁" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▌ \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▇▇▇" + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎ " + WhitePieceBackgroundColourEscape + "▄▄▄▄" + LightSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + " \x1b[0m",
	}
	WhiteRookLightSquare9 = []string{
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎" + WhitePieceBackgroundColourEscape + "▊" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▄" + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▊" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▄" + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▊" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▄  \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎" + WhitePieceBackgroundColourEscape + "▊▁▁▁▁▁" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▎ \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▌" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▇▇▇▌  \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎ " + WhitePieceBackgroundColourEscape + "▄▄▄▄▄" + LightSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + " \x1b[0m",
	}
	WhiteRookLightSquare10 = []string{
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "│" + WhitePieceForegroundColourEscape + " ▃▖▃▃▗▃  \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "│" + WhitePieceBackgroundColourEscape + "▊" + WhitePieceForegroundColourEscape + "▎" + LightSquareForegroundColourEscape + "▝  ▘" + WhitePieceForegroundColourEscape + "▊" + LightSquareBackgroundColourEscape + "▎ \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "│" + WhitePieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▅━━━━▅" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "│" + WhitePieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▊" + WhitePieceForegroundColourEscape + "    " + LightSquareBackgroundColourEscape + "▎  \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎ " + WhitePieceForegroundColourEscape + "▝" + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▅▅▅▅" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▘ " + LightSquareForegroundColourEscape + " \x1b[0m",
	}
	WhiteRookLightSquare11 = []string{
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "│ " + WhitePieceForegroundColourEscape + "▃▃ ▃ ▃▃" + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "│▊" + WhitePieceForegroundColourEscape + "" + WhitePieceBackgroundColourEscape + "▎ " + LightSquareBackgroundColourEscape + "▅" + WhitePieceBackgroundColourEscape + " " + LightSquareBackgroundColourEscape + "▅" + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + "▊" + LightSquareForegroundColourEscape + "" + LightSquareBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "│" + WhitePieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▅━━━━━▅" + LightSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "│" + WhitePieceForegroundColourEscape + "  " + WhitePieceBackgroundColourEscape + "     " + LightSquareBackgroundColourEscape + "   \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + "  " + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▅▅▅▅▅" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  " + LightSquareForegroundColourEscape + " \x1b[0m",
	}
	WhiteRookLightSquare12 = []string{
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "│" + WhitePieceForegroundColourEscape + " ▂▂ ▂▂ ▂▂" + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "│" + WhitePieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + "┗" + WhitePieceForegroundColourEscape + "▎▊" + LightSquareForegroundColourEscape + "┛" + WhitePieceForegroundColourEscape + " " + LightSquareBackgroundColourEscape + "▊  \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "│" + WhitePieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▎▁▁▁▁▁▁" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▊" + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "│" + WhitePieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▌" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▆▆▆▆▌" + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "  \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "│" + WhitePieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + "    " + LightSquareForegroundColourEscape + " " + LightSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + "  \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▅▅▅▅▅▅" + LightSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + " \x1b[0m",
	}
	WhiteKingLightSquare8 = []string{
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + "  ▄▄" + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "  \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + "▃▄" + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▜▛" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▄▄" + LightSquareForegroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▊" + WhitePieceForegroundColourEscape + "   " + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "  " + LightSquareBackgroundColourEscape + "▎\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎" + WhitePieceBackgroundColourEscape + "▙▃▃▃▃▟" + LightSquareBackgroundColourEscape + " \x1b[0m",
	}
	WhiteKingLightSquare9 = []string{
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "│" + WhitePieceForegroundColourEscape + "  ▗▅▖   \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + "▂▄▃" + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▆" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▃▄▂" + LightSquareForegroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▊" + WhitePieceForegroundColourEscape + "   " + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "   " + LightSquareBackgroundColourEscape + "▎\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎" + WhitePieceBackgroundColourEscape + "▙▃▃▃▃▃▟" + LightSquareBackgroundColourEscape + " \x1b[0m",
	}
	WhiteKingLightSquare10 = []string{
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "│" + WhitePieceForegroundColourEscape + "   ▄▄" + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "   \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "│ " + WhitePieceForegroundColourEscape + "▁ " + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▄▄" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + " ▁" + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▘" + WhitePieceForegroundColourEscape + "▇ " + LightSquareBackgroundColourEscape + "▙▟" + WhitePieceBackgroundColourEscape + "  " + LightSquareForegroundColourEscape + "▝" + LightSquareBackgroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▖" + WhitePieceForegroundColourEscape + "      " + LightSquareForegroundColourEscape + " " + LightSquareBackgroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎" + WhitePieceBackgroundColourEscape + "▇▄▄▄▄▄▄▇" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + " \x1b[0m",
	}
	WhiteKingLightSquare11 = []string{
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "│" + WhitePieceForegroundColourEscape + "   ▗▄▖" + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "   \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "│ " + WhitePieceForegroundColourEscape + "▁▁" + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▙▃▟" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▁▁" + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + "▟" + WhitePieceBackgroundColourEscape + "  " + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▄▇" + WhitePieceBackgroundColourEscape + "  " + LightSquareForegroundColourEscape + "▝" + LightSquareBackgroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎" + WhitePieceBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + "       " + LightSquareForegroundColourEscape + "▗" + LightSquareBackgroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▄▄▄▄▄▄▄" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  \x1b[0m",
	}
	WhiteKingLightSquare12 = []string{
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "│" + WhitePieceForegroundColourEscape + "    ▃▃     \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "│" + WhitePieceForegroundColourEscape + "  " + LightSquareForegroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▊▁▁" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▎" + LightSquareForegroundColourEscape + "  " + WhitePieceForegroundColourEscape + "  \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + "▗▅▆▅▖▗▅▆▅▖" + LightSquareForegroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎" + WhitePieceBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + "   " + LightSquareForegroundColourEscape + "  " + WhitePieceForegroundColourEscape + "   " + LightSquareBackgroundColourEscape + "▊" + LightSquareForegroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎" + WhitePieceBackgroundColourEscape + "▌" + WhitePieceForegroundColourEscape + "        " + LightSquareForegroundColourEscape + "▟" + LightSquareBackgroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "▎ " + WhitePieceBackgroundColourEscape + "▅▅▅▅▅▅▅▅" + LightSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + " \x1b[0m",
	}
	WhiteQueenLightSquare8 = []string{
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + "      " + LightSquareForegroundColourEscape + " \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▝" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▃" + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "  " + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▃" + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▘" + LightSquareBackgroundColourEscape + " \x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + "    " + LightSquareBackgroundColourEscape + "▊ \x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + " ▝" + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▃▃▃▃" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▘" + LightSquareForegroundColourEscape + " \x1b[0m",
	}
	WhiteQueenLightSquare9 = []string{
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "  " + WhitePieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "▁" + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + "  \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▄▟" + WhitePieceBackgroundColourEscape + " " + LightSquareBackgroundColourEscape + "▙▄▊" + LightSquareForegroundColourEscape + " \x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▌" + WhitePieceForegroundColourEscape + "     " + LightSquareBackgroundColourEscape + "▌" + LightSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + " ▝" + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▃▃▃▃▃" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▘" + LightSquareForegroundColourEscape + " \x1b[0m",
	}
	WhiteQueenLightSquare10 = []string{
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + "  " + WhitePieceForegroundColourEscape + "      " + LightSquareForegroundColourEscape + " ▊\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▎" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▖ " + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▘▝" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + " ▗▊" + LightSquareForegroundColourEscape + " \x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▌" + WhitePieceForegroundColourEscape + " " + LightSquareBackgroundColourEscape + "▇" + WhitePieceBackgroundColourEscape + " ▇" + LightSquareBackgroundColourEscape + "▇" + WhitePieceBackgroundColourEscape + " " + LightSquareBackgroundColourEscape + "▌" + LightSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▊" + WhitePieceForegroundColourEscape + "      " + LightSquareBackgroundColourEscape + "▎" + LightSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "  " + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▄▄▄▄▄▄" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + " \x1b[0m",
	}
	WhiteQueenLightSquare11 = []string{
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "  " + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + " " + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "  " + LightSquareForegroundColourEscape + "  \x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + "▜▂" + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "▗" + WhitePieceBackgroundColourEscape + " " + LightSquareBackgroundColourEscape + "▖" + LightSquareForegroundColourEscape + " " + WhitePieceForegroundColourEscape + "▂▛" + LightSquareForegroundColourEscape + " \x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▊" + WhitePieceForegroundColourEscape + " " + LightSquareBackgroundColourEscape + "▇" + WhitePieceBackgroundColourEscape + "   " + LightSquareBackgroundColourEscape + "▇" + WhitePieceBackgroundColourEscape + " " + LightSquareBackgroundColourEscape + "▌" + LightSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▊" + WhitePieceForegroundColourEscape + "       " + LightSquareBackgroundColourEscape + "▎" + LightSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▄▄▄▄▄▄▄" + LightSquareBackgroundColourEscape + "  \x1b[0m",
	}
	WhiteQueenLightSquare12 = []string{
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "           " + LightSquareForegroundColourEscape + "▊\x1b[0m",
		LightSquareForegroundColourEscape + LightSquareBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + "▗▁" + LightSquareForegroundColourEscape + "  " + WhitePieceBackgroundColourEscape + "▘▝" + LightSquareBackgroundColourEscape + "  " + WhitePieceForegroundColourEscape + "▁▖" + LightSquareForegroundColourEscape + " \x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▊ " + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▄" + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▘" + WhitePieceForegroundColourEscape + "  " + LightSquareForegroundColourEscape + "▝" + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▄" + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + "▎" + LightSquareForegroundColourEscape + " \x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + WhitePieceBackgroundColourEscape + "▊" + WhitePieceForegroundColourEscape + "        " + LightSquareBackgroundColourEscape + "▎" + LightSquareForegroundColourEscape + "▊\x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + " " + WhitePieceBackgroundColourEscape + " " + WhitePieceForegroundColourEscape + "      " + LightSquareForegroundColourEscape + " " + LightSquareBackgroundColourEscape + " ▊\x1b[0m",
		WhitePieceForegroundColourEscape + LightSquareBackgroundColourEscape + " " + LightSquareForegroundColourEscape + " " + WhitePieceBackgroundColourEscape + "▅▅▅▅▅▅▅▅" + LightSquareBackgroundColourEscape + "  \x1b[0m",
	}
)
