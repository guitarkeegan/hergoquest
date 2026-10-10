package board

import "strings"

type SquareState int

type Coord struct {
	Row, Col int
}

const (
	empty SquareState = iota
	occupiedBarbarian
	trap
	furnature
	stares
	chest
)

var squareStateView = map[SquareState]string{
	empty:             "[]",
	occupiedBarbarian: "Ba",
	trap:              "XX",
	furnature:         "MM",
	stares:            "Ss",
	chest:             "CH",
}

type Board [][]SquareState

func Make(rows, cols int) Board {
	board := make(Board, rows)
	for i := range board {
		board[i] = make([]SquareState, cols)
	}
	return board
}

func (b Board) Remove(coord Coord) Board {
	b[coord.Row][coord.Col] = empty
	return b
}

func (b Board) Place(sqrState SquareState, coord Coord) Board {
	b[coord.Row][coord.Col] = sqrState
	return b
}

func (b Board) Get(coord Coord) SquareState {
	return b[coord.Row][coord.Col]
}

func (b Board) Move(from, to Coord) {
	squareState := b.Get(from)
	b.Remove(from)
	b.Place(squareState, to)
}

func (b Board) View() string {
	var brd strings.Builder
	for i := range b {
		for j := 0; j < len(b[0]); j++ {
			brd.WriteString(squareStateView[b[i][j]])
			brd.WriteString(" ")
		}
		brd.WriteString("\n")
	}
	return brd.String()
}
