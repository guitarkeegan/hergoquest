package board

import "errors"

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

func (p Piece) AvailableMovement() int {
	return 0
}

type Occupier interface {
	Glyph() string
	AvailableMovement() int
}

type Board struct {
	cells [][]Occupier
}

var (
	ErrOutOfBounds = errors.New("coord is out of bounds")
	ErrNoOccupant  = errors.New("no occupant to move")
)

func New(rows, cols int) *Board {
	cells := make([][]Occupier, rows)
	for i := range cells {
		cells[i] = make([]Occupier, cols)
		for j := range cells[i] {
			cells[i][j] = vacant
		}
	}
	return &Board{cells: cells}
}

func (b *Board) Rows() int {
	return len(b.cells)
}

func (b *Board) Cols() int {
	if len(b.cells) == 0 {
		return 0
	}
	return len(b.cells[0])
}

func (b *Board) inBounds(coord Coord) bool {
	return coord.Row >= 0 && coord.Row < b.Rows() &&
		coord.Col >= 0 && coord.Col < b.Cols()
}

func (b *Board) Remove(coord Coord) {
	b.cells[coord.Row][coord.Col] = vacant
}

func (b *Board) Place(o Occupier, coord Coord) {
	b.cells[coord.Row][coord.Col] = o
}

func (b *Board) Get(coord Coord) Occupier {
	return b.cells[coord.Row][coord.Col]
}

// move to algos pkg
// TODO: will need to update when traps, doors, etc. exist
func isValidPath[T any](matrix [][]T, start, end [2]int, movesRemaining int) bool {
	ROWS := len(matrix)
	COLS := len(matrix[0])

	type Step struct {
		pos   [2]int
		moves int
	}

	q := []Step{
		{pos: start, moves: 0},
	}

	DIRECTIONS := [][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	seen := make(map[[2]int]bool)
	seen[start] = true

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		for _, dir := range DIRECTIONS {
			nextRow := cur.pos[0] + dir[0]
			nextCol := cur.pos[1] + dir[1]

			next := [2]int{nextRow, nextCol}

			if nextRow < 0 || nextRow >= ROWS ||
				nextCol < 0 || nextCol >= COLS ||
				seen[next] {
				continue
			}

			nextMoves := cur.moves + 1

			if nextMoves > movesRemaining {
				continue
			}

			if next == end {
				return true
			}

			seen[next] = true

			q = append(q, Step{
				pos:   next,
				moves: nextMoves,
			})
		}
	}

	return false
}

func (b *Board) isValidMove(from, to Coord) bool {
	occupant := b.Get(from)
	return isValidPath(b.cells, [2]int{from.Row, from.Col}, [2]int{to.Row, to.Col}, occupant.AvailableMovement())
}

// Move reports whether the occupant at from could reach to, and moves it if so.
// An error is returned when the move can't be attempted at all.
// TODO: don't forget about occupied cells
func (b *Board) Move(from, to Coord) (bool, error) {
	if !b.inBounds(from) || !b.inBounds(to) {
		return false, ErrOutOfBounds
	}
	occupant := b.Get(from)
	if occupant == vacant {
		return false, ErrNoOccupant
	}
	if !b.isValidMove(from, to) {
		return false, nil
	}
	b.Remove(from)
	b.Place(occupant, to)
	return true, nil
}
