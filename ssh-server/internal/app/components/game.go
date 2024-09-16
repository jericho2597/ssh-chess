package components

import (
	"log"
	"ssh-server/config"
	"ssh-server/internal/app/board_render"
	"ssh-server/internal/app/model"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/notnil/chess"
)

type Game struct {
	State    *chess.Game

	PlayerColor chess.Color
	SelectedPiece *chess.Square
}

func NewGame() *Game {
	return &Game{
		State: chess.NewGame(),

		PlayerColor: chess.White,
	}
}

func (game *Game) View(state model.TuiState) string {

	var sb strings.Builder

	if(game.State.Outcome() == chess.Draw) {
		sb.WriteString(game.renderOutcome("Draw!"))
	} else if (game.State.Outcome() == chess.WhiteWon) {
		sb.WriteString(game.renderOutcome("Winner!"))
	} else if (game.State.Outcome() == chess.BlackWon) {
		sb.WriteString(game.renderOutcome("Loser!"))
	} else {
		sb.WriteString(game.renderBoard(state.SquareSize))
	}

	return sb.String()
}

func (game *Game) renderOutcome(outcomeMessage string) string {
	return outcomeMessage
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
			squareColor := (rank + file + 1) % 2
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
		return board_render.GetEmptySquareRenderingFunc(isLightSquare, squareSize)()
	}

	pieceType := piece.Type().String()
	pieceColor := piece.Color() == chess.White

	return board_render.GetPieceRenderingFunc(pieceColor, isLightSquare, squareSize, pieceType)()
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
		case "n":
			// Process move
			game.State = chess.NewGame()
			game.SelectedPiece = nil
			game.PlayerColor = chess.White
		}
	}

	return nil
}

func (game *Game) handleMouseClick(x int, y int, tuiState model.TuiState) {

	squareSize := tuiState.SquareSize

	// gets the file from 0 (left most file) to 7 (right most file)
	file := ((x - ((tuiState.Width - (squareSize * 8)) / 2)) / squareSize)
	// gets the rank from 0 (top most rank) to 7 (bottom rank)
	rank := 7 - ((y - config.HeaderHeight - ((tuiState.Height - config.HeaderHeight - config.FooterHeight - squareSize*4) / 2)) / (squareSize / 2))

	if file < 0 || file > 7 || rank < 0 || rank > 7 {
		return // Click was outside the board
	}

	square := chess.Square(rank*8 + file)

	piece := game.State.Position().Board().Piece(square)
	if piece != chess.NoPiece && piece.Color() == game.PlayerColor {
		game.SelectedPiece = &square

	} else if game.SelectedPiece != nil {
		
		validMoves := game.State.ValidMoves()

		for _, move := range validMoves {

			if move.S1() == *game.SelectedPiece && move.S2() == square {

				err := game.State.Move(move)
				if err != nil {
					log.Println("Invalid move: ", err)
					return
				}

				// Reset the selected piece after the move is made
				game.SelectedPiece = nil
				if game.PlayerColor == chess.White {
					game.PlayerColor = chess.Black
				} else {
					game.PlayerColor = chess.White
				}

				return
			}
		}
	}
}
