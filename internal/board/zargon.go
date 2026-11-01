package board

import "math"

// how does zargon work...

func (b *Board) zargonMove() {

	for _, m := range b.monsters {
		if !m.IsDead() && m.IsVisible() {
			b.moveToHero(m)
		}
	}

}

// returns the number of moves needed to reach the combatant
// returns -1 of no combantant is found, which should never happen
func (b *Board) distance(start, end Coord) int {
	ROWS := len(b.cells)
	COLS := len(b.cells[0])

	type Step struct {
		pos   Coord
		moves int
	}

	q := []Step{
		{pos: start, moves: 0},
	}

	seen := make(map[Coord]bool)
	seen[start] = true

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		for _, dir := range DIRECTIONS {
			nextRow := cur.pos.Row + dir[0]
			nextCol := cur.pos.Col + dir[1]

			next := Coord{Row: nextRow, Col: nextCol}

			if nextRow < 0 || nextRow >= ROWS ||
				nextCol < 0 || nextCol >= COLS ||
				b.IsWall(cur.pos, next) ||
				seen[next] {
				continue
			}

			nextMoves := cur.moves + 1

			if next == end {
				// WARN: check this
				return nextMoves
			}

			seen[next] = true

			q = append(q, Step{
				pos:   next,
				moves: nextMoves,
			})
		}
	}

	return -1
}

func (b *Board) findClosestHero(monster Combatant) Coord {

	minMoves := math.MaxInt
	var closest Coord

	for _, hero := range b.heroes {
		h, err := b.GetCombatant(hero)
		if err != nil {
			panic("no heros")
		}
		m, err := b.GetCombatant(monster)
		if err != nil {
			panic("no monsters")
		}
		d := b.distance(m, h)
		if d < minMoves {
			minMoves = d
			closest = h
		}
	}

	return closest
}

func (b *Board) moveToHero(monster Combatant) {
	// find closest hero
	// heroCoord := b.findClosestHero(monster)
	// move towards them until movement runs out

	// if adjacent, attack
}

// for every monster
// if monster is visible
// move to hero
// attack the closest hero
// or attack the weakest hero
