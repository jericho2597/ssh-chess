package components

import (
	"fmt"
	"ssh-server/internal/app/board_render"
	"ssh-server/internal/app/model"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/notnil/chess"
)

type Game struct {
	State    *chess.Game
	Input    string
	ErrorMsg string
}

func NewGame() *Game {
	return &Game{
		State: chess.NewGame(),
	}
}

func (game *Game) View(state model.TuiState) string {

	var sb strings.Builder

	// Display the chess board
	sb.WriteString(game.renderBoard())
	sb.WriteString("\n\n")

	// Display the current player's turn
	sb.WriteString(fmt.Sprintf("Current turn: %s\n", game.State.Position().Turn()))
	sb.WriteString("\n")

	// Display the move input field
	sb.WriteString(fmt.Sprintf("Enter your move: %s\n", game.Input))

	// Display any error messages
	if game.ErrorMsg != "" {
		sb.WriteString(fmt.Sprintf("Error: %s\n", game.ErrorMsg))
	}

	return sb.String()
}

func (game *Game) renderBoard() string {
	board := game.State.Position().Board()
	var sb strings.Builder

	for rank := 7; rank >= 0; rank-- {
		rankPieces := make([][]string, 6) // 6 lines per square
		for file := 0; file < 8; file++ {
			piece := board.Piece(chess.Square(rank*8 + file))
			squareColor := (rank + file) % 2
			pieceStrings := game.getPieceString(piece, squareColor == 0)
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

func (game *Game) getPieceString(piece chess.Piece, isLightSquare bool) []string {

	if piece.Type() == chess.NoPieceType {
		return board_render.GetEmptySquareRendering(isLightSquare, 12)
	}

	pieceType := piece.Type().String()
	pieceColor := piece.Color() == chess.White
	size := 12

	return board_render.GetPieceRendering(pieceColor, isLightSquare, size, pieceType)
}

func (game *Game) Update(msg tea.Msg) tea.Cmd {

	switch msg := msg.(type) {
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
