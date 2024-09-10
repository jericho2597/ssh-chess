package components

import (
	"fmt"
	"log"
	"ssh-server/config"
	"ssh-server/internal/app/board_render"
	"ssh-server/internal/app/model"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/notnil/chess"
)

type Game struct {
	State    *chess.Game
	Input    string
	ErrorMsg string

	SelectedPiece *chess.Square

	BlackPieceColor  lipgloss.Color
	WhitePieceColor  lipgloss.Color
	DarkSquareColor  lipgloss.Color
	LightSquareColor lipgloss.Color
}

func NewGame() *Game {
	return &Game{
		State: chess.NewGame(),

		BlackPieceColor:  lipgloss.Color("#00050f"),
		WhitePieceColor:  lipgloss.Color("#ffffff"),
		DarkSquareColor:  lipgloss.Color("#3d2f0a"),
		LightSquareColor: lipgloss.Color("#bdaf8d"),
	}
}

func (game *Game) View(state model.TuiState) string {

	var sb strings.Builder

	sb.WriteString(game.renderBoard(state.SquareSize))

	if game.ErrorMsg != "" {
		sb.WriteString(fmt.Sprintf("Error: %s\n", game.ErrorMsg))
	}

	style1 := lipgloss.NewStyle().Foreground(lipgloss.Color("#00050f")).Background(lipgloss.Color("#bdaf8d"))
	style2 := lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff")).Background(lipgloss.Color("#3d2f0a"))

	sb.WriteString(style1.Render("Hello there\n"))
	sb.WriteString(style2.Render("Hello there\n"))

	return sb.String()
}

func (game *Game) renderBoard(squareSize int) string {
	board := game.State.Position().Board()
	var sb strings.Builder

	// terminal cells are twice as tall as they are weight, so for
	// a square with 'squareSize' cells (in width), the height (in cells)
	// is that over 2
	squareHeightInCells := squareSize / 2

	for rank := 7; rank >= 0; rank-- {
		rankPieces := make([][]string, squareHeightInCells)
		for file := 0; file < 8; file++ {
			piece := board.Piece(chess.Square(rank*8 + file))
			squareColor := (rank + file) % 2
			pieceStrings := game.getPieceString(piece, squareColor == 0, squareSize)
			for i, line := range pieceStrings {
				rankPieces[i] = append(rankPieces[i], line)
			}
		}
		for _, line := range rankPieces {
			sb.WriteString(strings.Join(line, ""))
			sb.WriteString("\n")
		}
	}

	return sb.String()
}

func (game *Game) getPieceString(piece chess.Piece, isLightSquare bool, squareSize int) []string {

	if piece.Type() == chess.NoPieceType {
		return board_render.GetEmptySquareRenderingFunc(isLightSquare, squareSize)(game.WhitePieceColor, game.BlackPieceColor, game.LightSquareColor, game.DarkSquareColor)
	}

	pieceType := piece.Type().String()
	pieceColor := piece.Color() == chess.White

	return board_render.GetPieceRenderingFunc(pieceColor, isLightSquare, squareSize, pieceType)(game.WhitePieceColor, game.BlackPieceColor, game.LightSquareColor, game.DarkSquareColor)
}

func (game *Game) Update(msg tea.Msg, tuiState model.TuiState) tea.Cmd {

	switch msg := msg.(type) {
	case tea.MouseMsg:
		if tea.MouseEvent(msg).Action == tea.MouseActionRelease {
			game.handleMouseClick(msg.X, msg.Y, tuiState)
		}
		return nil
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			// Process move
			move, err := chess.UCINotation{}.Decode(game.State.Position(), game.Input)
			if err != nil {
				game.ErrorMsg = fmt.Sprintf("Invalid move: %v", err)
			} else if err := game.State.Move(move); err != nil {
				game.ErrorMsg = fmt.Sprintf("Illegal move: %v", err)
			} else {
				game.Input = ""
				game.ErrorMsg = ""
			}
		case "backspace":
			if len(game.Input) > 0 {
				game.Input = game.Input[:len(game.Input)-1]
			}
		default:
			if len(msg.String()) == 1 {
				game.Input += msg.String()
			}
		}
	}

	return nil
}

func (game *Game) handleMouseClick(x int, y int, tuiState model.TuiState) {

	squareSize := tuiState.SquareSize

	// gets the file from 0 (left most file) to 7 (right most file)
	file := ((x - ((tuiState.Width - (squareSize * 8)) / 2)) / squareSize)
	// gets the rank from 0 (top most rank) to 7 (bottom rank)
	rank := ((y - config.HeaderHeight - ((tuiState.Height - config.HeaderHeight - squareSize*4) / 2)) / (squareSize / 2))

	if file < 0 || file > 7 || rank < 0 || rank > 7 {
		return // Click was outside the board
	}

	square := chess.Square(rank*8 + file)

	// Now you can use this square with the chess package
	piece := game.State.Position().Board().Piece(square)
	if piece != chess.NoPiece {
		// A piece was clicked
		game.SelectedPiece = &square

		log.Println(piece.String())
	} else if game.SelectedPiece != nil {
		// // An empty square was clicked after a piece was selected, attempt to make a move
		// move := chess.Move{
		// 	s1: *game.SelectedPiece,
		// 	s2: square,
		// }
		// err := game.State.Move(&move)
		// if err != nil {
		// 	// Invalid move, handle error (e.g., display message to user)
		// }
		// game.SelectedPiece = nil // Clear selection after move attempt
	}
}
