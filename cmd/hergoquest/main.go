package main

import (
	"fmt"

	b "github.com/guitarkeegan/hergoquest/internal/board"
	h "github.com/guitarkeegan/hergoquest/internal/hero"
	v "github.com/guitarkeegan/hergoquest/internal/view"
)

// A mini game.
// one hero, with one randomly placed monster
// hero moves first, and the hero and monster take turns
// moving and/or attacking, until someone wins
// battle will be automated

func main() {
	board := b.New(5, 6)
	barbarian := h.NewDefault(h.Barbarian)
	board.Place(barbarian, b.Coord{Row: 0, Col: 0})
	v.Render(board)
	barbarian.AddMovementQuota(4)
	moved, err := board.Move(b.Coord{Row: 0, Col: 0}, b.Coord{Row: 2, Col: 2})
	if err != nil {
		fmt.Println("move failed:", err)
		return
	}
	if !moved {
		fmt.Println("not enough movement to reach destination")
	}
	v.Render(board)

}
