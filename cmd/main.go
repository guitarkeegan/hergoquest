package main

import (
	b "github.com/guitarkeegan/musings/cmd/board"
	h "github.com/guitarkeegan/musings/cmd/hero"
	v "github.com/guitarkeegan/musings/cmd/view"
)

// how to move a hero from 0,0 to 2,2

func main() {
	board := b.Make(5, 6)
	barbarian := h.NewDefault(h.Barbarian)
	v.Render(board.Place(barbarian, b.Coord{Row: 0, Col: 0}))
}
