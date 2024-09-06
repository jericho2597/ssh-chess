package components

import (
	"fmt"
	"ssh-server/internal/app/model"
	"ssh-server/internal/app/styles"
	s "ssh-server/internal/app/styles"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
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

	sb.WriteString("\n")
	sb.WriteString(styles.BlackBishopLightSquare8)
	sb.WriteString("\n")
	sb.WriteString(styles.WhiteBishopLightSquare8)
	sb.WriteString("\n")
	sb.WriteString(styles.BlackBishopDarkSquare8)
	sb.WriteString("\n")
	sb.WriteString(styles.WhiteBishopDarkSquare8)
	sb.WriteString("\n")

	return sb.String()
}

func (game *Game) renderBoard() string {
	board := game.State.Position().Board()

	table := table.New().Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(s.Border))

	for rank := 7; rank >= 0; rank-- {
		row := make([]string, 8)
		for file := 0; file < 8; file++ {
			piece := board.Piece(chess.Square(rank*8 + file))
			row[file] = game.getPieceChar(piece)
		}
		table = table.Row(row...)
	}

	return table.Render()
}

func (game *Game) getPieceChar(piece chess.Piece) string {
	switch piece.Type() {
	case chess.King:
		if piece.Color() == chess.White {
			return "♔"
		}
		return "♚"
	case chess.Queen:
		if piece.Color() == chess.White {
			return "♕"
		}
		return "♛"
	case chess.Rook:
		if piece.Color() == chess.White {
			return "♖"
		}
		return "♜"
	case chess.Bishop:
		if piece.Color() == chess.White {
			return "♗"
		}
		return "♝"
	case chess.Knight:
		if piece.Color() == chess.White {
			return "♘"
		}
		return "♞"
	case chess.Pawn:
		if piece.Color() == chess.White {
			return "♙"
		}
		return "♟"
	default:
		return "·"
	}
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
