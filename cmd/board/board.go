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

func (p Piece) AvailableMovement() int {
	return 0
}

type Occupier interface {
	Glyph() string
	AvailableMovement() int
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

func (b Board) isValidMove(from, to Coord) bool {
	// same call?
	occupant := b.Get(from)
	if isValidPath(b, [2]int{from.Row, from.Col}, [2]int{to.Row, to.Col}, occupant.AvailableMovement()) {
		return true
	}
	return false
}

// return whether or not the move was possible
func (b Board) Move(from, to Coord) bool {
	occupant := b.Get(from)
	if b.isValidMove(from, to) {
		b.Remove(from)
		b.Place(occupant, to)
		return true
	}
	return false
}
