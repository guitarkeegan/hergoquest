package board

import (
	"errors"
	"github.com/guitarkeegan/hergoquest/internal/dice"
	"github.com/guitarkeegan/hergoquest/internal/log"
	"math"
)

// TODO: should probably live here
// so heros only need to know the number of dice
// for a given action. 2 for move, 3 for attact etc.

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

type Combatant interface {
	Occupier
	Attack() int
	Defend() int
	IsDead() bool
	TakeDamage(int)
	RemainingHP() int
	Moved(...bool) bool
	// attacked, searched for traps, treasure, cast a spell
	Acted(...bool) bool
	// TODO: will need to think on this one
	// Spell() string
}

type turnType string

const (
	heroTurn    turnType = "hero"
	monsterTurn turnType = "monster"
)

type Board struct {
	cells [][]Occupier
	// INFO: where to store the state?
	heros              []Combatant
	monsters           []Combatant
	combatantPositions map[Combatant]Coord
	turn               turnType
	dice               *dice.Rollist
}

// turn is derived from the heros available movements
// INFO: what if a hero decided not to move?
// we must make them exhast their movemnts in that case.
func (b *Board) Turn() turnType {
	for _, h := range b.heros {
		if h.AvailableMovement() > 0 {
			return heroTurn
		}
	}
	return monsterTurn
}

func (b *Board) startMonstersTurn() {}

func (b *Board) startHerosTurn() {
	// for _, h := range b.heros {
	// TODO: reset available movements
	// }
}

func (b *Board) endTurn() {
	currentTurn := b.Turn()
	if currentTurn == heroTurn {
		// set monsters turn
	}
}

var (
	ErrOutOfBounds       = errors.New("coord is out of bounds")
	ErrNoOccupant        = errors.New("no occupant to move")
	ErrNonCombatant      = errors.New("not a combatant interface")
	ErrCombatantNotFound = errors.New("combatant not found")
)

// maybe add a config for the start of games...
func New(rows, cols int) *Board {
	cells := make([][]Occupier, rows)
	for i := range cells {
		cells[i] = make([]Occupier, cols)
		for j := range cells[i] {
			cells[i][j] = vacant
		}
	}
	return &Board{
		cells:              cells,
		heros:              []Combatant{},
		monsters:           []Combatant{},
		combatantPositions: make(map[Combatant]Coord),
		dice:               dice.New(),
	}
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

func (b *Board) PlaceHero(c Combatant, coord Coord) {
	b.cells[coord.Row][coord.Col] = c
	b.heros = append(b.heros, c)
	b.combatantPositions[c] = coord
}

func (b *Board) PlaceMonster(c Combatant, coord Coord) {
	b.cells[coord.Row][coord.Col] = c
	b.monsters = append(b.monsters, c)
	b.combatantPositions[c] = coord
}

func (b *Board) Place(o Occupier, coord Coord) {
	b.cells[coord.Row][coord.Col] = o
}

func (b *Board) PlaceCombatant(c Combatant, coord Coord) {
	b.cells[coord.Row][coord.Col] = c
}

// returns the Occupier given a coord, to get the coord
// of a combatant, use GetCombatant
func (b *Board) Get(coord Coord) Occupier {
	return b.cells[coord.Row][coord.Col]
}

func (b *Board) GetCombatant(c Combatant) (Coord, error) {
	if coord, ok := b.combatantPositions[c]; ok {
		return coord, nil
	}
	// TODO: add to errors
	return Coord{}, ErrCombatantNotFound
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

func (b *Board) Attack(attackCoord, defendCoord Coord) {

	// if close enough
	isAdjacent := math.Abs(float64(attackCoord.Row)-float64(defendCoord.Row)) <= 1 &&
		math.Abs(float64(attackCoord.Col)-float64(defendCoord.Col)) <= 1

	if !isAdjacent {
		announce.Log("no enemy in range")
		return
	}

	a, ok := b.Get(attackCoord).(Combatant)
	if !ok {
		//TODO: handle non combatant
		return
	}

	d, ok := b.Get(defendCoord).(Combatant)
	if !ok {
		//TODO: handle non combatant
		return
	}

	attack, err := b.dice.RollCombat(a.Attack())
	if err != nil {
		panic("handle roll combat err")
	}

	// updated attackers turn
	a.Acted(true)

	defend, err := b.dice.RollCombat(d.Defend())
	if err != nil {
		panic("handle roll combat err")
	}

	// hero or monster, derive from turn
	var difference int
	if b.turn == heroTurn {
		difference = attack.Attack - defend.EvilDefend
	} else {
		difference = attack.Attack - defend.GoodDefend
	}

	if difference > 0 {
		d.TakeDamage(difference)
		announce.Log("took damage, %d remaining.", d.RemainingHP())
	} else {
		announce.Log("no damage, %d remaining.", d.RemainingHP())
	}

	// TODO:  always call clean up?
	b.update()

}

// return a list of living heros/monsters
func filterOutFallen[T comparable](fallen T, pile []T) []T {
	out := make([]T, 0, len(pile))
	for _, body := range pile {
		if fallen != body {
			out = append(out, body)
		}
	}
	return out
}

// can be generalized, but essentially clean up the
// fallen combatants
// TODO: add turn check
func (b *Board) update() {
	announce.Log("updating...")
	for _, h := range b.heros {
		if h.IsDead() {
			if coord, err := b.GetCombatant(h); err == nil {
				b.Remove(coord)
				delete(b.combatantPositions, h)
				b.heros = filterOutFallen(h, b.heros)
			}
		}
	}
	for _, m := range b.monsters {
		if m.IsDead() {
			if coord, err := b.GetCombatant(m); err == nil {
				b.Remove(coord)
				delete(b.combatantPositions, m)
				b.monsters = filterOutFallen(m, b.monsters)
			}
		}
	}
	announce.Log("heros remaining: %d", len(b.heros))
	announce.Log("monsters remaining: %d", len(b.monsters))

	switch b.turn {
	case heroTurn:
		var allDone bool
		for _, h := range b.heros {
			if h.Acted() && h.Moved() {
				allDone = true
			} else {
				allDone = false
				break
			}
		}
		if allDone {
			// TODO: restart with the next turn
			b.turn = monsterTurn
		}
	case monsterTurn:
		var allDone bool
		for _, m := range b.monsters {
			if m.Acted() && m.Moved() {
				allDone = true
			} else {
				allDone = false
				break
			}
		}
		if allDone {
			// TODO: restart with the next turn
			b.turn = heroTurn
		}
	}

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
	oc, ok := occupant.(Combatant)
	if !ok {
		return false, ErrNonCombatant
	}
	b.combatantPositions[oc] = to
	oc.Moved(true)
	b.Remove(from)
	b.Place(occupant, to)
	return true, nil
}
