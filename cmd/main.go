package main

import (
	b "github.com/guitarkeegan/musings/cmd/board"
	h "github.com/guitarkeegan/musings/cmd/hero"
	v "github.com/guitarkeegan/musings/cmd/view"
)

// how to move a hero from 0,0 to 2,2
// iron out ref vs val for board
// how should be update?

func main() {
	board := b.New(5, 6)
	barbarian := h.NewDefault(h.Barbarian)
	v.Render(board.Place(barbarian, b.Coord{Row: 0, Col: 0}))
	barbarian.AddMovementQuota(4)
	board = board.Move(b.Coord{Row: 0, Col: 0}, b.Coord{Row: 2, Col: 2})
	v.Render(*board)

}
