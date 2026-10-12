package board

type Coord struct {
	Row, Col int
}

type Piece int

const (
	vacant Piece = iota
	closedDoor
	openDoor
)

func (p Piece) Glyph() string {
	switch p {
	case vacant:
		return "[]"
	case closedDoor:
		return "--"
	case openDoor:
		return "()"
	default:
		return "??"
	}
}

type Occupier interface {
	Glyph() string
}

type Board [][]Occupier

func Make(rows, cols int) Board {
	board := make(Board, rows)
	for i := range board {
		board[i] = make([]Occupier, cols)
	}
	for i := range board {
		for j := range board[i] {
			board[i][j] = vacant
		}
	}
	return board
}

func (b Board) Remove(coord Coord) Board {
	b[coord.Row][coord.Col] = vacant
	return b
}

func (b Board) Place(o Occupier, coord Coord) Board {
	b[coord.Row][coord.Col] = o
	return b
}

func (b Board) Get(coord Coord) Occupier {
	return b[coord.Row][coord.Col]
}

func (b Board) Move(from, to Coord) {
	occupant := b.Get(from)
	b.Remove(from)
	b.Place(occupant, to)
}
