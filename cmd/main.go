package main

import (
	b "github.com/guitarkeegan/musings/cmd/board"
	h "github.com/guitarkeegan/musings/cmd/hero"
	v "github.com/guitarkeegan/musings/cmd/view"
)

// hero quest
// intro
// 10x10 grid
// 1 hero
// hero has hp, mp, attack die, defence die
// can roll to move
// can move and action once per turn
// grid has 1 monster
// monster as set hp, mp, attack, defense

func main() {
	board := b.Make(5, 6)
	barbarian := h.NewDefault(h.Barbarian)
	v.Render(board.Place(b.FromHero(barbarian.Typ), b.Coord{Row: 2, Col: 4}))
}
